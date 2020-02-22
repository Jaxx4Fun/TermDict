# 项目说明
TermDict是一个命令行词典项目。
目前使用http请求HTML后，直接解析DOM，而后使用yaml库进行结构体转string的操作（偷懒）。

## 目录说明
```bash
TermDict
├── base                                    # 存放基本数据结构定义，如Word
│   ├── base.go                                     
│   └── base_test.go
├── td                                      # td=TerminalDict                           
│   └── main.go                             # main入口
├── go.mod
├── go.sum
├── README.md
├── util                                    # 工具目录
│   └── lru                                 # lru（暂时未使用）
│       ├── lru.go
│       ├── lru_test.go
│       └── node.go
└── youdao                                  # 有道
    ├── parser.go                           # 页面解析/JSON解析
    ├── parser_test.go          
    ├── youdao.go                           # http访问youdao
    └── youdao_test.go
```


## 安装
### 前提
```bash
# 设置环境变量GOBIN，后续install在该目录下
export $GOBIN=<放二进制文件的目录>
# 将GOBIN目录添加到PATH方便查找
export $PATH=$PATH:$GOBIN
```
### 安装
```bash
cd TermDict
go install ./td/
 ```
## 使用
```shell script
# 查询单词, 格式td <word>，如：
➤ td hello
单词: hello
解释:
- 释: int. 喂；哈罗，你好，您好
- 释: n. 表示问候， 惊奇或唤起注意时的用语
- 释: n. (Hello) 人名；（法）埃洛
例句:
- 英: Hello, who's speaking, please?
  中: 喂， 请问你是谁呀?
- 英: 'Visitor: Hello! Where should I park my car?'
  中: 参观者：你好！请问我的车该停哪儿？
- 英: “Hello?” he said.
  中: “你好吗？” 他说。
来源: Youdao
```