package platforms

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/elamre/attractive_defense/world"
	"github.com/hajimehoshi/ebiten/v2"
)

type BasicPlatform struct {
	image          *ebiten.Image
	x, y           int
	pixelX, pixelY int
	grid           *world.Grid
	options        *ebiten.DrawImageOptions
	destroy        bool
}

func (b *BasicPlatform) SetForDeletion(g *world.Grid) {
	b.destroy = true
	g.SetGrid(b.x, b.y, world.GridLevelPlatform, nil)

}

func (b *BasicPlatform) Update(g *world.Grid) {

}

func (b *BasicPlatform) Draw(image *ebiten.Image) {
	image.DrawImage(b.image, b.options)
}

func NewPlatformAt(x, y int, grid *world.Grid) *BasicPlatform {
	b := BasicPlatform{x: x, y: y, grid: grid}
	b.options = &ebiten.DrawImageOptions{}
	b.pixelY = y * 64
	b.pixelX = x * 64
	b.options.GeoM.Translate(float64(b.pixelX), float64(b.pixelY))
	b.image = assets.Get[*ebiten.Image](assets.AssetsPlatformImage)
	grid.SetGrid(x, y, world.GridLevelPlatform, &b)
	grid.SetGrid(x, y, world.GridLevelGui, nil)
	checkSurroundings := []struct{ x, y int }{{x: -1, y: 0}, {x: 1, y: 0}, {x: 0, y: 1}, {x: 0, y: -1}}
	for _, s := range checkSurroundings {
		cX, cY := x+s.x, y+s.y
		if grid.OutOfBounds(cX, cY) {
			continue
		}
		if grid.GetGridEntity(cX, cY, world.GridLevelPlatform) == nil {
			if ent := grid.GetGridEntity(cX, cY, world.GridLevelGui); ent != nil {
				if _, ok := ent.(*CanPurchasePlatform); ok {
					grid.SetGrid(cX, cY, world.GridLevelGui, NewPurchasePlatform(cX, cY))
				}
			}
		}
	}
	return &b
}
