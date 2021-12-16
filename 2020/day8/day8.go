package day8

import (
	"fmt"
	"github.com/stundzia/adventofcode/utils"
	"strconv"
	"strings"
)

func parseInstruction(inst string) (op string, arg int) {
	instSlice := strings.Split(inst, " ")
	op = instSlice[0]
	arg, _ = strconv.Atoi(instSlice[1])
	return
}

func handleInstruction(instructions []string, acc int, next int) (accNew, nextNew int) {
	op, arg := parseInstruction(instructions[next])
	switch op {
	case "nop":
		accNew = acc
		nextNew = next + 1
		break
	case "jmp":
		accNew = acc
		nextNew = next + arg
		break
	case "acc":
		accNew = acc + arg
		nextNew = next + 1
		break
	default:
		break
	}
	return accNew, nextNew
}

func DoSilver() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2020, 8, "\n")
	runInstructions := map[int]interface{}{}
	var next int
	var acc int
	for {
		if _, ok := runInstructions[next]; ok {
			return strconv.Itoa(acc)
		} else {
			runInstructions[next] = struct{}{}
		}
		acc, next = handleInstruction(input, acc, next)
	}
	return strconv.Itoa(len(input))
}

func runUntilLoopCountReached(instructions []string, change int, loopCount int) (acc int, cleanTerminate bool) {
	op, _ := parseInstruction(instructions[change])
	if op == "nop" {
		instructions[change] = strings.Replace(instructions[change], "nop", "jmp", 1)
	} else if op == "jmp" {
		instructions[change] = strings.Replace(instructions[change], "jmp", "nop", 1)
	}
	runInstructions := map[int]int{}
	var next int
	for {
		if next == len(instructions) {
			fmt.Println("Clean terminate")
			cleanTerminate = true
			return acc, cleanTerminate
		}
		if run, _ := runInstructions[next]; run == loopCount {
			return -1, false
		} else {
			runInstructions[next]++
		}
		acc, next = handleInstruction(instructions, acc, next)
	}
	return -1, false
}

func DoGold() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2020, 8, "\n")
	for i, _ := range input {
		instructions := make([]string, len(input))
		copy(instructions, input)
		acc, cleanTerminate := runUntilLoopCountReached(instructions, i, 1)
		if cleanTerminate {
			return strconv.Itoa(acc)
		}
	}
	return "Fuck it"
}
