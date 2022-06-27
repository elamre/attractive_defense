package world

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/elamre/go_helpers/pkg/slice_helpers"
	"github.com/elamre/tentsuyu/tentsuyutils"
	"github.com/hajimehoshi/ebiten/v2"
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

type GridPos struct {
	PosX, PosY, PosZ int
}

func (g *GridPos) compare(other interface{}) int {
	return 0
}

type TriggerAble interface {
	Trigger(x, y int, other interface{})
}

type Grid struct {
	buildings                    *assets.EntityManager[*GridEntity]
	Width, Height, levels        int
	fields                       []GridEntity
	triggerFields                [][]*TriggerAble
	entityToField                map[GridEntity]GridPos
	entityToEntityPtr            map[GridEntity]*GridEntity
	magnetism                    []int
	triggered                    []int
	SelectedGridX, SelectedGridY int

	selectionImage      []*ebiten.Image
	selectionImageCount int

	ProjectoryMng *ProjectoryManager

	GridChangeCallback func(x, y, z int, entity GridEntity)
}

func NewGrid(width, height, levels int) Grid {
	return Grid{
		Width:               width,
		Height:              height,
		levels:              levels,
		buildings:           assets.NewEntityManager[*GridEntity](),
		fields:              make([]GridEntity, width*height*levels),
		triggerFields:       make([][]*TriggerAble, width*height),
		entityToField:       make(map[GridEntity]GridPos),
		entityToEntityPtr:   make(map[GridEntity]*GridEntity),
		magnetism:           make([]int, width*height),
		triggered:           make([]int, width*height),
		SelectedGridX:       -1,
		SelectedGridY:       -1,
		selectionImage:      assets.Get[[]*ebiten.Image](assets.AssetsGuiSelectAnim),
		selectionImageCount: 0,
	}
}

func (g *Grid) AddTriggerFunc(x, y int, trigger TriggerAble) {
	idx := (g.Width * y) + x
	if g.triggerFields[idx] == nil {
		g.triggerFields[idx] = make([]*TriggerAble, 0)
	}
	g.triggerFields[idx] = append(g.triggerFields[idx], &trigger)
}

func (g *Grid) RemoveTrigger(x, y int, trigger TriggerAble) {
	idx := (g.Width * y) + x
	g.triggerFields[idx] = slice_helpers.RemoveFromList[*TriggerAble](&trigger, g.triggerFields[idx])
}

func (g *Grid) TestTrigger(x, y int, ent interface{}) bool {
	idx := (g.Width * y) + x
	if idx >= len(g.triggerFields) {
		return false
	}
	if ar := g.triggerFields[idx]; ar != nil {
		for i := range ar {
			g.triggered[idx]++

			(*ar[i]).Trigger(x, y, ent)
		}
		return true
	}
	return false
}

func (g *Grid) AddMagnetism(posX, posY int) {
	g.magnetism[posY*g.Width+posX]++
}

func (g *Grid) RemoveMagnetism(posX, posY int) {
	g.magnetism[posY*g.Width+posX]--

	if g.magnetism[posY*g.Width+posX] == 0 {
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
	if mouseX < 0 || mouseX > g.Width-1 || mouseY < 0 || mouseY > g.Height-1 {
		return -1, -1
	}
	return mouseX, mouseY
}

func (g *Grid) OutOfBounds(x, y int) bool {
	return x < 0 || x > g.Width-1 || y < 0 || y > g.Height-1
}

func (g *Grid) ClosestBuilding(gridX, gridY int) GridEntity {
	var closest GridEntity
	distance := float64(1000000)

	for e := range g.buildings.Entities {
		ee := *g.buildings.Entities[e]
		gg := g.entityToField[ee]
		dst := tentsuyutils.Distance(float64(gridX), float64(gridY), float64(gg.PosX), float64(gg.PosY))
		if dst < distance {
			distance = dst
			closest = ee
		}
	}
	return closest
}

func (g *Grid) SetGrid(x, y, z int, entity GridEntity) {
	idx := ((g.Height * g.Width * z) + g.Width*y) + x
	if g.GridChangeCallback != nil {
		g.GridChangeCallback(x, y, z, entity)
	}
	if z == GridLevelStructures {
		if entity == nil {
			ee := g.fields[idx]
			g.buildings.SetForRemoval(g.entityToEntityPtr[g.fields[idx]])
			g.buildings.CleanDeadEntities()
			delete(g.entityToField, ee)
			delete(g.entityToEntityPtr, g.fields[idx])

		} else {
			g.entityToField[entity] = GridPos{
				PosX: x,
				PosY: y,
				PosZ: z,
			}
			g.entityToEntityPtr[entity] = &entity // TODO this is super shit
			g.buildings.AddEntity(&entity)

		}
	}
	g.fields[idx] = entity
}

func (g *Grid) GetGridEntity(x, y, z int) GridEntity {
	idx := (g.Height*g.Width*z + g.Width*y) + x
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
		for y := 0; y < g.Height; y++ {
			for x := 0; x < g.Width; x++ {
				idx := (g.Height*g.Width*z + g.Width*y) + x
				if e := g.fields[idx]; e != nil {
					e.Update(g)
				}
			}
		}
	}
}

func (g *Grid) DrawGrid(screen *ebiten.Image) {
	for z := 0; z < g.levels; z++ {
		for y := 0; y < g.Height; y++ {
			for x := 0; x < g.Width; x++ {
				idx := (g.Height*g.Width*z + g.Width*y) + x
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
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			//ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%d", g.magnetism[y*g.Width+x]), x*64, y*64)
			//ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%d", len(g.triggerFields[y*g.Width+x])), x*64, y*64+12)
			//ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%d", (g.triggered[y*g.Width+x])), x*64+24, y*64+12)
		}
	}
}
