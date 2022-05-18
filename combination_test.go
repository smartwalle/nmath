package nmath_test

import (
	"github.com/smartwalle/nmath"
	"testing"
)

func TestCombination(t *testing.T) {
	//var l1 = []interface{}{"A", "K", "Q", "J", "10", "9", "8", "7", "6", "5", "4", "3", "2"}
	//var l2 = []interface{}{"♠", "♥", "♦", "♣"}
	var l1 = []string{"1", "2", "3", "4"}
	var l2 = []string{"A", "B", "C", "D"}
	var l3 = []string{"a", "b"}

	var p = [][]string{l1, l2, l3}

	t.Log(nmath.Combination[string](p))
}
