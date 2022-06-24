package buildings

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/elamre/attractive_defense/world"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

type BasicTurret struct {
	turretBase           *ebiten.Image
	turretGun            *ebiten.Image
	x, y, pixelX, pixelY int
	turretBaseOpt        *ebiten.DrawImageOptions
	turretGunOpt         *ebiten.DrawImageOptions
	destroy              bool
}

func (b *BasicTurret) SetForDeletion(g *world.Grid) {
	b.destroy = true
}

func (b *BasicTurret) Update(g *world.Grid) {
	if b.destroy {
		g.SetGrid(b.x, b.y, world.GridLevelStructures, nil)
	}
}

func (b *BasicTurret) Draw(image *ebiten.Image) {
	image.DrawImage(b.turretBase, b.turretBaseOpt)
	mouseX, mouseY := ebiten.CursorPosition()

	op := &ebiten.DrawImageOptions{}
	mouseXFloat := float64(mouseX) - float64(b.pixelX+32)
	mouseYFloat := float64(mouseY) - float64(b.pixelY+32)

	angle := math.Atan2(mouseYFloat, mouseXFloat)

	op.GeoM.Translate(-64/2, -64/2)
	op.GeoM.Rotate(angle)
	op.GeoM.Translate(64/2, 64/2)
	// Place at correct position
	op.GeoM.Translate(float64(b.pixelX), float64(b.pixelY))
	image.DrawImage(b.turretGun, op)
}

func NewBasicTurret(x, y int) *BasicTurret {
	b := BasicTurret{x: x, y: y, pixelY: y * 64, pixelX: x * 64,
		turretBase: assets.Get[*ebiten.Image](assets.AssetsTurretBase_1),
		turretGun:  assets.Get[*ebiten.Image](assets.AssetsTurretGun_light_1),
	}
	b.turretBaseOpt = &ebiten.DrawImageOptions{}
	b.turretGunOpt = &ebiten.DrawImageOptions{}
	b.turretGunOpt.GeoM.Translate(float64(b.pixelX), float64(b.pixelY))
	b.turretBaseOpt.GeoM.Translate(float64(b.pixelX), float64(b.pixelY))
	return &b
}
