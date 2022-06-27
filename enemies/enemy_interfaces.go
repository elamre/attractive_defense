package enemies

import (
	"github.com/elamre/attractive_defense/game"
	"github.com/elamre/attractive_defense/world"
	"github.com/hajimehoshi/ebiten/v2"
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

type EnemyHullSpecifications struct {
	image     *ebiten.Image
	maxSpeed  float64
	width     int
	height    int
	maxHealth int
}

type EnemyTurretSpecifications struct {
	reloadSpeed int
	targetRange float64
	shoot       func(pixelX, pixelY, targetX, targetY float64, manager *world.ProjectoryManager)
}
