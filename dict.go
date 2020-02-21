package TermDict

type Dict interface {
	LookUp(word string) *Word
}

type Word struct {
	Spell    string
	Trans    []Translation
	Examples []string
	From     int
}

type Translation struct {
	POS   string //词性 position of speech
	Trans string //目标语言中释义
}

// 词性
const (
	Verb = iota
)

// constant for Spell from
const (
	Youdao = iota
	Baidu
)

type TermDict struct {

}
