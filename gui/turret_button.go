package gui

import (
	"github.com/elamre/attractive_defense/buildings"
	"github.com/elamre/attractive_defense/game"
	"github.com/elamre/attractive_defense/world"
	"github.com/hajimehoshi/ebiten/v2"
)

type BuyTurretButton struct {
	img *ebiten.Image
}

func NewBuyTurretButton() *BuyTurretButton {
	return &BuyTurretButton{ /*img: assets.Get[*ebiten.Image](assets.AssetsGuiTurret)*/ }
}

func (c *BuyTurretButton) Draw(screen *ebiten.Image, location *ebiten.DrawImageOptions) {
	screen.DrawImage(c.img, location)
}

func (c *BuyTurretButton) Selected(p *game.Player, b *BottomGui, g *world.Grid) {
	g.SetGrid(g.SelectedGridX, g.SelectedGridY, world.GridLevelStructures, buildings.NewBasicTurret(g.SelectedGridX, g.SelectedGridY))
	p.Money -= 100
	b.SetSelectedBuildingButtons()
}
