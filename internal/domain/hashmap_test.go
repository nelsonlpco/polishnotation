package domain

import (
	"reflect"
	"testing"
)

func Test_MakeAnValidHashMap(t *testing.T) {
	simpleExpression := "1+2-38*(233*4409)"

	expectedHashMap := make(map[Element]string)
	expectedHashMap['A'] = "1"
	expectedHashMap['B'] = "2"
	expectedHashMap['C'] = "38"
	expectedHashMap['D'] = "233"
	expectedHashMap['E'] = "4409"

	hashMapManager := NewHashMap(simpleExpression)

	isEqual := reflect.DeepEqual(expectedHashMap, hashMapManager.value)

	if !isEqual {
		t.Errorf("want %v s equals %v", expectedHashMap, hashMapManager.value)
	}
}

func Test_GetHashValue(t *testing.T) {

	tests := []struct {
		name    string
		operand Element
		want    float32
	}{
		{
			name:    "B = 2",
			operand: 'B',
			want:    2,
		},
		{
			name:    "C = 38",
			operand: 'C',
			want:    38,
		},
		{
			name:    "when receive an invalid Operand returns 0",
			operand: 'F',
			want:    0,
		},
		{
			name:    "when find an operand without numbers 0",
			operand: 'G',
			want:    0,
		},
	}

	for _, test := range tests {
		simpleExpression := "1+2-38*(233*4409)"

		t.Run(test.name, func(t *testing.T) {
			hashMapManager := NewHashMap(simpleExpression)
			hashMapManager.value['G'] = ""
			value := hashMapManager.GetValue(test.operand)

			if value != test.want {
				t.Errorf("want %v but receive %v", test.want, value)
			}
		})
	}

}
