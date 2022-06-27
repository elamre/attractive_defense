package world

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type ProjectoryManager struct {
	playerProjectiles *assets.EntityManager[*ProjectoryInterface]
	enemyProjectiles  *assets.EntityManager[*ProjectoryInterface]
}

func NewProjectoryManager() *ProjectoryManager {
	p := ProjectoryManager{
		playerProjectiles: assets.NewEntityManager[*ProjectoryInterface](),
		enemyProjectiles:  assets.NewEntityManager[*ProjectoryInterface](),
	}
	return &p
}

func (p *ProjectoryManager) AddEnemyProjectile(projectile ProjectoryInterface) {
	p.enemyProjectiles.AddEntity(&projectile)
}

func (p *ProjectoryManager) AddPlayerProjectile(projectile ProjectoryInterface) {
	p.playerProjectiles.AddEntity(&projectile)
}

func (p *ProjectoryManager) Update(g *Grid) {
	for i := range p.enemyProjectiles.Entities {
		e := *p.enemyProjectiles.Entities[i]
		e.Update()
		x, y := int(e.GetHitBox().X), int(e.GetHitBox().Y)
		gridX, gridY := x/64, y/64
		if !e.IsAlive() {
			p.enemyProjectiles.SetForRemoval(p.enemyProjectiles.Entities[i])
		} else if ee := g.GetGridEntity(x/64, y/64, GridLevelStructures); ee != nil {
			trg := ee.(BuildingInterface)
			if x > gridX*64+16 && x < gridX*64+64-16 {
				if y > gridY*64+16 && y < gridY*64+64-16 {
					trg.InflictDamage(e.GetProjectileEffect().Damage)
					p.enemyProjectiles.SetForRemoval(p.enemyProjectiles.Entities[i])
				}
			}
		}
	}
	for i := range p.playerProjectiles.Entities {
		e := *p.playerProjectiles.Entities[i]
		e.Update()
		x, y := int(e.GetHitBox().X), int(e.GetHitBox().Y)
		if !e.IsAlive() {
			p.playerProjectiles.SetForRemoval(p.playerProjectiles.Entities[i])
		}
		_, _ = x, y
		//TODO player guns
	}
	p.playerProjectiles.CleanDeadEntities()
	p.enemyProjectiles.CleanDeadEntities()
}

func (p *ProjectoryManager) Draw(screen *ebiten.Image) {
	for i := range p.enemyProjectiles.Entities {
		e := *p.enemyProjectiles.Entities[i]
		e.Draw(screen)
	}
	for i := range p.playerProjectiles.Entities {
		e := *p.playerProjectiles.Entities[i]
		e.Draw(screen)
	}
}
