package messup

import (
	"math"
	"math/rand"
	"strings"
	"time"
	"unicode"

	"github.com/JayceChant/drivel/pkg/fenci"

)

func init() {
	rand.Seed(time.Now().Unix())
}

// Trans messup the character order.
// If wordSegment is true, the character swap won't occur between words.
func Trans(text string, wordSegment bool) string {
	if len(text) < 2 { // prune
		return text
	}

	var words []string
	sep := ""
	if wordSegment {
		// Chinese word segmentation
		words = fenci.Split(text)
	} else {
		// English word segmentation
		words = strings.Split(text, " ")
		sep = " "
	}

	for i := range words {
		words[i] = transImpl(words[i])
	}

	return strings.Join(words, sep)
}

func transImpl(text string) string {
	chars := []rune(text)
	if len(chars) < 2 { // prune
		return text
	}

	// changed-required-validate will lead to dead-loop if
	// 1. not enough characters allow to move, e.g. 'A----', '## 三'
	// 2. all movable characters are the same, e.g. 'oo', '哈哈哈'
	// TODO: improve the validate
	var shuffled string
	//for {
	messup(chars)
	shuffled = string(chars)
	//if shuffled != text {
	//	break
	//}
	//}

	return shuffled
}

func messup(chars []rune) {
	n := len(chars)
	var lo, hi = 0, 0
	for hi <= n {
		if hi == n || !isLatinOrHan(chars[hi]) {
			if hi-lo >= 2 {
				messupImpl(chars[lo:hi])
			}
			lo = hi + 1
		}
		hi++
	}
}

func messupImpl(chars []rune) {
	n := len(chars)
	act := int(math.Sqrt(float64(n - 1))) // swap_times = sqrt(len - 1)
	if act < 1 {
		return
	}

	// swap adjacent characters randomly
	for i := 0; i < act; i++ {
		left := rand.Intn(n - 1)
		chars[left], chars[left+1] = chars[left+1], chars[left]
	}
}

func isLatinOrHan(r rune) bool {
	return unicode.In(r, unicode.Latin, unicode.Han)
}
