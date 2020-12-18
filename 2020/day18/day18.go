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


func parenthesizeAddition(expression string) string {
	done := false
	for i, e := range expression {
		if string(e) == "+" {
			beforeIndex := i - 2
			afterIndex := i + 2
			bf := string(expression[beforeIndex])
			aft := string(expression[afterIndex])

			if aft == "(" {
				closing := 0
				for t := i+2; t < len(expression); t++ {
					if string(expression[t]) == "(" {
						closing--
					}
					if string(expression[t]) == ")" {
						closing++
					}
					if closing == 1 {
						afterIndex = t
						fmt.Println(aft)
						aft = string(expression[afterIndex])
						fmt.Println(aft)
						break
					}
				}
			}

			beforeParenthesis := false
			afterParenthesis := false
			if i - 3 >= 0 && string(expression[i - 3]) == "(" {
				beforeParenthesis = true
			}
			if afterIndex + 1 < len(expression) && string(expression[afterIndex + 1]) == ")" {
				afterParenthesis = true
			}
			if bf != "(" && bf != ")" && aft != "(" && (!afterParenthesis || !beforeParenthesis) {
				expression = expression[:i - 2] + "(" + expression[i - 2:]
				expression = expression[:afterIndex + 2] + ")" + expression[afterIndex + 2:]
				done = true
				break
			}
		}
	}
	if done {
		return parenthesizeAddition(expression)
	}
	fmt.Println(expression)
	return expression
}

func evaluateExpression2(expression string) int {
	expression = parenthesizeAddition(expression)
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

func parenthesiseExpression(expression string) string {
	return "(" + expression + ")"
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
	sum := 0
	for _, exp := range input {
		sum += evaluateExpression2(exp)
	}

	fmt.Println(evaluateExpression2("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"))
	fmt.Println(parenthesizeAddition("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"))
	// Not 27531981444417
	return fmt.Sprintf("%d", sum)
}
