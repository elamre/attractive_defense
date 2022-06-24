package turrets

import (
	"github.com/elamre/attractive_defense/game"
	"github.com/elamre/attractive_defense/world"
	"github.com/hajimehoshi/ebiten/v2"
)

type TurretBaseInterface interface {
	CanUpgrade() bool
	Upgrade()
	Draw(dst *ebiten.DrawImageOptions, screen *ebiten.Image)
	TakeDamage(damage int) bool
}

type TurretGunInterface interface {
	CanUpgrade() bool
	Upgrade()
	Update(target game.Targetable)
	Draw(dst *ebiten.DrawImageOptions, screen *ebiten.Image)
}

type Turret struct {
	gun     TurretGunInterface
	base    TurretBaseInterface
	opts    *ebiten.DrawImageOptions
	destroy bool
	x, y    int
}

func NewTurret(locX, locY int, gun TurretGunInterface, base TurretBaseInterface) *Turret {
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(locX*64), float64(locY*64))
	return &Turret{
		gun:  gun,
		base: base,
		opts: &opts,
		x:    locX,
		y:    locY,
	}
}

func (t *Turret) Damage(damage int) {
	if t.base.TakeDamage(damage) {
		// We are destroyed most likely
	}
}

func (t *Turret) Update(g *world.Grid) {
	if t.destroy {
		g.SetGrid(t.x, t.y, world.GridLevelStructures, nil)
	}
}

func (t *Turret) SetForDeletion(g *world.Grid) {
	t.destroy = true
}

func (t *Turret) Draw(screen *ebiten.Image) {
	t.base.Draw(t.opts, screen)
	t.gun.Draw(t.opts, screen)
}
