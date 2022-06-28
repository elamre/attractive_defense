package gui

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/elamre/attractive_defense/buildings"
	"github.com/elamre/attractive_defense/game"
	"github.com/elamre/attractive_defense/world"
	"github.com/hajimehoshi/ebiten/v2"
)

func NewRepairUpgradeButton(lab *buildings.ResearchLab) *Button {
	return &Button{
		description: "Doubles the repair speed, for the same cost!",
		image:       assets.Get[*ebiten.Image](assets.AssetsGuiResearchRepair),
		cost:        2000,
		selected: func(p *game.Player, gui *SideGui, g *world.Grid) bool {
			if p.FastRepairBeing || lab.ResearchInProgress() {
				return false
			}
			p.FastRepairBeing = true
			gui.SetResearchContext(p, lab)
			lab.StartResearch(1200, func() {
				p.FastRepair = true
			})
			return true
		},
	}
}

func NewResearchBeam(lab *buildings.ResearchLab) *Button {
	return &Button{
		description: "Unlocks the beam turret, extra strong",
		image:       assets.Get[*ebiten.Image](assets.AssetsGuiResearchBeam),
		cost:        1000,
		selected: func(p *game.Player, gui *SideGui, g *world.Grid) bool {
			if p.BeamBeingResearch || lab.ResearchInProgress() {
				return false
			}
			p.BeamBeingResearch = true
			gui.SetResearchContext(p, lab)
			lab.StartResearch(800, func() {
				p.BeamResearch = true
				if gui.showingBuildings {
					gui.SetBuildingContext(p)
				}
			})
			return true
		},
	}
}

func NewResearchRocket(lab *buildings.ResearchLab) *Button {
	return &Button{
		description: "Unlocks the rocket turret, strongest of them all, but slow",
		image:       assets.Get[*ebiten.Image](assets.AssetsGuiResearchRocket),
		cost:        1500,
		selected: func(p *game.Player, gui *SideGui, g *world.Grid) bool {
			if p.RocketBeingResearch || lab.ResearchInProgress() {
				return false
			}
			p.RocketBeingResearch = true
			gui.SetResearchContext(p, lab)
			lab.StartResearch(800, func() {
				p.RocketResearch = true
				if gui.showingBuildings {
					gui.SetBuildingContext(p)
				}
			})
			return true
		},
	}
}

func NewResearchMoney(lab *buildings.ResearchLab) *Button {
	return &Button{
		description: "Get double money for EVERYTHING",
		image:       assets.Get[*ebiten.Image](assets.AssetsGuiResearchMoney),
		cost:        5000,
		selected: func(p *game.Player, gui *SideGui, g *world.Grid) bool {
			if p.DoubleBeingMoney || lab.ResearchInProgress() {
				return false
			}
			p.DoubleBeingMoney = true
			gui.SetResearchContext(p, lab)
			lab.StartResearch(1800, func() {
				p.DoubleMoney = true
			})
			return true
		},
	}
}
