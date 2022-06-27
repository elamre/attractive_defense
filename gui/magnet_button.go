package gui

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/elamre/attractive_defense/buildings"
	"github.com/elamre/attractive_defense/game"
	"github.com/elamre/attractive_defense/world"
	"github.com/hajimehoshi/ebiten/v2"
)

func NewMagnetButton() *Button {
	return &Button{
		image:       assets.Get[*ebiten.Image](assets.AssetsGuiMagnet),
		cost:        MagnetCost,
		description: "Use this to expand your base",
		selected: func(p *game.Player, gui *SideGui, g *world.Grid) bool {
			if IsBuildable(g) {
				p.Risk++
				gui.NoticeLevel = p.Risk
				g.SetGrid(g.SelectedGridX, g.SelectedGridY, world.GridLevelStructures, buildings.NewBasicMagnet(g.SelectedGridX, g.SelectedGridY, g))
				gui.NoButtonsContext()
				return true
			}
			return false
		},
	}
}
