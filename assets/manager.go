package assets

import (
	"github.com/elamre/tentsuyu"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	_ "image/png"
	"log"
)

const ()

const (
	ImageFolder = "assets/images/"
	FontFolder  = "assets/fonts/"
)

const (
	AssetsPlatformImage          = "AssetsPlatformImage"
	AssetsPlusSymbol             = "AssetsPlusSymbol"
	AssetsTurretGun_light_bullet = "AssetsTurretGun_light_bullet"
	AssetsTurretGun_light_1      = "AssetsTurretGun_light_1"
	AssetsTurretGun_light_2      = "AssetsTurretGun_light_2"
	AssetsTurretGun_light_3      = "AssetsTurretGun_light_3"
	AssetsTurretGun_light_4      = "AssetsTurretGun_light_4"
	AssetsTurretGun_heavy_bullet = "AssetsTurretGun_heavy_bullet"
	AssetsTurretGun_heavy_1      = "AssetsTurretGun_heavy_1"
	AssetsTurretGun_heavy_2      = "AssetsTurretGun_heavy_2"
	AssetsTurretGun_heavy_3      = "AssetsTurretGun_heavy_3"
	AssetsTurretGun_heavy_4      = "AssetsTurretGun_heavy_4"
	AssetsTurretBase_1           = "AssetsTurretBase_1"
	AssetsTurretBase_2           = "AssetsTurretBase_2"
	AssetsTurretBase_3           = "AssetsTurretBase_3"
	AssetsTurretBase_4           = "AssetsTurretBase_4"
	AssetsMagnet                 = "AssetsMagnet"
	AssetsMagnetBase             = "AssetsMagnetBase"
	AssetsEnemy                  = "AssetsEnemy"

	AssetsGuiEmpty  = "AssetsGuiEmpty"
	AssetsGuiCancel = "AssetsGuiCancel"
	AssetsGuiSell   = "AssetsGuiSell"

	AssetsGuiLightTurret        = "AssetsGuiLightTurret"
	AssetsGuiLightTurretUpgrade = "AssetsGuiLightTurretUpgrade"
	AssetsGuiHeavyTurret        = "AssetsGuiHeavyTurret"
	AssetsGuiHeavyTurretUpgrade = "AssetsGuiHeavyTurretUpgrade"
	AssetsGuiBaseUpgrade        = "AssetsGuiBaseUpgrade"
	AssetsGuiMagnet             = "AssetsGuiMagnet"

	AssetsGuiSelectAnim = "AssetsGuiSelectAnim"

	AssetsGuiTopPart            = "AssetsGuiTopPart"
	AssetsGuiBottomBorder       = "AssetsGuiBottomBorder"
	AssetsGuiMiddle             = "AssetsMiddle"
	AssetsGuiWarningLevelLow    = "AssetsGuiWarningLevelLow"
	AssetsGuiWarningLevelMedium = "AssetsGuiWarningLevelMedium"
	AssetsGuiWarningLevelHigh   = "AssetsGuiWarningLevelHigh"
)

var manager *tentsuyu.AssetsManager
var deferList []func()

type Animation struct {
	Images []*ebiten.Image
}

func GetManager() *tentsuyu.AssetsManager {
	if manager == nil {
		deferList = make([]func(), 0)
		manager = tentsuyu.NewAssetsManager()
		spriteset, _, err := ebitenutil.NewImageFromFile(ImageFolder + "spritesheet_complete.png")
		if err != nil {
			panic(err)
		}

		manager.AssetMap[AssetsMagnetBase] = spriteset.SubImage(image.Rect(9*64, 0, 64+9*64, 64)).(*ebiten.Image)
		manager.AssetMap[AssetsMagnet] = spriteset.SubImage(image.Rect(8*64, 0, 64+8*64, 64)).(*ebiten.Image)
		manager.AssetMap[AssetsPlatformImage] = spriteset.SubImage(image.Rect(6*64, 0, 64+6*64, 64)).(*ebiten.Image)
		manager.AssetMap[AssetsPlusSymbol] = spriteset.SubImage(image.Rect(7*64, 0, 64+7*64, 64)).(*ebiten.Image)
		manager.AssetMap[AssetsTurretGun_light_bullet] = spriteset.SubImage(image.Rect(4*64, 1*64, 64+4*64, 1*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsTurretGun_light_1] = spriteset.SubImage(image.Rect(0*64, 1*64, 64+0*64, 1*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsTurretGun_light_2] = spriteset.SubImage(image.Rect(1*64, 1*64, 64+1*64, 1*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsTurretGun_light_3] = spriteset.SubImage(image.Rect(2*64, 1*64, 64+2*64, 1*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsTurretGun_light_4] = spriteset.SubImage(image.Rect(3*64, 1*64, 64+3*64, 1*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsTurretGun_heavy_bullet] = spriteset.SubImage(image.Rect(4*64, 2*64, 64+4*64, 2*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsTurretGun_heavy_1] = spriteset.SubImage(image.Rect(0*64, 2*64, 64+0*64, 2*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsTurretGun_heavy_2] = spriteset.SubImage(image.Rect(1*64, 2*64, 64+1*64, 2*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsTurretGun_heavy_3] = spriteset.SubImage(image.Rect(2*64, 2*64, 64+2*64, 2*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsTurretGun_heavy_4] = spriteset.SubImage(image.Rect(3*64, 2*64, 64+3*64, 2*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsTurretBase_1] = spriteset.SubImage(image.Rect(0*64, 0*64, 64+0*64, 0*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsTurretBase_2] = spriteset.SubImage(image.Rect(1*64, 0*64, 64+1*64, 0*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsTurretBase_3] = spriteset.SubImage(image.Rect(2*64, 0*64, 64+2*64, 0*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsTurretBase_4] = spriteset.SubImage(image.Rect(3*64, 0*64, 64+3*64, 0*64+64)).(*ebiten.Image)

		manager.AssetMap[AssetsGuiSelectAnim] = make([]*ebiten.Image, 0)
		for i := 6; i < 10; i++ {
			img := spriteset.SubImage(image.Rect(i*64, 7*64, 64+i*64, 7*64+64)).(*ebiten.Image)
			mm := manager.AssetMap[AssetsGuiSelectAnim].([]*ebiten.Image)
			mm = append(mm, img)
			manager.AssetMap[AssetsGuiSelectAnim] = mm
		}

		manager.AssetMap[AssetsGuiTopPart] = spriteset.SubImage(image.Rect(10*64, 6*64, 64+11*64, 7*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsGuiBottomBorder] = spriteset.SubImage(image.Rect(10*64, 11*64, 64+11*64, 11*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsGuiMiddle] = spriteset.SubImage(image.Rect(10*64, 8*64, 64+11*64, 8*64+64)).(*ebiten.Image)

		manager.AssetMap[AssetsGuiWarningLevelLow] = spriteset.SubImage(image.Rect(10*64, 5*64, 10*64+5, 5*64+16)).(*ebiten.Image)
		manager.AssetMap[AssetsGuiWarningLevelMedium] = spriteset.SubImage(image.Rect(10*64+5, 5*64, 10*64+10, 5*64+16)).(*ebiten.Image)
		manager.AssetMap[AssetsGuiWarningLevelHigh] = spriteset.SubImage(image.Rect(10*64+10, 5*64, 10*64+15, 5*64+16)).(*ebiten.Image)

		manager.AssetMap[AssetsGuiLightTurretUpgrade] = spriteset.SubImage(image.Rect(0*64, 7*64, 64+0*64, 7*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsGuiHeavyTurretUpgrade] = spriteset.SubImage(image.Rect(1*64, 7*64, 64+1*64, 7*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsGuiLightTurret] = spriteset.SubImage(image.Rect(0*64, 7*64, 64+0*64, 7*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsGuiHeavyTurret] = spriteset.SubImage(image.Rect(1*64, 7*64, 64+1*64, 7*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsGuiBaseUpgrade] = spriteset.SubImage(image.Rect(2*64, 7*64, 64+2*64, 7*64+64)).(*ebiten.Image)

		manager.AssetMap[AssetsGuiMagnet] = spriteset.SubImage(image.Rect(0*64, 8*64, 64+0*64, 8*64+64)).(*ebiten.Image)

		deferList = append(deferList, func() {

		})
		log.Printf("Manager: %+v", manager)

	}
	return manager
}

func Get[t any](asset string) t {
	return GetManager().AssetMap[asset].(t)
}

func CleanUp() {
	if manager == nil {
		return
	}
	mng := GetManager()
	for _, f := range deferList {
		f()
	}
	_ = mng
}
