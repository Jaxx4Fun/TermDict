package youdao_test

import (
	"github.com/Johnny4Fun/TermDict/base"
	"github.com/Johnny4Fun/TermDict/youdao"
	"net/http"
	"testing"
)

func TestPageParser(t *testing.T) {
	t.Run("look up \"what\" from youdao and parse it", func(t *testing.T) {
		url := "http://www.youdao.com/w/what"
		resp := mustGetResponse(t, url)
		parser := youdao.ParserFunc(youdao.ParseHTML)
		wd, err := parser.Parse(resp.Body)

		assertNoError(t, err)
		assertTranslationNum(t, wd, 4)
		if len(wd.Examples) <= 0{
			t.Log("got no examples")
		}
	})

}

func assertTranslationNum(t *testing.T, wd *base.Word, want int) {
	t.Helper()
	if len(wd.Trans) != want {
		t.Errorf("explain got %d, want %d", len(wd.Trans), want)
	}
}

func mustGetResponse(t *testing.T, url string) *http.Response {
	t.Helper()
	resp, err := http.Get(url)
	if err != nil {
		t.Fatalf("failed to GET %q, %v", url, err)
	}
	return resp
}
func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("expected no error, but %v", err)
	}
}
