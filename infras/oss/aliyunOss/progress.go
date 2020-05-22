package aliyunOss

import (
	"fmt"
	aliOss "github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// 定义进度条监听器。
type OssProgressListener struct {
}

// 定义进度变更事件处理函数。
func (listener *OssProgressListener) ProgressChanged(event *aliOss.ProgressEvent) {
	switch event.EventType {
	case aliOss.TransferStartedEvent:
		fmt.Printf("Transfer Started, ConsumedBytes: %d, TotalBytes %d.\n",
			event.ConsumedBytes, event.TotalBytes)
	case aliOss.TransferDataEvent:
		fmt.Printf("\rTransfer Data, ConsumedBytes: %d, TotalBytes %d, %d%%.",
			event.ConsumedBytes, event.TotalBytes, event.ConsumedBytes*100/event.TotalBytes)
	case aliOss.TransferCompletedEvent:
		fmt.Printf("\nTransfer Completed, ConsumedBytes: %d, TotalBytes %d.\n",
			event.ConsumedBytes, event.TotalBytes)
	case aliOss.TransferFailedEvent:
		fmt.Printf("\nTransfer Failed, ConsumedBytes: %d, TotalBytes %d.\n",
			event.ConsumedBytes, event.TotalBytes)
	default:
	}
}

// 上传使用进度条
func ProgressUpload(bucketName, objectKeyName, localFilePath string) error {
	// 获取存储空间
	bucket, err := oss.Aliyun.Bucket(bucketName)
	if !e.Ec(err) {
		return err
	}

	// 带进度条的上传。
	err = bucket.PutObjectFromFile(objectKeyName, localFilePath, aliOss.Progress(&OssProgressListener{}))
	if !e.Ec(err) {
		return err
	}

	return nil
}

// 下载使用进度条
func ProgressDownload(bucketName, objectKeyName, dstFilePath string) error {
	// 获取存储空间
	bucket, err := oss.Aliyun.Bucket(bucketName)
	if !e.Ec(err) {
		return err
	}

	// 带进度条的下载。
	err = bucket.GetObjectToFile(objectKeyName, dstFilePath, aliOss.Progress(&OssProgressListener{}))
	if !e.Ec(err) {
		return err
	}
	return nil
}
