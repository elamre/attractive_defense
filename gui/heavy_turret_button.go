package gui

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/elamre/attractive_defense/buildings/turrets"
	"github.com/elamre/attractive_defense/game"
	"github.com/elamre/attractive_defense/world"
	"github.com/hajimehoshi/ebiten/v2"
)

func NewHeavyTurretButton() *Button {
	return &Button{
		image:       assets.Get[*ebiten.Image](assets.AssetsGuiHeavyTurret),
		cost:        HeavyTurretCost,
		description: "Heavy turret. Shoots slow but powerful",
		selected: func(p *game.Player, gui *SideGui, g *world.Grid) bool {
			if IsBuildable(g) {
				e := turrets.NewHeavyTurret(g.SelectedGridX, g.SelectedGridY, g)
				g.SetGrid(g.SelectedGridX, g.SelectedGridY, world.GridLevelStructures, e)
				gui.SetBuildingSelectedContext(p, e)
				return true
			}
			return false
		},
	}
}
