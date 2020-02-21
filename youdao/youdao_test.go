package youdao_test

import (
	"github.com/quii/learn-go-with-tests/mytest/TermDict"
	"github.com/quii/learn-go-with-tests/mytest/TermDict/youdao"
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

func checkWordSpell(t *testing.T, word *TermDict.Word, spell string) {
	t.Helper()
	if word.Spell != spell{
		t.Errorf("got result for %q, expect %q", word.Spell, spell)
	}
}
