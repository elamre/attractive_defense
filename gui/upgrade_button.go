package gui

import (
	"github.com/elamre/attractive_defense/buildings"
	"github.com/elamre/attractive_defense/game"
	"github.com/elamre/attractive_defense/world"
)

func NewUpgradableButton(upgradable buildings.UpgradeAble) *Button {
	return &Button{
		image: upgradable.GetUpgradeButton(),
		cost:  upgradable.UpgradeCost(),
		selected: func(p *game.Player, gui *SideGui, g *world.Grid) bool {
			upgradable.Upgrade()
			return true
		},
	}
}
