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
	AssetsTurretGun_beam_1       = "AssetsTurretGun_beam_1"
	AssetsTurretGun_beam_2       = "AssetsTurretGun_beam_2"
	AssetsTurretGun_beam_3       = "AssetsTurretGun_beam_3"
	AssetsTurretGun_beam_4       = "AssetsTurretGun_beam_4"
	AssetsTurretGun_rocket_1     = "AssetsTurretGun_rocket_1"
	AssetsTurretGun_rocket_2     = "AssetsTurretGun_rocket_2"
	AssetsTurretGun_rocket_3     = "AssetsTurretGun_rocket_3"
	AssetsTurretGun_rocket_4     = "AssetsTurretGun_rocket_4"
	AssetsTurretBase_1           = "AssetsTurretBase_1"
	AssetsTurretBase_2           = "AssetsTurretBase_2"
	AssetsTurretBase_3           = "AssetsTurretBase_3"
	AssetsTurretBase_4           = "AssetsTurretBase_4"
	AssetsMagnet                 = "AssetsMagnet"
	AssetsMagnetBase             = "AssetsMagnetBase"

	AssetsGuiEmpty  = "AssetsGuiEmpty"
	AssetsGuiCancel = "AssetsGuiCancel"

	AssetsGuiSell         = "AssetsGuiSell"
	AssetsGuiRepair       = "AssetsGuiRepair"
	AssetsGuiRepairWrench = "AssetsGuiRepairWrench"

	AssetsGuiLightTurret         = "AssetsGuiLightTurret"
	AssetsGuiLightTurretUpgrade  = "AssetsGuiLightTurretUpgrade"
	AssetsGuiHeavyTurret         = "AssetsGuiHeavyTurret"
	AssetsGuiHeavyTurretUpgrade  = "AssetsGuiHeavyTurretUpgrade"
	AssetsGuiBeamTurret          = "AssetsGuiBeamTurret"
	AssetsGuiBeamTurretUpgrade   = "AssetsGuiBeamTurretUpgrade"
	AssetsGuiRocketTurret        = "AssetsGuiRocketTurret"
	AssetsGuiRocketTurretUpgrade = "AssetsGuiRocketTurretUpgrade"
	AssetsGuiBaseUpgrade         = "AssetsGuiBaseUpgrade"
	AssetsGuiMagnet              = "AssetsGuiMagnet"

	AssetsGuiSelectAnim = "AssetsGuiSelectAnim"
	AssetsBuildAnim     = "AssetsBuildAnim"
	AssetsUpgradeAnim   = "AssetsUpgradeAnim"

	AssetsGuiTopPart            = "AssetsGuiTopPart"
	AssetsGuiBottomBorder       = "AssetsGuiBottomBorder"
	AssetsGuiMiddle             = "AssetsMiddle"
	AssetsGuiWarningLevelLow    = "AssetsGuiWarningLevelLow"
	AssetsGuiWarningLevelMedium = "AssetsGuiWarningLevelMedium"
	AssetsGuiWarningLevelHigh   = "AssetsGuiWarningLevelHigh"

	AssetsPlayerCrystalAnim = "AssetsPlayerCrystal"

	AssetsEnemy              = "AssetsEnemy"
	AssetsEnemyScoutLight    = "AssetsEnemyScoutLight"
	AssetsEnemyScoutHeavy    = "AssetsEnemyScoutHeavy"
	AssetsEnemyScoutShielded = "AssetsEnemyScoutShielded"
	AssetsEnemyScoutElite    = "AssetsEnemyScoutElite"

	AssetsEnemyMediumLight    = "AssetsEnemyMediumLight"
	AssetsEnemyMediumHeavy    = "AssetsEnemyMediumHeavy"
	AssetsEnemyMediumShielded = "AssetsEnemyMediumShielded"
	AssetsEnemyMediumElite    = "AssetsEnemyMediumElite"
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
		manager.AssetMap[AssetsTurretGun_light_bullet] = spriteset.SubImage(image.Rect(4*64+27, 1*64+29, 4*64+35, 1*64+34)).(*ebiten.Image)
		manager.AssetMap[AssetsTurretGun_light_1] = spriteset.SubImage(image.Rect(0*64, 1*64, 64+0*64, 1*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsTurretGun_light_2] = spriteset.SubImage(image.Rect(1*64, 1*64, 64+1*64, 1*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsTurretGun_light_3] = spriteset.SubImage(image.Rect(2*64, 1*64, 64+2*64, 1*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsTurretGun_light_4] = spriteset.SubImage(image.Rect(3*64, 1*64, 64+3*64, 1*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsTurretGun_heavy_bullet] = spriteset.SubImage(image.Rect(4*64, 2*64, 64+4*64, 2*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsTurretGun_heavy_1] = spriteset.SubImage(image.Rect(0*64, 2*64, 64+0*64, 2*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsTurretGun_heavy_2] = spriteset.SubImage(image.Rect(1*64, 2*64, 64+1*64, 2*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsTurretGun_heavy_3] = spriteset.SubImage(image.Rect(2*64, 2*64, 64+2*64, 2*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsTurretGun_heavy_4] = spriteset.SubImage(image.Rect(3*64, 2*64, 64+3*64, 2*64+64)).(*ebiten.Image)

		manager.AssetMap[AssetsTurretGun_beam_1] = spriteset.SubImage(image.Rect(5*64, 1*64, 64+5*64, 1*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsTurretGun_beam_2] = spriteset.SubImage(image.Rect(6*64, 1*64, 64+6*64, 1*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsTurretGun_beam_3] = spriteset.SubImage(image.Rect(7*64, 1*64, 64+7*64, 1*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsTurretGun_beam_4] = spriteset.SubImage(image.Rect(8*64, 1*64, 64+8*64, 1*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsTurretGun_rocket_1] = spriteset.SubImage(image.Rect(5*64, 2*64, 64+5*64, 2*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsTurretGun_rocket_2] = spriteset.SubImage(image.Rect(6*64, 2*64, 64+6*64, 2*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsTurretGun_rocket_3] = spriteset.SubImage(image.Rect(7*64, 2*64, 64+7*64, 2*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsTurretGun_rocket_4] = spriteset.SubImage(image.Rect(8*64, 2*64, 64+8*64, 2*64+64)).(*ebiten.Image)

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

		manager.AssetMap[AssetsPlayerCrystalAnim] = make([]*ebiten.Image, 0)
		for i := 0; i < 8; i++ {
			img := spriteset.SubImage(image.Rect(i*64, 11*64, 64+i*64, 11*64+64)).(*ebiten.Image)
			mm := manager.AssetMap[AssetsPlayerCrystalAnim].([]*ebiten.Image)
			mm = append(mm, img)
			manager.AssetMap[AssetsPlayerCrystalAnim] = mm
		}

		manager.AssetMap[AssetsBuildAnim] = make([]*ebiten.Image, 0)
		for i := 0; i < 10; i++ {
			img := spriteset.SubImage(image.Rect(i*64, 12*64, 64+i*64, 12*64+64)).(*ebiten.Image)
			mm := manager.AssetMap[AssetsBuildAnim].([]*ebiten.Image)
			mm = append(mm, img)
			manager.AssetMap[AssetsBuildAnim] = mm
		}
		manager.AssetMap[AssetsUpgradeAnim] = make([]*ebiten.Image, 0)
		for i := 0; i < 12; i++ {
			img := spriteset.SubImage(image.Rect(i*64, 13*64, 64+i*64, 13*64+64)).(*ebiten.Image)
			mm := manager.AssetMap[AssetsUpgradeAnim].([]*ebiten.Image)
			mm = append(mm, img)
			manager.AssetMap[AssetsUpgradeAnim] = mm
		}

		manager.AssetMap[AssetsGuiSell] = spriteset.SubImage(image.Rect(6*64, 6*64, 64+6*64, 6*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsGuiRepair] = spriteset.SubImage(image.Rect(7*64, 6*64, 64+7*64, 6*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsGuiRepairWrench] = spriteset.SubImage(image.Rect(4*64, 5*64, 64+4*64, 5*64+64)).(*ebiten.Image)

		manager.AssetMap[AssetsGuiTopPart] = spriteset.SubImage(image.Rect(10*64, 6*64, 64+11*64, 7*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsGuiBottomBorder] = spriteset.SubImage(image.Rect(10*64, 11*64, 64+11*64, 11*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsGuiMiddle] = spriteset.SubImage(image.Rect(10*64, 8*64, 64+11*64, 8*64+64)).(*ebiten.Image)

		manager.AssetMap[AssetsGuiWarningLevelLow] = spriteset.SubImage(image.Rect(10*64, 5*64, 10*64+5, 5*64+16)).(*ebiten.Image)
		manager.AssetMap[AssetsGuiWarningLevelMedium] = spriteset.SubImage(image.Rect(10*64+5, 5*64, 10*64+10, 5*64+16)).(*ebiten.Image)
		manager.AssetMap[AssetsGuiWarningLevelHigh] = spriteset.SubImage(image.Rect(10*64+10, 5*64, 10*64+15, 5*64+16)).(*ebiten.Image)

		manager.AssetMap[AssetsGuiLightTurretUpgrade] = spriteset.SubImage(image.Rect(0*64, 6*64, 59+0*64, 6*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsGuiHeavyTurretUpgrade] = spriteset.SubImage(image.Rect(1*64, 6*64, 59+1*64, 6*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsGuiBaseUpgrade] = spriteset.SubImage(image.Rect(2*64, 6*64, 59+2*64, 6*64+64)).(*ebiten.Image)

		manager.AssetMap[AssetsGuiLightTurret] = spriteset.SubImage(image.Rect(0*64, 7*64, 59+0*64, 7*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsGuiHeavyTurret] = spriteset.SubImage(image.Rect(1*64, 7*64, 59+1*64, 7*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsGuiBeamTurret] = spriteset.SubImage(image.Rect(2*64, 7*64, 59+2*64, 7*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsGuiRocketTurret] = spriteset.SubImage(image.Rect(3*64, 7*64, 59+3*64, 7*64+64)).(*ebiten.Image)

		manager.AssetMap[AssetsGuiMagnet] = spriteset.SubImage(image.Rect(0*64, 8*64, 59+0*64, 8*64+64)).(*ebiten.Image)

		manager.AssetMap[AssetsEnemyScoutLight] = spriteset.SubImage(image.Rect(0*64, 3*64, 64+0*64, 3*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsEnemyScoutHeavy] = spriteset.SubImage(image.Rect(1*64, 3*64, 64+1*64, 3*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsEnemyScoutShielded] = spriteset.SubImage(image.Rect(2*64, 3*64, 64+2*64, 3*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsEnemyScoutElite] = spriteset.SubImage(image.Rect(3*64, 3*64, 64+3*64, 3*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsEnemyMediumLight] = spriteset.SubImage(image.Rect(4*64, 3*64, 64+4*64, 3*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsEnemyMediumHeavy] = spriteset.SubImage(image.Rect(5*64, 3*64, 64+5*64, 3*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsEnemyMediumShielded] = spriteset.SubImage(image.Rect(6*64, 3*64, 64+6*64, 3*64+64)).(*ebiten.Image)
		manager.AssetMap[AssetsEnemyMediumElite] = spriteset.SubImage(image.Rect(7*64, 3*64, 64+7*64, 3*64+64)).(*ebiten.Image)

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
