package fenci

import (
	gjb "github.com/yanyiwu/gojieba"

)

// Split cut text into Chinese words
func Split(text string) []string {
	// adapter that wrap GoJjieba
	eng := gjb.NewJieba()
	defer eng.Free()
	return eng.Cut(text, true)
}
