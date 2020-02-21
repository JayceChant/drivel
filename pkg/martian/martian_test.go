package martian

import (
	"testing"
)

var (
	transCases = []struct {
		text string
		want string
	}{
		{"测试", "恻鉽"},
		{"这是一句火星文", "適湜①呴焱暒妏"},
		{"让我们尝试更长的句子，譬如加上标点符号。", "讓莪們甞鉽浭萇哋呴ふ，譬洳咖仩摽點苻呺。"},
		{"天地玄黄宇宙洪荒日月盈昃辰宿列张寒来暑往秋收冬藏闰余成岁律吕调阳云腾致雨露结为霜", "兲哋玆曂荢宙葓巟ㄖ仴盁昃宸蹜烮張寒唻濐暀偢荍笗蔵潤悇荿嵗侓焒蜩陽囩駦臸雨蕗結潙灀"},
		{"人之初性本善性相近习相远苟不教性乃迁教之道贵以专", "亾と初悻夲僐悻楿菦習楿逺苟芣嘋悻釢迁嘋と檤貴姒抟"},
		{"天对地雨对风大陆对长空山花对海树赤日对苍穹雷隐隐雾蒙蒙日下对天中", "兲怼哋雨怼颩汏陸怼萇涳屾埖怼嗨樹哧ㄖ怼芲穹檑陻陻霚懞懞ㄖ芐怼兲狆"},
		{"最罪尊遵昨左佐柞做作坐座", "朂嶵澊噂葃咗佐柞莋莋唑蓙"},
	}
)

func TestTrans(t *testing.T) {
	for _, tt := range transCases {
		t.Run(string([]rune(tt.text)[:2]), func(t *testing.T) {
			if got := Trans(tt.text); got != tt.want {
				t.Errorf("Trans() = %v, want %v", got, tt.want)
			}
		})
	}
}
