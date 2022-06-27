package world

import (
	"github.com/elamre/tentsuyu"
	"github.com/hajimehoshi/ebiten/v2"
)

type ProjectoryInterface interface {
	GetHitBox() *tentsuyu.Rectangle
	IsAlive() bool
	Update()
	Draw(image *ebiten.Image)
	GetProjectileEffect() *ProjectileEffect
	Impact()
}

type ProjectileEffect struct {
	Homing             bool
	Damage             int
	DamageRadius       int
	SlowDownPercentage int
	SlowDownTime       int
	FreezeTime         int
}

type BaseProjectile struct {
	Effect *ProjectileEffect
	Speed  int
}
