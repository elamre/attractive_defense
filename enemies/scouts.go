package enemies

import (
	"github.com/elamre/attractive_defense/assets"
	"github.com/elamre/attractive_defense/world"
	"github.com/hajimehoshi/ebiten/v2"
)

var lightScoutProjectile = world.ProjectileEffect{
	Speed:  5,
	Damage: 5,
}

var heavyScoutProjectile = world.ProjectileEffect{
	Speed:  4,
	Damage: 15,
}

var eliteScoutProjectile = world.ProjectileEffect{
	Speed:  6,
	Damage: 30,
}

var scoutEnemyTurret = EnemyTurretSpecifications{
	reloadSpeed: 10,
	targetRange: 60,
	shoot: func(pixelX, pixelY, targetX, targetY float64, manager *world.ProjectoryManager) {
		manager.AddEnemyProjectile(world.NewSmallProjectile(pixelX, pixelY, targetX, targetY, &lightScoutProjectile, 100))
	},
}

var heavyScoutEnemyTurret = EnemyTurretSpecifications{
	reloadSpeed: 20,
	targetRange: 80,
	shoot: func(pixelX, pixelY, targetX, targetY float64, manager *world.ProjectoryManager) {
		manager.AddEnemyProjectile(world.NewSmallProjectile(pixelX, pixelY, targetX, targetY, &heavyScoutProjectile, 100))
	},
}

var eliteScoutEnemyTurret = EnemyTurretSpecifications{
	reloadSpeed: 30,
	targetRange: 100,
	shoot: func(pixelX, pixelY, targetX, targetY float64, manager *world.ProjectoryManager) {
		manager.AddEnemyProjectile(world.NewSmallProjectile(pixelX, pixelY, targetX, targetY, &eliteScoutProjectile, 100))
	},
}

var scoutEnemyHull = EnemyHullSpecifications{
	maxSpeed:  3,
	width:     32,
	height:    32,
	maxHealth: 20,
	maxShield: 0,
	reward:    25,
}

var shieldedScoutEnemyHull = EnemyHullSpecifications{
	maxSpeed:  3,
	width:     32,
	height:    32,
	maxHealth: 10,
	maxShield: 40,
	reward:    50,
}

var heavyScoutEnemyHull = EnemyHullSpecifications{
	maxSpeed:  2,
	width:     32,
	height:    32,
	maxHealth: 50,
	maxShield: 20,
	reward:    75,
}

var eliteScoutEnemyHull = EnemyHullSpecifications{
	maxSpeed:  2.5,
	width:     32,
	height:    32,
	maxHealth: 100,
	maxShield: 50,
	reward:    100,
}

func NewScoutEnemy(pixelX, pixelY float64) EnemyInterface {
	return NewBasicEnemy(pixelX, pixelY, &scoutEnemyTurret, &scoutEnemyHull)
}

func NewShieldedScoutEnemy(pixelX, pixelY float64) EnemyInterface {
	return NewBasicEnemy(pixelX, pixelY, &scoutEnemyTurret, &shieldedScoutEnemyHull)
}

func NewHeavyScoutEnemy(pixelX, pixelY float64) EnemyInterface {
	return NewBasicEnemy(pixelX, pixelY, &heavyScoutEnemyTurret, &heavyScoutEnemyHull)
}

func NewEliteScoutEnemy(pixelX, pixelY float64) EnemyInterface {
	return NewBasicEnemy(pixelX, pixelY, &eliteScoutEnemyTurret, &eliteScoutEnemyHull)
}

func initScoutEnemyImages() {
	scoutEnemyHull.image = assets.Get[*ebiten.Image](assets.AssetsEnemyScoutLight)
	shieldedScoutEnemyHull.image = assets.Get[*ebiten.Image](assets.AssetsEnemyScoutShielded)
	heavyScoutEnemyHull.image = assets.Get[*ebiten.Image](assets.AssetsEnemyScoutHeavy)
	eliteScoutEnemyHull.image = assets.Get[*ebiten.Image](assets.AssetsEnemyScoutElite)
}
