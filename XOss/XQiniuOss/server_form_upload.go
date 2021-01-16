package XQiniuOss

import (
	"bytes"
	"context"
	"github.com/qiniu/api.v7/v7/storage"
)

func (client *QnClient) FormUploadWithLocalFile(localFilePath, fileKey string) (storage.PutRet, error) {
	return client.FormUploadWithLocalFileToBucket(client.cfg.DefaultBucket, localFilePath, fileKey)
}

/*
表单上传
@param bucket string 指定上传的bucket
@param fileKey string 文件唯一key
@param localFilePath string 本地文件路径
*/
func (client *QnClient) FormUploadWithLocalFileToBucket(bucket, localFilePath, fileKey string) (storage.PutRet, error) {
	putPolicy := storage.PutPolicy{
		Scope:      bucket,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
		Expires:    uint64(client.cfg.UpTokenExpires),
		MimeLimit:  client.cfg.MimeLimit,
		FsizeMin:   int64(client.cfg.FsizeMin),
		FsizeLimit: int64(client.cfg.FsizeMax),
	}
	upToken := putPolicy.UploadToken(client.mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = client.cfg.UseHTTPS
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = client.cfg.UseCdnDomains
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)

	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{
		// Params: map[string]string{
		// 	"x:name": "github logo",
		// },
	}
	err := formUploader.PutFile(context.Background(), &ret, upToken, fileKey, localFilePath, &putExtra)
	return ret, err
}

func (client *QnClient) FormUploadWithByteSlice(fileKey string, data []byte) (storage.PutRet, error) {
	return client.FormUploadWithByteSliceToBucket(client.cfg.DefaultBucket, fileKey, data)
}

// 字节数组上传
func (client *QnClient) FormUploadWithByteSliceToBucket(bucket, fileKey string, data []byte) (storage.PutRet, error) {
	putPolicy := storage.PutPolicy{
		Scope:      bucket,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
		Expires:    uint64(client.cfg.UpTokenExpires),
		MimeLimit:  client.cfg.MimeLimit,
		FsizeMin:   int64(client.cfg.FsizeMin),
		FsizeLimit: int64(client.cfg.FsizeMax),
	}
	upToken := putPolicy.UploadToken(client.mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = client.cfg.UseHTTPS
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = client.cfg.UseCdnDomains
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)

	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{
		// Params: map[string]string{
		// 	"x:name": "github logo",
		// },
	}
	dataLen := int64(len(data))
	err := formUploader.Put(context.Background(), &ret, upToken, fileKey, bytes.NewReader(data), dataLen, &putExtra)
	return ret, err
}
