package main

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/elamre/attractive_defense/buildings/turrets"
	"github.com/elamre/attractive_defense/enemies"
	"github.com/elamre/attractive_defense/game"
	"github.com/elamre/attractive_defense/gui"
	"github.com/elamre/attractive_defense/world"
	"github.com/elamre/gameutil"
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

	projectoryManager *world.ProjectoryManager
	enemyManager      *enemies.EnemyManager
}

var stars game.StarBackground

func NewAttractiveDefense(width, height int) *AD {

	grid := world.NewGrid(width, height, 3)
	projectory := world.NewProjectoryManager()
	enemyManager := enemies.NewEnemyManager()
	player := game.NewPlayer()
	sideGui := gui.NewSideGui(800-128, 0)
	camera := gui.Camera{ViewPort: f64.Vec2{800, 600}}

	grid.GridChangeCallback = func(x, y, z int, entity world.GridEntity) {
		if z == world.GridLevelStructures {
			if entity != nil {
				if turret, ok := entity.(*turrets.Turret); ok {
					for i := range enemyManager.Entities {
						turret.CheckAndSetTarget(*enemyManager.Entities[i])
					}
				}
			}
		}
	}

	ad := AD{
		g:                 &grid,
		p:                 player,
		gui2:              sideGui,
		world:             ebiten.NewImage(width*64, height*64),
		camera:            camera,
		projectoryManager: projectory,
		enemyManager:      enemyManager,
	}

	return &ad
}

func (ad *AD) Update() error {
	stars.Update()
	ad.camera.Update()
	ad.g.UpdateGrid()
	ad.p.UpdatePlayer(ad.g)
	ad.gui2.Update(ad.p, ad.g, &ad.camera)
	ad.projectoryManager.Update(ad.g)
	ad.enemyManager.Update(ad.g, ad.p, ad.projectoryManager)
	x, y := ebiten.CursorPosition()

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		if ad.enemyManager.ShouldSpawn() {
			ad.enemyManager.Spawn(ad.g, ad.gui2.NoticeLevel)
		}
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		cMx, cMy := ad.camera.ScreenToWorld(x, y)
		ad.enemyManager.AddEnemy(enemies.NewScoutEnemy(cMx, cMy), ad.g)
	}
	ad.gui2.InGui(x, y)
	return nil
}

func (ad *AD) Draw(screen *ebiten.Image) {
	stars.Draw(ad.world)
	ad.g.DrawGrid(ad.world)
	ad.enemyManager.Draw(ad.world)
	ad.projectoryManager.Draw(ad.world)

	ad.camera.Render(ad.world, screen)
	ad.p.DrawPlayer(screen)
	ad.gui2.Draw(screen)
	ad.world.Clear()
}

func (ad *AD) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	gameutil.InitPixel()
	log.SetFlags(log.Lshortfile)
	assets.GetManager()
	stars = game.NewStarBackground(300, 20*64, 16*64)
	defer assets.CleanUp()
	ebiten.SetWindowSize(800, 600)
	if err := ebiten.RunGame(NewAttractiveDefense(20, 16)); err != nil {
		panic(err)
	}
}
