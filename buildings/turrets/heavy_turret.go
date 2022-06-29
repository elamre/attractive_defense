package turrets

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/elamre/attractive_defense/enemies"
	"github.com/elamre/attractive_defense/world"
	"github.com/hajimehoshi/ebiten/v2"
)

type heavyTurretGun struct {
	image         *ebiten.Image
	upgradeButton *ebiten.Image
	level         int
	bulletEffects world.ProjectileEffect
	findRange     int
}

func (l *heavyTurretGun) Range() int {
	return l.findRange
}
func (l *heavyTurretGun) Fire(x, y, tX, tY float64, enemyInterface enemies.EnemyInterface, manager *world.ProjectoryManager) {
	p := world.NewHeavyProjectile(x, y, tX, tY, &l.bulletEffects, 250)
	manager.AddPlayerProjectile(p)
}

func (l *heavyTurretGun) ReloadTime() int {
	if l.level >= 3 {
		return 15
	}
	return 30
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
	l.findRange += 10
	switch l.level {
	case 1:
		l.bulletEffects.Damage *= 1.5
		l.image = assets.Get[*ebiten.Image](assets.AssetsTurretGun_heavy_2)
	case 2:
		l.image = assets.Get[*ebiten.Image](assets.AssetsTurretGun_heavy_3)
	case 3:
		l.bulletEffects.Speed *= 2
		l.image = assets.Get[*ebiten.Image](assets.AssetsTurretGun_heavy_4)
	default:
		panic("Should never get here")
	}
	l.level++
}
func (d *heavyTurretGun) Description() string {
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
		bulletEffects: world.ProjectileEffect{Damage: 40, Speed: 5},
		findRange:     3 * 64,
	}
}

func NewHeavyTurret(x, y int, g *world.Grid) *Turret {
	b := newDefaultBase()
	return NewTurret(x, y, newHeavyTurretGun(), b, g)
}
