package gui

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/elamre/attractive_defense/buildings"
	"github.com/elamre/attractive_defense/game"
	"github.com/elamre/attractive_defense/world"
	"github.com/hajimehoshi/ebiten/v2"
)

type BuyMagnetButton struct {
	img *ebiten.Image
}

func NewBuyMagnetButton() *BuyMagnetButton {
	return &BuyMagnetButton{img: assets.Get[*ebiten.Image](assets.AssetsGuiMagnet)}
}

func (c *BuyMagnetButton) Draw(screen *ebiten.Image, location *ebiten.DrawImageOptions) {
	screen.DrawImage(c.img, location)
}

func (c *BuyMagnetButton) Selected(p *game.Player, b *BottomGui, g *world.Grid) {
	g.SetGrid(g.SelectedGridX, g.SelectedGridY, world.GridLevelStructures, buildings.NewBasicMagnet(g.SelectedGridX, g.SelectedGridY, g))
	p.Money -= 200
	b.SetSelectedBuildingButtons()
}
