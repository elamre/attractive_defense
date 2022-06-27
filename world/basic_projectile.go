package world

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/elamre/tentsuyu"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

type BasicProjectile struct {
	damages        *ProjectileEffect
	pixelX, pixelY float64
	dX, dY         float64
	angle          float64
	dst            ebiten.DrawImageOptions
	image          *ebiten.Image
	hitBox         *tentsuyu.Rectangle
	lifeCounter    int
}

func NewSmallProjectile(startPixelX, startPixelY, targetPixelX, targetPixelY float64, effect *ProjectileEffect, lifeCounter int) *BasicProjectile {
	b := BasicProjectile{lifeCounter: lifeCounter, image: assets.Get[*ebiten.Image](assets.AssetsTurretGun_light_bullet), damages: effect}
	b.calculateProjectory(startPixelX, startPixelY, targetPixelX, targetPixelY)
	return &b
}

func NewBasicProjectile(startPixelX, startPixelY, targetPixelX, targetPixelY float64) *BasicProjectile {
	b := BasicProjectile{
		damages:     &ProjectileEffect{Damage: 5, Speed: 5},
		image:       assets.Get[*ebiten.Image](assets.AssetsTurretGun_light_bullet),
		lifeCounter: 100,
	}
	b.calculateProjectory(startPixelX, startPixelY, targetPixelX, targetPixelY)
	return &b
}

func (b *BasicProjectile) calculateProjectory(startPixelX, startPixelY, targetPixelX, targetPixelY float64) {
	deltaX := startPixelX - targetPixelX
	deltaY := startPixelY - targetPixelY

	b.angle = math.Atan2(deltaY, deltaX)

	b.pixelX = startPixelX - 3
	b.pixelY = startPixelY
	b.hitBox = &tentsuyu.Rectangle{X: b.pixelX, Y: b.pixelY, W: 9, H: 9}
	b.dst.GeoM.Rotate(b.angle)
	b.dst.GeoM.Translate(b.pixelX, b.pixelY)
	//
	b.dX = math.Cos(b.angle)
	b.dY = math.Sin(b.angle)
}

func (b *BasicProjectile) GetHitBox() *tentsuyu.Rectangle {
	return b.hitBox
}
func (b *BasicProjectile) IsAlive() bool {
	return b.lifeCounter > 0
}
func (b *BasicProjectile) Update() {
	travelX := -(b.dX * b.damages.Speed)
	travelY := -(b.dY * b.damages.Speed)
	b.pixelX += travelX
	b.pixelY += travelY
	b.dst.GeoM.Translate(travelX, travelY)
	b.lifeCounter--
	b.hitBox.X = b.pixelX
	b.hitBox.Y = b.pixelY
}
func (b *BasicProjectile) Draw(image *ebiten.Image) {
	image.DrawImage(b.image, &b.dst)
}
func (b *BasicProjectile) GetProjectileEffect() *ProjectileEffect {
	return b.damages
}

func (b *BasicProjectile) Impact() {
	b.lifeCounter = 0
}
