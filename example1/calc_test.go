package main

import "testing"

func TestAdd(t *testing.T) {
	if ans := Add(1, 2); ans != 3 {
		t.Errorf("1 + 2 expected be 3, but %d got", ans)
	}

	if ans := Add(-10, -20); ans != -30 {
		t.Errorf("-10 + -20 expected be -30, but %d got", ans)
	}

}

func TestMul(t *testing.T) {
	createMulTestCase(t, &calcCase{2, 3, 6})
	createMulTestCase(t, &calcCase{2, -3, -6})
	createMulTestCase(t, &calcCase{2, 0, 0})
}

type calcCase struct {
	A, B, Expected int
}

func createMulTestCase(t *testing.T, c *calcCase) {
	t.Helper()
	if ans := Mul(c.A, c.B); ans != c.Expected {
		t.Fatalf("%d * %d expected %d, but got %d", c.A, c.B, c.Expected, ans)
	}
}
