// 直接访问翻译网站，解析html后输出
// 单机版本，不需要cache
package main

import (
	"fmt"
	"github.com/Johnny4Fun/TermDict/youdao"
	"os"
)

func directLookup(word string) {

	dict := youdao.NewDict()
	result := dict.LookUp(word)
	fmt.Fprint(os.Stdout, result)
}

func main() {
	if len(os.Args) < 2 {
		return
	}
	word := os.Args[1]
	directLookup(word)
}
