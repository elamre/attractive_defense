package gui

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/elamre/attractive_defense/game"
	"github.com/elamre/attractive_defense/world"
	"github.com/hajimehoshi/ebiten/v2"
)

type CancelButton struct {
	img *ebiten.Image
}

func NewCancelButton() *CancelButton {
	return &CancelButton{img: assets.Get[*ebiten.Image](assets.AssetsGuiCancel)}
}

func (c *CancelButton) Draw(screen *ebiten.Image, location *ebiten.DrawImageOptions) {
	screen.DrawImage(c.img, location)
}

func (c *CancelButton) Selected(p *game.Player, b *SideGui, g *world.Grid) {
	//b.SetIdleButtons()
	g.SetSelectedPos(-1, -1)
	// TODO set player idle here as well
}
