package world

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/colornames"
)

type BuildingInterface interface {
	Targetable
	InflictDamage(damage float64)
	Alive() bool
	GetBuilding() *Building
}

type Building struct {
	PixelX, PixelY        float64
	GridX, GridY          int
	Repairing             bool
	Health                float64
	MaxHealth             float64
	repairCounter         int
	repairTransparency    float64
	cntRepairTransparency float64
}

func (b *Building) RepairPerTick() float64 {
	return b.MaxHealth / 100
}

func (b *Building) RepairTick() bool {
	if b.Health < 1 {
		b.Repairing = false
		return false
	}
	b.repairCounter++
	if b.repairCounter >= 10 {
		b.repairCounter -= 10

		b.Health += b.RepairPerTick()
		if b.Health >= b.MaxHealth {
			b.Health = b.MaxHealth
			b.Repairing = false
		}
		return true
	}
	return false
}

func (b *Building) GetPixelCoordinates() (int, int) {
	return int(b.PixelX), int(b.PixelY)
}

func (b *Building) DrawGui(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, b.PixelX+2, b.PixelY+58, 60, 4, colornames.Red)
	ebitenutil.DrawRect(screen, b.PixelX+2, b.PixelY+58, (b.Health/b.MaxHealth)*60, 4, colornames.Green)
	if b.Repairing {
		//b.repairTransparency
		// This is bad performance
		opt := ebiten.DrawImageOptions{}
		opt.GeoM.Translate(b.PixelX, b.PixelY)
		b.repairTransparency += b.cntRepairTransparency
		if b.repairTransparency <= 0.1 {
			b.cntRepairTransparency = 0.05
			b.repairTransparency = 0.1
		} else if b.repairTransparency >= 0.9 {
			b.repairTransparency = 0.9
			b.cntRepairTransparency = -0.05
		}
		opt.ColorM.Translate(0, 0, 0, -b.repairTransparency)
		img := assets.Get[*ebiten.Image](assets.AssetsGuiRepairWrench)
		screen.DrawImage(img, &opt)
	}

}
