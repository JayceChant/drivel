package fenci

import (
	"strings"
	"testing"
)

var (
	splitCases = []struct {
		text string
		want string
	}{
		{"我来到北京清华大学", "我/来到/北京/清华大学"},
		{"他来到了网易杭研大厦", "他/来到/了/网易/杭研/大厦"},
		{"小明硕士毕业于中国科学院计算所，后在日本京都大学深造", "小明/硕士/毕业/于/中国科学院/计算所/，/后/在/日本京都大学/深造"},
		// {"", ""},
		// {"", ""},
		// {"", ""},
		// {"", ""},
	}
)

func TestSplit(t *testing.T) {
	for _, tt := range splitCases {
		t.Run(string([]rune(tt.text)[:2]), func(t *testing.T) {
			if got := strings.Join(Split(tt.text), "/"); got != tt.want {
				t.Errorf("Split() = %v, want %v", got, tt.want)
			}
		})
	}
}
