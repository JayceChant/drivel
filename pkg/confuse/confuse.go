package confuse

import (
	"math/rand"
	"strings"

	"github.com/JayceChant/drivel/pkg/fenci"

)

// Trans make the words in (Chinese) text confusing by shuffle character order
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
	act := l / 4 // TODO: chose a better swap times
	if act < 1 {
		act = 1
	}
	for i := 0; i < act; i++ {
		var j int
		for {
			j = rand.Intn(l)
			if i != j {
				break
			}
		}
		chars[i], chars[j] = chars[j], chars[i]
	}
	return chars
}
