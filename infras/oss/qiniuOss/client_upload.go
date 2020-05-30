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
		Expires:    uint64(QiniuOssCfg().UpTokenExpires),
		MimeLimit:  QiniuOssCfg().MimeLimit,
		FsizeMin:   int64(QiniuOssCfg().FsizeMin),
		FsizeLimit: int64(QiniuOssCfg().FsizeMax),
	}
	upToken = putPolicy.UploadToken(QiniuOssClient())
	return
}

// 覆盖上传，返回上传凭证,给客户端上传
func OverwriteUpload(bucket, keyToOverwrite string) (upToken string) {
	// keyToOverwrite为需要覆盖的文件名
	putPolicy := storage.PutPolicy{
		Scope:      fmt.Sprintf("%s:%s", bucket, keyToOverwrite),
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
		Expires:    uint64(QiniuOssCfg().UpTokenExpires),
		MimeLimit:  QiniuOssCfg().MimeLimit,
		FsizeMin:   int64(QiniuOssCfg().FsizeMin),
		FsizeLimit: int64(QiniuOssCfg().FsizeMax),
	}
	upToken = putPolicy.UploadToken(QiniuOssClient())
	return
}

// 带回调上传，返回上传凭证,给客户端上传
func CallbackUpload(bucket string) (upToken string) {
	putPolicy := storage.PutPolicy{
		Scope:            bucket,
		CallbackURL:      QiniuOssCfg().CallbackURL,
		CallbackBody:     `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
		CallbackBodyType: QiniuOssCfg().CallbackBodyType,
		Expires:          uint64(QiniuOssCfg().UpTokenExpires),
		MimeLimit:        QiniuOssCfg().MimeLimit,
		FsizeMin:         int64(QiniuOssCfg().FsizeMin),
		FsizeLimit:       int64(QiniuOssCfg().FsizeMax),
	}
	upToken = putPolicy.UploadToken(QiniuOssClient())
	return
}
