package aliyunOss

import (
	"GoWebScaffold/infras"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var aliyunOssClient *oss.Client

func AliyunOssComponent() *oss.Client {
	infras.Check(aliyunOssClient)
	return aliyunOssClient
}

func SetComponent(c *oss.Client) {
	aliyunOssClient = c
}
