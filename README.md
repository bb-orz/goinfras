

### 目录结构

```
goinfras       项目
|-- core            核心业务：服务层业务逻辑、领域层、持久层
|-- apis            用户接口层
|-- build           构建程序文件：makefile、dockerfile、sh、ini... 
|-- docs            项目文档
|-- infras          基础设施层：算法、utils、应用初始化（配置加载、数据库连接、日志、web框架、验证器、...）、负载均衡器、RPC...
|-- services        应用层接口：应用服务
```

#### apis
 - 文件名称为可以描述其业务含义的单词
 - 定义外部交互逻辑和交互形式：如webUI渲染、RESTful、RPC等
 - 不涉及任何业务，随时可以替换为其他形式的交互方式 （JSON、protobuf等）
 - 使用services层提供的业务应用接口:services各模块的构造和初始化 
 
#### services
 - 文件名称使用为可以描述其业务含义的单词
 - 需要对外暴露的：
    - DTO、service interface 接口方法
    - 常量枚举、常数等 

#### core 
 - core层为services层的具体实现；
 - 分为应用层service、领域层domain、数据访问层dao；
 - 应用层service定义服务提供的接口实现、领域层domain实现具体的业务领域逻辑、数据访问层dao实现具体的数据持久化操作
 - 各层级的错误传输需统一定义，且最终于service层接收错误处理；
 - 关于各层之间的数据交换使用DTO数据传输对象；
 - 关于数据访问层DAO，可根据需要使用orm、sql builder、sql原生驱动实现；
 - 文件名称需满足可以描述其业务+分层名称；
 - 
 
#### infras
基础设施层
infra包可成独立库，包含应用启动器，应用组件（即拿即用，组件已封装好初始化，包括数据库通用dao，etcd，日志，mq，oss，oauth2等常用组件），组件注册，算法实现


- algo 算法实现
- boot 应用初始化相关
- lb   负载均衡相关
- httpclient http通信相关
- rpc RPC通信相关

#### build
应用构建相关脚本
- ini 配置文件
- makefile
- dockerfile
- sh
