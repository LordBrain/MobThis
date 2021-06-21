package utils

import (
	"github.com/gen2brain/beeep"
)

func MobNotify(title, message string) error {
	err := beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
	if err != nil {
		return err
	}
	err = beeep.Notify(title, message, "assets/information.png")
	if err != nil {
		return err
	}
	return nil
}

func CheckMobbers(current, updated []string) ([]string, []string) {
	var addMobber []string
	var removeMobber []string

	//Check for a new mobber
	for _, mobber := range updated {
		for _, oldMobber := range current {
			if mobber != oldMobber {
				addMobber = append(addMobber, mobber)
			}
		}

	}
	for _, mobber := range current {
		for _, formerMobber := range updated {
			if formerMobber != mobber {
				removeMobber = append(removeMobber, mobber)
			}

		}

	}
	// fmt.Println("Add Mobbers: ", addMobber)
	// fmt.Println("Remove Mobbers: ", removeMobber)
	return addMobber, removeMobber
}
