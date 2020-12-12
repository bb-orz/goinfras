package XEsOfficial

import (
	esv8 "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/estransport"
	"net/http"
)

var esClient *esv8.Client

// 创建默认的ES Client
func CreateDefaultESClient() error {
	var err error
	esClient, err = esv8.NewDefaultClient()
	return err
}

/**
 * @Description:  				创建可配置的ES Client
 * @param config  				用户配置信息
 * @param httpHeader 			设置API HTTP Header
 * @param transport 			设置API HTTP transport object
 * @param logger 				设置logger object
 * @param selector				设置selector object
 * @param retryBackoffFunc  	设置可选的退避持续时间处理函数
 * @param connectionPoolFunc	设置连接池处理函数
 * @return *esv8.Client
 * @return error
 */
func NewESClient(config *Config, httpHeader http.Header, httpTransport http.RoundTripper, logger estransport.Logger, selector estransport.Selector, retryBackoffFunc RetryBackoffFunc, connectionPoolFunc ConnectionPoolFunc) (*esv8.Client, error) {
	var err error
	var esClient *esv8.Client

	if config == nil {
		esClient, err = esv8.NewDefaultClient()
		return esClient, err
	}

	esConfig := genESConfig(config)

	// 设置其他高级配置项
	if httpHeader != nil {
		esConfig.Header = httpHeader
	}
	if httpTransport != nil {
		esConfig.Transport = httpTransport
	}
	if logger != nil {
		esConfig.Logger = logger
	}
	if selector != nil {
		esConfig.Selector = selector
	}
	if retryBackoffFunc != nil {
		esConfig.RetryBackoff = retryBackoffFunc
	}
	if connectionPoolFunc != nil {
		esConfig.ConnectionPoolFunc = connectionPoolFunc
	}

	esClient, err = esv8.NewClient(esConfig)
	if err != nil {
		return nil, err
	}

	return esClient, nil
}
