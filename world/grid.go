package world

import (
	"fmt"
	"github.com/elamre/attractive_defense/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	GridLevelPlatform   = 0
	GridLevelStructures = 1
	GridLevelGui        = 2
	GridLevelsTotal     = 3
	GridWidth           = 64
	GridHeight          = 48
)

type GridEntity interface {
	Update(g *Grid)
	SetForDeletion(g *Grid)
	Draw(*ebiten.Image)
}

type Grid struct {
	width, height, levels        int
	fields                       []GridEntity
	magnetism                    []int
	SelectedGridX, SelectedGridY int

	selectionImage      []*ebiten.Image
	selectionImageCount int
}

func NewGrid(width, height, levels int) Grid {
	return Grid{
		width:               width,
		height:              height,
		levels:              levels,
		fields:              make([]GridEntity, width*height*levels),
		magnetism:           make([]int, width*height),
		SelectedGridX:       -1,
		SelectedGridY:       -1,
		selectionImage:      assets.Get[[]*ebiten.Image](assets.AssetsGuiSelectAnim),
		selectionImageCount: 0,
	}
}

func (g *Grid) AddMagnetism(posX, posY int) {
	g.magnetism[posY*g.width+posX]++
}

func (g *Grid) RemoveMagnetism(posX, posY int) {
	g.magnetism[posY*g.width+posX]--

	if g.magnetism[posY*g.width+posX] == 0 {
		if e := g.GetGridEntity(posX, posY, GridLevelPlatform); e != nil {
			e.SetForDeletion(g)
		}
		if e := g.GetGridEntity(posX, posY, GridLevelStructures); e != nil {
			e.SetForDeletion(g)
		}
		if e := g.GetGridEntity(posX, posY, GridLevelGui); e != nil {
			g.SetGrid(posX, posY, GridLevelGui, nil)
		}
	}
}

func (g *Grid) SetSelectedPos(posX, posY int) {
	g.SelectedGridY = posY
	g.SelectedGridX = posX
}

func (g *Grid) MouseToGridPos(mouseX, mouseY int) (int, int) {
	mouseX /= 64
	mouseY /= 64
	if mouseX < 0 || mouseX > g.width-1 || mouseY < 0 || mouseY > g.height-1 {
		return -1, -1
	}
	return mouseX, mouseY
}

func (g *Grid) OutOfBounds(x, y int) bool {
	return x < 0 || x > g.width-1 || y < 0 || y > g.height-1
}

func (g *Grid) SetGrid(x, y, z int, entity GridEntity) {
	idx := ((g.height * g.width * z) + g.width*y) + x
	g.fields[idx] = entity
}

func (g *Grid) GetGridEntity(x, y, z int) GridEntity {
	idx := (g.height*g.width*z + g.width*y) + x
	return g.fields[idx]
}

func (g *Grid) IsEmpty(x, y int) bool {
	for l := 0; l < GridLevelsTotal; l++ {
		if g.GetGridEntity(x, y, l) != nil {
			return false
		}
	}
	return true
}

func (g *Grid) UpdateGrid() {
	for z := 1; z < g.levels; z++ {
		for y := 0; y < g.height; y++ {
			for x := 0; x < g.width; x++ {
				idx := (g.height*g.width*z + g.width*y) + x
				if e := g.fields[idx]; e != nil {
					e.Update(g)
				}
			}
		}
	}
}

func (g *Grid) DrawGrid(screen *ebiten.Image) {
	for z := 0; z < g.levels; z++ {
		for y := 0; y < g.height; y++ {
			for x := 0; x < g.width; x++ {
				idx := (g.height*g.width*z + g.width*y) + x
				if e := g.fields[idx]; e != nil {
					e.Draw(screen)
				}
			}
		}
	}
	if g.SelectedGridY != -1 && g.SelectedGridX != -1 {
		opt := ebiten.DrawImageOptions{}
		g.selectionImageCount++
		if g.selectionImageCount >= 40 {
			g.selectionImageCount = 0
		}
		opt.GeoM.Translate(float64(g.SelectedGridX*64), float64(g.SelectedGridY*64))
		screen.DrawImage(g.selectionImage[g.selectionImageCount/10], &opt)
	}
	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%d", g.magnetism[y*g.width+x]), x*64, y*64)
		}
	}
}
