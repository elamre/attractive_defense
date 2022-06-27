package gui

import (
	"github.com/elamre/attractive_defense/buildings"
	"github.com/elamre/attractive_defense/buildings/turrets"
	"github.com/elamre/attractive_defense/game"
	"github.com/elamre/attractive_defense/world"
)

func NewUpgradableButton(upgradable buildings.UpgradeAble, e world.GridEntity) *Button {
	return &Button{
		description: upgradable.Description(),
		image:       upgradable.GetUpgradeButton(),
		cost:        float64(upgradable.UpgradeCost()),
		selected: func(p *game.Player, gui *SideGui, g *world.Grid) bool {
			if upgradable.UpgradeCost() > 0 {
				upgradable.Upgrade()
				if t, ok := e.(*turrets.Turret); ok {
					t.Upgraded(upgradable)
				}
				gui.SetBuildingSelectedContext(p, e)
				return true
			}
			return false
		},
	}
}
