package day18

import (
	"fmt"
	"github.com/stundzia/adventofcode/utils"
	"strconv"
)


func evaluateExpression(expression string) int {
	//(9 * (7 * 7 * 2 + 3 + 8 + 8) + 5 + 5) + 9 * 6
	operation := "+"
	result := 0
	for i:=0; i < len(expression);i++ {
		switch string(expression[i]) {
		case " ":
			continue
		case "(":
			parenth := ""
			closing := 0
			i++
			mainFor:
			for ;closing < 1; i++ {
				switch string(expression[i]) {
				case "(":
					closing--
					parenth += string(expression[i])
					break
				case ")":
					if closing == 0 {
						break mainFor
					}
					closing++
					parenth += string(expression[i])
					break
				default:
					parenth += string(expression[i])
					break
				}
				if closing == 1 {
					break
				}
			}
			if operation == "+" {result += evaluateExpression(parenth)}
			if operation == "*" {result *= evaluateExpression(parenth)}
			operation = ""
			break
		case "*":
			operation = "*"
		case "+":
			operation = "+"
		default:
			num, _ := strconv.Atoi(string(expression[i]))
			if operation == "+" {result += num}
			if operation == "*" {result *= num}
			operation = ""
		}
	}
	return result
}


func DoSilver() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2020, 18, "\n")
	sum := 0
	for _, exp := range input {
		sum += evaluateExpression(exp)
	}
	return fmt.Sprintf("%d", sum)
}

func DoGold() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2020, 18, "\n")
	return fmt.Sprintf("%d", len(input))
}
