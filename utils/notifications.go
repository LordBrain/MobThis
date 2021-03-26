package utils

import "github.com/gen2brain/beeep"

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
