package gui

var (
	cancelButton    *Button
	buyTurretButton *Button
	sellButton      *Button

	buyMagnet      *Button
	buyLightTurret *Button
	buyHeavyTurret *Button
)

var (
	LightTurretButton = "LightTurretButton"
	HeavyTurretButton = "HeavyTurretButton"
	MagnetButton      = "MagnetButton"
)

func InitButtons() map[string]bool {
	unlocked := make(map[string]bool)
	unlocked[MagnetButton] = true
	unlocked[LightTurretButton] = true
	unlocked[HeavyTurretButton] = true

	buyMagnet = NewMagnetButton()
	buyLightTurret = NewLightTurretButton()
	buyHeavyTurret = NewHeavyTurretButton()

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
	return retButtons
}

func GetSelectedBuildingButtons(buttonsToShow map[string]bool) []*Button {
	return nil
}
