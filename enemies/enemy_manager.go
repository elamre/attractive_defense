package enemies

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/elamre/attractive_defense/game"
	"github.com/elamre/attractive_defense/world"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"math/rand"
)

type EnemyManager struct {
	*assets.EntityManager[*EnemyInterface]
	targetToEnemy map[world.BuildingInterface][]EnemyInterface
	waveNumber    int
}

func NewEnemyManager() *EnemyManager {
	InitEnemyImages()
	initHeaviesEnemyImages()
	e := EnemyManager{EntityManager: assets.NewEntityManager[*EnemyInterface](), targetToEnemy: make(map[world.BuildingInterface][]EnemyInterface)}
	return &e
}

func (e *EnemyManager) assignTarget(enemy EnemyInterface, g *world.Grid) {
	x, y := enemy.GetPixelPosition()
	closest := g.ClosestBuilding(x/64, y/64)
	if closest == nil {
		return
	}
	enemy.SetTarget(closest.(world.Targetable))
}

func (e *EnemyManager) AddEnemy(enemy EnemyInterface, g *world.Grid) {
	e.assignTarget(enemy, g)
	e.AddEntity(&enemy)

}

func (e *EnemyManager) ShouldSpawn() bool {
	return len(e.Entities) == 0
}

func (e *EnemyManager) Spawn(g *world.Grid, difficulty int) int {
	assets.StaticSoundManager.PlayNewWave()
	wave := make([]WaveContent, 0)
	maxWave := len(SoutWaves) + len(heaviesWave)
	extra := e.waveNumber - maxWave
	if extra < 0 {
		extra = 0
	}
	if e.waveNumber < len(SoutWaves) {
		for i := range SoutWaves[e.waveNumber].content {
			wave = append(wave, SoutWaves[e.waveNumber].content[i])
		}
	} else if e.waveNumber < (len(SoutWaves) + len(heaviesWave)) {
		for i := range heaviesWave[e.waveNumber-len(SoutWaves)].content {
			wave = append(wave, heaviesWave[e.waveNumber-len(SoutWaves)].content[i])
		}
	} else {
		for i := range SoutWaves[e.waveNumber%len(SoutWaves)].content {
			wave = append(wave, SoutWaves[e.waveNumber%len(SoutWaves)].content[i])
		}
		for i := range heaviesWave[e.waveNumber%len(heaviesWave)].content {
			wave = append(wave, heaviesWave[e.waveNumber%len(heaviesWave)].content[i])
		}
	}
	for i := range wave {
		log.Printf("Wave %d amount %d difficulty %d extra %d", i, wave[i].amount, difficulty, extra)
		amount := wave[i].amount + (difficulty - 1) + extra
		for s := 0; s < amount; s++ {
			if (difficulty - 1) > 0 {
				difficulty--
			}
			side := rand.Int() % 4
			rX := rand.Float64() * float64(g.Width*64)
			rY := rand.Float64() * float64(g.Height*64)
			if side == 0 {
				rY = 0
			} else if side == 1 {
				rX = 0
			} else if side == 2 {
				rY = float64(g.Height) * 64
			} else {
				rX = float64(g.Width) * 64
			}
			t := wave[i].spawn(rX, rY)
			e.assignTarget(t, g)
			e.AddEnemy(t, g)
		}
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
			p.AddMoney(float64(enemy.GetReward()))
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

func (e *EnemyManager) WaveNr() int {
	return e.waveNumber
}
