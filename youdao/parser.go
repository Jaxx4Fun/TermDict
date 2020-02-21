package youdao

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/quii/learn-go-with-tests/mytest/TermDict"
	"io"
	"strings"
)

type Parser struct {
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(r io.Reader) (*TermDict.Word, error) {
	wd := new(TermDict.Word)

	document, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, fmt.Errorf("failed to use goquery extract html, %v", err)
	}

	resultContainer := document.Find("div#results-contents").First()

	//translations
	resultContainer.Find("#phrsListTab>.trans-container>ul>li").Each(func(i int, selection *goquery.Selection) {
		text := selection.Text()
		trans := splitTranslation(text)

		if trans != nil {
			wd.Trans = append(wd.Trans, *trans)
		}

	})

	// example sentences
	//resultContainer.Find("div.phrsListTab > ")

	return wd, nil

}

// split translation into 2 parts: POS & translation
func splitTranslation(text string) *TermDict.Translation {

	var pos string
	var trans string

	transs := strings.SplitN(text, " ", 1)

	if l:= len(transs); l >= 2 {
		pos = transs[0]
		trans = transs [1]
	}else if l == 1 {
		trans = transs[0]
	}else {
		return nil
	}
	ts := TermDict.Translation{
		POS:   pos,
		Trans: trans,
	}
	return &ts
}

