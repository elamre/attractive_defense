package turrets

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type defaultBase struct {
	image  *ebiten.Image
	health int
	level  int
}

func newDefaultBase() *defaultBase {
	return &defaultBase{
		image: assets.Get[*ebiten.Image](assets.AssetsTurretBase_1),
		level: 0,
	}
}

func (d *defaultBase) CanUpgrade() bool {
	return d.level < 4
}

func (d *defaultBase) Upgrade() {
	switch d.level {
	case 1:
		d.image = assets.Get[*ebiten.Image](assets.AssetsTurretBase_2)
	case 2:
		d.image = assets.Get[*ebiten.Image](assets.AssetsTurretBase_3)
	case 3:
		d.image = assets.Get[*ebiten.Image](assets.AssetsTurretBase_4)
	default:
		panic("Should never get here")
	}
	d.level++
}

func (d *defaultBase) Draw(dst *ebiten.DrawImageOptions, screen *ebiten.Image) {
	screen.DrawImage(d.image, dst)
}
func (d *defaultBase) TakeDamage(damage int) bool {
	return false // we can not get destroyed yet
}
