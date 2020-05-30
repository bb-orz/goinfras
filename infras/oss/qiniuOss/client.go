package qiniuOss

import (
	qiuniuOss "github.com/qiniu/api.v7/v7/auth/qbox"
)

func NewQiniuOssMac(cfg *qiniuOssConfig) *qiuniuOss.Mac {
	// 七牛云存储初始化
	return qiuniuOss.NewMac(cfg.AccessKey, cfg.SecretKey)
}
