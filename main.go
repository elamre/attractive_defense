package main

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/elamre/attractive_defense/buildings"
	"github.com/elamre/attractive_defense/game"
	"github.com/elamre/attractive_defense/gui"
	"github.com/elamre/attractive_defense/projectory"
	"github.com/elamre/attractive_defense/world"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/math/f64"
	"log"
)

type AD struct {
	g      *world.Grid
	p      *game.Player
	gui2   *gui.SideGui
	world  *ebiten.Image
	camera gui.Camera

	projectoryManager *projectory.ProjectoryManager
}

func (ad *AD) Update() error {
	ad.camera.Update()
	ad.g.UpdateGrid()
	ad.p.UpdatePlayer(ad.g)
	ad.gui2.Update(ad.p, ad.g, &ad.camera)
	ad.projectoryManager.Update(ad.g)
	x, y := ebiten.CursorPosition()
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		cMx, cMy := ad.camera.ScreenToWorld(x, y)
		ad.projectoryManager.AddPlayerProjectile(projectory.NewBasicProjectile(cMx, cMy, 5*64+32, 5*64+32))
	}
	ad.gui2.InGui(x, y)
	return nil
}

func (ad *AD) Draw(screen *ebiten.Image) {
	ad.g.DrawGrid(ad.world)
	ad.camera.Render(ad.world, screen)
	ad.p.DrawPlayer(screen)
	ad.projectoryManager.Draw(screen)
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
	g.SetGrid(5, 5, world.GridLevelStructures, buildings.NewLifeCrystal(5, 5, &g))
	ad := AD{&g, game.NewPlayer(), gui.NewSideGui(800-128, 0), ebiten.NewImage(20*64, 16*64), gui.Camera{ViewPort: f64.Vec2{800, 600}}, projectory.NewProjectoryManager()}

	defer assets.CleanUp()
	ebiten.SetWindowSize(800, 600)
	if err := ebiten.RunGame(&ad); err != nil {
		panic(err)
	}
}
