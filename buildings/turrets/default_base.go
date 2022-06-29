package turrets

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type defaultBase struct {
	image         *ebiten.Image
	upgradeButton *ebiten.Image
	maxHealth     float64
	maxShield     int
	level         int
	armor         float64
}

func newDefaultBase() *defaultBase {
	return &defaultBase{
		maxHealth:     50,
		image:         assets.Get[*ebiten.Image](assets.AssetsTurretBase_1),
		upgradeButton: assets.Get[*ebiten.Image](assets.AssetsGuiBaseUpgrade),
		level:         1,
		armor:         1,
	}
}

func (d *defaultBase) Description() string {
	switch d.level {
	case 1:
		return "Extra health"
	case 2:
		return "Adding armor"
	case 3:
		return "Very slow auto heal"
	default:
		return ""
	}
}

func (d *defaultBase) GetMaxHealth() float64 {
	return d.maxHealth
}

func (d *defaultBase) GetUpgradeButton() *ebiten.Image {
	return d.upgradeButton
}

func (d *defaultBase) UpgradeCost() int {
	if d.level == 4 {
		return -1
	}
	return 200
}

func (d *defaultBase) Upgrade() {
	switch d.level {
	case 1:
		d.maxHealth *= 2
		d.image = assets.Get[*ebiten.Image](assets.AssetsTurretBase_2)
	case 2:
		d.maxShield = 30
		d.image = assets.Get[*ebiten.Image](assets.AssetsTurretBase_3)
	case 3:
		d.armor = 0.8
		d.image = assets.Get[*ebiten.Image](assets.AssetsTurretBase_4)
	default:
		panic("Should never get here")
	}
	d.level++
}

func (d *defaultBase) Draw(dst *ebiten.DrawImageOptions, screen *ebiten.Image) {
	screen.DrawImage(d.image, dst)

}

func (d *defaultBase) TakeDamage(damage float64) float64 {
	return damage / d.armor
}
func (d *defaultBase) GetMaxShield() int {
	return d.maxShield
}
