package fenci

import (
	gjb "github.com/yanyiwu/gojieba"

)

// Split cut text into Chinese words
// at current revision, it's just an adapter that wrap GoJjieba
func Split(text string) []string {
	// wrap GoJjieba
	eng := gjb.NewJieba()
	defer eng.Free()
	return eng.Cut(text, true)
}
