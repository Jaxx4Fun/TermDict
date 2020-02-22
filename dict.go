package TermDict

import "github.com/Johnny4Fun/TermDict/base"

type Dict interface {
	LookUp(word string) *base.Word
}

// constant for Spell from
const (
	Youdao = iota
	Baidu
)

type TermDict struct {

}
