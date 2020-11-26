package qiniuOss

import (
	qiniuOss "github.com/qiniu/api.v7/v7/auth/qbox"
)

func NewQiniuOssMac(cfg *Config) *qiniuOss.Mac {
	// 七牛云存储初始化
	return qiniuOss.NewMac(cfg.AccessKey, cfg.SecretKey)
}
