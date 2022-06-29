package gui

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/elamre/attractive_defense/game"
	"github.com/elamre/attractive_defense/world"
	"github.com/hajimehoshi/ebiten/v2"
)

const MagnetCost = 1000
const ResearchCost = 1500
const LightTurretCost = 200
const HeavyTurretCost = 600
const BeamTurretCost = 1000
const RocketTurretCost = 1500

type Button struct {
	cost        float64
	selected    func(p *game.Player, gui *SideGui, g *world.Grid) bool
	image       *ebiten.Image
	description string
}

func (b *Button) Draw(screen *ebiten.Image, location *ebiten.DrawImageOptions) {
	if b.cost > 0 {
		screen.DrawImage(b.image, location)
	}
}

func (b *Button) Selected(p *game.Player, gui *SideGui, g *world.Grid) {
	if p.Money >= b.cost && b.cost > 0 {
		if b.selected(p, gui, g) {
			p.Money -= b.cost
		}
	} else if p.Money < b.cost {
		assets.StaticSoundManager.PlayInsufficientFunds()
	}
}

func IsBuildable(g *world.Grid) bool {
	if g.OutOfBounds(g.SelectedGridX, g.SelectedGridY) {
		return false
	}
	if g.GetGridEntity(g.SelectedGridX, g.SelectedGridY, world.GridLevelPlatform) == nil {
		return false
	}
	if g.GetGridEntity(g.SelectedGridX, g.SelectedGridY, world.GridLevelStructures) != nil {
		return false
	}
	return true
}
