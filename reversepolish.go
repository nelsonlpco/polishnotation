package reversepolish

import (
	"fmt"
	"github.com/nelsonlpco/reversepolish/internal/domain"
)

type Manager struct {
	expression *domain.Expression
}

func (m *Manager) Resolve(mathExpression string) float32 {
	expression := domain.NewExpression(mathExpression)
	postFixedExpression := postFixExpression(expression.Simplified())

	fmt.Print(string(postFixedExpression))

	operationStack := domain.NewStack[float32]()
	for _, element := range postFixedExpression {
		if element.IsAnOperand() {
			operationStack.Push(expression.HashMap().GetValue(element))
		} else {
			y := operationStack.Pop()
			x := operationStack.Pop()

			switch element {
			case '+':
				operationStack.Push(x + y)
			case '-':
				operationStack.Push(x - y)
			case '*':
				operationStack.Push(x * y)
			case '/':
				operationStack.Push(x / y)
			}
		}
	}

	return operationStack.Pop()
}

func NewReversePolishManager() *Manager {
	return &Manager{}
}

func postFixExpression(simplifiedExpression []domain.Element) []domain.Element {
	var postFixedExpression string
	var currentOperatorPriority uint16
	priorityStack := domain.NewStack[domain.PriorityModel]()

	for _, element := range simplifiedExpression {
		if element.IsAnOperand() {
			postFixedExpression += string(element)
		} else if element.IsAnOperator() {
			currentOperatorPriority, _ = domain.GetOperatorPriority(element)

			if priorityStack.Length() > 0 && priorityStack.Value().Priority >= currentOperatorPriority && element != '(' {
				current := priorityStack.Pop()
				postFixedExpression += string(current.Element)
			}

			priorityStack.Push(domain.PriorityModel{Priority: currentOperatorPriority, Element: element})
		} else if element == ')' {
			item := priorityStack.Pop()
			for item.Element != '(' {
				postFixedExpression += string(item.Element)
				item = priorityStack.Pop()
			}
		}
	}

	for priorityStack.Length() > 0 {
		postFixedExpression += string(priorityStack.Pop().Element)
	}

	return []domain.Element(postFixedExpression)
}
