package game

import (
	"fmt"
	"github.com/elamre/attractive_defense/world"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Player struct {
	Money      float64
	Risk       int
	repairing  map[*world.Building]bool
	removeList []*world.Building
	//gui   *gui.BottomGui
}

func NewPlayer() *Player {

	return &Player{Money: 5000, Risk: 1, repairing: make(map[*world.Building]bool), removeList: make([]*world.Building, 0)}
}

func (p *Player) AddRepairing(b *world.Building) {
	p.repairing[b] = true
}

func (p *Player) UpdatePlayer(g *world.Grid) {
	for k, _ := range p.repairing {
		if k.Repairing {
			if p.Money > k.RepairPerTick() {
				if k.RepairTick() {
					p.Money -= k.RepairPerTick()
					if !k.Repairing {
						// We are repaired
						p.removeList = append(p.removeList, k)
					}
				}
			} else {
				k.Repairing = false
				p.removeList = append(p.removeList, k)
			}
		}
	}
	for i := range p.removeList {
		delete(p.repairing, p.removeList[i])
	}
	p.removeList = p.removeList[:0]
}

func (p *Player) DrawPlayer(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("$$$ %d", int(p.Money)), 0, 0)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("fps: %.2f", ebiten.CurrentFPS()), 0, 20)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("tps: %.2f", ebiten.CurrentTPS()), 0, 40)
}
