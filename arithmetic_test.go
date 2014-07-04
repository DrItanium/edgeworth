package edgeworth

import "testing"

func Test_TooManyArithmeticOperations_1(t *testing.T) {
	/* we need to make sure that we have at most 16 operations */
	if ArithmeticOperationCount > ImmediateFormOffset {
		t.Errorf("Too many arithmetic operations defined. Max: %d, Actual: %d", ImmediateFormOffset, ArithmeticOperationCount)
	} else if ArithmeticOperationCount == ImmediateFormOffset {
		t.Log("There are no free arithmetic operation slots left!")
	} else {
		t.Logf("There are %d free arithmetic opertion slots available.", ImmediateFormOffset-ArithmeticOperationCount)
	}
}
