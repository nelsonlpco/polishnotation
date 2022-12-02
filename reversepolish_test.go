package reversepolish

import (
	"github.com/nelsonlpco/reversepolish/internal/domain"
	"reflect"
	"testing"
)

func TestNewReversePolishManager(t *testing.T) {
	manager := NewReversePolishManager()

	if manager == nil {
		t.Error("expected to create an new instance of reversepolish.Manager")
	}
}

func Test_postFixExpression_valid_operations(t *testing.T) {
	tests := []struct {
		name           string
		mathExpression string
		want           string
	}{
		{
			name:           "2+2*3 postfix to ABC*+",
			mathExpression: "2+2*3",
			want:           "ABC*+",
		},
		{
			name:           "2-4/2*3 postfix to ABC/D*-",
			mathExpression: "2-4/2*3",
			want:           "ABC/D*-",
		},
		{
			name:           "2+2*(3*2) postfix to ABCD**+",
			mathExpression: "2+2*(3*2)",
			want:           "ABCD**+",
		},
		{
			name:           "2+2*(3*2)/2 postfix to ABCD**E/+",
			mathExpression: "2+2*(3*2)/2",
			want:           "ABCD**E/+",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			expectedPostFixedExpression := []domain.Element(test.want)
			expression := domain.NewExpression(test.mathExpression)
			postFixedExpression := postFixExpression(expression.Simplified())

			if !reflect.DeepEqual(expectedPostFixedExpression, postFixedExpression) {
				t.Errorf("want %v equals to %v", string(expectedPostFixedExpression), string(postFixedExpression))
			}
		})
	}

}

func TestManager_Resolve(t *testing.T) {
	tests := []struct {
		name           string
		mathExpression string
		want           float32
	}{
		{
			name:           "2+2*3 = 8",
			mathExpression: "2+2*3",
			want:           float32(8),
		},
		{
			name:           "2-4/2*3 = -4",
			mathExpression: "2-4/2*3",
			want:           float32(-4),
		},
		{
			name:           "2+2*(3*2) = 14",
			mathExpression: "2+2*(3*2)",
			want:           float32(14),
		},
		{
			name:           "2+2*(3*2)/2 = 8",
			mathExpression: "2+2*(3*2)/2",
			want:           float32(8),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			manager := NewReversePolishManager()
			result := manager.Resolve(test.mathExpression)

			if result != test.want {
				t.Errorf("want %v = %v, but %v", test.mathExpression, test.want, result)
			}
		})
	}
}
