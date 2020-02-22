package base

import (
	"gopkg.in/yaml.v2"
	"log"
)

var (
	// 是否显示例句来源
	OptShowExampleSrc = false
	// 是否显示查询结果的来源
	OptShowResultSrc = true
	// 词性和解释是否分行显示
	// true : <pos> <trans>
	// false: - 词性: <pos>
	//			释义: <trans>
	OptSeparatePosAndTrans = false
)

// 单词
type Word struct {
	Spell    string        `yaml:"单词"`
	Trans    []Translation `yaml:"解释"`
	Examples []Example     `yaml:"例句,omitempty"`
	From     WordSrc       `yaml:"来源,omitempty"`
}

func (w Word) String() string {
	return wordToYamlString(&w)
}

// 例句，包括英文、中文以及例句出处
type Example struct {
	Sentence    string     `yaml:"英"`
	Translation string     `yaml:"中"`
	From        ExampleSrc `yaml:"源,omitempty"`
}

type Translation struct {
	//词性 position of speech
	POS string `yaml:"词性,omitempty"`
	//目标语言释义
	Trans string `yaml:"释"`
}

func (t Translation) String() string {
	return t.POS + " " + t.Trans
}

type WordSrc string

func (w WordSrc) IsZero() bool {
	return !OptShowResultSrc && w != ""
}

type ExampleSrc string

func (f ExampleSrc) IsZero() bool {
	return !OptShowExampleSrc
}

func wordToYamlString(word *Word) string {
	b, err := yaml.Marshal(word)
	if err != nil {
		log.Printf("failed to marshal Word %v, %v", word.Spell, err)
	}
	return string(b)
}
