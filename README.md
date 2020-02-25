# 项目说明
TermDict是一个命令行词典项目。
目前使用http请求HTML后，直接解析DOM，而后使用yaml库进行结构体转string的操作（偷懒）。

## 目录说明
```bash
.
.
├── app                                     //可执行文件的目录
│   ├── td                                  //简易版本，直接通过http查词
│   ├── td_client                           //rpc版本的客户端
│   └── td_server                           //rpc服务
├── base                                    //基础数据结构、常量
│   ├── base.go
│   ├── base_test.go
│   ├── const.go
│   └── dict.go
├── go.mod
├── go.sum
├── README.md
├── util                                    //工具
│   ├── cache                               //缓存
│   │   ├── cache.go                        //缓存的抽象定义
│   │   ├── lru.go                          //lru，非线程安全
│   │   ├── lru_test.go
│   │   ├── lru_threadsafe.go               //包装后的lru，线程安全
│   │   ├── lru_threadsafe_test.go
│   │   └── node.go
│   └── util.go
└── youdao                                  //有道词典的请求与解析
    ├── parser.go                           //目前仅解析http，后续考虑直接访问api，解析json
    ├── parser_test.go                  
    ├── youdao.go                           
    └── youdao_test.go
```


## 简易版本安装
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
## TODO: C/S版本安装
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