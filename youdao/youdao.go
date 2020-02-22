package youdao

import (
	"fmt"
	"github.com/Johnny4Fun/TermDict/base"
	"log"
	"net/http"
)

type Dict struct {
}
func NewDict() *Dict{
	return &Dict{}
}
//
const queryURLPattern = "http://www.youdao.com/w/%s"

var responseParser = ParserFunc(ParseHTML)

func (y *Dict) LookUp(word string) *base.Word {
	url := fmt.Sprintf(queryURLPattern, word)
	resp, err := http.Get(url)
	if err != nil {
		// handle error
		log.Fatalf("GET %q Error, %v", url, err)
		return nil
	}
	defer resp.Body.Close()
	wd, err := responseParser.Parse(resp.Body)
	if err != nil {
		log.Fatalf("parse resp body error, %v", err)
	}

	wd.Spell = word
	wd.From = base.Youdao

	return wd
}
