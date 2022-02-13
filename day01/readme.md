# go语言简介

### go语言特点
1. 背靠大厂, 有google 背书, 可靠
2. 天生支持并发(最显著特点)
3. 语法简单容易上手
4. 内置runtime, 支持GC(垃圾回收)
5. 丰富的标准库
6. 部署维护成本低 (知乎网站初始使用python开发, 由于部署维护成本太高后续换成go, 节省成本)

### 应用领域

1. 服务器编程
2. 开发云平台 (docker k8s都是go开发)
3. 区块链
4. 分布式网络
5. 网络编程

### 环境安装
windows 直接下载 msi安装包 然后安装

需要添加安装目录的 bin文件夹到 Path变量中
如
> D:\Go1.17.7\bin

需要添加环境变量
> GOPATH
值为
> D:\Go1.17.7\bin

之后在命令行输入 
> go version

其他系统过程也是类似

### vscode 开发环境配置

1. 安装扩展程序 go


2. 安装go开发工具集
   1. 添加环境变量 GO111MODULE 为 on; GOPROXY 为 https://goproxy.cn
   2. 重启vscode
   3. ctrl + shift + p 搜索 go install
   4. 全选工具集安装
> https://github.com/goproxy/goproxy.cn/blob/master/README.zh-CN.md


tips: 将vscode 的 安装目录的 bin 文件夹 添加到系统变量的path中然后就可以 使用 code . 命令用vs code 打开当前所在文件夹

### 开始开发

在系统文件夹下执行
> go mod init ./day01 初始化
> go mod tidy 解决其他依赖包的问题
> go get http://githubxxxxxxx 安装第三方依赖 
> go get -u github.com/gin-gonic/gin // 安装gin go的web 框架

代码片段快捷键
pkgm  main包, 生成main 函数
fp   输出 fmt.Println()
ff   输出 fmt.Printf("%v", var1)
