package turrets

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/elamre/attractive_defense/world"
	"github.com/hajimehoshi/ebiten/v2"
)

type heavyTurretGun struct {
	image         *ebiten.Image
	upgradeButton *ebiten.Image
	level         int
}

func (l *heavyTurretGun) GetUpgradeButton() *ebiten.Image {
	return l.upgradeButton
}

func (l *heavyTurretGun) UpgradeCost() int {
	if l.level == 4 {
		return -1
	}
	return 200 * l.level
}

func (l *heavyTurretGun) Upgrade() {
	switch l.level {
	case 1:
		l.image = assets.Get[*ebiten.Image](assets.AssetsTurretGun_heavy_2)
	case 2:
		l.image = assets.Get[*ebiten.Image](assets.AssetsTurretGun_heavy_3)
	case 3:
		l.image = assets.Get[*ebiten.Image](assets.AssetsTurretGun_heavy_4)
	default:
		panic("Should never get here")
	}
	l.level++
}
func (l *heavyTurretGun) Update(target world.Targetable) {

}
func (l *heavyTurretGun) Draw(dst *ebiten.DrawImageOptions, screen *ebiten.Image) {

	screen.DrawImage(l.image, dst)

}

func newHeavyTurretGun() *heavyTurretGun {
	return &heavyTurretGun{
		image:         assets.Get[*ebiten.Image](assets.AssetsTurretGun_heavy_1),
		upgradeButton: assets.Get[*ebiten.Image](assets.AssetsGuiHeavyTurretUpgrade),
		level:         1,
	}
}

func NewHeavyTurret(x, y int, g *world.Grid) *Turret {
	return NewTurret(x, y, newHeavyTurretGun(), newDefaultBase(), g)
}
