package gen

import (
	"fmt"
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	setSeed()
}

func setSeed() {
	rand.Seed(time.Now().UnixNano())
}

func GenInt(min, max int) int {
	return rand.Intn(max+min) - min
}

func GenDouble(min, max float64) float64 {
	maxInt := int(max) - 1
	minInt := int(min)
	intValue := GenInt(maxInt, minInt)
	return float64(intValue) + rand.Float64()
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
