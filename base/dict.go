package base

type Dict interface {
	LookUp(word string) *Word
}
