package gui

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/elamre/attractive_defense/buildings"
	"github.com/elamre/attractive_defense/game"
	"github.com/elamre/attractive_defense/world"
	"github.com/hajimehoshi/ebiten/v2"
)

func NewLightTurretButton() *Button {
	return &Button{
		image: assets.Get[*ebiten.Image](assets.AssetsGuiLightTurret),
		cost:  LightTurretCost,
		selected: func(p *game.Player, gui *SideGui, g *world.Grid) bool {
			if IsBuildable(g) {
				g.SetGrid(g.SelectedGridX, g.SelectedGridY, world.GridLevelStructures, buildings.NewBasicTurret(g.SelectedGridX, g.SelectedGridY))
				return true
			}
			return false
		},
	}
}
