package buildings

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/elamre/attractive_defense/world"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/colornames"
)

type ResearchLab struct {
	*world.Building
	anim                []*ebiten.Image
	animCnt             int
	startAnimCnt        int
	drawOpt             *ebiten.DrawImageOptions
	researchCb          func()
	researchDuration    int
	researchMaxDuration int
	setForDeletion      bool
	deleted             bool
}

func (b *ResearchLab) SetForDeletion(g *world.Grid) {
	b.setForDeletion = true
}

func (b *ResearchLab) ResearchInProgress() bool {
	return b.researchDuration > 0
}

func (b *ResearchLab) StartResearch(duration int, callBack func()) {
	b.researchCb = callBack
	b.researchDuration = duration
	b.researchMaxDuration = duration
}

func (b *ResearchLab) Update(g *world.Grid) {
	if !b.Alive() {
		b.SetForDeletion(g)
	}
	if !b.deleted && b.setForDeletion {
		g.SetGrid(b.GridX, b.GridY, world.GridLevelStructures, nil)
		b.deleted = true
	}

	if b.researchDuration > 0 {
		if b.startAnimCnt < 4 {
			b.startAnimCnt++
		}
		if b.startAnimCnt >= 4 {
			b.startAnimCnt = 0
			b.animCnt++
		}
		if b.animCnt == 7 {
			b.animCnt = 0
		}
		b.researchDuration--
		if b.researchDuration <= 0 {
			if b.researchCb != nil {
				b.researchCb()
				b.researchDuration = 0
			}
		}
	} else {
		b.animCnt = 0
		b.startAnimCnt = 0
	}
}

func (d *ResearchLab) GetBuilding() *world.Building {
	return d.Building
}

func (b *ResearchLab) InflictDamage(damage float64) {
	b.Health -= damage
}
func (b *ResearchLab) Alive() bool {
	return b.Health >= 1
}

func (b *ResearchLab) Draw(image *ebiten.Image) {
	image.DrawImage(b.anim[b.animCnt], b.drawOpt)
	if b.researchDuration > 0 {
		ebitenutil.DrawRect(image, b.PixelX+2, b.PixelY+54, 60-((float64(b.researchDuration)/float64(b.researchMaxDuration))*60), 4, colornames.Blue)
	}
	b.DrawGui(image)
}

func NewResearchLab(x, y int) *ResearchLab {
	l := ResearchLab{
		Building: &world.Building{
			PixelX:    float64(x * 64),
			PixelY:    float64(y * 64),
			GridX:     x,
			GridY:     y,
			Repairing: false,
			Health:    200,
			MaxHealth: 200,
		},
		anim: assets.Get[[]*ebiten.Image](assets.AssetsResearchLabAnim),
	}

	l.drawOpt = &ebiten.DrawImageOptions{}
	l.drawOpt.GeoM.Translate(l.PixelX, l.PixelY)

	return &l
}
