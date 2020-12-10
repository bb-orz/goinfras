# ElasticSearch 客户端

> 基于 github.com/elastic/go-elasticsearch/v8 包

### ElasticSearch Documentation
> Documentation https://www.elastic.co/guide/cn/elasticsearch/guide/current/index.html


### XEs Starter Usage
```
// ...

// Defind OptionalConfig
optCfg := new(OptionalConfig)
optCfg.HttpHeader = ...                 // 设置API HTTP Header
optCfg.HttpTransport = ...              // 设置API HTTP transport object
optCfg.Logger  = ...                     // 设置logger object
optCfg.Selector = ...                   // 设置selector object
optCfg.RetryBackoffFunc = ...             // 设置可选的退避持续时间处理函数
optCfg.ConnectionPoolFunc = ...           // 设置连接池处理函数
 

goinfras.Register(XEs.NewStarter(optCfg))

// ...
```

### XEs Config Setting
```
// 自建服务的配置信息
Addresses []string // es服务的集群节点设置.
Username  string   // 基于 HTTP Basic Authentication 的用户名
Password  string   // 基于 HTTP Basic Authentication 的密码鉴权

// elastic.io 云服务的配置信息，覆盖自建服务的配置
CloudID string // elastic.io服务的CloudID (https://elastic.co/cloud)，需注册后获取
APIKey  string // Base64-encoded token for authorization; if set, overrides username and password.

// PEM-encoded 加密传输鉴权令牌.
// When set, an empty certificate pool will be created, and the certificates will be appended to it.
// The option is only valid when the transport is not specified, or when it's http.Transport.
CACert []byte

// API请求重试配置设置
RetryOnStatus        []int // List of status codes for retry. Default: 502, 503, 504.
DisableRetry         bool  // 禁用重试，Default: false.
EnableRetryOnTimeout bool  // 启用重试超时设置，Default: false.
MaxRetries           int   // 最多重试次数，Default: 3.

// 服务节点查找周期设置
DiscoverNodesOnStart  bool          // 是否初始化客户端时查找节点. Default: false.
DiscoverNodesInterval time.Duration // 是否周期型查找节点. Default: disabled.

// Metrics度量：聚合查询设置
EnableMetrics     bool // 是否启用Metrics度量度量，类聚合查询
EnableDebugLogger bool // 是否在debug时启用日志记录


```


### XEs Usage
```


```