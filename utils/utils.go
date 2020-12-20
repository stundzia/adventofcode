package utils

import (
	"fmt"
	"io/ioutil"
	"math"
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

func SlicesIntEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}


func RotateCoordinates(x, y float64, degrees float64) (newX, newY float64) {
	// y' = y*cos(a) - x*sin(a)
	// x' = y*sin(a) + x*cos(a)
	sinA, cosA := math.Sincos(degrees * math.Pi / 180)
	newY = y * cosA - x * sinA
	newX = y * sinA + x * cosA
	return newX, newY
}

func RunWithTimeMetricsAndPrintOutput(solver AocSolver) {
	start := time.Now()
	fmt.Println("Solution is: ", solver())
	fmt.Println("Solution took: ", time.Now().Sub(start))
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}


func RemoveFromIntSlice(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}


func RemoveFrom2DIntSlice(s [][]int, i int) [][]int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func StringSliceContains(ss []string, s string) bool {
	for _, sVal := range ss {
		if sVal == s {
			return true
		}
	}
	return false
}

func ReverseString(s string) string {
	newS := ""
	for i := len(s) - 1; i >= 0; i-- {
		newS += string(s[i])
	}
	return newS
}

func ReverseStringSlice(s []string) []string {
	newS := []string{}
	for i := len(s) - 1; i >= 0; i-- {
		newS = append(newS, s[i])
	}
	return newS
}

func ReverseStringSliceSlice(ss [][]string) [][]string {
	newS := [][]string{}
	for i := len(ss) - 1; i >= 0; i-- {
		newS = append(newS, ss[i])
	}
	return newS
}

func SlicesStringEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}