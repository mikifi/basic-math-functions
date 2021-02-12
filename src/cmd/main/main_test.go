package main

import (
	rpn "internal/reverspolishnotation"
	"math"
	"strings"
	"testing"
)

const float64EqualityThreshold = 1e-9

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

func ErrorContains(out error, want string) bool {
	if out == nil {
		return want == ""
	}
	if want == "" {
		return false
	}
	return strings.Contains(out.Error(), want)
}

var rpnTests = []struct {
	expr     string  // input
	expected float64 // expected result
}{
	{"1 1 +", 2},
	{"3 2 *", 6},
	{"7 3 -", 4},
	{"9 3 /", 3},
}

func TestRpn(t *testing.T) {
	for _, tt := range rpnTests {
		actual, _ := rpn.ReversePolishNotation(tt.expr)
		if !almostEqual(actual, tt.expected) {
			t.Errorf("ReversePolishNotation(%s): expected %f, actual %f", tt.expr, tt.expected, actual)
		}
	}
}

func TestInvalid(t *testing.T) {
	_, err := rpn.ReversePolishNotation("abc")
	if !ErrorContains(err, "invalid token") {
		t.Errorf("unexpected error: %v", err)
	}
}
