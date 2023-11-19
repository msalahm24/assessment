package util

import (
	"math/rand"
	"strconv"
	"strings"
)

func GenerateOTP() string {
	numbers := make([]string, 4)

	for i := 0; i < 4; i++ {
		numbers[i] = strconv.Itoa(rand.Intn(9))
	}

	return strings.Join(numbers, "")
}
