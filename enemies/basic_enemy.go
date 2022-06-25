package enemies

import (
	"github.com/elamre/attractive_defense/game"
	"github.com/elamre/attractive_defense/projectory"
	"github.com/elamre/attractive_defense/world"
	"github.com/elamre/tentsuyu/tentsuyutils"
	"github.com/hajimehoshi/ebiten/v2"
)

type BasicEnemy struct {
	shootRange     float64
	x, y           int
	pixelX, pixelY int
	target         game.Targetable
}

func (b *BasicEnemy) SetTarget(target game.Targetable) {
	b.target = target
}

func (b *BasicEnemy) Update(g *world.Grid, p *game.Player, projectoryManager *projectory.ProjectoryManager) {
	if b.target != nil {
		tX, tY
		tentsuyutils.Distance()
	}
}

func (b *BasicEnemy) IsAlive() bool {
	return true
}

func (b *BasicEnemy) Draw(screen *ebiten.Image) {

}

func (b *BasicEnemy) GetPixelPosition() (int, int) {
	return b.pixelX, b.pixelY
}
