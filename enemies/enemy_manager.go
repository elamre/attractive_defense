package enemies

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/elamre/attractive_defense/game"
	"github.com/elamre/attractive_defense/projectory"
	"github.com/elamre/attractive_defense/world"
	"github.com/hajimehoshi/ebiten/v2"
)

type EnemyInterface interface {
	SetTarget(target game.Targetable)
	Update(g *world.Grid, p *game.Player, projectoryManager *projectory.ProjectoryManager)
	IsAlive() bool
	Draw(screen *ebiten.Image)
}

type EnemyManager struct {
	assets.EntityManager[EnemyInterface]
}

func (e *EnemyManager) Update(g *world.Grid, p *game.Player, projectoryManager *projectory.ProjectoryManager) {
	for i := range e.Entities {
		if e.Entities[i].IsAlive() {
			e.Entities[i].Update(g, p, projectoryManager)
		} else {
			e.SetForRemoval(e.Entities[i])
		}
	}
	e.CleanDeadEntities()
}

func (e *EnemyManager) Draw(screen *ebiten.Image) {
	for i := range e.Entities {
		if e.Entities[i].IsAlive() {
			e.Entities[i].Draw(screen)
		}
	}
}
