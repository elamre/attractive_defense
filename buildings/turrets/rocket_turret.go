package turrets

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/elamre/attractive_defense/world"
	"github.com/hajimehoshi/ebiten/v2"
)

type rocketTurretGun struct {
	image         *ebiten.Image
	upgradeButton *ebiten.Image
	level         int
	bulletEffects world.ProjectileEffect
}

func (l *rocketTurretGun) Fire(x, y, tX, tY float64, manager *world.ProjectoryManager) {
	p := world.NewSmallProjectile(x, y, tX, tY, &l.bulletEffects, 250)
	manager.AddPlayerProjectile(p)
}

func (l *rocketTurretGun) ReloadTime() int {
	if l.level >= 3 {
		return 15
	}
	return 30
}

func (l *rocketTurretGun) GetUpgradeButton() *ebiten.Image {
	return l.upgradeButton
}

func (l *rocketTurretGun) UpgradeCost() int {
	if l.level == 4 {
		return -1
	}
	return 200 * l.level
}

func (l *rocketTurretGun) Upgrade() {
	switch l.level {
	case 1:
		l.bulletEffects.Damage *= 1.5
		l.image = assets.Get[*ebiten.Image](assets.AssetsTurretGun_rocket_2)
	case 2:
		l.image = assets.Get[*ebiten.Image](assets.AssetsTurretGun_rocket_3)
	case 3:
		l.bulletEffects.Speed *= 2
		l.image = assets.Get[*ebiten.Image](assets.AssetsTurretGun_rocket_4)
	default:
		panic("Should never get here")
	}
	l.level++
}
func (d *rocketTurretGun) Description() string {
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
func (l *rocketTurretGun) Update(target world.Targetable) {

}
func (l *rocketTurretGun) Draw(dst *ebiten.DrawImageOptions, screen *ebiten.Image) {
	screen.DrawImage(l.image, dst)
}

func newrocketTurretGun() *rocketTurretGun {
	return &rocketTurretGun{
		image:         assets.Get[*ebiten.Image](assets.AssetsTurretGun_rocket_1),
		upgradeButton: assets.Get[*ebiten.Image](assets.AssetsGuiRocketTurretUpgrade),
		level:         1,
		bulletEffects: world.ProjectileEffect{Damage: 40, Speed: 3},
	}
}

func NewRocketTurret(x, y int, g *world.Grid) *Turret {
	b := newDefaultBase()
	return NewTurret(x, y, newrocketTurretGun(), b, g)
}
