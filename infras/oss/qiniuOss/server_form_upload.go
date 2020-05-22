package qiniuOss

import (
	"bytes"
	"context"
	"github.com/qiniu/api.v7/v7/storage"
)

/*
表单上传
@param bucket string 指定上传的bucket
@param fileKey string 文件唯一key
@param localFilePath string 本地文件路径
*/
func FormUploadWithLocalFile(bucket, localFilePath, fileKey string) (storage.PutRet, error) {
	putPolicy := storage.PutPolicy{
		Scope:      bucket,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
		Expires:    uint64(config.OssConf.Qiniu.UpTokenExpires),
		MimeLimit:  config.OssConf.Qiniu.MimeLimit,
		FsizeMin:   int64(config.OssConf.Qiniu.FsizeMin),
		FsizeLimit: int64(config.OssConf.Qiniu.FsizeMax),
	}
	upToken := putPolicy.UploadToken(oss.Qiniu)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = config.OssConf.Qiniu.UseHTTPS
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = config.OssConf.Qiniu.UseCdnDomains
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

// 字节数组上传
func FormUploadWithByteSlice(bucket, fileKey string, data []byte) (storage.PutRet, error) {
	putPolicy := storage.PutPolicy{
		Scope:      bucket,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
		Expires:    uint64(config.OssConf.Qiniu.UpTokenExpires),
		MimeLimit:  config.OssConf.Qiniu.MimeLimit,
		FsizeMin:   int64(config.OssConf.Qiniu.FsizeMin),
		FsizeLimit: int64(config.OssConf.Qiniu.FsizeMax),
	}
	upToken := putPolicy.UploadToken(oss.Qiniu)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = config.OssConf.Qiniu.UseHTTPS
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = config.OssConf.Qiniu.UseCdnDomains
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
