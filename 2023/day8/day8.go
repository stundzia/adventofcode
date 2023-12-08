package day8

import (
	"fmt"
	"github.com/stundzia/adventofcode/utils"
	"log"
	"strings"
)

type travel struct {
	currentNode     *mapNode
	stepsTaken      int
	nextInstruction int
	instructions    []string
	nodes           map[string]*mapNode
	lastZ           int
	zSteps          int
	stepsToFirstZ   int
}

type mapNode struct {
	name     string
	previous *mapNode
	left     *mapNode
	right    *mapNode
}

func parseTravelThing(instructions string, nodes []string) *travel {
	t := &travel{
		currentNode:     nil,
		nextInstruction: 0,
		instructions:    strings.Split(instructions, ""),
		nodes:           map[string]*mapNode{},
	}
	for _, n := range nodes {
		parts := strings.Split(n, " = ")
		t.nodes[parts[0]] = &mapNode{
			name:     parts[0],
			previous: nil,
			left:     nil,
			right:    nil,
		}
	}
	t.parseNodeConns(nodes)
	t.currentNode = t.nodes["AAA"]
	return t
}

func (t *travel) parseNodeConns(nodesInfo []string) {
	for _, n := range nodesInfo {
		parts := strings.Split(n, " = ")
		leftRightParts := strings.Split(parts[1], ", ")
		t.nodes[parts[0]].left = t.nodes[leftRightParts[0][1:]]
		t.nodes[parts[0]].right = t.nodes[leftRightParts[1][:3]]
	}
}

func (t *travel) doStep(final string) bool {
	if t.nextInstruction >= len(t.instructions) {
		t.nextInstruction = 0
	}
	switch t.instructions[t.nextInstruction] {
	case "R":
		t.currentNode = t.currentNode.right
	case "L":
		t.currentNode = t.currentNode.left
	default:
		log.Fatal("wtf is this: ")
	}
	t.stepsTaken++
	t.nextInstruction = t.nextInstruction + 1

	return t.currentNode.name == final
}

func (t *travel) doStepP2() (bool, int) {
	if t.zSteps != 0 {
		t.stepsTaken += t.zSteps
		//t.zStepsMap[t.stepsTaken] = struct{}{}
		return true, t.stepsTaken
	}
	if t.nextInstruction >= len(t.instructions) {
		t.nextInstruction = 0
	}
	switch t.instructions[t.nextInstruction] {
	case "R":
		t.currentNode = t.currentNode.right
	case "L":
		t.currentNode = t.currentNode.left
	default:
		log.Fatal("wtf is this: ")
	}
	t.stepsTaken++
	t.nextInstruction = t.nextInstruction + 1

	if t.currentNode.name[2] == 'Z' {
		if t.lastZ == 0 {
			t.lastZ = t.stepsTaken
			t.stepsToFirstZ = t.stepsTaken
		} else {
			t.zSteps = t.stepsTaken - t.lastZ
		}
	}

	return t.currentNode.name[2] == 'Z', t.stepsTaken
}

func DoSilver() string {
	input, _ := utils.ReadInputFileContentsAsStringSliceLines(2023, 8)
	t := parseTravelThing(input[0][0], input[1])

	for t.doStep("ZZZ") != true {

	}

	res := t.stepsTaken

	return fmt.Sprintf("Solution: %d", res)
}

func DoGold() string {
	input, _ := utils.ReadInputFileContentsAsStringSliceLines(2023, 8)

	ts := map[string]*travel{}
	for _, node := range input[1] {
		if node[2] == 'A' {
			t := parseTravelThing(input[0][0], input[1])
			t.currentNode = t.nodes[node[:3]]
			ts[node] = t
		}
	}

	var biggestStepper *travel
	var biggestStep *travel
	for i := 0; ; i++ {
		toContinue := false
		for _, t := range ts {
			t.doStepP2()
			if t.zSteps == 0 {
				toContinue = true
			}
			if !toContinue {
				if biggestStepper == nil || t.zSteps > biggestStepper.zSteps {
					biggestStepper = t
				}
				if biggestStep == nil || t.stepsTaken > biggestStep.stepsTaken {
					biggestStep = t
				}
			}
		}
		if toContinue {
			continue
		}
		if biggestStepper != nil && biggestStepper == biggestStep {
			break
		}
	}
mf:
	for {
		_, step := biggestStepper.doStepP2()
		for _, t := range ts {
			if t == biggestStepper {
				continue
			}
			_, otherStep := t.doStepP2()
			for otherStep < step {
				_, otherStep = t.doStepP2()
			}
			if otherStep != step {
				continue mf
			}
		}
		return fmt.Sprintf("Solution: %d", step)
	}
}
