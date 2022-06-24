package game

import (
	"fmt"
	"github.com/elamre/attractive_defense/world"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Player struct {
	Money int
	risk  int
	//gui   *gui.BottomGui
}

func NewPlayer() *Player {
	return &Player{Money: 10000, risk: 1}
}

func (p *Player) UpdatePlayer(g *world.Grid) {

}

func (p *Player) DrawPlayer(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("$$$ %d", p.Money), 0, 0)
}
