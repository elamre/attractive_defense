package enemies

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/elamre/attractive_defense/game"
	"github.com/elamre/attractive_defense/world"
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand"
)

type EnemyInterface interface {
	SetTarget(target world.Targetable)
	GetTarget() world.Targetable
	Update(g *world.Grid, p *game.Player, projectoryManager *world.ProjectoryManager)
	CheckCollision(projectoryInterface world.ProjectoryInterface) bool
	IsAlive() bool
	Draw(screen *ebiten.Image)
	GetPixelPosition() (int, int)
}

type EnemyManager struct {
	*assets.EntityManager[*EnemyInterface]
	targetToEnemy map[world.BuildingInterface][]EnemyInterface
	waveNumber    int
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

func (e *EnemyManager) ShouldSpawn() bool {
	return len(e.Entities) == 0
}

func (e *EnemyManager) Spawn(g *world.Grid, difficulty int) int {
	for i := 0; i < 4; i++ {
		rX := rand.Float64() * float64(g.Width*64)
		rY := rand.Float64() * float64(g.Height*64)
		if i == 0 {
			rY = 0
		} else if i == 1 {
			rX = 0
		} else if i == 2 {
			rY = float64(g.Height) * 64
		} else {
			rX = float64(g.Width) * 64
		}
		t := NewBasicEnemy(rX, rY, nil)
		e.assignTarget(t, g)
		e.AddEnemy(t, g)
	}

	e.waveNumber++
	return e.waveNumber // The wave
}

func (e *EnemyManager) Update(g *world.Grid, p *game.Player, projectoryManager *world.ProjectoryManager) {
	for i := range e.Entities {
		enemy := *e.Entities[i]
		if enemy.IsAlive() {
			enemy.Update(g, p, projectoryManager)
			for pp := range projectoryManager.PlayerProjectiles.Entities {
				if enemy.CheckCollision(*projectoryManager.PlayerProjectiles.Entities[pp]) {

				}
			}
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
