package turrets

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/elamre/attractive_defense/enemies"
	"github.com/elamre/attractive_defense/world"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/colornames"
)

type beamTurretGun struct {
	image         *ebiten.Image
	upgradeButton *ebiten.Image
	level         int
	bulletEffects world.ProjectileEffect
	findRange     int
	tX, tY        float64
	pX, pY        float64
	beam          *world.LaserProjectile
	enemy         enemies.EnemyInterface
}

func (l *beamTurretGun) Range() int {
	return l.findRange
}

func (l *beamTurretGun) Fire(x, y, tX, tY float64, enemyInterface enemies.EnemyInterface, manager *world.ProjectoryManager) {
	l.pX = x
	l.pY = y
	l.beam.SetTarget(tX+32, tY+32)
	l.enemy = enemyInterface
	if enemyInterface.CheckCollision(l.beam) {
	}
}

func (l *beamTurretGun) ReloadTime() int {
	return 1
}

func (l *beamTurretGun) GetUpgradeButton() *ebiten.Image {
	return l.upgradeButton
}

func (l *beamTurretGun) UpgradeCost() int {
	if l.level == 4 {
		return -1
	}
	return 800 * l.level
}

func (l *beamTurretGun) Upgrade() {
	b := l.beam.GetProjectileEffect()
	switch l.level {
	case 1:
		b.Damage *= 2
		l.image = assets.Get[*ebiten.Image](assets.AssetsTurretGun_beam_2)
	case 2:
		b.SlowDownTime = 60
		b.SlowDownPercentage = 0.5
		l.image = assets.Get[*ebiten.Image](assets.AssetsTurretGun_beam_3)
	case 3:
		b.SlowDownPercentage = 0.8
		l.image = assets.Get[*ebiten.Image](assets.AssetsTurretGun_beam_4)
	default:
		panic("Should never get here")
	}
	l.beam.UpdateProjectilEffect(b)
	l.level++
}
func (d *beamTurretGun) Description() string {
	switch d.level {
	case 1:
		return "More damage"
	case 2:
		return "Slower"
	case 3:
		return "Even slower"
	default:
		return ""
	}
}
func (l *beamTurretGun) Update(target world.Targetable) {
	if l.enemy != nil {

	}

}
func (l *beamTurretGun) Draw(dst *ebiten.DrawImageOptions, screen *ebiten.Image) {
	screen.DrawImage(l.image, dst)
	if l.enemy != nil {
		x, y := l.enemy.GetPixelPosition()
		ebitenutil.DrawLine(screen, l.pX, l.pY, float64(x)+32, float64(y)+32, colornames.Blue)
		if !l.enemy.IsAlive() {
			l.enemy = nil
		}
	}
}

func newbeamTurretGun() *beamTurretGun {
	b := &beamTurretGun{
		image:         assets.Get[*ebiten.Image](assets.AssetsTurretGun_beam_1),
		upgradeButton: assets.Get[*ebiten.Image](assets.AssetsGuiBeamTurretUpgrade),
		level:         1,
		bulletEffects: world.ProjectileEffect{Damage: 1, Speed: 10, SlowDownPercentage: 0.2},
		findRange:     6 * 64,
		beam:          world.NewLaserProjectile(0, 0, 0, 0),
	}
	b.beam.UpdateProjectilEffect(&b.bulletEffects)
	return b
}

func NewBeamTurretGun(x, y int, g *world.Grid) *Turret {
	b := newDefaultBase()
	return NewTurret(x, y, newbeamTurretGun(), b, g)
}
