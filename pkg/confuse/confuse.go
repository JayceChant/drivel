package confuse

import (
	"strings"

	"github.com/JayceChant/drivel/pkg/martian"
	"github.com/JayceChant/drivel/pkg/messup"

)

// Trans make the words in (Chinese) text confusing by shuffling character order.
func Trans(text string, marker string, wordSegment bool, useMartian bool) string {
	sections := strings.Split(text, marker)
	c := len(sections)
	for i := 1; i < c; i = i + 2 {
		sections[i] = messup.Trans(sections[i], wordSegment)
		if useMartian {
			sections[i] = martian.Trans(sections[i])
		}
	}
	return strings.Join(sections, "")
}
