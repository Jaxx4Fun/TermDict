package youdao

import (
	"fmt"
	"github.com/Johnny4Fun/TermDict/base"
	"github.com/PuerkitoBio/goquery"
	"io"
	"strings"
)

type Parser interface {
	Parse(r io.Reader) (*base.Word, error)
}

type ParserFunc func(r io.Reader) (*base.Word, error)

func (f ParserFunc) Parse(r io.Reader) (*base.Word, error) {
	return f(r)
}

const (
	domIndexEnglish = iota
	domIndexChinese
	domIndexFrom
)

func ParseHTML(html io.Reader) (*base.Word, error) {
	wd := new(base.Word)

	document, err := goquery.NewDocumentFromReader(html)
	if err != nil {
		return nil, fmt.Errorf("failed to use goquery extract html, %v", err)
	}

	resultContainer := document.Find("div#results-contents").First()

	//translations
	resultContainer.Find("#phrsListTab>.trans-container>ul>li").Each(func(i int, selection *goquery.Selection) {
		text := selection.Text()
		trans := getTranslation(text)

		if trans != nil {
			wd.Trans = append(wd.Trans, *trans)
		}

	})

	// example sentences
	liSlice := resultContainer.Find("#examplesToggle > #bilingual > ul > li")
	size := liSlice.Size()
	wd.Examples = make([]base.Example, size)

	liSlice.Each(func(i int, li *goquery.Selection) {

		sel := li.Find("p")

		eng := trimBlankToken(sel.Eq(domIndexEnglish).Text())
		chn := trimBlankToken(sel.Eq(domIndexChinese).Text())
		from := trimBlankToken(sel.Eq(domIndexFrom).Text())

		example := base.Example{
			Sentence:    eng,
			Translation: chn,
			From:        base.ExampleSrc(from),
		}

		wd.Examples[i] = example
	})

	return wd, nil
}

// split translation into 2 parts: POS & translation
func getTranslation(text string) *base.Translation {

	var pos string
	var trans string
	if base.OptSeparatePosAndTrans {
		var posAndTrans = strings.SplitN(text, " ", 2)

		if l := len(posAndTrans); l >= 2 {
			pos = posAndTrans[0]
			trans = posAndTrans [1]
		} else if l == 1 {
			trans = posAndTrans[0]
		} else {
			return nil
		}
	} else {
		trans = text
	}
	ts := base.Translation{
		POS:   pos,
		Trans: trans,
	}
	return &ts
}

func trimBlankToken(s string) string {
	return strings.TrimFunc(s, func(r rune) bool {
		switch r {
		case '\n', '\t', ' ':
			return true
		}
		return false
	})
}
