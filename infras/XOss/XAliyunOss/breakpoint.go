package XAliyunOss

import (
	aliOss "github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type BreakPointOss struct{}

func NewBreakPointOss() *BreakPointOss {
	return new(BreakPointOss)
}

func (*BreakPointOss) BreakPointUpload(bucketName, objectKeyName, localFilePath string) error {
	// 获取存储空间。
	bucket, err := AliyunOssComponent().Bucket(bucketName)
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

func (*BreakPointOss) BreakPointDownload(bucketName, objectKeyName, dstFilePath string) error {
	// 获取存储空间。
	bucket, err := AliyunOssComponent().Bucket(bucketName)
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
