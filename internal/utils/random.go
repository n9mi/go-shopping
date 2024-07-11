package utils

import "math/rand"

func RandomNumberWithRange(min, max int) int {
	return rand.Intn(max-min) + min
}
