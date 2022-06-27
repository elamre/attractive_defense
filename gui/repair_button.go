package gui

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/elamre/attractive_defense/game"
	"github.com/elamre/attractive_defense/world"
	"github.com/hajimehoshi/ebiten/v2"
)

type RepairButton struct {
	img *ebiten.Image
}

func NewRepairButton() *RepairButton {
	return &RepairButton{img: assets.Get[*ebiten.Image](assets.AssetsGuiRepair)}
}

func (c *RepairButton) Draw(screen *ebiten.Image, location *ebiten.DrawImageOptions) {
	screen.DrawImage(c.img, location)
}

func (c *RepairButton) Selected(p *game.Player, b *SideGui, g *world.Grid) {
	if e := g.GetGridEntity(g.SelectedGridX, g.SelectedGridY, world.GridLevelStructures); e != nil {
		building := e.(world.BuildingInterface).GetBuilding()
		if building.Health < building.MaxHealth {
			building.Repairing = true
			p.AddRepairing(building)
		}
	}
	// TODO sell
}
