package messup

import (
	"math"
	"math/rand"
	"strings"

	"github.com/JayceChant/drivel/pkg/fenci"

)

// Trans messup the character order.
// If wordSegment is true, the character swap won't occur between words.
func Trans(text string, wordSegment bool) string {
	if wordSegment {
		words := fenci.Split(text)
		for i := range words {
			words[i] = transImpl(words[i])
		}
		return strings.Join(words, "")
	}

	return transImpl(text)
}

func transImpl(text string) string {
	chars := []rune(text)
	l := len(chars)

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
