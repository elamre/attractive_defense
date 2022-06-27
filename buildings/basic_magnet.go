package buildings

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/elamre/attractive_defense/platforms"
	"github.com/elamre/attractive_defense/world"
	"github.com/hajimehoshi/ebiten/v2"
)

type BasicMagnet struct {
	*world.Building
	magnetBase     *ebiten.Image
	Magnet         *ebiten.Image
	MagnetBaseOpt  *ebiten.DrawImageOptions
	magnetOpt      *ebiten.DrawImageOptions
	rotation       float64
	deleted        bool
	setForDeletion bool
}

func (b *BasicMagnet) InflictDamage(damage float64) {
	b.Health -= damage
	if !b.Alive() {
		b.Health = 0
		b.setForDeletion = true
	}
}
func (b *BasicMagnet) Alive() bool {
	return b.Health >= 1
}

func (b *BasicMagnet) SetForDeletion(g *world.Grid) {
	b.setForDeletion = true
}

func (b *BasicMagnet) Update(g *world.Grid) {
	if !b.deleted && b.setForDeletion {
		g.SetGrid(b.GridX, b.GridY, world.GridLevelStructures, nil)
		for _, s := range assets.Surroundings5 {
			cX, cY := b.GridX+s.X, b.GridY+s.Y

			if g.OutOfBounds(cX, cY) {
				continue
			}
			g.RemoveMagnetism(cX, cY)
		}
		g.RemoveMagnetism(b.GridX, b.GridY)
		b.deleted = true
	}
}

func (b *BasicMagnet) Draw(image *ebiten.Image) {
	image.DrawImage(b.magnetBase, b.MagnetBaseOpt)

	b.magnetOpt.GeoM.Reset()
	b.magnetOpt.GeoM.Translate(-64/2, -64/2)
	b.magnetOpt.GeoM.Rotate(b.rotation)
	b.magnetOpt.GeoM.Translate(64/2, 64/2)
	// Place at correct position
	b.magnetOpt.GeoM.Translate(b.PixelX, b.PixelY)
	image.DrawImage(b.Magnet, b.magnetOpt)
	b.rotation += 0.05
	b.DrawGui(image)
}

func (d *BasicMagnet) GetBuilding() *world.Building {
	return d.Building
}
func NewBasicMagnet(x, y int, g *world.Grid) *BasicMagnet {
	b := BasicMagnet{
		Building: &world.Building{
			PixelX:    float64(x * 64),
			PixelY:    float64(y * 64),
			GridX:     x,
			GridY:     y,
			Repairing: false,
			Health:    50,
			MaxHealth: 50,
		},
		magnetBase: assets.Get[*ebiten.Image](assets.AssetsMagnetBase),
		Magnet:     assets.Get[*ebiten.Image](assets.AssetsMagnet),
	}
	b.MagnetBaseOpt = &ebiten.DrawImageOptions{}
	b.magnetOpt = &ebiten.DrawImageOptions{}
	b.MagnetBaseOpt.GeoM.Translate(b.PixelX, b.PixelY)
	g.AddMagnetism(x, y)

	for _, s := range assets.Surroundings5 {
		cX, cY := x+s.X, y+s.Y

		if g.OutOfBounds(cX, cY) {
			continue
		}
		g.AddMagnetism(cX, cY)
		if g.IsEmpty(cX, cY) {
			boolCanPurchase := false
			for _, cc := range assets.Surroundings5[:4] {
				ccX, ccY := cX+cc.X, cY+cc.Y
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
