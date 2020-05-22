

### 目录结构

```
GoWebScaffold       项目
|-- core            核心业务：应用层、领域层、持久层
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
 - core层为services层的具体实现
 - 分为应用层、领域层、数据访问层
 - 文件名称需满足可以描述其业务+分层名称
 - 每个业务模块在一个文件夹，里面实现该模块的Domain、DAO、PO
 
#### infras
基础设施层

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
