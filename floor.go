package math4go

import (
	"math"
)

// Round 保留 n 位小数，并对小数部分进行四舍五入
func Round(f float64, n int) float64 {
	pow10n := math.Pow10(n)
	return math.Trunc((f+0.5/pow10n)*pow10n) / pow10n
}

// Floor 保留 n 位小数，并对小数部分进行向下取整
func Floor(f float64, n int) float64 {
	pow10n := math.Pow10(n)
	if f < 0 {
		return (math.Trunc((f*-1+0.9/pow10n)*pow10n) / pow10n) * -1
	}
	return math.Trunc(f*pow10n) / pow10n
}

// Ceil 保留 n 位小数，并对小数部分进行向上取整
func Ceil(f float64, n int) float64 {
	return Floor(f*-1, n) * -1
}
