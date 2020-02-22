package youdao_test

import (
	"github.com/Johnny4Fun/TermDict/base"
	"github.com/Johnny4Fun/TermDict/youdao"
	"testing"
)

func TestYoudaoDict(t *testing.T) {
	yd := youdao.NewDict()
	t.Run("test look a word", func(t *testing.T) {
		word := "fool"
		got := yd.LookUp(word)

		checkWordSpell(t, got, word)
	})
}

func checkWordSpell(t *testing.T, word *base.Word, spell string) {
	t.Helper()
	if word.Spell != spell{
		t.Errorf("got result for %q, expect %q", word.Spell, spell)
	}
}
