package day23

import "fmt"

type CupGame struct {
	CurrentCup *Cup
}


type Cup struct {
	Next *Cup
	Previous *Cup
	Label int
}


func NewCupGame(cupLabels []int) *CupGame {
	cg := &CupGame{}
	cups := []*Cup{}
	for _, cl := range cupLabels {
		cups = append(cups, &Cup{
			Next:     nil,
			Previous: nil,
			Label:    cl,
		})
	}
	for i, cup := range cups {
		if i == 0 {
			cup.Previous = cups[len(cups) - 1]
		} else {
			cup.Previous = cups[i - 1]
		}
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
	fmt.Println("Picked up: ", cups[0].Label, cups[1].Label, cups[2].Label)
	cups[0].Previous = nil
	cg.CurrentCup.Next = cups[len(cups) - 1].Next
	cg.CurrentCup.Next.Previous = cg.CurrentCup
	cups[len(cups) - 1].Next.Previous = cg.CurrentCup
	cups[len(cups) - 1].Next = nil
	cg.PrintCups()
	destCup := cg.getDestinationCup()
	fmt.Println("destination: ", destCup.Label)
	destCup.Next.Previous = cups[len(cups) - 1]
	cups[len(cups) - 1].Next = destCup.Next
	destCup.Next = cups[0]
	cups[0].Previous = destCup
	cg.CurrentCup = cg.CurrentCup.Next
	cg.PrintCups()
	fmt.Println("cc: ", cg.CurrentCup.Label)
}

func (cg *CupGame) PrintCups() {
	cc := cg.CurrentCup
	fmt.Printf(" %d <- %d -> %d ", cc.Previous.Label, cc.Label, cc.Next.Label)
	cc = cc.Next
	for ; cc != cg.CurrentCup; {
		fmt.Printf(" %d <- %d -> %d ", cc.Previous.Label, cc.Label, cc.Next.Label)
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