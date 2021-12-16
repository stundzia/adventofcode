package day21

import (
	"fmt"
	"github.com/stundzia/adventofcode/utils"
	"sort"
	"strings"
)

func DoSilver() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2020, 21, "\n")
	menu := NewMenu(input)
	for i := 0; i < 5; i++ {
		menu.formIngredientAllergenMaps()
	}
	return fmt.Sprintf("%d", menu.countUnmappedIngredients()) // 2265
}

func DoGold() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2020, 21, "\n")
	menu := NewMenu(input)
	for i := 0; i < 5; i++ {
		menu.formIngredientAllergenMaps()
	}
	alsIngs := []string{}
	for k, v := range menu.IngredientsAllergensMap {
		alsIngs = append(alsIngs, fmt.Sprintf("%s-%s", v, k))
	}
	sort.Strings(alsIngs)
	res := []string{}
	for _, ai := range alsIngs {
		res = append(res, strings.Split(ai, "-")[1])
	}
	return strings.Join(res, ",")
}
