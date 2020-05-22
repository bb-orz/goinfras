package aliyunOss

import (
	aliOss "github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gofuncchan/ginger/util/e"
)

func BreakPointUpload(bucketName, objectKeyName, localFilePath string) error {
	// 获取存储空间。
	bucket, err := oss.Aliyun.Bucket(bucketName)
	if !e.Ec(err) {
		return err
	}

	// 分片大小100K，3个协程并发上传分片，使用断点续传。
	err = bucket.UploadFile(objectKeyName, localFilePath, 100*1024, aliOss.Routines(3), aliOss.Checkpoint(true, ""))
	if !e.Ec(err) {
		return err
	}

	return nil
}

func BreakPointDownload(bucketName, objectKeyName, dstFilePath string) error {
	// 获取存储空间。
	bucket, err := oss.Aliyun.Bucket(bucketName)
	if !e.Ec(err) {
		return err
	}

	// 分片下载。3个协程并发下载分片，开启断点续传下载。
	// 其中"<yourObjectName>"为objectKey， "LocalFile"为filePath，100*1024为partSize。
	err = bucket.DownloadFile(objectKeyName, dstFilePath, 100*1024, aliOss.Routines(3), aliOss.Checkpoint(true, ""))
	if !e.Ec(err) {
		return err
	}

	return nil
}
