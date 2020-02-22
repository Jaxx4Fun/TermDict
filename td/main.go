package main

import (
	"fmt"
	"github.com/Johnny4Fun/TermDict/youdao"
	"os"
)
var (
	//word  = flag.String("word", "", "word to loop up")
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	word := &os.Args[1]
	dict := youdao.NewDict()
	result := dict.LookUp(*word)
	fmt.Fprint(os.Stdout,result)
}
