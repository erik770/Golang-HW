package calc

import (
	"errors"
	"github.com/erik770/Golang-HW/pkg/stack"
	"strconv"
	"strings"
)

type Priority int

const (
	lowest Priority = -2
	low    Priority = 1
	medium Priority = 2
	high   Priority = 3
)

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
	operands := map[string]struct{}{
		"(": {},
		")": {},
		"+": {},
		"-": {},
		"*": {},
		"/": {},
	}
	if _, isExist := operands[element]; !isExist {
		return false
	}
	return true
}

func getOperatorPriority(operand string) Priority {
	operandsPriority := map[string]Priority{
		"(": lowest,
		"/": low,
		"*": low,
		"+": medium,
		"-": medium,
		")": high,
	}
	return operandsPriority[operand]
}

func checkExpression(expression string) error {
	var (
		openingBracket int
		closingBracket int
	)
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

	var resExpression []string
	sliceExpr := strings.Split(expression, "")

	for i := 0; i < len(sliceExpr); i++ {
		switch {
		case isOperator(sliceExpr[i]):
			resExpression = append(resExpression, sliceExpr[i])
		case isDigit(sliceExpr[i]):
			var value string
			for ; isDigit(sliceExpr[i]); i++ {
				value += sliceExpr[i]
			}
			resExpression = append(resExpression, value)
			i--
		}
	}

	return resExpression
}

func operatorHandler(operator string, operatorsStack, operandsStack *stack.Stack, isLast bool) (res string, err error) {
	currentPriority, lastPriority := getOperatorPriority(operator), getOperatorPriority(operatorsStack.Top())
	for currentPriority >= lastPriority && currentPriority+lastPriority > 0 {
		switch {
		case operandsStack.GetLen() == 1 && isLast:
			return operandsStack.Pop(), nil
		case operatorsStack.Top() == "(":
			operatorsStack.Pop()
			currentPriority = 0
			continue
		}
		rightOperand, leftOperand := operandsStack.Pop(), operandsStack.Pop()
		exp, err := binaryEvaluate(leftOperand, rightOperand, operatorsStack.Pop())
		if err != nil {
			return "", err
		}
		operandsStack.Push(exp)
		lastPriority = getOperatorPriority(operatorsStack.Top())
	}
	if currentPriority != 0 {
		operatorsStack.Push(operator)
		return "", err
	}
	return "", err
}

func evaluate(expression []string) (string, error) {
	var (
		operatorsStack stack.Stack
		operandsStack  stack.Stack
	)

	operatorsStack.Push(expression[0])
	expression = expression[1:]
	for i := range expression {
		switch {
		case isOperator(expression[i]):
			res, err := operatorHandler(expression[i], &operatorsStack, &operandsStack, len(expression)-1 == i)
			if err != nil || res != "" {
				return res, err
			}
		case isDigit(expression[i]):
			operandsStack.Push(expression[i])
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
