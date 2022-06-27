package enemies

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/elamre/attractive_defense/game"
	"github.com/elamre/attractive_defense/world"
	"github.com/elamre/tentsuyu"
	"github.com/elamre/tentsuyu/tentsuyutils"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

type BasicEnemy struct {
	*EnemyHullSpecifications
	*EnemyTurretSpecifications
	curSpeed       float64
	pixelX, pixelY float64
	target         world.Targetable
	moveVec        tentsuyu.Vector2d
	dst            ebiten.DrawImageOptions
	hitbox         tentsuyu.Rectangle
	health         int
	distLeft       float64
	prevX, prevY   int
	shootCounter   int
}

func InitEnemyImages() {
	scoutEnemyHull.image = assets.Get[*ebiten.Image](assets.AssetsEnemy)
}

var scoutEnemyTurretEasy EnemyTurretSpecifications = EnemyTurretSpecifications{
	reloadSpeed: 10,
	targetRange: 60,
	shoot: func(pixelX, pixelY, targetX, targetY float64, manager *world.ProjectoryManager) {
		manager.AddEnemyProjectile(world.NewBasicProjectile(pixelX, pixelY, targetX, targetY))
	},
}

var scoutEnemyHull EnemyHullSpecifications = EnemyHullSpecifications{
	image:     nil,
	maxSpeed:  3,
	width:     0,
	height:    0,
	maxHealth: 0,
}

func NewScoutEnemy(pixelX, pixelY float64) EnemyInterface {
	return NewBasicEnemy(pixelX, pixelY, &scoutEnemyTurretEasy, &scoutEnemyHull)
}

func NewBasicEnemy(pixelX, pixelY float64, turret *EnemyTurretSpecifications, hull *EnemyHullSpecifications) EnemyInterface {
	b := BasicEnemy{
		EnemyHullSpecifications:   hull,
		EnemyTurretSpecifications: turret,
		pixelX:                    pixelX,
		pixelY:                    pixelY,
		prevX:                     int(pixelX / 64),
		prevY:                     int(pixelY / 64),
		curSpeed:                  3,
		health:                    20,
		hitbox:                    tentsuyu.Rectangle{W: 32, H: 32},
	}
	return &b
}
func (b *BasicEnemy) CheckCollision(projectoryInterface world.ProjectoryInterface) bool {
	if projectoryInterface.IsAlive() {
		c := projectoryInterface.GetHitBox()
		if b.hitbox.Contains(c.X, c.Y) {
			projectoryInterface.Impact()
			b.health -= projectoryInterface.GetProjectileEffect().Damage
			return true
		}
	}
	return false
}
func (b *BasicEnemy) GetTarget() world.Targetable {
	return b.target
}
func (b *BasicEnemy) SetTarget(target world.Targetable) {
	b.target = target
	if target == nil {
		return
	}
	tX, tY := target.GetPixelCoordinates()
	difX, difY := float64(tX)-b.pixelX, float64(tY)-b.pixelY

	angle := math.Atan2(difY, difX)

	//b.hitBox = &tentsuyu.Rectangle{X: b.pixelX, Y: b.pixelY, W: 9, H: 9}
	b.dst.GeoM.Reset()
	b.dst.GeoM.Translate(-64/2, -64/2)
	b.dst.GeoM.Rotate(angle)
	b.dst.GeoM.Translate(64/2, 64/2)
	b.dst.GeoM.Translate(b.pixelX, b.pixelY)
	b.moveVec.Y = math.Sin(angle)
	b.moveVec.X = math.Cos(angle)
	b.distLeft = tentsuyutils.Distance(b.pixelX, b.pixelY, float64(tX), float64(tY))
	b.curSpeed = b.maxSpeed
}

func (b *BasicEnemy) Update(g *world.Grid, p *game.Player, projectoryManager *world.ProjectoryManager) {
	if b.target == nil {
		return
	}
	travelX := b.moveVec.X * b.curSpeed
	travelY := b.moveVec.Y * b.curSpeed
	b.pixelX += travelX
	b.pixelY += travelY
	b.hitbox.X = b.pixelX + 16
	b.hitbox.Y = b.pixelY + 16
	b.dst.GeoM.Translate(travelX, travelY)
	b.distLeft -= b.moveVec.Length() * b.maxSpeed
	if b.distLeft < b.targetRange {
		b.curSpeed = 0
		b.shootCounter++
		if b.shootCounter >= b.reloadSpeed {
			b.shootCounter = 0
			x, y := b.target.GetPixelCoordinates()
			x += 32
			y += 32

			projectoryManager.AddEnemyProjectile(world.NewBasicProjectile(b.pixelX+32, b.pixelY+32, float64(x), float64(y)))
		}
	} else {
		if newX, newY := int((b.pixelX+32)/64), int((b.pixelY+32)/64); newX != b.prevX || newY != b.prevY {
			b.prevX = newX
			b.prevY = newY
			if g.TestTrigger(b.prevX, b.prevY, b) {
			}
		}
	}
}

func (b *BasicEnemy) IsAlive() bool {
	return b.health > 0
}

func (b *BasicEnemy) Draw(screen *ebiten.Image) {
	screen.DrawImage(b.image, &b.dst)
}

func (b *BasicEnemy) GetPixelPosition() (int, int) {
	return int(b.pixelX), int(b.pixelY)
}
