package XAliyunOss

import (
	aliOss "github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var aliyunOssClient *aliOss.Client

// 创建一个默认配置的Manager
func CreateDefaultClient(config *Config) error {
	var err error
	if config == nil {
		config = DefaultConfig()
	}
	aliyunOssClient, err = NewClient(config)
	return err
}

func NewClient(cfg *Config) (*aliOss.Client, error) {
	// Aliyun OSS初始化
	return aliOss.New(
		cfg.Endpoint,
		cfg.AccessKeyId,
		cfg.AccessKeySecret,
		aliOss.Timeout(int64(cfg.ConnTimeout), int64(cfg.RWTimeout)),
		aliOss.UseCname(cfg.UseCname),
		aliOss.EnableCRC(cfg.EnableCRC),
	)
}
