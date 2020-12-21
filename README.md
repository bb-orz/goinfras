# Goinfras 

Goinfras是一个后端应用基础设施的资源组件启动器，其实现了一些后端常用的组件客户端或工具，并提供可执行的实例。
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

待更新...


### 用例

使用启动器项目，您只需做如下步骤：

##### Step 1：注册您所需要的启动器

```
import (
	"github.com/bb-orz/goinfras"
	"github.com/bb-orz/goinfras/XCron"
	"github.com/bb-orz/goinfras/XEtcd"
	"github.com/bb-orz/goinfras/XLogger"
	"github.com/bb-orz/goinfras/XMQ/XNats"
	"github.com/bb-orz/goinfras/XMQ/XRedisPubSub"
	"github.com/bb-orz/goinfras/XOAuth"
	"github.com/bb-orz/goinfras/XOss/XAliyunOss"
	"github.com/bb-orz/goinfras/XOss/XQiniuOss"
	"github.com/bb-orz/goinfras/XStore/XGorm"
	"github.com/bb-orz/goinfras/XStore/XMongo"
	"github.com/bb-orz/goinfras/XStore/XRedis"
	"github.com/bb-orz/goinfras/XValidate"
	"github.com/gin-gonic/gin"

	_ "github.com/bb-orz/goinfras-sample/simple/restful" // 自动载入Restful API模块
	"github.com/bb-orz/goinfras/XGin"
)

// 注册应用组件启动器，把基础设施各资源组件化
func RegisterStarter() {
	
	goinfras.RegisterStarter(XLogger.NewStarter())

	// 注册Cron定时任务
	// 可以自定义一些定时任务给starter启动
	goinfras.RegisterStarter(XCron.NewStarter())

	// 注册ETCD
	goinfras.RegisterStarter(XEtcd.NewStarter())

	// 注册mongodb启动器
	goinfras.RegisterStarter(XMongo.NewStarter())

	// 注册mysql启动器
	goinfras.RegisterStarter(XGorm.NewStarter())
	// 注册Redis连接池
	goinfras.RegisterStarter(XRedis.NewStarter())
	// 注册Oss
	goinfras.RegisterStarter(XAliyunOss.NewStarter())
	goinfras.RegisterStarter(XQiniuOss.NewStarter())
	// 注册Mq
	goinfras.RegisterStarter(XNats.NewStarter())
	goinfras.RegisterStarter(XRedisPubSub.NewStarter())
	// 注册Oauth Manager
	goinfras.RegisterStarter(XOAuth.NewStarter())


	// 注册gin web 服务启动器
	// TODO add your gin middlewares
	middlewares := make([]gin.HandlerFunc, 0)
	goinfras.RegisterStarter(XGin.NewStarter(middlewares...))

	// 注册验证器
	goinfras.RegisterStarter(XValidate.NewStarter())

	// 对资源组件启动器进行排序
	goinfras.SortStarters()

}
```


##### Step 2：选择您的web引擎：gin/echo,定义相应的接口并在包初始化时注册接口路由

```
func init() {
	// 初始化时自动注册该API到Gin Engine
	XGin.RegisterApi(new(SimpleApi))
}

type SimpleApi struct {
	service1 services.IService1
}

// SetRouter由Gin Engine 启动时调用
func (s *SimpleApi) SetRoutes() {
	s.service1 = services.GetService1()

	engine := XGin.XEngine()

	engine.GET("simple/foo", s.Foo)
	engine.GET("simple/bar", s.Bar)
}

func (s *SimpleApi) Foo(ctx *gin.Context) {
	email := ctx.Param("email")
	// 调用服务
	err := s.service1.Foo(services.InDTO{Email: email})

	// 处理错误
	fmt.Println(err)
}

func (s *SimpleApi) Bar(ctx *gin.Context) {
	email := ctx.Param("email")
	// 调用服务
	err := s.service1.Bar(services.InDTO{Email: email})

	// 处理错误
	fmt.Println(err)
}


```



##### Step 3：创建应用并启动

```
var app *goinfras.Application // 应用实例

func main() {
	// 初始化Viper配置加载器，导入配置，启动参数由命令行flag输入
	fmt.Println("Viper Config Loading  ......")
	viperCfg := goinfras.ViperLoader()

	// 注册应用组件启动器
	fmt.Println("Register Starters  ......")
	RegisterStarter()

	// 创建应用程序启动管理器
	app = goinfras.NewApplication(viperCfg)

	// 运行应用,启动已注册的资源组件
	fmt.Println("Application Starting ......")
	app.Up()
}
```


##### Step4：确定你的配置信息，可通过环境变量、配置文件或远程配置中心设置
> 本项目提供模板配置文件供参考：example.yaml

运行goinfras的项目需注意：

- 使用goinfras.ViperLoader()，载入初始viper配置实例时，默认接收以下命令行参数，获取viper实例初始配置：
   -  -f ：Config file flag,like: -f ../config/config.yaml
   -  -P : Remote K/V config flag, system provider，support etcd/consul. like: -P=etcd
   -  -E : Remote K/V config flag, system endpoint，etcd requires http://ip:port  consul requires ip:port
   -  -K : Remote K/V config flag, k is the path in the k/v store to retrieve configuration,like: -K /configs/myapp.json"
   -  -T : Remote K/V config flag, upport: 'json', 'toml', 'yaml', 'yml', 'properties', 'props', 'prop', 'env', 'dotenv'. like: -T=json
   -  -D : Remote K/V config flag, Currently, only tested with etcd support
   -  -a : ENV config flag, enable automatic, like: -a=true
   -  -e : ENV config flag, allow env  empty,like: -e=false
   -  -p : ENV config flag, env prefix,like: -p=goinfras_
   -  -k : ENV config flag, env keys,like: -k=aaa -k=bbb
   
   
### 工具

待更新...