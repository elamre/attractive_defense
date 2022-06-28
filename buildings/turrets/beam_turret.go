package turrets

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/elamre/attractive_defense/world"
	"github.com/hajimehoshi/ebiten/v2"
)

type beamTurretGun struct {
	image         *ebiten.Image
	upgradeButton *ebiten.Image
	level         int
	bulletEffects world.ProjectileEffect
}

func (l *beamTurretGun) Fire(x, y, tX, tY float64, manager *world.ProjectoryManager) {
	p := world.NewSmallProjectile(x, y, tX, tY, &l.bulletEffects, 250)
	manager.AddPlayerProjectile(p)
}

func (l *beamTurretGun) ReloadTime() int {
	if l.level >= 3 {
		return 1
	}
	return 3
}

func (l *beamTurretGun) GetUpgradeButton() *ebiten.Image {
	return l.upgradeButton
}

func (l *beamTurretGun) UpgradeCost() int {
	if l.level == 4 {
		return -1
	}
	return 200 * l.level
}

func (l *beamTurretGun) Upgrade() {
	switch l.level {
	case 1:
		l.bulletEffects.Damage *= 1.5
		l.image = assets.Get[*ebiten.Image](assets.AssetsTurretGun_beam_2)
	case 2:
		l.image = assets.Get[*ebiten.Image](assets.AssetsTurretGun_beam_3)
	case 3:
		l.bulletEffects.Speed *= 2
		l.image = assets.Get[*ebiten.Image](assets.AssetsTurretGun_beam_4)
	default:
		panic("Should never get here")
	}
	l.level++
}
func (d *beamTurretGun) Description() string {
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
func (l *beamTurretGun) Update(target world.Targetable) {

}
func (l *beamTurretGun) Draw(dst *ebiten.DrawImageOptions, screen *ebiten.Image) {
	screen.DrawImage(l.image, dst)
}

func newbeamTurretGun() *beamTurretGun {
	return &beamTurretGun{
		image:         assets.Get[*ebiten.Image](assets.AssetsTurretGun_beam_1),
		upgradeButton: assets.Get[*ebiten.Image](assets.AssetsGuiBeamTurretUpgrade),
		level:         1,
		bulletEffects: world.ProjectileEffect{Damage: 2, Speed: 8},
	}
}

func NewBeamTurretGun(x, y int, g *world.Grid) *Turret {
	b := newDefaultBase()
	return NewTurret(x, y, newbeamTurretGun(), b, g)
}
