package gui

import (
	"github.com/elamre/attractive_defense/game"
	"github.com/elamre/attractive_defense/world"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type BottomGuiButton interface {
	Selected(p *game.Player, gui *BottomGui, g *world.Grid)
	Draw(screen *ebiten.Image, location *ebiten.DrawImageOptions)
}

var (
	cancelButton    *CancelButton
	buyTurretButton *BuyTurretButton
	sellButton      *SellButton
	buyMagnet       *BuyMagnetButton
)

type BottomGui struct {
	opt       ebiten.DrawImageOptions
	emptyTile *ebiten.Image
	buttons   []BottomGuiButton
}

func NewBottomGui() *BottomGui {
	b := &BottomGui{
		//emptyTile: assets.Get[*ebiten.Image](assets.AssetsGuiEmpty),
		buttons: make([]BottomGuiButton, 10),
	}
	//cancelButton = NewCancelButton()
	//buyTurretButton = NewBuyTurretButton()
	//sellButton = NewSellButton()
	//buyMagnet = NewBuyMagnetButton()
	return b
}

func (b *BottomGui) SetSelectedBuildingButtons() {
	for i := range b.buttons {
		b.buttons[i] = nil
	}
	//b.buttons[0] = cancelButton
	//b.buttons[1] = sellButton
}

func (b *BottomGui) SetBuildingsButtons() {
	for i := range b.buttons {
		b.buttons[i] = nil
	}
	//b.buttons[0] = cancelButton
	//b.buttons[1] = buyMagnet
	//b.buttons[2] = buyTurretButton
}

func (b *BottomGui) SetPowersButtons() {
	for i := range b.buttons {
		b.buttons[i] = nil
	}
	//b.buttons[0] = cancelButton
}

func (b *BottomGui) SetIdleButtons() {
	for i := range b.buttons {
		b.buttons[i] = nil
	}
	//b.buttons[0] = cancelButton
}

func (b *BottomGui) Update(p *game.Player, g *world.Grid) {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()
		if b.MouseOnBottom(mx, my) {
			if button := b.buttons[mx/64]; button != nil {
				button.Selected(p, b, g)
			}
		} else {
			x, y := g.MouseToGridPos(mx, my)
			g.SetSelectedPos(x, y)
			b.SetIdleButtons()
			if x != -1 && y != -1 {
				if g.GetGridEntity(x, y, world.GridLevelStructures) != nil {
					b.SetSelectedBuildingButtons()
				} else if g.GetGridEntity(x, y, world.GridLevelPlatform) != nil {
					b.SetBuildingsButtons()
				}
			}
		}
	} else if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		g.SetSelectedPos(-1, -1)
	}
}

func (b *BottomGui) MouseOnBottom(mx, my int) bool {
	return mx < 640 && my > 1024-64
}

func (b *BottomGui) Draw(screen *ebiten.Image) {
	if true {
		return
	}
	opt := ebiten.DrawImageOptions{}
	opt.GeoM.Translate(0, 1024-64)
	for i := 0; i < 10; i++ {
		screen.DrawImage(b.emptyTile, &opt)
		if button := b.buttons[i]; button != nil {
			button.Draw(screen, &opt)
		}
		opt.GeoM.Translate(float64(64), 0)
	}
}
