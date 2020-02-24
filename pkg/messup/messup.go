package messup

import (
	"math"
	"math/rand"

)

// Trans ...
func Trans(text string, n int) string {
	chars := []rune(text)
	l := len(chars)
	if l < n {
		return text
	}

	var shuffled string
	for {
		shuffled = string(messup(chars, l))
		if shuffled != text {
			break
		}
	}

	return shuffled
}

func messup(chars []rune, l int) []rune {
	act := int(math.Sqrt(float64(l - 1))) // swap_times = sqrt(len - 1)
	if act < 1 {
		act = 1
	}
	// swap adjacent characters randomly
	for i := 0; i < act; i++ {
		j := rand.Intn(l - 1)
		chars[j], chars[j+1] = chars[j+1], chars[j]
	}
	return chars
}
