package confuse

import (
	"strings"

	"github.com/JayceChant/drivel/pkg/martian"
	"github.com/JayceChant/drivel/pkg/messup"

)

// Trans make the words in (Chinese) text confusing by shuffling character order.
// n is the minimal length of the words to be processed,
// which means, words that length < n will be skipped.
func Trans(text string, marker string, n int, useMartian bool) string {
	if n < 2 {
		n = 2
	}
	nn := n * 2 // Chinese character take at least twice room in []byte

	//words := fenci.Split(text)
	words := strings.Split(text, marker)
	c := len(words)
	for i := 1; i < c; i = i + 2 {
		if len(words[i]) < nn {
			continue // early prune
		}

		words[i] = messup.Trans(words[i], n)
		if useMartian {
			words[i] = martian.Trans(words[i])
		}
	}
	return strings.Join(words, "")
}
