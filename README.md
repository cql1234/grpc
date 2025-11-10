# gRPC + Protocol Buffers 简单示例
这是一个简单的 gRPC 和 Protocol Buffers 学习示例，包含一个问候服务。
## 项目结构
```
.
├── proto/
│   ├── greeter.proto          # Protocol Buffers 定义文件
│   ├── greeter.pb.go          # 生成的消息代码
│   └── greeter_grpc.pb.go     # 生成的 gRPC 服务代码
├── server/
│   └── main.go                # gRPC 服务器实现
├── client/
│   └── main.go                # gRPC 客户端实现
├── go.mod                      # Go 模块文件
├── generate.bat               # 生成 proto 代码的脚本
├── start_server.bat           # 启动服务器脚本
├── start_client.bat           # 启动客户端脚本
└── README.md                   # 说明文档
```
## 核心概念
### 1. Protocol Buffers (.proto 文件)
- 定义服务接口和消息格式
- 跨语言的数据序列化格式
- 比 JSON/XML 更小、更快、更简单
### 2. gRPC 服务类型
本示例包含两种 RPC 类型：
- **简单 RPC**: 客户端发送一个请求，服务器返回一个响应
- **服务器流式 RPC**: 客户端发送一个请求，服务器返回多个响应流
## 快速开始
### 方法一：使用批处理脚本（推荐）
1. **启动服务器**  
   双击运行 `start_server.bat`
2. **启动客户端**  
   在另一个终端双击运行 `start_client.bat`
### 方法二：手动运行
1. **启动服务器**
```bash
go run server/main.go
```
2. **启动客户端**（在新终端）
```bash
go run client/main.go
```
## 完整使用步骤
### 1. 安装 Protocol Buffers 编译器
Windows 用户可以从以下地址下载：
https://github.com/protocolbuffers/protobuf/releases
下载后将 protoc.exe 添加到系统 PATH 环境变量中。
### 2. 安装 Go 插件
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
确保 GOPATH/bin 在系统 PATH 中。
### 3. 下载 Go 依赖
```bash
go mod tidy
```
### 4. 生成 gRPC 代码（如果修改了 .proto 文件）
运行 `generate.bat` 或手动执行：
```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/greeter.proto
```
这会生成：
- `proto/greeter.pb.go` - 消息类型定义
- `proto/greeter_grpc.pb.go` - gRPC 服务接口
## 学习要点
### Proto 文件解析 (proto/greeter.proto)
```protobuf
syntax = "proto3";                    // 使用 proto3 语法
package greeter;                      // 包名
option go_package = "grpc-demo/proto"; // Go 包路径
service Greeter {                     // 定义服务
  rpc SayHello (HelloRequest) returns (HelloResponse) {}
  rpc SayHelloMultipleTimes (HelloRequest) returns (stream HelloResponse) {}
}
message HelloRequest {                // 定义请求消息
  string name = 1;                    // 字段编号从 1 开始
}
message HelloResponse {               // 定义响应消息
  string message = 1;
  int32 count = 2;
}
```
### 服务器端实现要点
1. **实现服务接口**
   - 继承 `UnimplementedGreeterServer`
   - 实现 proto 中定义的方法
2. **创建并启动服务器**
   - 监听 TCP 端口
   - 创建 gRPC 服务器实例
   - 注册服务实现
   - 开始提供服务
### 客户端实现要点
1. **建立连接**
   - 使用 `grpc.Dial()` 连接服务器
2. **调用远程方法**
   - 简单 RPC：直接调用，获取单个响应
   - 流式 RPC：循环接收多个响应
## 运行示例
### 预期输出
**服务器端：**
```
2025/11/10 21:00:00 gRPC 服务器启动，监听端口 :50051
2025/11/10 21:00:05 收到请求: name=张三
2025/11/10 21:00:06 收到流式请求: name=李四
2025/11/10 21:00:06 发送消息 1/5
2025/11/10 21:00:07 发送消息 2/5
...
```
**客户端：**
```
===== 测试 1: 简单的 SayHello =====
收到响应: 你好, 张三! 欢迎学习 gRPC 和 Protocol Buffers (count: 1)
===== 测试 2: 服务器流式 RPC =====
收到流式响应: 第 1 次问候: 你好, 李四! (count: 1)
收到流式响应: 第 2 次问候: 你好, 李四! (count: 2)
...
流式响应接收完毕
客户端演示完成!
```
## 扩展练习
1. **添加更多 RPC 方法**
   - 客户端流式 RPC（客户端发送多个请求）
   - 双向流式 RPC（双向同时发送）
2. **增强消息类型**
   - 添加更多字段（时间戳、状态等）
   - 使用嵌套消息
   - 使用枚举类型
3. **错误处理**
   - 实现超时控制
   - 添加重试机制
   - 使用 gRPC 状态码
4. **安全性**
   - 添加 TLS 加密
   - 实现认证机制
## 常见问题
**Q: 为什么需要 Protocol Buffers？**  
A: 提供强类型接口、自动代码生成、高效序列化。
**Q: gRPC 相比 REST API 的优势？**  
A: 
- 性能更好（HTTP/2 + 二进制协议）
- 支持流式传输
- 强类型接口
- 自动生成客户端代码
**Q: 修改 .proto 文件后需要做什么？**  
A: 重新运行 `generate.bat` 生成新的 Go 代码。
**Q: 如何调试 gRPC？**  
A: 可以使用 grpcurl 工具或 Postman 的 gRPC 功能。
## 参考资料
- [gRPC 官方文档](https://grpc.io/)
- [Protocol Buffers 文档](https://protobuf.dev/)
- [gRPC Go 教程](https://grpc.io/docs/languages/go/quickstart/)
- [Proto3 语法指南](https://protobuf.dev/programming-guides/proto3/)
