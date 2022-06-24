package buildings

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/elamre/attractive_defense/platforms"
	"github.com/elamre/attractive_defense/world"
	"github.com/hajimehoshi/ebiten/v2"
)

var checkSurroundings = []struct{ x, y int }{
	{x: -1, y: 0},
	{x: 1, y: 0},
	{x: 0, y: 1},
	{x: 0, y: -1},
	{x: -2, y: 0},
	{x: 2, y: 0},
	{x: 0, y: 2},
	{x: 0, y: -2},
	{x: 1, y: 1},
	{x: -1, y: 1},
	{x: 1, y: -1},
	{x: -1, y: -1},
}

type BasicMagnet struct {
	magnetBase           *ebiten.Image
	Magnet               *ebiten.Image
	x, y, pixelX, pixelY int
	MagnetBaseOpt        *ebiten.DrawImageOptions
	magnetOpt            *ebiten.DrawImageOptions
	rotation             float64
	deleted              bool
}

func (b *BasicMagnet) SetForDeletion(g *world.Grid) {
	if b.deleted {
		return
	}
	b.deleted = true

	// We only deduct the magnetism
	for _, s := range checkSurroundings {
		cX, cY := b.x+s.x, b.y+s.y

		if g.OutOfBounds(cX, cY) {
			continue
		}
		g.RemoveMagnetism(cX, cY)
	}
	g.RemoveMagnetism(b.x, b.y)

}

func (b *BasicMagnet) Update(g *world.Grid) {
	if b.deleted {
		g.SetGrid(b.x, b.y, world.GridLevelStructures, nil)
	}
}

func (b *BasicMagnet) Draw(image *ebiten.Image) {
	image.DrawImage(b.magnetBase, b.MagnetBaseOpt)

	b.magnetOpt.GeoM.Reset()
	b.magnetOpt.GeoM.Translate(-64/2, -64/2)
	b.magnetOpt.GeoM.Rotate(b.rotation)
	b.magnetOpt.GeoM.Translate(64/2, 64/2)
	// Place at correct position
	b.magnetOpt.GeoM.Translate(float64(b.pixelX), float64(b.pixelY))
	image.DrawImage(b.Magnet, b.magnetOpt)
	b.rotation += 0.05
}

func NewBasicMagnet(x, y int, g *world.Grid) *BasicMagnet {
	b := BasicMagnet{x: x, y: y, pixelY: y * 64, pixelX: x * 64,
		magnetBase: assets.Get[*ebiten.Image](assets.AssetsMagnetBase),
		Magnet:     assets.Get[*ebiten.Image](assets.AssetsMagnet),
	}
	b.MagnetBaseOpt = &ebiten.DrawImageOptions{}
	b.magnetOpt = &ebiten.DrawImageOptions{}
	b.MagnetBaseOpt.GeoM.Translate(float64(b.pixelX), float64(b.pixelY))
	g.AddMagnetism(x, y)

	for _, s := range checkSurroundings {
		cX, cY := x+s.x, y+s.y

		if g.OutOfBounds(cX, cY) {
			continue
		}
		g.AddMagnetism(cX, cY)
		if g.IsEmpty(cX, cY) {
			boolCanPurchase := false
			for _, cc := range checkSurroundings[:4] {
				ccX, ccY := cX+cc.x, cY+cc.y
				if g.OutOfBounds(ccX, ccY) {
					continue
				}
				if g.GetGridEntity(ccX, ccY, world.GridLevelPlatform) != nil {
					boolCanPurchase = true
				}
			}
			if boolCanPurchase {
				g.SetGrid(cX, cY, world.GridLevelGui, platforms.NewPurchasePlatform(cX, cY))
			} else {
				g.SetGrid(cX, cY, world.GridLevelGui, platforms.NewCanPurchasePlatform(cX, cY))
			}
		}
	}
	return &b
}