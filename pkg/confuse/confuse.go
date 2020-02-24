package confuse

import (
	"math"
	"math/rand"
	"strings"

	"github.com/JayceChant/drivel/pkg/fenci"

)

// Trans make the words in (Chinese) text confusing by shuffling character order.
// n is the minimal length of the words to be processed,
// which means, words that length < n will be skipped.
func Trans(text string, n int) string {
	if n < 2 {
		n = 2
	}
	nn := n * 2 // Chinese character take at least twice room in []byte

	words := fenci.Split(text)
	for i := range words {
		if len(words[i]) < nn {
			continue // early prune
		}

		words[i] = shuffle(words[i], n)
	}
	return strings.Join(words, "")
}

func shuffle(text string, n int) string {
	chars := []rune(text)
	l := len(chars)
	if l < n {
		return text
	}

	var shuffled string
	for {
		shuffled = string(shuffleImpl(chars, l))
		if shuffled != text {
			break
		}
	}

	return shuffled
}

func shuffleImpl(chars []rune, l int) []rune {
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
