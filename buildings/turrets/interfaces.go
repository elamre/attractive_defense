package turrets

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/elamre/attractive_defense/enemies"
	"github.com/elamre/attractive_defense/world"
	"github.com/elamre/tentsuyu/tentsuyutils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/colornames"
	"math"
)

type TurretBaseInterface interface {
	UpgradeCost() int
	Upgrade()
	GetUpgradeButton() *ebiten.Image
	Draw(dst *ebiten.DrawImageOptions, screen *ebiten.Image)
	TakeDamage(damage int)
	GetHealth() int
}

type TurretGunInterface interface {
	UpgradeCost() int
	Upgrade()
	GetUpgradeButton() *ebiten.Image
	Update(target world.Targetable)
	Draw(dst *ebiten.DrawImageOptions, screen *ebiten.Image)
}

type Turret struct {
	Gun            TurretGunInterface
	Base           TurretBaseInterface
	EnemyTarget    enemies.EnemyInterface
	opts           *ebiten.DrawImageOptions
	gunOpts        *ebiten.DrawImageOptions
	destroy        bool
	x, y           int
	pixelX, pixelY float64
	shootTimer     int
}

func (b *Turret) InflictDamage(damage int) {
	b.Base.TakeDamage(damage)
}

func (b *Turret) Alive() bool {
	return b.Base.GetHealth() > 0
}

func (t *Turret) GetPixelCoordinates() (x, y int) {
	return t.x * 64, t.y * 64
}

func (t *Turret) Trigger(x, y int, other interface{}) {
	t.EnemyTarget = other.(enemies.EnemyInterface)
}

func NewTurret(locX, locY int, gun TurretGunInterface, base TurretBaseInterface, g *world.Grid) *Turret {
	opts := ebiten.DrawImageOptions{}
	gunOpts := ebiten.DrawImageOptions{}

	t := &Turret{
		Gun:     gun,
		Base:    base,
		opts:    &opts,
		gunOpts: &gunOpts,
		x:       locX,
		y:       locY,
		pixelX:  float64(locX * 64),
		pixelY:  float64(locY * 64),
	}

	opts.GeoM.Translate(float64(locX*64), float64(locY*64))
	gunOpts.GeoM.Translate(float64(locX*64), float64(locY*64))

	for _, s := range assets.Surroundings3x3 {
		cX, cY := locX+s.X, locY+s.Y

		if g.OutOfBounds(cX, cY) {
			continue
		}
		g.AddTriggerFunc(cX, cY, t)
	}
	return t
}

func (t *Turret) Damage(damage int) {

}

func (t *Turret) Update(g *world.Grid) {
	if !t.Alive() {
		t.destroy = true
	}
	if t.destroy {
		t.EnemyTarget = nil
		for _, s := range assets.Surroundings3x3 {
			cX, cY := t.x+s.X, t.y+s.Y

			if g.OutOfBounds(cX, cY) {
				continue
			}
			g.RemoveTrigger(cX, cY, t)

		}
		g.SetGrid(t.x, t.y, world.GridLevelStructures, nil)
	}
	if t.EnemyTarget != nil {
		t.gunOpts.GeoM.Reset()
		enemyX, enemyY := t.EnemyTarget.GetPixelPosition()
		eX, eY := float64(enemyX)+32, float64(enemyY)+32
		mouseXFloat := eX - (t.pixelX + 32)
		mouseYFloat := eY - (t.pixelY + 32)

		angle := math.Atan2(mouseYFloat, mouseXFloat)

		t.gunOpts.GeoM.Translate(-64/2, -64/2)
		t.gunOpts.GeoM.Rotate(angle)
		t.gunOpts.GeoM.Translate(64/2, 64/2)
		t.gunOpts.GeoM.Translate(t.pixelX, t.pixelY)
		if tentsuyutils.Distance(t.pixelX+32, t.pixelY+32, eX, eY) < 200 {
			if t.shootTimer == 10 {
				t.shootTimer = 0
				g.ProjectoryMng.AddPlayerProjectile(world.NewBasicProjectile(t.pixelX+32, t.pixelY+32, eX, eY))
			}
		}
	}
	if t.shootTimer < 10 {
		t.shootTimer++
	}
}

func (t *Turret) SetForDeletion(g *world.Grid) {
	t.destroy = true
}

func (t *Turret) Draw(screen *ebiten.Image) {
	t.Base.Draw(t.opts, screen)
	t.Gun.Draw(t.gunOpts, screen)
	ebitenutil.DrawRect(screen, float64(t.x*64+2), float64(t.y*64+58), 60, 4, colornames.Red)
	ebitenutil.DrawRect(screen, float64(t.x*64+2), float64(t.y*64+58), (float64(t.Base.GetHealth())/50)*60, 4, colornames.Green)
}
