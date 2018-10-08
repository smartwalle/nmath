package math4go

import (
	"fmt"
	"testing"
)

func TestRound(t *testing.T) {
	var a = 10.0666
	fmt.Println(Round(a, 2))

	a = 10.0444
	fmt.Println(Round(a, 2))
}
