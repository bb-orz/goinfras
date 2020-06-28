package aliyunOss

import (
	aliOss "github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func NewClient(cfg *AliyunOssConfig) (*aliOss.Client, error) {
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
