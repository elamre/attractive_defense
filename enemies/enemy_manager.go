package enemies

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/elamre/attractive_defense/game"
	"github.com/elamre/attractive_defense/world"
	"github.com/hajimehoshi/ebiten/v2"
)

type EnemyInterface interface {
	SetTarget(target world.Targetable)
	GetTarget() world.Targetable
	Update(g *world.Grid, p *game.Player, projectoryManager *world.ProjectoryManager)
	IsAlive() bool
	Draw(screen *ebiten.Image)
	GetPixelPosition() (int, int)
}

type EnemyManager struct {
	*assets.EntityManager[*EnemyInterface]
	targetToEnemy map[world.BuildingInterface][]EnemyInterface
}

func NewEnemyManager() *EnemyManager {
	e := EnemyManager{EntityManager: assets.NewEntityManager[*EnemyInterface](), targetToEnemy: make(map[world.BuildingInterface][]EnemyInterface)}
	return &e
}

func (e *EnemyManager) assignTarget(enemy EnemyInterface, g *world.Grid) {
	x, y := enemy.GetPixelPosition()
	target := g.ClosestBuilding(x/64, y/64).(world.BuildingInterface)

	enemy.SetTarget(target.(world.Targetable))
}

func (e *EnemyManager) AddEnemy(enemy EnemyInterface, g *world.Grid) {
	e.assignTarget(enemy, g)
	e.AddEntity(&enemy)
}

func (e *EnemyManager) Update(g *world.Grid, p *game.Player, projectoryManager *world.ProjectoryManager) {
	for i := range e.Entities {
		enemy := *e.Entities[i]
		if enemy.IsAlive() {
			enemy.Update(g, p, projectoryManager)
			if enemy.GetTarget() != nil {
				if !enemy.GetTarget().(world.BuildingInterface).Alive() {
					e.assignTarget(enemy, g)
				}
			}
		} else {

			e.SetForRemoval(e.Entities[i])
		}
	}

	e.CleanDeadEntities()
}

func (e *EnemyManager) Draw(screen *ebiten.Image) {
	for i := range e.Entities {
		enemy := *e.Entities[i]
		if enemy.IsAlive() {
			enemy.Draw(screen)
		}
	}
}
