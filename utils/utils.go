package utils

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type AocSolver = func() string

func ReadInputFileBytes(year int, day int) ([]byte, error) {
	absPath, err := filepath.Abs(fmt.Sprintf("%d/day%d/input.txt", year, day))
	fileBytes, err := ioutil.ReadFile(absPath)
	return fileBytes, err
}

func ReadInputFileContentsAsString(year int, day int) (string, error) {
	fileBytes, err := ReadInputFileBytes(year, day)
	return string(fileBytes), err
}

func ReadInputFileContentsAsStringSlice(year int, day int, sep string) ([]string, error) {
	fileBytes, err := ReadInputFileBytes(year, day)
	fileStr := string(fileBytes)
	return strings.Split(fileStr, sep), err
}

func ReadInputFileContentsAsIntSlice(year int, day int, sep string) ([]int, error) {
	strSlice, err := ReadInputFileContentsAsStringSlice(year, day, sep)
	if err != nil {
		return nil, err
	}
	res := make([]int, len(strSlice))
	for i, val := range strSlice {
		num, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		res[i] = num
	}
	return res, err
}

// SumIntSlice returns the sum of all integers in an integer slice.
func SumIntSlice(slice []int) int {
	var res int
	for _, val := range slice {
		res += val
	}
	return res
}

func SliceStringToInt(slice []string) []int {
	var res []int
	for _, val := range slice {
		v, _ := strconv.Atoi(val)
		res = append(res, v)
	}
	return res
}


func RunWithTimeMetricsAndPrintOutput(solver AocSolver) {
	start := time.Now()
	fmt.Println("Solution is: ", solver())
	fmt.Println("Solution took: ", time.Now().Sub(start))
}