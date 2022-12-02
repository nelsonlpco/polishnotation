package domain

import "fmt"

type Element rune

var ErrorInvalidOperator = "%v is an invalid operator"

var operators = []Element{'+', '-', '*', '/', '('}

func (e Element) IsAnOperator() bool {
	for _, op := range operators {
		if op == e {
			return true
		}
	}

	return false
}

func (e Element) IsAnOperand() bool {
	return e >= 'a' && e <= 'z' || e >= 'A' && e <= 'Z'
}

func GetOperatorPriority(element Element) (uint16, error) {
	switch element {
	case '(':
		return 1, nil
	case '+', '-':
		return 2, nil
	case '*', '/':
		return 3, nil
	}
	return 0, fmt.Errorf(ErrorInvalidOperator, element)
}
