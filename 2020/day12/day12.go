package day12

import (
	"fmt"
	"github.com/stundzia/adventofcode/utils"
	"strconv"
)


func parseNavCommand(c string) (direction string, value int) {
	direction = string(c[0])
	value, _ = strconv.Atoi(c[1:])
	return direction, value
}

func DoSilver() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2020, 12, "\n")
	ship := NewShip(90)
	for _, command := range input {
		dir, val := parseNavCommand(command)
		ship.HandleNavCommand(dir, val)
	}
	return fmt.Sprintf(strconv.Itoa(ship.getManhattenDistance()))
}

func DoGold() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2020, 12, "\n")
	return fmt.Sprintf(strconv.Itoa(len(input)))
}
