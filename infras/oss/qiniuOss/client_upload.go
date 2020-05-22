package qiniuOss

import (
	"fmt"
	"github.com/qiniu/api.v7/v7/storage"
)

// 简单上传，返回上传凭证,给客户端上传
func SimpleUpload(bucket string) (upToken string) {
	putPolicy := storage.PutPolicy{
		Scope:      bucket,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
		Expires:    uint64(config.OssConf.Qiniu.UpTokenExpires),
		MimeLimit:  config.OssConf.Qiniu.MimeLimit,
		FsizeMin:   int64(config.OssConf.Qiniu.FsizeMin),
		FsizeLimit: int64(config.OssConf.Qiniu.FsizeMax),
	}
	upToken = putPolicy.UploadToken(oss.Qiniu)
	return
}

// 覆盖上传，返回上传凭证,给客户端上传
func OverwriteUpload(bucket, keyToOverwrite string) (upToken string) {
	// keyToOverwrite为需要覆盖的文件名
	putPolicy := storage.PutPolicy{
		Scope:      fmt.Sprintf("%s:%s", bucket, keyToOverwrite),
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
		Expires:    uint64(config.OssConf.Qiniu.UpTokenExpires),
		MimeLimit:  config.OssConf.Qiniu.MimeLimit,
		FsizeMin:   int64(config.OssConf.Qiniu.FsizeMin),
		FsizeLimit: int64(config.OssConf.Qiniu.FsizeMax),
	}
	upToken = putPolicy.UploadToken(oss.Qiniu)
	return
}

// 带回调上传，返回上传凭证,给客户端上传
func CallbackUpload(bucket string) (upToken string) {
	putPolicy := storage.PutPolicy{
		Scope:            bucket,
		CallbackURL:      config.OssConf.Qiniu.CallbackURL,
		CallbackBody:     `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
		CallbackBodyType: config.OssConf.Qiniu.CallbackBodyType,
		Expires:          uint64(config.OssConf.Qiniu.UpTokenExpires),
		MimeLimit:        config.OssConf.Qiniu.MimeLimit,
		FsizeMin:         int64(config.OssConf.Qiniu.FsizeMin),
		FsizeLimit:       int64(config.OssConf.Qiniu.FsizeMax),
	}
	upToken = putPolicy.UploadToken(oss.Qiniu)
	return
}
