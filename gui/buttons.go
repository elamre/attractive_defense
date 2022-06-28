package gui

var (
	cancelButton    *Button
	buyTurretButton *Button
	sellButton      *Button

	buyMagnet       *Button
	buyLightTurret  *Button
	buyHeavyTurret  *Button
	buyBeamTurret   *Button
	buyRocketTurret *Button
	buyResearch     *Button
)

var (
	LightTurretButton  = "LightTurretButton"
	HeavyTurretButton  = "HeavyTurretButton"
	RocketTurretButton = "RocketTurretButton"
	BeamTurretButton   = "BeamTurretButton"
	MagnetButton       = "MagnetButton"
)

func InitButtons() map[string]bool {
	unlocked := make(map[string]bool)
	unlocked[MagnetButton] = true
	unlocked[LightTurretButton] = true
	unlocked[HeavyTurretButton] = true
	unlocked[RocketTurretButton] = true
	unlocked[BeamTurretButton] = true

	buyMagnet = NewMagnetButton()
	buyLightTurret = NewLightTurretButton()
	buyHeavyTurret = NewHeavyTurretButton()
	buyBeamTurret = NewBeamTurretButton()
	buyRocketTurret = NewRocketTurretButton()
	buyResearch = NewResearchBuildingButton()

	return unlocked
}

func GetBuildingButtons(unlockedMap map[string]bool) []*Button {
	retButtons := make([]*Button, 0)
	if unlockedMap[MagnetButton] {
		retButtons = append(retButtons, buyMagnet)
	}
	if unlockedMap[LightTurretButton] {
		retButtons = append(retButtons, buyLightTurret)
	}
	if unlockedMap[HeavyTurretButton] {
		retButtons = append(retButtons, buyHeavyTurret)
	} else {
		//maybe show that it exist but can't be pressed?
	}
	if unlockedMap[BeamTurretButton] {
		retButtons = append(retButtons, buyBeamTurret)
	}
	if unlockedMap[RocketTurretButton] {
		retButtons = append(retButtons, buyRocketTurret)
	}

	return retButtons
}

func GetSelectedBuildingButtons(buttonsToShow map[string]bool) []*Button {
	return nil
}
