package platforms

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/elamre/attractive_defense/world"
	"github.com/hajimehoshi/ebiten/v2"
)

type CanPurchasePlatform struct {
	PurchasePlatform
}

func NewCanPurchasePlatform(x, y int) *CanPurchasePlatform {
	c := CanPurchasePlatform{PurchasePlatform: *NewPurchasePlatform(x, y)}
	c.options.ColorM.Translate(0, 0, 0, -0.5)
	return &c
}

func (p *CanPurchasePlatform) Update(g *world.Grid) {
	// We do nothing
}

func (p *CanPurchasePlatform) Draw(image *ebiten.Image) {
	image.DrawImage(p.image, p.options)
}

type PurchasePlatform struct {
	image          *ebiten.Image
	options        *ebiten.DrawImageOptions
	pixelX, pixelY int
	x, y           int
}

func (p *PurchasePlatform) SetForDeletion(g *world.Grid) {
	// Nothign happens in the gui layer
}

func (p *PurchasePlatform) Update(g *world.Grid) {
	/*if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if x > p.pixelX && y > p.pixelY {
			if x < p.pixelX+64 && y < p.pixelY+64 {
				NewPlatformAt(p.x, p.y, g)
				g.SetGrid(p.x, p.y, world.GridLevelGui, nil)
			}
		}
	}*/
}

func (p *PurchasePlatform) Draw(image *ebiten.Image) {
	image.DrawImage(p.image, p.options)
}

func NewPurchasePlatform(x, y int) *PurchasePlatform {
	opt := ebiten.DrawImageOptions{}
	opt.GeoM.Translate(float64(x*64), float64(y*64))
	return &PurchasePlatform{options: &opt, image: assets.Get[*ebiten.Image](assets.AssetsPlusSymbol), pixelX: x * 64, pixelY: y * 64, x: x, y: y}
}
