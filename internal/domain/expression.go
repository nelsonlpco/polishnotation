package domain

import (
	"strconv"
	"strings"
)

type Expression struct {
	value      string
	simplified []Element
	hashMap    *HashMap
}

func (e Expression) Expression() string {
	return e.value
}

func (e Expression) Simplified() []Element {
	return e.simplified
}

func (e Expression) HashMap() *HashMap {
	return e.hashMap
}

func NewExpression(expression string) *Expression {
	return &Expression{
		value:      expression,
		simplified: simplifyExpression(expression),
		hashMap:    NewHashMap(expression),
	}
}

func simplifyExpression(expression string) []Element {
	var simplifiedExpression string
	var operand Element = 'A'
	isNumber := false
	splitedExpression := strings.Split(expression, "")

	for _, element := range splitedExpression {
		if _, err := strconv.Atoi(element); err != nil {
			simplifiedExpression += element
			isNumber = false
		} else {
			if !isNumber {
				simplifiedExpression += string(operand)
				operand++
			}
			isNumber = true
		}
	}

	return []Element(simplifiedExpression)
}
