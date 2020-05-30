package qiniuOss

import (
	"fmt"
	"github.com/qiniu/api.v7/v7/storage"
)

// 简单上传，返回上传凭证,给客户端上传
func (client *QnClient) SimpleUpload(bucket string) (upToken string) {
	putPolicy := storage.PutPolicy{
		Scope:      bucket,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
		Expires:    uint64(client.cfg.UpTokenExpires),
		MimeLimit:  client.cfg.MimeLimit,
		FsizeMin:   int64(client.cfg.FsizeMin),
		FsizeLimit: int64(client.cfg.FsizeMax),
	}
	upToken = putPolicy.UploadToken(client.mac)
	return
}

// 覆盖上传，返回上传凭证,给客户端上传
func (client *QnClient) OverwriteUpload(bucket, keyToOverwrite string) (upToken string) {
	// keyToOverwrite为需要覆盖的文件名
	putPolicy := storage.PutPolicy{
		Scope:      fmt.Sprintf("%s:%s", bucket, keyToOverwrite),
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
		Expires:    uint64(client.cfg.UpTokenExpires),
		MimeLimit:  client.cfg.MimeLimit,
		FsizeMin:   int64(client.cfg.FsizeMin),
		FsizeLimit: int64(client.cfg.FsizeMax),
	}
	upToken = putPolicy.UploadToken(client.mac)
	return
}

// 带回调上传，返回上传凭证,给客户端上传
func (client *QnClient) CallbackUpload(bucket string) (upToken string) {
	putPolicy := storage.PutPolicy{
		Scope:            bucket,
		CallbackURL:      client.cfg.CallbackURL,
		CallbackBody:     `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
		CallbackBodyType: client.cfg.CallbackBodyType,
		Expires:          uint64(client.cfg.UpTokenExpires),
		MimeLimit:        client.cfg.MimeLimit,
		FsizeMin:         int64(client.cfg.FsizeMin),
		FsizeLimit:       int64(client.cfg.FsizeMax),
	}
	upToken = putPolicy.UploadToken(client.mac)
	return
}
