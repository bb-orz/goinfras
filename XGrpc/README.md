# XGrpc 客户端
grpc + protobuf 客户端组件
> 基于  https://github.com/grpc/grpc-go 和 https://github.com/golang/protobuf

> GRPC  中文文档 https://www.kancloud.cn/adapa/go-grpc/1109311

> Grpc Official Documentation https://grpc.io/docs/languages/go/quickstart/

> Protobuf Official Documentation https://developers.google.com/protocol-buffers/docs/gotutorial

> Protobuf 中文翻译解析： https://www.jianshu.com/nb/27416806

### Protobuf 生成命令

#### 1.安装protoc

#### 2.定义proto协议文件

#### 3.生成go端代码
```
自动生成命名详解
protoc --proto_path=path/to/protofile/rootdir/ --go_out=path/to/gofileout/ --micro_out=path/to/gomicrofileout {proto_file.proto}
--proto_path= 编译proto文件的所在的目录
--go_out=生成 编译go文件输出的目录
--micro_out=生成micro风格的文件目录
最后一个是目标文件

示例：
protoc --proto_path= --go_out=plugins=grpc:. ./test.proto
```