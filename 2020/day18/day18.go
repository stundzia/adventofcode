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
	for i := 0; i < len(expression); i++ {
		switch string(expression[i]) {
		case " ":
			continue
		case "(":
			parenth := ""
			closing := 0
			i++
		mainFor:
			for ; closing < 1; i++ {
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
			if operation == "+" {
				result += evaluateExpression(parenth)
			}
			if operation == "*" {
				result *= evaluateExpression(parenth)
			}
			operation = ""
			break
		case "*":
			operation = "*"
		case "+":
			operation = "+"
		default:
			num, _ := strconv.Atoi(string(expression[i]))
			if operation == "+" {
				result += num
			}
			if operation == "*" {
				result *= num
			}
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
				for t := i + 3; t < len(expression); t++ {
					if string(expression[t]) == "(" {
						closing--
					}
					if string(expression[t]) == ")" {
						closing++
					}

					if closing == 1 {
						afterIndex = t
						if afterIndex >= len(expression) {
							aft = ""
						} else {
							aft = string(expression[afterIndex])
						}
						break
					}
				}
			}
			if bf == ")" {
				closing := 0
				for t := i - 3; t >= 0; t-- {
					if string(expression[t]) == "(" {
						closing++
					}
					if string(expression[t]) == ")" {
						closing--
					}
					if closing == 1 {
						beforeIndex = t
						bf = string(expression[beforeIndex])
						break
					}
				}
			}

			beforeParenthesis := false
			afterParenthesis := false
			if beforeIndex-1 >= 0 && string(expression[beforeIndex-1]) == "(" {
				beforeParenthesis = true
			}
			if afterIndex+1 < len(expression) && string(expression[afterIndex+1]) == ")" {
				afterParenthesis = true
			}
			if !afterParenthesis || !beforeParenthesis {
				if beforeIndex-1 < 0 {
					expression = "(" + expression
				} else {
					expression = expression[:beforeIndex] + "(" + expression[beforeIndex:]
				}
				if afterIndex+2 >= len(expression) {
					expression = expression + ")"
				} else {
					expression = expression[:afterIndex+2] + ")" + expression[afterIndex+2:]
				}
				done = true
				break
			}
		}
	}
	if done {
		return parenthesizeAddition(expression)
	}
	return expression
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
		exp = parenthesizeAddition(exp)
		sum += evaluateExpression(exp)
	}
	return fmt.Sprintf("%d", sum)
}
