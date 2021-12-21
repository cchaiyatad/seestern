package gen

import (
	"fmt"
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func init() {
	setSeed()
}

func setSeed() {
	rand.Seed(time.Now().UnixNano())
}

// [,)
func GenInt(min, max int) int {
	validateMin, validateMax, sign := genNumValidate(min, max)
	return (rand.Intn(validateMax-validateMin) + validateMin) * sign
}

func genNumValidate(min, max int) (int, int, int) {
	if min >= max {
		max = 100 + min
	}

	if max <= 0 && min <= 0 {
		return -1 * max, -1 * min, -1
	}

	return min, max, 1
}

// [,)
func GenDouble(min, max float64) float64 {
	if min >= max {
		max = 100 + min
	}
	return min + rand.Float64()*(max-min)
}

func GenBoolean() bool {
	return rand.Intn(2) == 1
}

func GenString(length int, prefix, suffix string) string {
	return fmt.Sprintf("%s%s%s", prefix, randomString(length), suffix)
}

func randomString(length int) string {
	if length <= 0 {
		length = 1
	}
	randLength := rand.Intn(length)

	b := make([]rune, randLength)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
