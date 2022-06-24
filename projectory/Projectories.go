package projectory

import (
	"github.com/elamre/attractive_defense/game"
	"github.com/elamre/tentsuyu"
	"github.com/hajimehoshi/ebiten/v2"
)

type ProjectoryFactory interface {
	CreateNewProjectory(target game.Targetable, level int) Projectable
}

type basicProjectoryFactory struct {
	image *ebiten.Image
}

func (b *basicProjectoryFactory) CreateNewProjectory(target game.Targetable, level int) Projectable {
	return nil
}

type basicProjectory struct {
	*BaseProjectile
	hitbox *tentsuyu.Rectangle
}

func (b *basicProjectory) GetHitBox() *tentsuyu.Rectangle {
	return b.hitbox
}
func (b *basicProjectory) Destroy() bool {
	return true
}
func (b *basicProjectory) Update() {

}
func (b *basicProjectory) Draw(image *ebiten.Image) {

}
func (b *basicProjectory) GetBaseProjectile() *BaseProjectile {
	return b.BaseProjectile
}
