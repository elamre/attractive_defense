package world

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/elamre/tentsuyu"
	"github.com/hajimehoshi/ebiten/v2"
)

type LaserProjectile struct {
	damages          *ProjectileEffect
	pixelX, pixelY   float64
	targetX, targetY float64
	angle            float64
	dst              ebiten.DrawImageOptions
	image            *ebiten.Image
	hitBox           *tentsuyu.Rectangle
}

func NewLaserProjectile(x, y, tX, tY float64) *LaserProjectile {
	b := LaserProjectile{image: assets.Get[*ebiten.Image](assets.AssetsTurretGun_heavy_bullet), damages: &ProjectileEffect{}, hitBox: &tentsuyu.Rectangle{X: tX, Y: tY, W: 128, H: 128}}
	b.targetY = tY
	b.targetX = tX
	return &b
}

func (b *LaserProjectile) GetHitBox() *tentsuyu.Rectangle {
	return b.hitBox
}
func (b *LaserProjectile) IsAlive() bool {
	return true
}

func (b *LaserProjectile) UpdateProjectilEffect(p *ProjectileEffect) {
	b.damages = p
}
func (b *LaserProjectile) SetTarget(x, y float64) {
	b.hitBox.X = x - 64
	b.hitBox.Y = y - 64
}
func (b *LaserProjectile) Update() {

}
func (b *LaserProjectile) Draw(image *ebiten.Image) {
}
func (b *LaserProjectile) GetProjectileEffect() *ProjectileEffect {
	return b.damages
}

func (b *LaserProjectile) Impact() {
}
