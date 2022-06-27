package gui

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/elamre/attractive_defense/buildings/turrets"
	"github.com/elamre/attractive_defense/game"
	"github.com/elamre/attractive_defense/world"
	"github.com/hajimehoshi/ebiten/v2"
)

func NewRocketTurretButton() *Button {
	return &Button{
		image:       assets.Get[*ebiten.Image](assets.AssetsGuiRocketTurret),
		cost:        RocketTurretCost,
		description: "Large range dealing massive damage with homing missiles",
		selected: func(p *game.Player, gui *SideGui, g *world.Grid) bool {
			if IsBuildable(g) {
				e := turrets.NewRocketTurret(g.SelectedGridX, g.SelectedGridY, g)
				g.SetGrid(g.SelectedGridX, g.SelectedGridY, world.GridLevelStructures, e)
				gui.SetBuildingSelectedContext(p, e)
				return true
			}
			return false
		},
	}
}
