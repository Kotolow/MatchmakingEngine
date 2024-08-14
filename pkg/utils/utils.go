package utils

import (
	"math"
)

func GetRadius(sDiff, lDiff float64) float64 {
	return math.Sqrt(math.Pow(sDiff, 2) + math.Pow(lDiff, 2))
}
