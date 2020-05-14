package gunFactory

import (
	"fmt"
	factory2 "pattern/creational/factory"
	gun2 "pattern/creational/factory/gun"
)

func GetGun(gunType string) (factory2.IGun, error) {
	if gunType == "ak47" {
		return gun2.NewAk47(), nil
	}
	if gunType == "maverick" {
		return gun2.NewMaverick(), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}
