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

	FastRepairBeing     bool
	FastRepair          bool
	DoubleBeingMoney    bool
	DoubleMoney         bool
	RocketBeingResearch bool
	RocketResearch      bool
	BeamBeingResearch   bool
	BeamResearch        bool
	Score               int
	WaveNumber          int
	//gui   *gui.BottomGui
}

func NewPlayer() *Player {

	return &Player{Money: 2000, Risk: 1, repairing: make(map[*world.Building]bool), removeList: make([]*world.Building, 0), BeamResearch: false, RocketResearch: false}
}

func (p *Player) AddRepairing(b *world.Building) {
	p.repairing[b] = true
}

func (p *Player) UpdatePlayer(g *world.Grid) {
	for k, _ := range p.repairing {
		if k.Repairing {
			if p.Money > k.RepairPerTick() {
				if p.FastRepair {
					k.RepairTick()
					k.RepairTick()
				}
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
	if p.WaveNumber > 0 {
		p.Money += float64(p.WaveNumber+(p.Risk-1)) * 0.1
	}
	p.removeList = p.removeList[:0]
}

func (p *Player) DrawPlayer(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("$$$ %d", int(p.Money)), 10, 10)
	/*	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("fps: %.2f", ebiten.CurrentFPS()), 0, 20)
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("tps: %.2f", ebiten.CurrentTPS()), 0, 40)*/
}

func (p *Player) WaveFinished(waveNumber, attentionLevel int) {
	p.Risk = attentionLevel
	p.WaveNumber = waveNumber
	p.AddMoney(float64(waveNumber+1) * 100)

}

func (p *Player) AddMoney(reward float64) {
	multiPlier := 1.0
	if p.Risk >= 8 {
		multiPlier += 0.5
	}
	if p.Risk >= 16 {
		multiPlier += 0.5
	}
	if p.Risk >= 23 {
		multiPlier += 0.5
	}
	p.Score += int(reward * multiPlier)
	p.Money += reward * multiPlier
	if p.DoubleMoney {
		p.Money += reward * multiPlier
	}
}
