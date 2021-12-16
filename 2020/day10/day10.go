package day10

import (
	"github.com/stundzia/adventofcode/utils"
	"sort"
	"strconv"
)

func handleAdapter(adapterMap map[int]map[string]int, adapter int) int {
	allowedDiffs := []int{1, 2, 3}
	if _, ok := adapterMap[adapter]["connected_to"]; ok {
		return -1
	}
	diffUsed := -1
	connectedTo := -1
	canConnectCount := 0
	for _, diff := range allowedDiffs {
		if _, ok := adapterMap[adapter+diff]; ok {
			if diffUsed == -1 && connectedTo == -1 {
				diffUsed = diff
				connectedTo = adapter + diff
			}
			canConnectCount++
		}
	}
	adapterMap[adapter]["connected_to"] = connectedTo
	adapterMap[adapter]["diff"] = diffUsed
	adapterMap[adapter]["can_connect_to"] = canConnectCount
	return connectedTo
}

func adapterMapFromInput(input []int) map[int]map[string]int {
	adapterMap := map[int]map[string]int{}
	for _, adapter := range input {
		adapterMap[adapter] = map[string]int{}
	}
	next := 1
	for next > 0 {
		next = handleAdapter(adapterMap, next)
	}
	return adapterMap
}

func adapterMatrixFromInput(input []int) [][]int {
	sort.Ints(input)
	adapterMatrix := make([][]int, len(input))
	for i, x := range input {
		adapterMatrix[i] = make([]int, len(input))
		for j, y := range input {
			if y > x && y-x <= 3 {
				adapterMatrix[i][j] = 1
			} else {
				adapterMatrix[i][j] = 0
			}
		}
	}
	return adapterMatrix
}

func DoSilver() string {
	input, _ := utils.ReadInputFileContentsAsIntSlice(2020, 10, "\n")
	adapterMap := adapterMapFromInput(input)
	count1Diff := 1
	count3Diff := 1
	for _, value := range adapterMap {
		if value["diff"] == 1 {
			count1Diff++
		}
		if value["diff"] == 3 {
			count3Diff++
		}
	}

	return strconv.Itoa(count1Diff * count3Diff)
}

func multiplyMatrices(matrix1 [][]int, matrix2 [][]int) [][]int {
	res := make([][]int, len(matrix1))
	for i, _ := range res {
		res[i] = make([]int, len(matrix1))
	}

	for i, x := range matrix1 {
		for j, _ := range x {
			for k := 0; k < len(matrix1); k++ {
				res[i][j] = res[i][j] + (matrix1[i][k] * matrix2[k][j])
			}
		}
	}
	return res
}

func DoGold() string {
	input, _ := utils.ReadInputFileContentsAsIntSlice(2020, 10, "\n")
	input = append(input, 0)
	max := 0
	for _, val := range input {
		if val > max {
			max = val
		}
	}
	input = append(input, max+3)
	adpMatrix := adapterMatrixFromInput(input)
	adpMatrixBase := make([][]int, len(adpMatrix))
	for i, _ := range adpMatrixBase {
		adpMatrixBase[i] = make([]int, len(adpMatrixBase))
	}
	copy(adpMatrixBase, adpMatrix)
	paths := 0
	for i := 0; i < len(input); i++ {
		adpMatrix = multiplyMatrices(adpMatrix, adpMatrixBase)
		paths += adpMatrix[0][len(adpMatrix)-1]
	}

	return strconv.Itoa(paths)
}
