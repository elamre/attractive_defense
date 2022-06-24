package main

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/elamre/attractive_defense/buildings"
	"github.com/elamre/attractive_defense/buildings/turrets"
	"github.com/elamre/attractive_defense/game"
	"github.com/elamre/attractive_defense/gui"
	"github.com/elamre/attractive_defense/platforms"
	"github.com/elamre/attractive_defense/world"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/math/f64"
	"log"
)

type AD struct {
	g      *world.Grid
	p      *game.Player
	gui2   *gui.SideGui
	world  *ebiten.Image
	camera gui.Camera
}

func (ad *AD) Update() error {
	ad.camera.Update()
	ad.g.UpdateGrid()
	ad.p.UpdatePlayer(ad.g)
	ad.gui2.Update(ad.p, ad.g, &ad.camera)
	x, y := ebiten.CursorPosition()
	ad.gui2.InGui(x, y)
	return nil
}

func (ad *AD) Draw(screen *ebiten.Image) {
	ad.g.DrawGrid(ad.world)
	ad.camera.Render(ad.world, screen)
	ad.p.DrawPlayer(screen)
	ad.gui2.Draw(screen)
	ad.world.Clear()
}

func (ad *AD) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	log.SetFlags(log.Lshortfile)
	assets.GetManager()
	g := world.NewGrid(20, 16, 3)
	g.SetGrid(4, 4, world.GridLevelStructures, buildings.NewBasicMagnet(4, 4, &g))
	platforms.NewPlatformAt(4, 4, &g)

	ad := AD{&g, game.NewPlayer(), gui.NewSideGui(800-128, 0), ebiten.NewImage(20*64, 16*64), gui.Camera{ViewPort: f64.Vec2{800, 600}}}
	g.SetGrid(5, 4, world.GridLevelPlatform, platforms.NewPlatformAt(5, 4, &g))
	g.SetGrid(5, 4, world.GridLevelStructures, turrets.NewLightTurret(5, 4))
	defer assets.CleanUp()
	ebiten.SetWindowSize(800, 600)
	if err := ebiten.RunGame(&ad); err != nil {
		panic(err)
	}
}
