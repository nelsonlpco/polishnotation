package domain

import (
	"strconv"
	"strings"
)

type HashMap struct {
	value map[Element]string
}

func (h HashMap) GetValue(i Element) float32 {
	if v, ok := h.value[i]; !ok {
		return 0
	} else {
		result, err := strconv.ParseFloat(v, 32)
		if err != nil {
			return 0
		}

		return float32(result)
	}
}

func NewHashMap(expression string) *HashMap {
	var hashMap = make(map[Element]string)

	var Element Element = 'A'

	splitedExpression := strings.Split(expression, "")
	expressionSize := len(splitedExpression)
	var numberElement string

	for expressionIndex := 0; expressionIndex < expressionSize; expressionIndex++ {
		element := splitedExpression[expressionIndex]
		if _, err := strconv.Atoi(element); err == nil {
			numberElement += element
		} else if numberElement != "" {
			hashMap[Element] = numberElement
			numberElement = ""
			Element++
		}
	}

	if numberElement != "" {
		hashMap[Element] = numberElement
	}

	return &HashMap{value: hashMap}
}
