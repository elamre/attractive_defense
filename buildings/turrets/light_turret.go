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
	bulletEffects world.ProjectileEffect
}

func (l *lightTurretGun) Fire(x, y, tX, tY float64, manager *world.ProjectoryManager) {
	p := world.NewSmallProjectile(x, y, tX, tY, &l.bulletEffects, 150)
	manager.AddPlayerProjectile(p)
}

func (l *lightTurretGun) ReloadTime() int {
	if l.level >= 3 {
		return 10
	}
	return 20
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
		l.bulletEffects.Damage *= 1.5
		l.image = assets.Get[*ebiten.Image](assets.AssetsTurretGun_light_2)
	case 2:
		l.image = assets.Get[*ebiten.Image](assets.AssetsTurretGun_light_3)
	case 3:
		l.bulletEffects.Speed *= 2
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
func (d *lightTurretGun) Description() string {
	switch d.level {
	case 1:
		return "More damage"
	case 2:
		return "Faster reload"
	case 3:
		return "Faster bullets"
	default:
		return ""
	}
}
func newLightTurretGun() *lightTurretGun {
	return &lightTurretGun{
		image:         assets.Get[*ebiten.Image](assets.AssetsTurretGun_light_1),
		upgradeButton: assets.Get[*ebiten.Image](assets.AssetsGuiLightTurretUpgrade),
		level:         1,
		bulletEffects: world.ProjectileEffect{Damage: 10, Speed: 3},
	}
}

func NewLightTurret(x, y int, g *world.Grid) *Turret {
	return NewTurret(x, y, newLightTurretGun(), newDefaultBase(), g)
}
