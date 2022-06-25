package projectory

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/elamre/tentsuyu"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

type BasicProjectile struct {
	speed          float64
	pixelX, pixelY float64
	dX, dY         float64
	angle          float64
	dst            ebiten.DrawImageOptions
	image          *ebiten.Image
	hitBox         *tentsuyu.Rectangle
	lifeCounter    int
}

func NewBasicProjectile(startPixelX, startPixelY, targetPixelX, targetPixelY float64) *BasicProjectile {
	b := BasicProjectile{
		speed: 3,
		image: assets.Get[*ebiten.Image](assets.AssetsTurretGun_light_bullet),
	}
	mouseXFloat := startPixelX - targetPixelX
	mouseYFloat := startPixelY - targetPixelY

	b.angle = math.Atan2(mouseYFloat, mouseXFloat)

	b.pixelX = startPixelX - 3
	b.pixelY = startPixelY
	b.hitBox = &tentsuyu.Rectangle{X: b.pixelX, Y: b.pixelY, W: 9, H: 9}
	b.dst.GeoM.Rotate(b.angle)
	b.dst.GeoM.Translate(b.pixelX, b.pixelY)
	//
	b.dX = math.Cos(b.angle)
	b.dY = math.Sin(b.angle)
	return &b
}

func (b *BasicProjectile) GetHitBox() *tentsuyu.Rectangle {
	return b.hitBox
}
func (b *BasicProjectile) IsAlive() bool {
	return b.lifeCounter < 100
}
func (b *BasicProjectile) Update() {
	travelX := -(b.dX * b.speed)
	travelY := -(b.dY * b.speed)
	b.pixelX += travelX
	b.pixelY += travelY
	b.dst.GeoM.Translate(travelX, travelY)
	b.lifeCounter++
}
func (b *BasicProjectile) Draw(image *ebiten.Image) {
	image.DrawImage(b.image, &b.dst)
}
func (b *BasicProjectile) GetProjectileEffect() *ProjectileEffect {
	return nil
}
