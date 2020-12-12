# Goinfras 

Goinfras是一个后端应用基础设施的资源组件启动器，其实现了一些后端常用的组件启动接口，并提供可执行的实例。
使用goinfras，你只需要在项目中注册所需要的组件，并编写所需要的配置信息，最后让Application启动即可使用。

### 实现的资源组件：
- XCron ：定时任务执行器，基于 https://github.com/robfig/cron/v3 包
- XEcho ：Echo Web 框架引擎，基于 https://github.com/labstack/echo/v4 包
- XEs ：ElasticSearch 客户端，基于 https://github.com/olivere/elastic/v7 或 https://github.com/elastic/go-elasticsearch/v8 官方包
- XEtcd ：分布式配置信息客户端，基于 https://go.etcd.io/etcd/clientv3 包
- XGin ：Gin Web 框架引擎，基于 https://github.com/gin-gonic/gin 包
- XJwt ：Json Web Token，令牌加解码工具，基于 https://github.com/dgrijalva/jwt-go 包
- XLogger ：高性能日志记录器，基于 https://go.uber.org/zap 包
- XMail ：网络邮件发送客户端，基于 https://gopkg.in/gomail.v2 包
- XMQ ：
    - XNats ：高性能nats消息队列连接池，基于 https://github.com/nats-io/nats.go 包
    - XRedisPubSub ：高效的Redis基础发布订阅连接池，https://github.com/gomodule/redigo/redis 包
- XOAuth ：QQ、微博、微信第三方OAuth登录鉴权管理器，实现后端逻辑部分
- XOSS ：
    - XAliyunOss ：阿里云对象存储客户端，基于 https://github.com/aliyun/aliyun-oss-go-sdk 包
    - XQiniuOss ：七牛云对象存储客户端，基于 https://github.com/qiniu/api.v7 包
- XSms ：
    - XAliyunSms ：阿里云短信服务客户端，基于 https://github.com/aliyun/alibaba-cloud-sdk-go 阿里云官方包
- XStore ：
    - XGorm ：高性能Orm存储DB连接池，基于 https://gorm.io/gorm 包
    - XMongo ：高性能Nosql存储连接池，基于 https://go.mongodb.org/mongo-driver 官方包
    - XRedis ：高性能redis缓存连接池，基于 https://github.com/gomodule/redigo 包
    - XSQLBuilder ：sql构造器DB实例，基于 https://github.com/didi/gendry 包
- XValidate ：高性能数据传输对象验证器，基于 https://gopkg.in/go-playground/validator.v9 包
- ...

### 架构图




### 用例



### 工具