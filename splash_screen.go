package main

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type SplashScreen struct {
	img          *ebiten.Image
	up           bool
	transparency float64
}

func NewSplashScreen() *SplashScreen {
	return &SplashScreen{img: assets.SplashImage, up: true, transparency: 0}
}
func (s *SplashScreen) Update() bool {
	if s.up {
		s.transparency += 0.03
		if s.transparency >= 3 {
			s.up = false
		}
	} else {
		s.transparency -= 0.03
		if s.transparency <= 0 {
			return true
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) || inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return true
	}
	return false
}

func (s *SplashScreen) Draw(screen *ebiten.Image) {
	opt := ebiten.DrawImageOptions{}
	trans := -1 + s.transparency
	if trans > 1 {
		trans = 1
	}
	opt.ColorM.Translate(0, 0, 0, trans)
	opt.GeoM.Scale(0.40, 0.40)
	opt.GeoM.Translate((800-(1920*0.4))/2, (600-(1080*0.4))/2)
	screen.DrawImage(s.img, &opt)
}
