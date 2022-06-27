package buildings

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/elamre/attractive_defense/platforms"
	"github.com/elamre/attractive_defense/world"
	"github.com/hajimehoshi/ebiten/v2"
)

type LifeCrystal struct {
	*world.Building
	anim             []*ebiten.Image
	animCnt          int
	startAnimCnt     int
	drawOpt          *ebiten.DrawImageOptions
	animCntDirection int
}

func (b *LifeCrystal) SetForDeletion(g *world.Grid) {

}

func (b *LifeCrystal) Update(g *world.Grid) {
	if b.animCntDirection == -1 && b.animCnt == 0 {
		b.animCntDirection = 1
		b.startAnimCnt = 0
	}
	if b.startAnimCnt < 100 {
		b.startAnimCnt++
	}
	if b.startAnimCnt == 100 {
		b.animCnt += b.animCntDirection
	}
	if b.animCnt == 8 {
		b.animCntDirection = -1
		b.animCnt--
	}

}

func (d *LifeCrystal) GetBuilding() *world.Building {
	return d.Building
}

func (b *LifeCrystal) InflictDamage(damage float64) {
	b.Health -= damage
	if b.Health < 0 {
		b.Health = 0
		// TODO
	}
}
func (b *LifeCrystal) Alive() bool {
	return b.Health >= 1
}

func (b *LifeCrystal) Draw(image *ebiten.Image) {
	image.DrawImage(b.anim[b.animCnt], b.drawOpt)
	b.DrawGui(image)
}

func NewLifeCrystal(x, y int, g *world.Grid) *LifeCrystal {
	l := LifeCrystal{
		Building: &world.Building{
			PixelX:    float64(x * 64),
			PixelY:    float64(y * 64),
			GridX:     x,
			GridY:     y,
			Repairing: false,
			Health:    1000,
			MaxHealth: 1000,
		},
		anim:             assets.Get[[]*ebiten.Image](assets.AssetsPlayerCrystalAnim),
		animCntDirection: 1,
	}

	l.drawOpt = &ebiten.DrawImageOptions{}
	l.drawOpt.GeoM.Translate(l.PixelX, l.PixelY)
	g.AddMagnetism(x, y)
	g.SetGrid(x, y, world.GridLevelPlatform, platforms.NewPlatformAt(x, y, g))

	for _, s := range assets.Surroundings5 {
		cX, cY := x+s.X, y+s.Y

		if g.OutOfBounds(cX, cY) {
			continue
		}
		g.AddMagnetism(cX, cY)
		if g.GetGridEntity(cX, cY, world.GridLevelPlatform) == nil {
			g.SetGrid(cX, cY, world.GridLevelPlatform, platforms.NewPlatformAt(cX, cY, g))
		}
	}
	return &l
}
