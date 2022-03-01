package calc

import (
	"errors"
	"log"
	"strconv"
)

const (
	lowestPriority = -2
	lowPriority    = iota
	mediumPriority
	highPriority
)

type stack struct {
	data   []string
	length int
}

func (s *stack) push(element string) {
	(*s).length++
	(*s).data = append((*s).data, element)
}

func (s *stack) pop() string {
	if (*s).length == 0 {
		log.Fatal("pop empty stack")
	}
	returnValue := (*s).data[(*s).length-1]
	(*s).data = (*s).data[:(*s).length-1]
	(*s).length--
	return returnValue
}

func (s stack) top() string {
	if s.length == 0 {
		log.Fatal("top empty stack")
	}
	return s.data[s.length-1]
}

func Calculate(expression string) (string, error) {
	err := checkExpression(expression)
	if err != nil {
		return "", err
	}
	formattedExpression := reformatExpression(expression)

	res, err := evaluate(formattedExpression)
	if err != nil {
		return "", err
	}
	return res, nil
}

func isDigit(element string) bool {
	if _, err := strconv.Atoi(element); err == nil {
		return true
	}
	return false
}

func isOperator(element string) bool {
	operands := map[string]bool{
		"(": true,
		")": true,
		"+": true,
		"-": true,
		"*": true,
		"/": true,
	}
	if _, isExist := operands[element]; !isExist {
		return false
	}
	return true
}

func getOperatorPriority(operand string) int {
	operandsPriority := map[string]int{
		"(": lowestPriority,
		"/": lowPriority,
		"*": lowPriority,
		"+": mediumPriority,
		"-": mediumPriority,
		")": highPriority,
	}
	return operandsPriority[operand]
}

func checkExpression(expression string) error {
	var openingBracket, closingBracket int
	for i := range expression {
		if !isDigit(string(expression[i])) && !isOperator(string(expression[i])) && expression[i] != ' ' {
			return errors.New("bad symbol in expression")
		}
		switch {
		case expression[i] == '(':
			openingBracket++
		case expression[i] == ')':
			closingBracket++
		}
	}
	if openingBracket != closingBracket {
		return errors.New("expression is incorrect")
	}

	return nil
}

func reformatExpression(expression string) []string {
	expression = "(" + expression + ")"
	//expression = expression + ")"
	var resExpression []string
	for i := 0; i < len(expression); i++ {
		switch {
		case isOperator(string(expression[i])):
			resExpression = append(resExpression, string(expression[i]))
		case isDigit(string(expression[i])):
			var value string
			for ; isDigit(string(expression[i])); i++ {
				value += string(expression[i])
			}
			resExpression = append(resExpression, value)
			i--
		}
	}

	return resExpression
}

func evaluate(expression []string) (string, error) {
	var operatorsStack, operandsStack stack

	operatorsStack.push(expression[0])
	expression = expression[1:]
	for i := range expression {
		switch {
		case isOperator(expression[i]):
			currentPriority, lastPriority := getOperatorPriority(expression[i]), getOperatorPriority(operatorsStack.top())
			for currentPriority >= lastPriority && currentPriority+lastPriority > 0 {
				switch {
				case operandsStack.length == 1 && len(expression)-1 == i:
					return operandsStack.pop(), nil
				case operatorsStack.top() == "(":
					operatorsStack.pop()
					currentPriority = 0
					continue
				}
				rightOperand, leftOperand := operandsStack.pop(), operandsStack.pop()
				exp, err := binaryEvaluate(leftOperand, rightOperand, operatorsStack.pop())
				if err != nil {
					return "", err
				}
				operandsStack.push(exp)
				lastPriority = getOperatorPriority(operatorsStack.top())
			}
			if currentPriority != 0 {
				operatorsStack.push(expression[i])
			}
		case isDigit(expression[i]):
			operandsStack.push(expression[i])
		}
	}

	return "", errors.New("problems with expression")
}

func binaryEvaluate(leftOperand, rightOperand, operator string) (string, error) {
	valLeft, _ := strconv.Atoi(leftOperand)
	valRight, _ := strconv.Atoi(rightOperand)
	switch operator {
	case "+":
		return strconv.Itoa(valLeft + valRight), nil
	case "-":
		return strconv.Itoa(valLeft - valRight), nil
	case "*":
		return strconv.Itoa(valLeft * valRight), nil
	case "/":
		if valRight == 0 {
			return "", errors.New("zero division")
		}
		return strconv.Itoa(valLeft / valRight), nil
	default:
		return "", errors.New("unknown symbol")
	}
}
