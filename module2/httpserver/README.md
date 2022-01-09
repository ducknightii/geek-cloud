
### 题目

编写一个 HTTP 服务器，大家视个人不同情况决定完成到哪个环节，但尽量把 1 都做完：

- 接收客户端 request，并将 request 中带的 header 写入 response header
- 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
- Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
- 当访问 localhost/healthz 时，应返回 200


### 答案

- 路由 /header 写入接收到的request中的请求
- VERSION=v1.0 go run main.go  路由 /herder 会写入VERSION 到 response header
- defer func() {requestInfo()}() 用于打印请求日志
- 构造一个啥也不干的路由 func 就可以吧
