package even

import (
	"testing"
)

func TestEven(t *testing.T) {
	if !Even(10) {
		t.Log("10 必须是偶数")
		t.Fail()
	}

	if Even(7) {
		t.Log("7 不是偶数")
		t.Fail()
	}
}

func TestOdd(t *testing.T) {
	if !Odd(11) {
		t.Log("11 must be odd!")
		t.Fail()
	}

	if Odd(10) {
		t.Log("10 不是奇数")
		t.Fail()
	}
}
