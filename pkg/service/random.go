package service

import (
	"math"
	"math/rand"
	"time"
)

func randomInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max - min + 1) + min
}

func randomFloat() float64 {
	return math.Round(4 * rand.Float64() * 100) / 100
}