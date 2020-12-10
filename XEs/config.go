package XEs

import (
	esv8 "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/estransport"
	"time"
)

type Config struct {
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
}

func DefaultConfig() *Config {
	return &Config{}
}

type RetryBackoffFunc func(attempt int) time.Duration
type ConnectionPoolFunc func([]*estransport.Connection, estransport.Selector) estransport.ConnectionPool

// 转换ES配置类型
func genESConfig(config *Config) esv8.Config {
	esCfg := esv8.Config{}

	// 设置服务主机及鉴权信息
	if config.CloudID != "" && config.APIKey != "" {
		esCfg.CloudID = config.CloudID
		esCfg.APIKey = config.APIKey
	} else if len(config.Addresses) != 0 {
		esCfg.Addresses = config.Addresses
		if config.Username != "" && config.Password != "" {
			esCfg.Username = config.Username
			esCfg.Password = config.Password
		}
	}

	// 安全令牌设置
	if config.CACert != nil {
		esCfg.CACert = config.CACert
	}

	// 重试设置,在没关闭重试和设置最大重试设置时
	if !config.DisableRetry && config.MaxRetries != 0 {
		esCfg.DisableRetry = config.DisableRetry
		esCfg.RetryOnStatus = config.RetryOnStatus
		esCfg.EnableRetryOnTimeout = config.EnableRetryOnTimeout
		esCfg.MaxRetries = config.MaxRetries
	}

	// 服务节点查找周期设置
	if config.DiscoverNodesOnStart {
		esCfg.DiscoverNodesOnStart = config.DiscoverNodesOnStart
	}
	if config.DiscoverNodesInterval != 0 {
		esCfg.DiscoverNodesInterval = config.DiscoverNodesInterval
	}

	if config.EnableMetrics {
		esCfg.EnableMetrics = config.EnableMetrics
	}

	if config.EnableDebugLogger {
		esCfg.EnableDebugLogger = config.EnableDebugLogger
	}
	return esCfg
}
