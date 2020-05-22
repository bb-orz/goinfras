package qiniuOss

import (
	"GoWebScaffold/infras/config"
	qiuniuOss "github.com/qiniu/api.v7/v7/auth/qbox"
)

func Init(appConf *config.AppConfig) *qiuniuOss.Mac {
	// 七牛云存储初始化
	if appConf.OssConf.Qiniu.Switch {
		return qiuniuOss.NewMac(appConf.OssConf.Qiniu.AccessKey, appConf.OssConf.Qiniu.SecretKey)
	}
	return nil
}
