package XQiniuOss

import (
	"context"
	"github.com/qiniu/api.v7/v7/storage"
)

/*
分块上传
@param bucket string 指定上传的bucket
@param fileKey string 文件唯一key
@param localFilePath string 本地文件路径
*/
func (client *QnClient) MultipartUpload(bucket, localFilePath, fileKey string) (storage.PutRet, error) {
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	upToken := putPolicy.UploadToken(client.mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = client.cfg.UseHTTPS
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = client.cfg.UseCdnDomains
	resumeUploader := storage.NewResumeUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.RputExtra{}
	err := resumeUploader.PutFile(context.Background(), &ret, upToken, fileKey, localFilePath, &putExtra)
	return ret, err
}

// 断点续传
