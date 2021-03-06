package gui

import (
	"fmt"
	"github.com/elamre/attractive_defense/assets"
	"github.com/elamre/attractive_defense/buildings"
	"github.com/elamre/attractive_defense/buildings/turrets"
	"github.com/elamre/attractive_defense/game"
	"github.com/elamre/attractive_defense/platforms"
	"github.com/elamre/attractive_defense/world"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type SideGui struct {
	top                *ebiten.Image
	bot                *ebiten.Image
	mid                *ebiten.Image
	noticeLevelImgLow  *ebiten.Image
	noticeLevelImgMed  *ebiten.Image
	noticeLevelImgHigh *ebiten.Image
	Rows               int
	NoticeLevel        int // max 24
	drawX, drawY       int

	tempCounter int
	buttonArray []*Button

	available map[string]bool

	showingBuildings bool

	sellButton       *SellButton
	repairButton     *RepairButton
	buildingSelected bool
	description      string
}

func NewSideGui(pixelX, pixelY int) *SideGui {
	s := &SideGui{
		drawY:              pixelY,
		drawX:              pixelX,
		top:                assets.Get[*ebiten.Image](assets.AssetsGuiTopPart),
		bot:                assets.Get[*ebiten.Image](assets.AssetsGuiBottomBorder),
		mid:                assets.Get[*ebiten.Image](assets.AssetsGuiMiddle),
		noticeLevelImgLow:  assets.Get[*ebiten.Image](assets.AssetsGuiWarningLevelLow),
		noticeLevelImgMed:  assets.Get[*ebiten.Image](assets.AssetsGuiWarningLevelMedium),
		noticeLevelImgHigh: assets.Get[*ebiten.Image](assets.AssetsGuiWarningLevelHigh),
		Rows:               8,
		NoticeLevel:        1,
		available:          InitButtons(),
		showingBuildings:   false,
		sellButton:         NewSellButton(),
		repairButton:       NewRepairButton(),
	}
	s.NoButtonsContext()
	return s
}

func (s *SideGui) InGui(x, y int) bool {
	if x >= s.drawX && x <= s.drawX+128 {
		if y >= s.drawY+19 && y <= (s.drawY+(s.Rows*64)+64+19) {
			return true
		}
	}
	return false
}

func (s *SideGui) NoButtonsContext() {
	s.buttonArray = []*Button{}
}

func (s *SideGui) SetBuildingContext(p *game.Player) {
	s.buttonArray = []*Button{}
	s.buttonArray = append(s.buttonArray, buyMagnet)
	s.buttonArray = append(s.buttonArray, buyResearch)
	s.buttonArray = append(s.buttonArray, buyLightTurret)
	s.buttonArray = append(s.buttonArray, buyHeavyTurret)
	if p.BeamResearch {
		s.buttonArray = append(s.buttonArray, buyBeamTurret)
	}
	if p.RocketResearch {
		s.buttonArray = append(s.buttonArray, buyRocketTurret)
	}
}

func (s *SideGui) SetResearchContext(p *game.Player, researchLab world.GridEntity) {
	lab := researchLab.(*buildings.ResearchLab)
	s.buttonArray = []*Button{}
	if !p.FastRepairBeing {
		s.buttonArray = append(s.buttonArray, NewRepairUpgradeButton(lab))
	}
	if !p.RocketBeingResearch {
		s.buttonArray = append(s.buttonArray, NewResearchRocket(lab))
	}
	if !p.BeamBeingResearch {
		s.buttonArray = append(s.buttonArray, NewResearchBeam(lab))
	}
	if !p.DoubleBeingMoney {
		s.buttonArray = append(s.buttonArray, NewResearchMoney(lab))
	}

}

func (s *SideGui) SetBuildingSelectedContext(p *game.Player, e world.GridEntity) {
	if t, ok := e.(*turrets.Turret); ok {
		s.buttonArray = []*Button{NewUpgradableButton(t.Gun, e), NewUpgradableButton(t.Base, e)}
	} else if _, ok := e.(*buildings.ResearchLab); ok {
		s.SetResearchContext(p, e)
	}
}

func (s *SideGui) Update(p *game.Player, g *world.Grid, camera *Camera) {
	//s.tempCounter++
	//s.NoticeLevel = (s.tempCounter / 10) % 25
	selectedEntity := g.GetGridEntity(g.SelectedGridX, g.SelectedGridY, world.GridLevelStructures)
	s.buildingSelected = selectedEntity != nil
	s.description = ""

	mx, my := ebiten.CursorPosition()
	if s.InGui(mx, my) {
		xPos := 0
		if mx > s.drawX+64 {
			xPos = 1
		}
		if my < 64+19 {
			if s.buildingSelected {
				if xPos == 0 {
					s.description = "Sell selected building"
				} else {
					s.description = "Repair selected building"
				}
			}
		} else {
			yIndex := (my - 19 - 64) / 64
			index := (yIndex * 2) + xPos
			if index < len(s.buttonArray) {
				if int(s.buttonArray[index].cost) > 0 {
					s.description = fmt.Sprintf("%s, cost: %d", s.buttonArray[index].description, int(s.buttonArray[index].cost))
				}
			}
		}
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {

		if s.InGui(mx, my) {
			xPos := 0
			if mx > s.drawX+64 {
				xPos = 1
			}
			if my < 64+19 {
				if s.buildingSelected {
					if xPos == 0 {
						s.sellButton.Selected(p, s, g)
					} else {
						s.repairButton.Selected(p, s, g)
					}
				}
				return
			}
			yIndex := (my - 19 - 64) / 64
			index := (yIndex * 2) + xPos
			if index < len(s.buttonArray) {
				s.buttonArray[index].Selected(p, s, g)
			}
		} else {
			s.NoButtonsContext()
			s.showingBuildings = false
			cMx, cMy := camera.ScreenToWorld(mx, my)
			x, y := g.MouseToGridPos(int(cMx), int(cMy))
			g.SetSelectedPos(x, y)
			if e := g.GetGridEntity(x, y, world.GridLevelGui); e != nil {
				if _, ok := e.(*platforms.PurchasePlatform); ok {
					// This should be much better TODO
					platforms.NewPlatformAt(x, y, g)
					g.SetGrid(x, y, world.GridLevelGui, nil)
				}
			}
			if x != -1 && y != -1 {
				if e := g.GetGridEntity(x, y, world.GridLevelStructures); e != nil {
					s.SetBuildingSelectedContext(p, e)
				} else if g.GetGridEntity(x, y, world.GridLevelPlatform) != nil {
					s.showingBuildings = true
					s.SetBuildingContext(p)
				}
			}
		}
	} else if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		s.NoButtonsContext()
		g.SetSelectedPos(-1, -1)
	}
	//s.showingBuildings = IsBuildable(g)
}

func (s *SideGui) Draw(screen *ebiten.Image) {
	noticeOpts := ebiten.DrawImageOptions{}

	noticeOpts.GeoM.Translate(float64(s.drawX), float64(s.drawY)+19-64)
	screen.DrawImage(s.top, &noticeOpts)
	noticeOpts.GeoM.Translate(0, 128)
	for i := 0; i < s.Rows; i++ {
		screen.DrawImage(s.mid, &noticeOpts)
		noticeOpts.GeoM.Translate(0, 64)
	}

	noticeOpts.GeoM.Reset()
	noticeOpts.GeoM.Translate(float64(s.drawX)+4, float64(s.drawY))
	for i := 0; i < s.NoticeLevel; i++ {
		if i < 8 {
			screen.DrawImage(s.noticeLevelImgLow, &noticeOpts)
		} else if i < 16 {
			screen.DrawImage(s.noticeLevelImgMed, &noticeOpts)
		} else {
			screen.DrawImage(s.noticeLevelImgHigh, &noticeOpts)
		}
		noticeOpts.GeoM.Translate(5, 0)
	}
	noticeOpts.GeoM.Reset()
	noticeOpts.GeoM.Translate(float64(s.drawX)+5, float64(s.drawY)+19+64)
	//if s.showingBuildings {
	//noticeOpts.ColorM.Translate(-0.5, -0.5, -0.5, 0)

	if s.buttonArray != nil {
		for i, ss := range s.buttonArray {
			ss.Draw(screen, &noticeOpts)
			noticeOpts.GeoM.Translate(60, 0)

			if i%2 == 1 {
				noticeOpts.GeoM.Translate(-120, 64)
			}
		}
	}
	//}
	noticeOpts.GeoM.Reset()
	noticeOpts.GeoM.Translate(float64(s.drawX)+1, float64(s.drawY)+20)
	if s.buildingSelected {
		screen.DrawImage(s.bot, &noticeOpts)
		s.sellButton.Draw(screen, &noticeOpts)
		noticeOpts.GeoM.Translate(62, 0)
		s.repairButton.Draw(screen, &noticeOpts)
	}
	ebitenutil.DebugPrintAt(screen, s.description, 12, 600-12-12)
}
