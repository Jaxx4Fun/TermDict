package base

import (
	"github.com/Johnny4Fun/TermDict/util"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
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
var (
	RegexCacheFileName, _ = regexp.Compile(`\w+`)
)
var (
	EnvTermDictRoot  string = "~/TermdictCache/"
	EnvTermDictCache string
)

func Initialize() {
	if env := os.Getenv(EnvTermDictRootKey); env != "" {
		EnvTermDictRoot = env
	}
	userCacheDir, _ := os.UserCacheDir()
	EnvTermDictRoot = strings.Replace(EnvTermDictRoot, "~", userCacheDir, 1)
	EnvTermDictCache = EnvTermDictRoot + "/cache"

	if err := util.MkdirIfNotExist(EnvTermDictCache); err != nil {
		log.Fatalf("failed to mkdir %s, %v", EnvTermDictCache, err)
	}

	file, err := os.OpenFile(EnvTermDictRoot+time.Now().Format("2006-01-02")+".log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln("failed to create log file")
	}

	log.SetOutput(file)
}

// 单词
type Word struct {
	Spell     string          `yaml:"单词"`
	Trans     []Translation   `yaml:"解释"`
	Pronounce []Pronunciation `yaml:"发音,omitempty,flow"`
	Examples  []Example       `yaml:"例句,omitempty"`
	Source    WordSrc         `yaml:"来源,omitempty"`
	From      FromStorage	  `yaml:"存储位置,omitempty"`
}

func (w *Word) Key() string {
	return w.Spell
}
func (w Word) String() string {
	return wordToYamlString(&w)
}

//存储位置
type FromStorage int

// 发音
type Pronunciation string

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

func (w FromStorage) IsZero() bool {
	return true
}
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
