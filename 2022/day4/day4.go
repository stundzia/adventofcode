package day4

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/stundzia/adventofcode/utils"
)

type assignmentPair struct {
	Amin int
	Amax int
	Bmin int
	Bmax int
}

func newAssignmentPair(l string) *assignmentPair {
	parts := strings.Split(l, ",")
	fmt.Println("parts: ", parts)
	rangeAStr := strings.Split(parts[0], "-")
	rangeBStr := strings.Split(parts[1], "-")
	aMin, _ := strconv.Atoi(rangeAStr[0])
	aMax, _ := strconv.Atoi(rangeAStr[1])
	bMin, _ := strconv.Atoi(rangeBStr[0])
	bMax, _ := strconv.Atoi(rangeBStr[1])
	return &assignmentPair{
		Amin: aMin,
		Amax: aMax,
		Bmin: bMin,
		Bmax: bMax,
	}
}

func (ap *assignmentPair) isCoveredTwice() bool {
	return (ap.Amin <= ap.Bmin && ap.Amax >= ap.Bmax) || (ap.Bmin <= ap.Amin && ap.Bmax >= ap.Amax)
}

func (ap *assignmentPair) overlaps() bool {
	return (ap.Amin >= ap.Bmin && ap.Amin <= ap.Bmax) || (ap.Bmin >= ap.Amin && ap.Bmin <= ap.Amax)
}

func DoSilver() string {
	lines, _ := utils.ReadInputFileContentsAsStringSlice(2022, 4, "\n")

	count := 0
	for _, l := range lines {
		ap := newAssignmentPair(l)
		if ap.isCoveredTwice() {
			count++
		}
	}

	return fmt.Sprintf("Solution: %d", count)
}

func DoGold() string {
	lines, _ := utils.ReadInputFileContentsAsStringSlice(2022, 4, "\n")

	count := 0
	for _, l := range lines {
		ap := newAssignmentPair(l)
		if ap.overlaps() {
			count++
		}
	}

	return fmt.Sprintf("Solution: %d", count)
}
