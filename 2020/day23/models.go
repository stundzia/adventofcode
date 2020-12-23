package day23

import (
	"fmt"
	"strconv"
)

type CupGame struct {
	CurrentCup *Cup
	MaxLabel int
	LabelMap map[int]*Cup
}


type Cup struct {
	Next *Cup
	Label int
}


func NewCupGame(cupLabels []int) *CupGame {
	cg := &CupGame{
		LabelMap: map[int]*Cup{},
	}
	cups := []*Cup{}
	for _, cl := range cupLabels {
		cup := &Cup{
			Next:     nil,
			Label:    cl,
		}
		cups = append(cups, cup)
		cg.LabelMap[cl] = cup
		if cl > cg.MaxLabel {
			cg.MaxLabel = cl
		}
	}
	for i, cup := range cups {
		if i == len(cups) - 1 {
			cup.Next = cups[0]
		} else {
			cup.Next = cups[i + 1]
		}
	}
	cg.CurrentCup = cups[0]
	return cg
}

func (cg *CupGame) getDestinationCup(pickedLabels [3]int) *Cup {
	label := cg.CurrentCup.Label - 1
	MainLoop:
	for ;label > 0; label-- {
		for _, pl := range pickedLabels {
			if label == pl {
				continue MainLoop
			}
		}
		return cg.LabelMap[label]
	}
	label = cg.MaxLabel
	MaxLabelLoop:
	for ;label > 0; label-- {
		for _, pl := range pickedLabels {
			if label == pl {
				continue MaxLabelLoop
			}
		}
		return cg.LabelMap[label]
	}
	return nil
}


func (cg *CupGame) doMove() {
	takenCups := []*Cup{}
	currentCup := cg.CurrentCup
	takenLabels := [3]int{}
	for i := 0; i < 3; i++ {
		takenCups = append(takenCups, currentCup.Next)
		takenLabels[i] = currentCup.Next.Label
		currentCup = currentCup.Next
	}
	cg.CurrentCup.Next = takenCups[len(takenCups) - 1].Next
	takenCups[len(takenCups) - 1].Next = nil
	destCup := cg.getDestinationCup(takenLabels)
	takenCups[len(takenCups) - 1].Next = destCup.Next
	destCup.Next = takenCups[0]
	cg.CurrentCup = cg.CurrentCup.Next
}

func (cg *CupGame) PrintCups() {
	cc := cg.CurrentCup
	fmt.Printf(" %d -> %d ", cc.Label, cc.Next.Label)
	cc = cc.Next
	for ; cc != cg.CurrentCup; {
		fmt.Printf(" %d -> %d ", cc.Label, cc.Next.Label)
		cc = cc.Next
	}
	cc = cg.CurrentCup
	fmt.Printf("\n %d ", cc.Label)
	cc = cc.Next
	for ; cc != cg.CurrentCup; {
		fmt.Printf(" %d ", cc.Label)
		cc = cc.Next
	}
	fmt.Printf("\n")
}

func (cg *CupGame) GetPart2Res() int {
	return cg.LabelMap[1].Next.Label * cg.LabelMap[1].Next.Next.Label
}

func (cg *CupGame) GetPart1Res() string {
	res := ""
	cup := cg.LabelMap[1].Next
	for ; cup != cg.LabelMap[1]; cup = cup.Next {
		res += strconv.Itoa(cup.Label)
	}
	return res
}