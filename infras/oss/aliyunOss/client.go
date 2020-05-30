package aliyunOss

import (
	"GoWebScaffold/infras"
	aliOss "github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func NewClient(cfg *aliyunOssConfig) *aliOss.Client {
	// Aliyun OSS初始化
	var err error
	var aoss *aliOss.Client
	aoss, err = aliOss.New(
		cfg.Endpoint,
		cfg.AccessKeyId,
		cfg.AccessKeySecret,
		aliOss.Timeout(int64(cfg.ConnTimeout), int64(cfg.RWTimeout)),
		aliOss.UseCname(cfg.UseCname),
		aliOss.EnableCRC(cfg.EnableCRC),
	)
	infras.FailHandler(err)
	return aoss

}
