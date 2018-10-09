package math4go

import (
	"fmt"
	"math"
	"testing"
)

func TestRound(t *testing.T) {
	fmt.Println("===== Round =====")
	var v = 10.0666
	fmt.Println(v, Round(v, 2))
	if Round(v, 2) != 10.07 {
		t.FailNow()
	}

	v = 10.0444
	fmt.Println(v, Round(v, 2))
	if Round(v, 2) != 10.04 {
		t.FailNow()
	}

	fmt.Println(math.Round(-10.9464), Round(-10.9464, 2))
	fmt.Println(math.Round(-10.9464), Round(-10.9444, 2))
	fmt.Println(math.Round(-10.9464), Round(-10.9464, 0))
	fmt.Println(math.Round(-10.4464), Round(-10.4464, 0))
}

func TestFloor(t *testing.T) {
	fmt.Println("===== Floor =====")
	var v = 10.999
	fmt.Println(v, Floor(v, 2))
	if Floor(v, 2) != 10.99 {
		t.FailNow()
	}

	v = 10.994
	fmt.Println(v, Floor(v, 2))
	if Floor(v, 2) != 10.99 {
		t.FailNow()
	}

	v = 10.911
	fmt.Println(v, Floor(v, 2))
	if Floor(v, 2) != 10.91 {
		t.FailNow()
	}

	fmt.Println(math.Floor(-10.1), Floor(-10.111, 2))
	fmt.Println(math.Floor(-10.1), Floor(-10.1, 0))
	fmt.Println(math.Floor(-10.9), Floor(-10.9, 0))
}

func TestCeil(t *testing.T) {
	fmt.Println("===== Ceil =====")
	var v = 10.119
	fmt.Println(v, Ceil(v, 2))
	if Ceil(v, 2) != 10.12 {
		t.FailNow()
	}

	v = 10.111
	fmt.Println(v, Ceil(v, 2))
	if Ceil(v, 2) != 10.12 {
		t.FailNow()
	}

	v = 10.11
	fmt.Println(v, Ceil(v, 2))
	if Ceil(v, 2) != 10.11 {
		t.FailNow()
	}

	fmt.Println(math.Ceil(-10.1), Ceil(-10.121, 2))
	fmt.Println(math.Ceil(-10.1), Ceil(-10.1, 0))
	fmt.Println(math.Ceil(-10.9), Ceil(-10.9, 0))
}
