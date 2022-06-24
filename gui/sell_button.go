package gui

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/elamre/attractive_defense/game"
	"github.com/elamre/attractive_defense/world"
	"github.com/hajimehoshi/ebiten/v2"
)

type SellButton struct {
	img *ebiten.Image
}

func NewSellButton() *SellButton {
	return &SellButton{img: assets.Get[*ebiten.Image](assets.AssetsGuiSell)}
}

func (c *SellButton) Draw(screen *ebiten.Image, location *ebiten.DrawImageOptions) {
	screen.DrawImage(c.img, location)
}

func (c *SellButton) Selected(p *game.Player, b *SideGui, g *world.Grid) {
	if e := g.GetGridEntity(g.SelectedGridX, g.SelectedGridY, world.GridLevelStructures); e != nil {
		//b.SetBuildingsButtons()
		e.SetForDeletion(g)
	} else if e = g.GetGridEntity(g.SelectedGridX, g.SelectedGridY, world.GridLevelPlatform); e != nil {
		e.SetForDeletion(g)
	}
	// TODO sell
}
