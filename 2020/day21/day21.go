package day21

import (
	"fmt"
	"github.com/stundzia/adventofcode/utils"
)



func DoSilver() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2020, 21, "\n")
	menu := NewMenu(input)
	menu.formIngredientAllergenMaps()
	menu.printStuff()
	return fmt.Sprintf("%d", menu.countUnmappedIngredients())  // 2149 < x <2335
}

func DoGold() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2020, 21, "\n")
	return fmt.Sprintf("%d", len(input))
}
