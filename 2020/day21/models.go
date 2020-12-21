package day21

import (
	"fmt"
	"strings"
)

type Menu struct {
	Items []*MenuItem
	AllergenicIngredients []string
	IngredientsAllergensMap map[string]string
	AllergenMustBeInIngredientsMap map[string]map[string]struct{}
	CanBeAllergensMap map[string]struct{}
	InertIngredientsMap map[string]struct{}
}

type MenuItem struct {
	id int
	menu *Menu
	ingredients map[string]struct{}
	allergens map[string]struct{}
}

func NewMenu(menuList []string) *Menu {
	menu := &Menu{
		Items:                 []*MenuItem{},
		AllergenicIngredients: []string{},
		IngredientsAllergensMap: map[string]string{},
		AllergenMustBeInIngredientsMap: map[string]map[string]struct{}{},
		CanBeAllergensMap: map[string]struct{}{},
		InertIngredientsMap: map[string]struct{}{},
	}
	for id, line := range menuList {
		menu.parseIngredientsAllergens(line, id)
	}
	return menu
}


func (m *Menu) parseIngredientsAllergens(line string, id int) {
	parts := strings.Split(line, " (")
	ingredients := strings.Split(parts[0], " ")
	allergens := strings.Split(parts[1][9:len(parts[1]) - 1], ", ")
	ingredientsMap := map[string]struct{}{}
	allergensMap := map[string]struct{}{}
	for _, ing := range ingredients {
		ingredientsMap[ing] = struct {}{}
	}
	for _, al := range allergens {
		allergensMap[al] = struct {}{}
	}
	mi := &MenuItem{
		id: id,
		menu: m,
		ingredients: ingredientsMap,
		allergens:  allergensMap,
	}
	m.Items = append(m.Items, mi)
}

func (m *Menu) updateIngredientsAllergensMap() {
	for al, ings := range m.AllergenMustBeInIngredientsMap {
		if len(ings) == 1 {
			for k, _ := range ings {
				m.IngredientsAllergensMap[k] = al
			}
		}
	}
}

func (m *Menu) countUnmappedIngredients() int {
	count := 0
	for _, mi := range m.Items {
		for ing, _ := range mi.ingredients {
			if _, ok := m.CanBeAllergensMap[ing]; !ok {
				m.InertIngredientsMap[ing] = struct{}{}
				count++
			}
		}
	}
	return count
}


func (mi *MenuItem) UpdateAllergensFromOtherItem(other *MenuItem) {
	if mi.id == other.id {
		return
	}
	for oi, _ := range other.ingredients {
		if _, ok := mi.ingredients[oi]; !ok {
			return
		}
	}
	for oa, _ := range other.allergens {
		mi.allergens[oa] = struct {}{}
	}
}

func (mi *MenuItem) FindSameAllergensAndIngredients(other *MenuItem) {
	if mi.id == other.id {
		return
	}
	sameAllergens := []string{}
	sameIngredients := []string{}
	for oa, _ := range other.allergens {
		if _, ok := mi.allergens[oa]; ok {
			sameAllergens = append(sameAllergens, oa)
		}
	}
	for oi, _ := range other.ingredients {
		if _, ok := mi.ingredients[oi]; ok {
			sameIngredients = append(sameIngredients, oi)
		}
	}
	if len(sameIngredients) == 1 && len(sameAllergens) == 1 {
		mi.menu.IngredientsAllergensMap[sameIngredients[0]] = sameAllergens[0]
	}
}

func (mi *MenuItem) unknownIngredientsAllergens() (ingredients, allergens map[string]struct{}) {
	ingredients = map[string]struct{}{}
	allergens = map[string]struct{}{}
	knownAls := map[string]interface{}{}
	for ing, _ := range mi.ingredients {
		if knownAl, ok := mi.menu.IngredientsAllergensMap[ing]; ok {
			knownAls[knownAl] = struct {}{}
			continue
		}
		ingredients[ing] = struct{}{}
	}
	for al, _ := range mi.allergens {
		if _, ok := knownAls[al]; ok {
			continue
		}
		allergens[al] = struct{}{}
	}
	return ingredients, allergens
}

func (mi *MenuItem) updateAllergenMustBeInIngredientsMap() {
	uIngs, uAls := mi.unknownIngredientsAllergens()
	for uAl, _ := range uAls {
		if _, ok := mi.menu.AllergenMustBeInIngredientsMap[uAl]; !ok {
			mi.menu.AllergenMustBeInIngredientsMap[uAl] = map[string]struct{}{}
			for uIng, _ := range uIngs {
				mi.menu.AllergenMustBeInIngredientsMap[uAl][uIng] = struct{}{}
			}
		} else {
			for mIng, _ := range mi.menu.AllergenMustBeInIngredientsMap[uAl] {
				if _, ok := uIngs[mIng]; !ok {
					delete(mi.menu.AllergenMustBeInIngredientsMap[uAl], mIng)
				}
			}
		}
	}
}


func (m *Menu) formIngredientAllergenMaps() {
	for _, mi := range m.Items {
		for _, omi := range m.Items {
			mi.UpdateAllergensFromOtherItem(omi)
			mi.FindSameAllergensAndIngredients(omi)
		}
	}
	for _, mi := range m.Items {
		mi.updateAllergenMustBeInIngredientsMap()
	}
	for i := 0; i < 2; i++ {
		m.updateIngredientsAllergensMap()
		for _, mi := range m.Items {
			mi.updateAllergenMustBeInIngredientsMap()
		}
	}
	for _, ingMap := range m.AllergenMustBeInIngredientsMap {
		for ing, _ := range ingMap {
			m.CanBeAllergensMap[ing] = struct{}{}
		}
	}
	for ing, _ := range m.IngredientsAllergensMap {
		m.CanBeAllergensMap[ing] = struct{}{}
	}
}

func (m *Menu) printStuff() {
	for _, mi := range m.Items {
		fmt.Printf("%d ings: %d alls: %d\n", mi.id, len(mi.ingredients), len(mi.allergens))
	}
	fmt.Println(m.AllergenMustBeInIngredientsMap)
	fmt.Println(m.IngredientsAllergensMap)
}