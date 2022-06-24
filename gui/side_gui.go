package gui

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/elamre/attractive_defense/game"
	"github.com/elamre/attractive_defense/world"
	"github.com/hajimehoshi/ebiten/v2"
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
}

func NewSideGui(pixelX, pixelY int) *SideGui {
	return &SideGui{
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
	}
}

func (s *SideGui) InGui(x, y int) bool {
	if x >= s.drawX && x <= s.drawX+128 {
		if y >= s.drawY+19 && y <= (s.drawY+(s.Rows*64)+64+19) {
			return true
		}
	}
	return false
}

func (s *SideGui) Update(*game.Player, *world.Grid) {
	s.tempCounter++
	s.NoticeLevel = (s.tempCounter / 10) % 25
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
}

func (s *SideGui) SetNoticeLevel(level int) {

}
