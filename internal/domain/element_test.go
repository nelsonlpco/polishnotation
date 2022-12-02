package domain

import "testing"

func TestElement_IsAnOperand(t *testing.T) {
	t.Run("should returns true when receive valid operands", func(t *testing.T) {
		validOperands := []Element{'A', 'b', 'Z', 'w', 'v', 'X', 'Y', 'c'}
		for _, op := range validOperands {
			if op.IsAnOperand() != true {
				t.Errorf("want %s is a valid operand", string(op))
			}
		}
	})

	t.Run("should returns false when receive invalid operands", func(t *testing.T) {
		invalidOperands := []Element{'1', '+', '-', '*', '(', ')'}
		for _, op := range invalidOperands {
			if op.IsAnOperand() != false {
				t.Errorf("want %s is a invalid operand", string(op))
			}
		}
	})
}

func TestElement_IsAnOperator(t *testing.T) {
	t.Run("should returns true when receive valid operators", func(t *testing.T) {
		validOperators := []Element{'+', '-', '*', '/', '('}
		for _, op := range validOperators {
			if op.IsAnOperator() != true {
				t.Errorf("want %s is a valid operator", string(op))
			}
		}
	})

	t.Run("should returns false when receive a valid operator", func(t *testing.T) {
		invalidOperators := []Element{'a', '^', ')', '1'}
		for _, op := range invalidOperators {
			if op.IsAnOperator() != false {
				t.Errorf("want %s is an invalid operator", string(op))
			}
		}
	})
}

func TestElement_GetOperatorPriority(t *testing.T) {
	tests := []struct {
		name    string
		o       Element
		want    uint16
		wantErr bool
	}{
		{
			name:    "operator ( is 1 priority",
			o:       '(',
			want:    uint16(1),
			wantErr: false,
		},
		{
			name:    "operator + is 2 priority",
			o:       '+',
			want:    uint16(2),
			wantErr: false,
		},
		{
			name:    "operator - is 2 priority",
			o:       '-',
			want:    uint16(2),
			wantErr: false,
		},
		{
			name:    "operator / is 3 priority",
			o:       '/',
			want:    uint16(3),
			wantErr: false,
		},
		{
			name:    "operator * is 3 priority",
			o:       '*',
			want:    uint16(3),
			wantErr: false,
		},
		{
			name:    "when receive an invalid operator should returns ErrorInvalidOperator",
			o:       'b',
			want:    uint16(0),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetOperatorPriority(tt.o)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPriority() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetPriority() got = %v, want %v", got, tt.want)
			}
		})
	}
}
