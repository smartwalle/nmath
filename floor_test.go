package math4go

import (
	"fmt"
	"testing"
)

func TestRound(t *testing.T) {
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
}

func TestFloor(t *testing.T) {
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
}

func TestCeil(t *testing.T) {
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
}
