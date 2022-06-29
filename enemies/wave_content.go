package enemies

type Wave struct {
	content []WaveContent
}

type WaveContent struct {
	amount int
	spawn  func(pixelX, pixelY float64) EnemyInterface
}

var SoutWaves = []Wave{
	{
		[]WaveContent{
			{
				6,
				func(pixelX, pixelY float64) EnemyInterface {
					return NewScoutEnemy(pixelX, pixelY)
				},
			},
		},
	},
	{
		[]WaveContent{
			{
				4,
				func(pixelX, pixelY float64) EnemyInterface {
					return NewScoutEnemy(pixelX, pixelY)
				},
			},
			{
				2,
				func(pixelX, pixelY float64) EnemyInterface {
					return NewShieldedScoutEnemy(pixelX, pixelY)
				},
			},
		},
	},
	{
		[]WaveContent{
			{
				2,
				func(pixelX, pixelY float64) EnemyInterface {
					return NewShieldedScoutEnemy(pixelX, pixelY)
				},
			},
			{
				2,
				func(pixelX, pixelY float64) EnemyInterface {
					return NewHeavyScoutEnemy(pixelX, pixelY)
				},
			},
		},
	},
	{
		[]WaveContent{
			{
				12,
				func(pixelX, pixelY float64) EnemyInterface {
					return NewScoutEnemy(pixelX, pixelY)
				},
			},
		},
	},
	{
		[]WaveContent{
			{
				4,
				func(pixelX, pixelY float64) EnemyInterface {
					return NewEliteScoutEnemy(pixelX, pixelY)
				},
			},
		},
	},
	{
		[]WaveContent{
			{
				18,
				func(pixelX, pixelY float64) EnemyInterface {
					return NewScoutEnemy(pixelX, pixelY)
				},
			},
		},
	},
	{
		[]WaveContent{
			{
				4,
				func(pixelX, pixelY float64) EnemyInterface {
					return NewEliteScoutEnemy(pixelX, pixelY)
				},
			},
			{
				4,
				func(pixelX, pixelY float64) EnemyInterface {
					return NewScoutEnemy(pixelX, pixelY)
				},
			},
		},
	},
	{
		[]WaveContent{
			{
				4,
				func(pixelX, pixelY float64) EnemyInterface {
					return NewEliteScoutEnemy(pixelX, pixelY)
				},
			},
			{
				4,
				func(pixelX, pixelY float64) EnemyInterface {
					return NewScoutEnemy(pixelX, pixelY)
				},
			},
			{
				4,
				func(pixelX, pixelY float64) EnemyInterface {
					return NewShieldedScoutEnemy(pixelX, pixelY)
				},
			},
		},
	},

	{
		[]WaveContent{
			{
				4,
				func(pixelX, pixelY float64) EnemyInterface {
					return NewEliteScoutEnemy(pixelX, pixelY)
				},
			},
			{
				4,
				func(pixelX, pixelY float64) EnemyInterface {
					return NewScoutEnemy(pixelX, pixelY)
				},
			},
			{
				4,
				func(pixelX, pixelY float64) EnemyInterface {
					return NewShieldedScoutEnemy(pixelX, pixelY)
				},
			},
			{
				4,
				func(pixelX, pixelY float64) EnemyInterface {
					return NewHeavyScoutEnemy(pixelX, pixelY)
				},
			},
		},
	},
	{
		[]WaveContent{
			{
				4,
				func(pixelX, pixelY float64) EnemyInterface {
					return NewEliteScoutEnemy(pixelX, pixelY)
				},
			},
			{
				6,
				func(pixelX, pixelY float64) EnemyInterface {
					return NewScoutEnemy(pixelX, pixelY)
				},
			},
			{
				6,
				func(pixelX, pixelY float64) EnemyInterface {
					return NewShieldedScoutEnemy(pixelX, pixelY)
				},
			},
			{
				6,
				func(pixelX, pixelY float64) EnemyInterface {
					return NewHeavyScoutEnemy(pixelX, pixelY)
				},
			},
		},
	},
	{
		[]WaveContent{
			{
				2,
				func(pixelX, pixelY float64) EnemyInterface {
					return NewEliteScoutEnemy(pixelX, pixelY)
				},
			},
			{
				8,
				func(pixelX, pixelY float64) EnemyInterface {
					return NewScoutEnemy(pixelX, pixelY)
				},
			},
			{
				8,
				func(pixelX, pixelY float64) EnemyInterface {
					return NewShieldedScoutEnemy(pixelX, pixelY)
				},
			},
			{
				8,
				func(pixelX, pixelY float64) EnemyInterface {
					return NewHeavyScoutEnemy(pixelX, pixelY)
				},
			},
		},
	},
}
var heaviesWave = []Wave{
	{
		[]WaveContent{
			{
				4,
				func(pixelX, pixelY float64) EnemyInterface {
					return NewHeaviesEnemy(pixelX, pixelY)
				},
			},
		},
	},
	{
		[]WaveContent{
			{
				4,
				func(pixelX, pixelY float64) EnemyInterface {
					return NewHeaviesEnemy(pixelX, pixelY)
				},
			},
			{
				4,
				func(pixelX, pixelY float64) EnemyInterface {
					return NewShieldedHeaviesEnemy(pixelX, pixelY)
				},
			},
		},
	},
	{
		[]WaveContent{
			{
				6,
				func(pixelX, pixelY float64) EnemyInterface {
					return NewShieldedHeaviesEnemy(pixelX, pixelY)
				},
			},
			{
				6,
				func(pixelX, pixelY float64) EnemyInterface {
					return NewHeavyHeaviesEnemy(pixelX, pixelY)
				},
			},
		},
	},
	{
		[]WaveContent{
			{
				20,
				func(pixelX, pixelY float64) EnemyInterface {
					return NewHeaviesEnemy(pixelX, pixelY)
				},
			},
		},
	},
	{
		[]WaveContent{
			{
				4,
				func(pixelX, pixelY float64) EnemyInterface {
					return NewEliteHeaviesEnemy(pixelX, pixelY)
				},
			},
		},
	},
	{
		[]WaveContent{
			{
				6,
				func(pixelX, pixelY float64) EnemyInterface {
					return NewHeaviesEnemy(pixelX, pixelY)
				},
			},
			{
				6,
				func(pixelX, pixelY float64) EnemyInterface {
					return NewScoutEnemy(pixelX, pixelY)
				},
			},
		},
	},
}
