package utils

import (
	"fmt"
	"gonum.org/v1/gonum/stat/combin"
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

// ReadInputFileContentsAsIntSliceLines - reads input as separated and grouped by blank lines ints, e.g.:
// ```
//
//	1000
//	2000
//	3000
//
//	4000
//
//	5000
//	6000
//
//	7000
//	8000
//	9000
//
//	10000
//
// ```
// Would return a 2D int slice [][]int{{1000, 2000, 3000}, {4000}, {5000, 6000}, {7000, 8000, 9000}, {10000}}
func ReadInputFileContentsAsIntSliceLines(year int, day int) ([][]int, error) {
	strSlice, err := ReadInputFileContentsAsStringSlice(year, day, "\n")
	if err != nil {
		return nil, err
	}
	res := [][]int{}
	l := []int{}
	for _, val := range strSlice {
		if val == "" {
			res = append(res, l)
			l = []int{}
			continue
		}
		num, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		l = append(l, num)
	}
	res = append(res, l)
	return res, err
}

func ReadInputFileContentsAsIntGrid(year int, day int) ([][]int, error) {
	strSlice, err := ReadInputFileContentsAsStringSlice(year, day, "\n")
	if err != nil {
		return nil, err
	}
	res := [][]int{}
	for _, line := range strSlice {
		l := []int{}
		for _, v := range line {
			val := string(v)
			num, err := strconv.Atoi(val)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			l = append(l, num)
		}
		res = append(res, l)
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

// SlicesIntEqual - returns true is a and b are equal int slices (e.g. if a = [2,3,1]
// then if b = [2,3,1] will return true, if b = [1,2,3] will return false, i.e. order is important).
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

// CoordsStrToInts - converts coords string of format "x,y" to ints and returns x,y as ints.
func CoordsStrToInts(coords string) (x, y int) {
	coordsSlice := strings.Split(coords, ",")
	x, _ = strconv.Atoi(coordsSlice[0])
	y, _ = strconv.Atoi(coordsSlice[1])
	return x, y
}

// CoordsIntsToStr - converts coords ints (x,y) to string format "x,y".
func CoordsIntsToStr(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

// Coords3DStrToInts - converts 3D coords string of format "x,y,z" to ints and returns x,y,z as ints.
func Coords3DStrToInts(coords string) (x, y, z int) {
	coordsSlice := strings.Split(coords, ",")
	x, _ = strconv.Atoi(coordsSlice[0])
	y, _ = strconv.Atoi(coordsSlice[1])
	z, _ = strconv.Atoi(coordsSlice[2])
	return x, y, z
}

// Coords3DIntsToStr - converts 3D coords to string format "x,y,z".
func Coords3DIntsToStr(x, y, z int) string {
	return fmt.Sprintf("%d,%d,%d", x, y, z)
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func GenUniqueCombos(n, k int) [][]int {
	res := [][]int{}
	for i := 0; i < combin.NumPermutations(n, k); i++ {
		comb := make([]int, k)
		combin.IndexToPermutation(comb, i, n, k)
		res = append(res, comb)
	}
	return res
}

func RotateCoordinates(x, y float64, degrees float64) (newX, newY float64) {
	// y' = y*cos(a) - x*sin(a)
	// x' = y*sin(a) + x*cos(a)
	sinA, cosA := math.Sincos(degrees * math.Pi / 180)
	newY = y*cosA - x*sinA
	newX = y*sinA + x*cosA
	return newX, newY
}

func RunWithTimeMetricsAndPrintOutput(solver AocSolver) {
	start := time.Now()
	fmt.Println("Solution is: ", solver())
	fmt.Println("Solution took: ", time.Now().Sub(start))
}

// GCD - returns greatest common divisor of given 2 ints.
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

// RemoveFromIntSlice - removes element at provided index from int slice.
func RemoveFromIntSlice(s []int, index int) []int {
	s[index] = s[len(s)-1]
	return s[:len(s)-1]
}

func RemoveFrom2DIntSlice(s [][]int, i int) [][]int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// StringSliceContains - returns true if slice of strings contains given string.
func StringSliceContains(ss []string, s string) bool {
	for _, sVal := range ss {
		if sVal == s {
			return true
		}
	}
	return false
}

// ReverseString - returns reversed string (i.e. given "abcd" returns "dcba").
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

// GetMaxFromIntSlice - returns the biggest positive integer from a slice, if no positive integer exists, returns 0.
func GetMaxFromIntSlice(nums []int) int {
	max := 0
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

// GetMinFromIntSlice - returns the smallest integer from a slice.
func GetMinFromIntSlice(nums []int) int {
	min := math.MaxInt
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

// GetMaxAbsFromIntSlice - returns the biggest absolute value of an integer from a slice.
func GetMaxAbsFromIntSlice(nums []int) int {
	max := 0
	for _, num := range nums {
		if AbsInt(num) > max {
			max = AbsInt(num)
		}
	}
	return max
}

// GetMinAbsFromIntSlice - returns the smallest absolute value of an integer from a slice.
func GetMinAbsFromIntSlice(nums []int) int {
	min := math.MaxInt
	for _, num := range nums {
		if AbsInt(num) < min {
			min = AbsInt(min)
		}
	}
	return min
}

// AbsInt - returns absolute value of an integer (if num < 0 returns -num).
func AbsInt(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
