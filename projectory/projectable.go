package projectory

import (
	"github.com/elamre/tentsuyu"
	"github.com/hajimehoshi/ebiten/v2"
)

type Projectable interface {
	GetHitBox() *tentsuyu.Rectangle
	Destroy() bool
	Update()
	Draw(image *ebiten.Image)
	GetBaseProjectile() *BaseProjectile
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
