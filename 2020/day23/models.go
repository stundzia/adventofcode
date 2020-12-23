package day23

import "fmt"

type CupGame struct {
	CurrentCup *Cup
}


type Cup struct {
	Next *Cup
	Label int
}


func NewCupGame(cupLabels []int) *CupGame {
	cg := &CupGame{}
	cups := []*Cup{}
	for _, cl := range cupLabels {
		cups = append(cups, &Cup{
			Next:     nil,
			Label:    cl,
		})
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

func (cg *CupGame) getDestinationCup() *Cup {
	label := cg.CurrentCup.Label - 1
	for ;label > 0; label-- {
		for cc := cg.CurrentCup.Next;cc != cg.CurrentCup; {
			if cc.Label == label {
				return cc
			}
			cc = cc.Next
		}
	}
	maxLabel := cg.CurrentCup.Label
	maxLabelCup := cg.CurrentCup
	for cc := cg.CurrentCup.Next; cc != cg.CurrentCup; {
		if cc.Label > maxLabel {
			maxLabel = cc.Label
			maxLabelCup = cc
		}
		cc = cc.Next
	}
	return maxLabelCup
}


func (cg *CupGame) doMove() {
	cups := []*Cup{}
	currentCup := cg.CurrentCup
	for i := 0; i < 3; i++ {
		cups = append(cups, currentCup.Next)
		currentCup = currentCup.Next
	}
	cg.CurrentCup.Next = cups[len(cups) - 1].Next
	cups[len(cups) - 1].Next = nil
	destCup := cg.getDestinationCup()
	cups[len(cups) - 1].Next = destCup.Next
	destCup.Next = cups[0]
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

func (cg *CupGame) PrintPart2Res() {
	cc := cg.CurrentCup
	cc = cc.Next
	for ; cc != cg.CurrentCup; {
		if cc.Label == 1 {
			fmt.Println("Result: ")
			fmt.Println(cc.Next.Label)
			fmt.Println(cc.Next.Next.Label)
		}
		cc = cc.Next
	}
}