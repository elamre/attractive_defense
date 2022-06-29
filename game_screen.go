package main

import (
	"github.com/elamre/attractive_defense/buildings"
	"github.com/elamre/attractive_defense/buildings/turrets"
	"github.com/elamre/attractive_defense/enemies"
	"github.com/elamre/attractive_defense/game"
	"github.com/elamre/attractive_defense/gui"
	"github.com/elamre/attractive_defense/world"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/math/f64"
)

type GameScreen struct {
	grid        *world.Grid
	player      *game.Player
	gui         *gui.SideGui
	worldImage  *ebiten.Image
	camera      gui.Camera
	textOverlay *TextOverlay

	projectoryManager *world.ProjectoryManager
	enemyManager      *enemies.EnemyManager
	lifeCrystal       *buildings.LifeCrystal
}

func NewGameScreen(screenWidth, screenHeight int, gridWidth, gridHeight int) *GameScreen {
	g := GameScreen{
		worldImage: ebiten.NewImage(gridWidth*64, gridHeight*64),
		camera:     gui.Camera{ViewPort: f64.Vec2{float64(screenWidth), float64(screenHeight)}},
	}
	g.Reset(screenWidth, screenHeight, gridWidth, gridHeight)
	return &g
}

func (g *GameScreen) Reset(screenWidth, screenHeight int, gridWidth, gridHeight int) {
	grid := world.NewGrid(gridWidth, gridHeight, 3)
	projectory := world.NewProjectoryManager()
	enemyManager := enemies.NewEnemyManager()
	player := game.NewPlayer()
	grid.ProjectoryMng = projectory
	sideGui := gui.NewSideGui(screenWidth-128, 0)
	lifeCrystal := buildings.NewLifeCrystal(gridWidth/2, gridHeight/2, &grid)
	grid.SetGrid(gridWidth/2, gridHeight/2, world.GridLevelStructures, lifeCrystal)

	g.grid = &grid
	g.gui = sideGui
	g.player = player
	g.projectoryManager = projectory
	g.enemyManager = enemyManager
	g.lifeCrystal = lifeCrystal
	g.textOverlay = NewTextOverlay()
	g.camera.Position[0] = 5 * 64
	g.camera.Position[1] = 4 * 64
	g.grid.GridChangeCallback = func(x, y, z int, entity world.GridEntity) {
		if z == world.GridLevelStructures {
			if entity != nil {
				if turret, ok := entity.(*turrets.Turret); ok {
					turret.SetClosest(enemyManager.Entities)
					turret.EnemyKilledCb = func() {
						turret.SetClosest(enemyManager.Entities)
					}
				}
			}
		}
	}
}

func (g *GameScreen) Update() {
	if !g.textOverlay.pause && g.lifeCrystal.Alive() {
		g.camera.Update()
		g.grid.UpdateGrid()
		g.player.UpdatePlayer(g.grid)
		g.gui.Update(g.player, g.grid, &g.camera)
		g.projectoryManager.Update(g.grid)
		g.enemyManager.Update(g.grid, g.player, g.projectoryManager)
		x, y := ebiten.CursorPosition()

		if g.enemyManager.ShouldSpawn() && !g.textOverlay.neverStarted {
			if g.textOverlay.Finished() {
				g.enemyManager.Spawn(g.grid, g.gui.NoticeLevel)
				g.textOverlay.Reset()
			} else if !g.textOverlay.started {
				g.textOverlay.StartCountdown()
				g.player.WaveFinished(g.enemyManager.WaveNr(), g.gui.NoticeLevel)
			}
		} else if g.textOverlay.neverStarted {
			if inpututil.IsKeyJustPressed(ebiten.KeySpace) && g.textOverlay.neverStarted {

				g.enemyManager.Spawn(g.grid, g.gui.NoticeLevel)
				g.textOverlay.neverStarted = false
			}
		}

		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
			cMx, cMy := g.camera.ScreenToWorld(x, y)
			g.enemyManager.AddEnemy(enemies.NewEliteHeaviesEnemy(cMx, cMy), g.grid)
		}
		g.gui.InGui(x, y)
		if !g.lifeCrystal.Alive() {
			g.textOverlay.SetGameOver(true, g.player.Score)
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		if g.textOverlay.gameOver {
			g.textOverlay.SetGameOver(false, 0)
			g.Reset(800, 600, 20, 16)
		} else {
			g.textOverlay.TogglePause()
		}
	}
	g.textOverlay.Update()
}

func (g *GameScreen) DrawGui(screen *ebiten.Image) {
	g.player.DrawPlayer(screen)
	g.gui.Draw(screen)
	g.worldImage.Clear()
	g.textOverlay.Draw(screen)
}

func (g *GameScreen) Draw(screen *ebiten.Image) {
	g.grid.DrawGrid(g.worldImage)
	g.enemyManager.Draw(g.worldImage)
	g.projectoryManager.Draw(g.worldImage)
}
