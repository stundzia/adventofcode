package day7

import (
	"fmt"
	"sync"
	"time"

	"github.com/stundzia/adventofcode/2019/computer"
	"github.com/stundzia/adventofcode/utils"
)

type Amps struct {
	A *Amp
	B *Amp
	C *Amp
	D *Amp
	E *Amp
}

type Amp struct {
	computer *computer.Computer
}

func (as *Amps) TrySequence(seq []int) int {
	go as.A.computer.Run()
	as.A.computer.InputPipe <- seq[0]
	as.A.computer.InputPipe <- 0
	go as.B.computer.Run()
	as.B.computer.InputPipe <- seq[1]
	as.B.computer.InputPipe <- <-as.A.computer.OutputPipe
	go as.C.computer.Run()
	as.C.computer.InputPipe <- seq[2]
	as.C.computer.InputPipe <- <-as.B.computer.OutputPipe
	go as.D.computer.Run()
	as.D.computer.InputPipe <- seq[3]
	as.D.computer.InputPipe <- <-as.C.computer.OutputPipe
	go as.E.computer.Run()
	as.E.computer.InputPipe <- seq[4]
	as.E.computer.InputPipe <- <-as.D.computer.OutputPipe
	res := <-as.E.computer.OutputPipe

	return res
}

func (as *Amps) TrySequenceWithFeedBackLoop(seq []int, resChannel chan int) int {
	AB := make(chan int, 5)
	BC := make(chan int, 5)
	CD := make(chan int, 5)
	DE := make(chan int, 5)
	EA := make(chan int, 5)

	as.A.computer.OutputPipe = AB
	as.B.computer.InputPipe = AB

	as.B.computer.OutputPipe = BC
	as.C.computer.InputPipe = BC

	as.C.computer.OutputPipe = CD
	as.D.computer.InputPipe = CD

	as.D.computer.OutputPipe = DE
	as.E.computer.InputPipe = DE

	as.E.computer.OutputPipe = EA
	as.A.computer.InputPipe = EA

	as.A.computer.FirstInputs = []int{seq[0] + 5, 0}
	as.B.computer.FirstInputs = []int{seq[1] + 5}
	as.C.computer.FirstInputs = []int{seq[2] + 5}
	as.D.computer.FirstInputs = []int{seq[3] + 5}
	as.E.computer.FirstInputs = []int{seq[4] + 5}
	go as.A.computer.Run()
	go as.B.computer.Run()
	go as.C.computer.Run()
	go as.D.computer.Run()
	go as.E.computer.Run()

	// TODO: bit hacky here, need to give time for the computers to spin up or the below condition might evaluate
	//  to true, because the goroutines haven't had a chance to start yet
	time.Sleep(10 * time.Millisecond)

	for as.E.computer.Running.Load() {
		time.Sleep(10 * time.Millisecond)
	}
	res := <-as.E.computer.OutputPipe

	resChannel <- res
	return res
}

func getBestSequence(opcodes []int) int {
	maxSignal := 0
	combos := utils.GenUniqueCombos(5, 5)

	for _, vals := range combos {
		ops := make([]int, len(opcodes))
		copy(ops, opcodes)
		amps := NewAmps(ops)
		res := amps.TrySequence(vals)
		if res > maxSignal {
			maxSignal = res
		}
	}

	return maxSignal
}

func getBestSequenceV2(opcodes []int) int {
	maxSignal := 0
	combos := utils.GenUniqueCombos(5, 5)
	resChannel := make(chan int)

	wg := sync.WaitGroup{}
	for _, vals := range combos {
		ops := make([]int, len(opcodes))
		copy(ops, opcodes)
		amps := NewAmps(ops)
		go amps.TrySequenceWithFeedBackLoop(vals, resChannel)
		wg.Add(1)
	}

	waitChan := make(chan struct{})
	go func() {
		wg.Wait()
		waitChan <- struct{}{}
	}()
Main:
	for {
		select {
		case res := <-resChannel:
			if res > maxSignal {
				maxSignal = res
			}
			wg.Done()
		case <-waitChan:
			break Main
		default:
		}
	}

	return maxSignal
}

func NewAmps(opcodes []int) *Amps {
	return &Amps{
		A: &Amp{computer: computer.NewComputer(opcodes)},
		B: &Amp{computer: computer.NewComputer(opcodes)},
		C: &Amp{computer: computer.NewComputer(opcodes)},
		D: &Amp{computer: computer.NewComputer(opcodes)},
		E: &Amp{computer: computer.NewComputer(opcodes)},
	}
}

func DoSilver() string {
	input, _ := utils.ReadInputFileContentsAsIntSlice(2019, 7, ",")
	return fmt.Sprintf("%d", getBestSequence(input))
}

func DoGold() string {
	input, _ := utils.ReadInputFileContentsAsIntSlice(2019, 7, ",")
	return fmt.Sprintf("%d", getBestSequenceV2(input))
}
