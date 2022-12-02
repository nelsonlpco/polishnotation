package domain

import (
	"reflect"
	"testing"
)

func TestNewExpression(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name string
		args args
		want *Expression
	}{
		{
			name: "simplify 1+222-231 to A+B-C",
			args: args{expression: "1+222-231"},
			want: &Expression{
				value:      "1+222-231",
				simplified: []Element("A+B-C"),
				hashMap: &HashMap{
					value: map[Element]string{
						'A': "1",
						'B': "222",
						'C': "231",
					},
				},
			},
		},
		{
			name: "simplify 2-4/2*3*(3+3)/2 to A-B/C*D*(E+F)/G",
			args: args{expression: "2-4/2*3*(3+3)/2"},
			want: &Expression{
				value:      "2-4/2*3*(3+3)/2",
				simplified: []Element("A-B/C*D*(E+F)/G"),
				hashMap: &HashMap{
					value: map[Element]string{
						'A': "2",
						'B': "4",
						'C': "2",
						'D': "3",
						'E': "3",
						'F': "3",
						'G': "2",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewExpression(tt.args.expression); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewExpression() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExpression_getters(t *testing.T) {
	expression := &Expression{
		value:      "1+222-231",
		simplified: []Element("A+B-C"),
		hashMap: &HashMap{
			value: map[Element]string{
				'A': "1",
				'B': "222",
				'C': "231",
			},
		},
	}

	if expression.Expression() != "1+222-231" {
		t.Errorf("want to get expression value")
	}

	if !reflect.DeepEqual(expression.Simplified(), []Element("A+B-C")) {
		t.Errorf("want to get simplified expression value")
	}

	if !reflect.DeepEqual(expression.HashMap(), &HashMap{
		value: map[Element]string{
			'A': "1",
			'B': "222",
			'C': "231",
		},
	}) {
		t.Errorf("want to get hashmap value")
	}
}
