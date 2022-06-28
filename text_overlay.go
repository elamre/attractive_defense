package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	"log"
)

type TextOverlay struct {
	countdownTimer    int
	waveFinishedTimer int
	goTimer           int
	neverStarted      bool
	started           bool
	mplusNormalFont   font.Face
}

func NewTextOverlay() *TextOverlay {

	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 72
	mplusNormalFont, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	t := TextOverlay{
		countdownTimer:    -1,
		waveFinishedTimer: -1,
		goTimer:           -1,
		neverStarted:      true,
		mplusNormalFont:   mplusNormalFont,
	}
	return &t
}

func (t *TextOverlay) StartCountdown() {
	t.neverStarted = false
	t.started = true
	t.waveFinishedTimer = 0
}

func (t *TextOverlay) Finished() bool {
	return t.countdownTimer == -1 && t.waveFinishedTimer == -1 && t.goTimer == -1 && !t.neverStarted && t.started
}

func (t *TextOverlay) Reset() {
	t.started = false
}

func (t *TextOverlay) Update() {
	if t.neverStarted {
		return
	}
	if t.waveFinishedTimer != -1 {
		t.waveFinishedTimer++
		if t.waveFinishedTimer >= 60*3 {
			t.waveFinishedTimer = -1
			t.countdownTimer = 0
		}
	} else if t.countdownTimer != -1 {
		t.countdownTimer++
		if t.countdownTimer >= 60*5 {
			t.countdownTimer = -1
			t.goTimer = 0
		}
	} else if t.goTimer != -1 {
		t.goTimer++
		if t.goTimer >= 60*2 {
			t.goTimer = -1
		}
	}
}

func (t *TextOverlay) Draw(screen *ebiten.Image) {

	ss := ""
	if t.goTimer != -1 {
		ss = "GO!"
	} else if t.countdownTimer != -1 {
		ss = fmt.Sprintf("%d", 5-(t.countdownTimer/60))
		//text.Draw(screen, fmt.Sprintf("%d", 5-(t.countdownTimer/60)), t.mplusNormalFont, 350, 360, color.White)
	} else if t.waveFinishedTimer != -1 {
		ss = "Wave Finished"
		//text.Draw(screen, "Wave Finished", t.mplusNormalFont, 350, 360, color.White)
	} else if t.neverStarted {
		ss = "Press space to start"
	}
	if !t.Finished() || t.neverStarted {
		tt := text.BoundString(t.mplusNormalFont, ss)
		text.Draw(screen, ss, t.mplusNormalFont, 400-tt.Size().X/2, 300-tt.Size().Y/2, color.White)
	}
}
