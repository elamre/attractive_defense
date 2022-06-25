package gui

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/elamre/attractive_defense/buildings/turrets"
	"github.com/elamre/attractive_defense/game"
	"github.com/elamre/attractive_defense/platforms"
	"github.com/elamre/attractive_defense/world"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type SideGui struct {
	top                *ebiten.Image
	bot                *ebiten.Image
	mid                *ebiten.Image
	noticeLevelImgLow  *ebiten.Image
	noticeLevelImgMed  *ebiten.Image
	noticeLevelImgHigh *ebiten.Image
	Rows               int
	NoticeLevel        int // max 24
	drawX, drawY       int

	tempCounter int
	buttonArray []*Button

	available map[string]bool

	showingBuildings bool
}

func NewSideGui(pixelX, pixelY int) *SideGui {
	s := &SideGui{
		drawY:              pixelY,
		drawX:              pixelX,
		top:                assets.Get[*ebiten.Image](assets.AssetsGuiTopPart),
		bot:                assets.Get[*ebiten.Image](assets.AssetsGuiBottomBorder),
		mid:                assets.Get[*ebiten.Image](assets.AssetsGuiMiddle),
		noticeLevelImgLow:  assets.Get[*ebiten.Image](assets.AssetsGuiWarningLevelLow),
		noticeLevelImgMed:  assets.Get[*ebiten.Image](assets.AssetsGuiWarningLevelMedium),
		noticeLevelImgHigh: assets.Get[*ebiten.Image](assets.AssetsGuiWarningLevelHigh),
		Rows:               8,
		NoticeLevel:        24,
		available:          InitButtons(),
		showingBuildings:   true,
	}
	s.buttonArray = GetBuildingButtons(s.available)
	return s
}

func (s *SideGui) InGui(x, y int) bool {
	if x >= s.drawX && x <= s.drawX+128 {
		if y >= s.drawY+19 && y <= (s.drawY+(s.Rows*64)+64+19) {
			return true
		}
	}
	return false
}

func (s *SideGui) SetBuildingContext() {
	s.buttonArray = GetBuildingButtons(s.available)
}

func (s *SideGui) SetBuildingSelectedContext(p *game.Player, e world.GridEntity) {
	if t, ok := e.(*turrets.Turret); ok {
		s.buttonArray = []*Button{NewUpgradableButton(t.Gun), NewUpgradableButton(t.Base)}
	}
}

func (s *SideGui) Update(p *game.Player, g *world.Grid, camera *Camera) {
	s.tempCounter++
	s.NoticeLevel = (s.tempCounter / 10) % 25

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()

		if s.InGui(mx, my) {
			xPos := 0
			if mx > s.drawX+64 {
				xPos = 1
			}
			yIndex := (my - 19 - 64) / 64
			index := (yIndex * 2) + xPos
			if index < len(s.buttonArray) {
				s.buttonArray[index].selected(p, s, g)
			}
			if e := g.GetGridEntity(g.SelectedGridX, g.SelectedGridY, world.GridLevelStructures); e != nil {
				s.SetBuildingSelectedContext(p, e)
			}
		} else {
			cMx, cMy := camera.ScreenToWorld(mx, my)
			x, y := g.MouseToGridPos(int(cMx), int(cMy))
			g.SetSelectedPos(x, y)
			if e := g.GetGridEntity(x, y, world.GridLevelGui); e != nil {
				if _, ok := e.(*platforms.PurchasePlatform); ok {
					// This should be much better TODO
					platforms.NewPlatformAt(x, y, g)
					g.SetGrid(x, y, world.GridLevelGui, nil)
				}
			}
			if x != -1 && y != -1 {
				if e := g.GetGridEntity(x, y, world.GridLevelStructures); e != nil {
					s.SetBuildingSelectedContext(p, e)
				} else if g.GetGridEntity(x, y, world.GridLevelPlatform) != nil {
					s.SetBuildingContext()
				}
			}
		}
	} else if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		g.SetSelectedPos(-1, -1)
	}
	//s.showingBuildings = IsBuildable(g)
}

func (s *SideGui) Draw(screen *ebiten.Image) {
	noticeOpts := ebiten.DrawImageOptions{}

	noticeOpts.GeoM.Translate(float64(s.drawX), float64(s.drawY)+19-64)
	screen.DrawImage(s.top, &noticeOpts)
	noticeOpts.GeoM.Translate(0, 128)
	for i := 0; i < s.Rows; i++ {
		screen.DrawImage(s.mid, &noticeOpts)
		noticeOpts.GeoM.Translate(0, 64)
	}
	screen.DrawImage(s.bot, &noticeOpts)
	noticeOpts.GeoM.Reset()
	noticeOpts.GeoM.Translate(float64(s.drawX)+4, float64(s.drawY))
	for i := 0; i < s.NoticeLevel; i++ {
		if i < 8 {
			screen.DrawImage(s.noticeLevelImgLow, &noticeOpts)
		} else if i < 16 {
			screen.DrawImage(s.noticeLevelImgMed, &noticeOpts)
		} else {
			screen.DrawImage(s.noticeLevelImgHigh, &noticeOpts)
		}
		noticeOpts.GeoM.Translate(5, 0)
	}
	noticeOpts.GeoM.Reset()
	noticeOpts.GeoM.Translate(float64(s.drawX)+5, float64(s.drawY)+19+64)
	if !s.showingBuildings {
		//noticeOpts.ColorM.Translate(-0.5, -0.5, -0.5, 0)
	}
	if s.buttonArray != nil {
		for i, ss := range s.buttonArray {
			ss.Draw(screen, &noticeOpts)
			noticeOpts.GeoM.Translate(60, 0)

			if i%2 == 1 {
				noticeOpts.GeoM.Translate(-120, 64)
			}
		}
	}
}

func (s *SideGui) SetNoticeLevel(level int) {

}
