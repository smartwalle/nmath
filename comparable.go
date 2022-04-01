package math4go

import "math"

type Comparable interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func Max[T Comparable](x, y T) T {
	return T(math.Max(float64(x), float64(y)))
}

func Min[T Comparable](x, y T) T {
	return T(math.Min(float64(x), float64(y)))
}
