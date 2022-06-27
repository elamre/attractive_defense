package gui

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/elamre/attractive_defense/buildings/turrets"
	"github.com/elamre/attractive_defense/game"
	"github.com/elamre/attractive_defense/world"
	"github.com/hajimehoshi/ebiten/v2"
)

func NewLightTurretButton() *Button {
	return &Button{
		image:       assets.Get[*ebiten.Image](assets.AssetsGuiLightTurret),
		cost:        LightTurretCost,
		description: "Light turret, cheap but weak",
		selected: func(p *game.Player, gui *SideGui, g *world.Grid) bool {
			if IsBuildable(g) {
				e := turrets.NewLightTurret(g.SelectedGridX, g.SelectedGridY, g)
				g.SetGrid(g.SelectedGridX, g.SelectedGridY, world.GridLevelStructures, e)
				gui.SetBuildingSelectedContext(p, e)
				return true
			}
			return false
		},
	}
}
