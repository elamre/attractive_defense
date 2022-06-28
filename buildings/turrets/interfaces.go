package turrets

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/elamre/attractive_defense/buildings"
	"github.com/elamre/attractive_defense/enemies"
	"github.com/elamre/attractive_defense/world"
	"github.com/elamre/tentsuyu/tentsuyutils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"math"
)

type TurretBaseInterface interface {
	UpgradeCost() int
	Upgrade()
	Description() string
	GetUpgradeButton() *ebiten.Image
	Draw(dst *ebiten.DrawImageOptions, screen *ebiten.Image)
	TakeDamage(damage float64) float64
	GetMaxHealth() float64
}

type TurretGunInterface interface {
	UpgradeCost() int
	Upgrade()
	Description() string
	GetUpgradeButton() *ebiten.Image
	Update(target world.Targetable)
	Draw(dst *ebiten.DrawImageOptions, screen *ebiten.Image)
	ReloadTime() int
	Fire(x, y, tX, tY float64, manager *world.ProjectoryManager)
}

type Turret struct {
	*world.Building
	Gun            TurretGunInterface
	Base           TurretBaseInterface
	EnemyTarget    enemies.EnemyInterface
	opts           *ebiten.DrawImageOptions
	gunOpts        *ebiten.DrawImageOptions
	upgradeEffect  []*ebiten.Image
	upgradeCounter float64
	destroy        bool
	shootTimer     int
	EnemyKilledCb  func()
}

func (d *Turret) Upgraded(able buildings.UpgradeAble) {
	if able == d.Base {
		d.Building.MaxHealth = d.Base.GetMaxHealth()
		d.Building.Health += d.Building.MaxHealth / 2
		if d.Building.Health > d.Building.MaxHealth {
			d.Building.Health = d.Building.MaxHealth
		}
	}
	d.upgradeCounter = 0
}

func (d *Turret) GetBuilding() *world.Building {
	return d.Building
}

func (b *Turret) InflictDamage(damage float64) {
	dmg := b.Base.TakeDamage(damage)
	b.Health -= dmg
}

func (b *Turret) Alive() bool {
	return b.Health >= 1
}

func (t *Turret) SetClosest(enemyInterface []*enemies.EnemyInterface) {
	closest := 9999999.0
	for i := range enemyInterface {
		e := *enemyInterface[i]
		x, y := e.GetPixelPosition()
		distance := tentsuyutils.Distance(t.PixelX+32, t.PixelY+32, float64(x), float64(y))
		if distance < closest {
			t.EnemyTarget = e
			closest = distance
		}
	}
}

func (t *Turret) CheckAndSetTarget(enemyInterface enemies.EnemyInterface) {
	x, y := enemyInterface.GetPixelPosition()
	distance := tentsuyutils.Distance(t.PixelX+32, t.PixelY+32, float64(x), float64(y))
	if t.EnemyTarget != nil {
		if !t.EnemyTarget.IsAlive() {
			t.EnemyTarget = enemyInterface
		}
		curX, curY := t.EnemyTarget.GetPixelPosition()
		curDistance := tentsuyutils.Distance(t.PixelX+32, t.PixelY+32, float64(curX), float64(curY))
		if distance < curDistance {
			t.EnemyTarget = enemyInterface
		}
		return
	} else {
		t.EnemyTarget = enemyInterface
	}

}

func (t *Turret) Trigger(x, y int, other interface{}) {
	if t.EnemyKilledCb != nil {
		t.EnemyKilledCb()
	}
}

func NewTurret(locX, locY int, gun TurretGunInterface, base TurretBaseInterface, g *world.Grid) *Turret {
	opts := ebiten.DrawImageOptions{}
	gunOpts := ebiten.DrawImageOptions{}

	t := &Turret{
		Gun:     gun,
		Base:    base,
		opts:    &opts,
		gunOpts: &gunOpts,
		Building: &world.Building{
			PixelX:    float64(locX * 64),
			PixelY:    float64(locY * 64),
			GridX:     locX,
			GridY:     locY,
			Repairing: false,
			Health:    base.GetMaxHealth(),
			MaxHealth: base.GetMaxHealth(),
		},
		upgradeEffect: assets.Get[[]*ebiten.Image](assets.AssetsUpgradeAnim),
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

func (t *Turret) Update(g *world.Grid) {
	if !t.Alive() {
		t.destroy = true
	}
	if t.destroy {
		t.EnemyTarget = nil
		for _, s := range assets.Surroundings3x3 {
			cX, cY := t.GridX+s.X, t.GridY+s.Y

			if g.OutOfBounds(cX, cY) {
				continue
			}
			g.RemoveTrigger(cX, cY, t)

		}
		g.SetGrid(t.GridX, t.GridY, world.GridLevelStructures, nil)
	}

	if t.EnemyTarget != nil {
		if !t.EnemyTarget.IsAlive() {
			t.EnemyTarget = nil
			if t.EnemyKilledCb != nil {
				t.EnemyKilledCb()
			}
		} else {
			t.gunOpts.GeoM.Reset()
			enemyX, enemyY := t.EnemyTarget.GetPixelPosition()
			eX, eY := float64(enemyX)+32, float64(enemyY)+32
			mouseXFloat := eX - (t.PixelX + 32)
			mouseYFloat := eY - (t.PixelY + 32)

			angle := math.Atan2(mouseYFloat, mouseXFloat)

			t.gunOpts.GeoM.Translate(-64/2, -64/2)
			t.gunOpts.GeoM.Rotate(angle)
			t.gunOpts.GeoM.Translate(64/2, 64/2)
			t.gunOpts.GeoM.Translate(t.PixelX, t.PixelY)
			if tentsuyutils.Distance(t.PixelX+32, t.PixelY+32, eX, eY) < 200 {
				if t.shootTimer >= t.Gun.ReloadTime() {
					t.shootTimer = 0
					t.Gun.Fire(t.PixelX+32, t.PixelY+32, eX, eY, g.ProjectoryMng)
				}
			}
		}
		t.shootTimer++
	}
	if int(t.upgradeCounter) < len(t.upgradeEffect) {
		t.upgradeCounter += 0.5
	} else {
		t.upgradeCounter = 999
	}
}

func (t *Turret) SetForDeletion(g *world.Grid) {
	t.destroy = true
}

func (t *Turret) Draw(screen *ebiten.Image) {
	t.Base.Draw(t.opts, screen)
	t.Gun.Draw(t.gunOpts, screen)
	if int(t.upgradeCounter) < len(t.upgradeEffect) {
		t.opts.GeoM.Translate(0, -16)
		screen.DrawImage(t.upgradeEffect[int(t.upgradeCounter)], t.opts)
		t.opts.GeoM.Translate(0, 16)
	}
	if t.EnemyTarget == nil {
		ebitenutil.DebugPrintAt(screen, "NNN", int(t.PixelX+2), int(t.PixelY))
	} else if !t.EnemyTarget.IsAlive() {
		ebitenutil.DebugPrintAt(screen, "DDD", int(t.PixelX+2), int(t.PixelY))
	} else {
		ebitenutil.DebugPrintAt(screen, "AAA", int(t.PixelX+2), int(t.PixelY))
	}
	t.DrawGui(screen)

}
