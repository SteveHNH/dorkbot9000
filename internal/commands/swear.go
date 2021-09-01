package commands

import (
	"fmt"
	"reflect"
)

func Swear(argument []string) (string, error) {
	if len(argument) == 0 {
		return "", fmt.Errorf("No argument provided")
	}
	switch argument[0] {
	case "set":
		if len(argument) < 2 {
			return "", fmt.Errorf("No argument provided")
		}
		value := reflect.TypeOf(argument[1])
		if value != int {
			return "", fmt.Errorf("Set must include a number")
		}
		return "Swear set to " + argument[1], nil
	case "+":
		return "Swears increased to by 1", nil
	case "-":
		return "Swears decreased to by 1", nil
	case "reset":
		return "Swears reset to 0", nil
	default:
		return "", fmt.Errorf("Invalid argument")
	}
}

func HelpSwear() string {
	return "Swear word"
}
