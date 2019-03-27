package _2_factory_method

import "testing"

func compute(factory OperatorFactory, a, b int) int {
	operator := factory.Create()
	operator.SetA(a)
	operator.SetB(b)
	return operator.Result()
}
func TestOperator(t *testing.T) {

	var factory OperatorFactory

	factory = PlusFactory{}

	if compute(factory, 2, 1) != 3 {
		t.Fatal("error with plus")
	}

	factory = MinusFactory{}
	if compute(factory, 2, 1) != 1 {
		t.Fatal("error with minus")
	}
}
