package XAliyunOss

import (
	aliOss "github.com/aliyun/aliyun-oss-go-sdk/oss"
)

/*
断点上传和断点下载通用操作
*/
type BreakPointOss struct {
	client *aliOss.Client
}

func (bp *BreakPointOss) BreakPointUpload(bucketName, objectKeyName, localFilePath string) error {
	// 获取存储空间。
	bucket, err := bp.client.Bucket(bucketName)
	if err != nil {
		return err
	}
	// 分片大小100K，3个协程并发上传分片，使用断点续传。
	err = bucket.UploadFile(objectKeyName, localFilePath, 100*1024, aliOss.Routines(3), aliOss.Checkpoint(true, ""))
	if err != nil {
		return err
	}
	return nil
}

func (bp *BreakPointOss) BreakPointDownload(bucketName, objectKeyName, dstFilePath string) error {
	// 获取存储空间。
	bucket, err := bp.client.Bucket(bucketName)
	if err != nil {
		return err
	}
	// 分片下载。3个协程并发下载分片，开启断点续传下载。
	// 其中"<yourObjectName>"为objectKey， "LocalFile"为filePath，100*1024为partSize。
	err = bucket.DownloadFile(objectKeyName, dstFilePath, 100*1024, aliOss.Routines(3), aliOss.Checkpoint(true, ""))
	if err != nil {
		return err
	}
	return nil
}
