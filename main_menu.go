package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
)

type MainMenu struct {
	InCredits   bool
	font        font.Face
	exitPressed bool
	playPressed bool
}

func NewMainMenu() *MainMenu {
	m := MainMenu{}
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		panic(err)
	}
	const dpi = 72
	mplusNormalFont, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		panic(err)
	}
	m.font = mplusNormalFont
	return &m
}

func (m *MainMenu) Draw(screen *ebiten.Image) {
	if m.InCredits {
		text.Draw(screen, "Voice Acting, Programming and Sprite adjusting: Elmar van Rijnswou\nEbitengine + Splash Screen: Hajime Hoshi\nSoundtrack: Roald Strauss\nSpace ship sprites: Michael Williams\n\nPress escape/space", m.font, 20, 20, color.White)
	} else {
		ebitenutil.DrawRect(screen, 300, 80, 200, 80, colornames.Red)
		ebitenutil.DrawRect(screen, 310, 90, 180, 60, colornames.Black)
		text.Draw(screen, "Play Game", m.font, 340, 125, color.White)

		ebitenutil.DrawRect(screen, 300, 240, 200, 80, colornames.Red)
		ebitenutil.DrawRect(screen, 310, 250, 180, 60, colornames.Black)
		text.Draw(screen, "Credits", m.font, 350, 240+47, color.White)

		ebitenutil.DrawRect(screen, 300, 400, 200, 80, colornames.Red)
		ebitenutil.DrawRect(screen, 310, 410, 180, 60, colornames.Black)
		text.Draw(screen, "Exit", m.font, 375, 400+50, color.White)

	}
}

func (m *MainMenu) Start() bool {
	if m.playPressed {
		m.playPressed = false
		return true
	}
	return false
}

func (m *MainMenu) Exit() bool {
	if m.exitPressed {
		m.exitPressed = false
		return true
	}
	return false
}

func (m *MainMenu) Update() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && !m.InCredits {
		x, y := ebiten.CursorPosition()
		if x > 300 && x < 500 {
			if y > 80 && y < 160 {
				m.playPressed = true
			} else if y > 240 && y < 320 {
				m.InCredits = true
			} else if y > 400 && y < 480 {
				m.exitPressed = true
			}
		}
	} else if m.InCredits {
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) || inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			m.InCredits = false
		}
	}
}
