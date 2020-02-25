package base_test

import (
	"fmt"
	"github.com/Johnny4Fun/TermDict/base"
	"testing"
)

func TestWord_String(t *testing.T) {
	word := &base.Word{
		Spell: "Hello",
		Trans: []base.Translation{
			{"adj", "你好"},
		},
		Examples: []base.Example{
			{"Hello world!",
				"你好世界",
				"Jun",},
		},
		Source: base.Youdao,
	}

	fmt.Println(word)
}
