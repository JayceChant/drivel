package confuse

import (
	"fmt"
	"strings"

	"github.com/JayceChant/drivel/pkg/martian"
	"github.com/JayceChant/drivel/pkg/messup"

)

// Trans make the words in (Chinese) text confusing by shuffling character order.
func Trans(text string, marker string, wordSegment bool, useMartian bool) (string, error) {
	sections := strings.Split(text, marker)
	c := len(sections)
	if c%2 == 0 { // if markers are in pairs, split result len should be odd
		return text, fmt.Errorf("target markers ('%s') are not matched in pairs", marker)
	}
	for i := 1; i < c; i = i + 2 {
		sections[i] = messup.Trans(sections[i], wordSegment)
		if useMartian {
			sections[i] = martian.Trans(sections[i])
		}
	}
	return strings.Join(sections, ""), nil
}
