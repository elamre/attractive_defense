package turrets

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/elamre/attractive_defense/world"
	"github.com/hajimehoshi/ebiten/v2"
)

type lightTurretGun struct {
	image         *ebiten.Image
	upgradeButton *ebiten.Image
	level         int
}

func (d *lightTurretGun) GetUpgradeButton() *ebiten.Image {
	return d.upgradeButton
}

func (l *lightTurretGun) UpgradeCost() int {
	if l.level == 4 {
		return -1
	}
	return 200
}
func (l *lightTurretGun) Upgrade() {
	switch l.level {
	case 1:
		l.image = assets.Get[*ebiten.Image](assets.AssetsTurretGun_light_2)
	case 2:
		l.image = assets.Get[*ebiten.Image](assets.AssetsTurretGun_light_3)
	case 3:
		l.image = assets.Get[*ebiten.Image](assets.AssetsTurretGun_light_4)
	default:
		panic("Should never get here")
	}
	l.level++
}
func (l *lightTurretGun) Update(target world.Targetable) {

}
func (l *lightTurretGun) Draw(dst *ebiten.DrawImageOptions, screen *ebiten.Image) {
	screen.DrawImage(l.image, dst)

}

func newLightTurretGun() *lightTurretGun {
	return &lightTurretGun{
		image:         assets.Get[*ebiten.Image](assets.AssetsTurretGun_light_1),
		upgradeButton: assets.Get[*ebiten.Image](assets.AssetsGuiLightTurretUpgrade),
		level:         1,
	}
}

func NewLightTurret(x, y int, g *world.Grid) *Turret {
	return NewTurret(x, y, newLightTurretGun(), newDefaultBase(), g)
}
