package lab2

import (
	"fmt"
	"regexp"
	"strings"
)

func StringToArray(inputS string) []string {
	re := regexp.MustCompile(`([0-9]\.*[0-9]*)+|[+*^/-]`)
	return re.FindAllString(inputS, -1)
}

func PostfixToPrefix(inputStr string) (string, error) {
	operator := regexp.MustCompile("[+*^/-]")
	badSigns, _ := regexp.MatchString(`[!@#$%&;'\/,]`, inputStr)
	if badSigns || len(inputStr) == 0 {
		return "", fmt.Errorf("invalid input expression")
	}

	var input = StringToArray(inputStr)

	operandCounter := 0
	operatorCounter := 0

	for _, e := range input {
		if operator.MatchString(e) {
			operatorCounter++
		} else {
			operandCounter++
		}
	}

	if operandCounter-1 > operatorCounter {
		return "", fmt.Errorf("error: too many operands")
	} else if operandCounter-1 < operatorCounter {
		return "", fmt.Errorf("error: too many operators")
	}

	stack := make([][]string, 0)
	for _, oper := range input {
		if !operator.MatchString(oper) {
			stack = append(stack, []string{oper})
		} else {
			op1 := stack[len(stack)-1]
			op2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			newLastElement := append([]string{oper}, op2...)
			newLastElement = append(newLastElement, op1...)
			stack = append(stack, newLastElement)
		}
	}
	return strings.Join(stack[0], " "), nil
}
