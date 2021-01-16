package XQiniuOss

import (
	"fmt"
	"github.com/qiniu/api.v7/v7/storage"
)

// 默认bucket简单上传，返回上传凭证,给客户端上传
func (client *QnClient) SimpleUpload() (upToken string) {
	return client.SimpleUploadToBucket(client.cfg.DefaultBucket)
}

// 简单上传，返回上传凭证给客户端上传
func (client *QnClient) SimpleUploadToBucket(bucket string) (upToken string) {
	putPolicy := storage.PutPolicy{
		Scope:      bucket,
		ReturnBody: client.cfg.ReturnBody,
		Expires:    uint64(client.cfg.UpTokenExpires),
		MimeLimit:  client.cfg.MimeLimit,
		FsizeMin:   int64(client.cfg.FsizeMin),
		FsizeLimit: int64(client.cfg.FsizeMax),
	}
	upToken = putPolicy.UploadToken(client.mac)
	return
}

// 默认bucket覆盖上传，返回上传凭证,给客户端上传
func (client *QnClient) OverwriteUpload(keyToOverwrite string) (upToken string) {
	return client.OverwriteUploadToBucket(client.cfg.DefaultBucket, keyToOverwrite)
}

// 覆盖上传，返回上传凭证,给客户端上传
func (client *QnClient) OverwriteUploadToBucket(bucket, keyToOverwrite string) (upToken string) {
	// keyToOverwrite为需要覆盖的文件名
	putPolicy := storage.PutPolicy{
		Scope:      fmt.Sprintf("%s:%s", bucket, keyToOverwrite),
		ReturnBody: client.cfg.ReturnBody,
		Expires:    uint64(client.cfg.UpTokenExpires),
		MimeLimit:  client.cfg.MimeLimit,
		FsizeMin:   int64(client.cfg.FsizeMin),
		FsizeLimit: int64(client.cfg.FsizeMax),
	}
	upToken = putPolicy.UploadToken(client.mac)
	return
}

// 默认bucket带回调上传，返回上传凭证,给客户端上传
func (client *QnClient) CallbackUpload() (upToken string) {
	return client.CallbackUploadToBucket(client.cfg.DefaultBucket)
}

// 带回调上传，返回上传凭证,给客户端上传
func (client *QnClient) CallbackUploadToBucket(bucket string) (upToken string) {
	putPolicy := storage.PutPolicy{
		Scope:            bucket,
		CallbackURL:      client.cfg.CallbackURL,
		CallbackBody:     client.cfg.CallbackBody,
		CallbackBodyType: client.cfg.CallbackBodyType,
		Expires:          uint64(client.cfg.UpTokenExpires),
		MimeLimit:        client.cfg.MimeLimit,
		FsizeMin:         int64(client.cfg.FsizeMin),
		FsizeLimit:       int64(client.cfg.FsizeMax),
	}
	upToken = putPolicy.UploadToken(client.mac)
	return
}
