package base

// 词性，暂时不用
const (
	Verb = iota
)

// 单词来源
const (
	Youdao = "Youdao"
	Baidu  = "Baidu"
)

const (
	EnvTermDictRootKey = "TERMDICT_DIR"
)

const (
	CacheCapacity = 10
)

// FromStorage
const (
	Online = FromStorage(iota)
	MemCache
)
