package buildings

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/elamre/attractive_defense/platforms"
	"github.com/elamre/attractive_defense/world"
	"github.com/hajimehoshi/ebiten/v2"
)

type LifeCrystal struct {
	anim             []*ebiten.Image
	animCnt          int
	startAnimCnt     int
	drawOpt          *ebiten.DrawImageOptions
	pixelX, pixelY   int
	x, y             int
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

func (b *LifeCrystal) Draw(image *ebiten.Image) {
	image.DrawImage(b.anim[b.animCnt], b.drawOpt)
}

func NewLifeCrystal(x, y int, g *world.Grid) *LifeCrystal {
	l := LifeCrystal{
		pixelY:           y * 64,
		pixelX:           x * 64,
		x:                x,
		y:                y,
		anim:             assets.Get[[]*ebiten.Image](assets.AssetsPlayerCrystalAnim),
		animCntDirection: 1,
	}

	l.drawOpt = &ebiten.DrawImageOptions{}
	l.drawOpt.GeoM.Translate(float64(l.pixelX), float64(l.pixelY))
	g.AddMagnetism(x, y)
	g.SetGrid(x, y, world.GridLevelPlatform, platforms.NewPlatformAt(x, y, g))

	for _, s := range checkSurroundings {
		cX, cY := x+s.x, y+s.y

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
