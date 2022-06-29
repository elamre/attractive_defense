package enemies

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/elamre/attractive_defense/world"
	"github.com/hajimehoshi/ebiten/v2"
)

var lightHeaviesProjectile = world.ProjectileEffect{
	Speed:  8,
	Damage: 30,
}

var heavyHeaviesProjectile = world.ProjectileEffect{
	Speed:  8,
	Damage: 40,
}

var eliteHeaviesProjectile = world.ProjectileEffect{
	Speed:  10,
	Damage: 60,
}

var heaviesEnemyTurret = EnemyTurretSpecifications{
	reloadSpeed: 50,
	targetRange: 100,
	shoot: func(pixelX, pixelY, targetX, targetY float64, manager *world.ProjectoryManager) {
		manager.AddEnemyProjectile(world.NewHeavyProjectile(pixelX, pixelY, targetX, targetY, &lightHeaviesProjectile, 200))
	},
}

var heavyHeaviesEnemyTurret = EnemyTurretSpecifications{
	reloadSpeed: 75,
	targetRange: 120,
	shoot: func(pixelX, pixelY, targetX, targetY float64, manager *world.ProjectoryManager) {
		manager.AddEnemyProjectile(world.NewHeavyProjectile(pixelX, pixelY, targetX, targetY, &heavyHeaviesProjectile, 200))
	},
}

var eliteHeaviesEnemyTurret = EnemyTurretSpecifications{
	reloadSpeed: 100,
	targetRange: 200,
	shoot: func(pixelX, pixelY, targetX, targetY float64, manager *world.ProjectoryManager) {
		manager.AddEnemyProjectile(world.NewHeavyProjectile(pixelX, pixelY, targetX, targetY, &eliteHeaviesProjectile, 400))
	},
}

var heaviesEnemyHull = EnemyHullSpecifications{
	maxSpeed:  3,
	width:     32,
	height:    32,
	maxHealth: 75,
	maxShield: 75,
	reward:    150,
}

var shieldedHeaviesEnemyHull = EnemyHullSpecifications{
	maxSpeed:  3,
	width:     32,
	height:    32,
	maxHealth: 100,
	maxShield: 50,
	reward:    175,
}

var heavyHeaviesEnemyHull = EnemyHullSpecifications{
	maxSpeed:  2.5,
	width:     32,
	height:    32,
	maxHealth: 100,
	maxShield: 100,
	reward:    200,
}

var eliteHeaviesEnemyHull = EnemyHullSpecifications{
	maxSpeed:  1,
	width:     32,
	height:    32,
	maxHealth: 200,
	maxShield: 200,
	reward:    300,
}

func NewHeaviesEnemy(pixelX, pixelY float64) EnemyInterface {
	return NewBasicEnemy(pixelX, pixelY, &heaviesEnemyTurret, &heaviesEnemyHull)
}

func NewShieldedHeaviesEnemy(pixelX, pixelY float64) EnemyInterface {
	return NewBasicEnemy(pixelX, pixelY, &heaviesEnemyTurret, &shieldedHeaviesEnemyHull)
}

func NewHeavyHeaviesEnemy(pixelX, pixelY float64) EnemyInterface {
	return NewBasicEnemy(pixelX, pixelY, &heavyHeaviesEnemyTurret, &heavyHeaviesEnemyHull)
}

func NewEliteHeaviesEnemy(pixelX, pixelY float64) EnemyInterface {
	return NewBasicEnemy(pixelX, pixelY, &eliteHeaviesEnemyTurret, &eliteHeaviesEnemyHull)
}

func initHeaviesEnemyImages() {
	heaviesEnemyHull.image = assets.Get[*ebiten.Image](assets.AssetsEnemyMediumLight)
	shieldedHeaviesEnemyHull.image = assets.Get[*ebiten.Image](assets.AssetsEnemyMediumShielded)
	heavyHeaviesEnemyHull.image = assets.Get[*ebiten.Image](assets.AssetsEnemyMediumHeavy)
	eliteHeaviesEnemyHull.image = assets.Get[*ebiten.Image](assets.AssetsEnemyMediumElite)
}
