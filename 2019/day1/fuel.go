package day1

import (
	"github.com/stundzia/adventofcode/utils"
	"strconv"
)


func getMassFuelReq(mass int) int {
	return mass / 3 - 2
}

func DoSilver() string {
	input, _ := utils.ReadInputFileContentsAsIntSlice(2019, 1, "\n")
	moduleFuelMasses := make([]int, len(input))
	for i, val := range input {
		moduleFuelMasses[i] = getMassFuelReq(val)
	}
	return strconv.Itoa(utils.SumIntSlice(moduleFuelMasses))
}

func DoGold() string {
	input, _ := utils.ReadInputFileContentsAsIntSlice(2019, 1, "\n")
	moduleFuelMasses := make([]int, len(input))
	for i, val := range input {
		res := getMassFuelReq(val)
		fuelMass := res
		for {
			fuelMass = getMassFuelReq(fuelMass)
			if fuelMass <= 0 {
				break
			}
			res += fuelMass
		}
		moduleFuelMasses[i] = res
	}
	return strconv.Itoa(utils.SumIntSlice(moduleFuelMasses))
}