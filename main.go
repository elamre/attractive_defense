package main

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/elamre/attractive_defense/game"
	"github.com/elamre/gameutil"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

const InSplashScreen = 1
const InMenuScreen = 2
const InGameScreen = 3

type AD struct {
	whichScreen int
	gameScreen  *GameScreen
	menuScreen  *MainMenu
	splayScreen *SplashScreen
}

var stars game.StarBackground

func NewAttractiveDefense(width, height int) *AD {
	assets.StaticSoundManager = assets.NewSoundManager()

	return &AD{gameScreen: NewGameScreen(800, 600, width, height), menuScreen: NewMainMenu(), whichScreen: InGameScreen, splayScreen: NewSplashScreen()}
}

func (ad *AD) Update() error {
	assets.StaticSoundManager.Update()
	switch ad.whichScreen {
	case InSplashScreen:
		if ad.splayScreen.Update() {
			ad.whichScreen = InMenuScreen
		}
	case InMenuScreen:
		stars.Update()
		ad.menuScreen.Update()
		if ad.menuScreen.Exit() {
			ad.splayScreen.up = true
			ad.splayScreen.transparency = 0
			ad.whichScreen = InSplashScreen
		} else if ad.menuScreen.Start() {
			ad.whichScreen = InGameScreen
		}
	case InGameScreen:
		stars.Update()
		ad.gameScreen.Update()
	}

	//
	return nil
}

func (ad *AD) Draw(screen *ebiten.Image) {
	switch ad.whichScreen {
	case InSplashScreen:
		ad.splayScreen.Draw(screen)
	case InMenuScreen:
		stars.Draw(ad.gameScreen.worldImage)
		ad.gameScreen.Draw(screen)
		ad.gameScreen.camera.Render(ad.gameScreen.worldImage, screen)
		ad.menuScreen.Draw(screen)
	case InGameScreen:
		stars.Draw(ad.gameScreen.worldImage)
		ad.gameScreen.Draw(screen)
		ad.gameScreen.camera.Render(ad.gameScreen.worldImage, screen)
		ad.gameScreen.DrawGui(screen)
	}

}

func (ad *AD) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	gameutil.InitPixel()
	log.SetFlags(log.Lshortfile)
	assets.GetManager()
	stars = game.NewStarBackground(600, 20*64, 16*64)
	defer assets.CleanUp()
	ebiten.SetWindowSize(800, 600)
	if err := ebiten.RunGame(NewAttractiveDefense(20, 16)); err != nil {
		panic(err)
	}
}
