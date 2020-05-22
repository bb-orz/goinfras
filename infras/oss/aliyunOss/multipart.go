package aliyunOss

import (
	"fmt"
	aliOss "github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

/*
1.分片上传（Multipart Upload）分为以下三个步骤：

2.初始化一个分片上传事件。
调用Bucket.InitiateMultipartUpload方法返回OSS创建的全局唯一的uploadId。

3.上传分片。
调用Bucket.UploadPart方法上传分片数据。

说明
对于同一个uploadId，分片号（partNumber）标识了该分片在整个文件内的相对位置。如果使用同一个分片号上传了新的数据，那么OSS上这个分片已有的数据将会被覆盖。
OSS将收到的分片数据的MD5值放在ETag头内返回给用户。
OSS计算上传数据的MD5值，并与SDK计算的MD5值比较，如果不一致则返回InvalidDigest错误码。
完成分片上传。

所有分片上传完成后，调用Bucket.CompleteMultipartUpload方法将所有分片合并成完整的文件。
*/

func MultipartUpload(bucketName, objectKeyName, localFilePath string) error {
	// 获取存储空间。
	bucket, err := oss.Aliyun.Bucket(bucketName)
	if !e.Ec(err) {
		return err
	}

	chunks, err := aliOss.SplitFileByPartNum(localFilePath, 3)
	fd, err := os.Open(localFilePath)
	defer fd.Close()

	// 指定存储类型为标准存储，缺省也为标准存储。
	storageType := aliOss.ObjectStorageClass(aliOss.StorageStandard)
	// 指定存储类型为归档存储。
	// storageType := oss.ObjectStorageClass(oss.StorageArchive)

	// 步骤1：初始化一个分片上传事件，指定存储类型为标准存储。
	imur, err := bucket.InitiateMultipartUpload(objectKeyName, storageType)
	// 步骤2：上传分片。
	var parts []aliOss.UploadPart
	for _, chunk := range chunks {
		fd.Seek(chunk.Offset, os.SEEK_SET)
		// 对每个分片调用UploadPart方法上传。
		part, err := bucket.UploadPart(imur, fd, chunk.Size, chunk.Number)
		if !e.Ec(err) {
			return err
		}
		parts = append(parts, part)
	}

	// 指定访问权限为公共读，缺省为继承bucket的权限。
	objectAcl := aliOss.ObjectACL(aliOss.ACLPublicRead)

	// 步骤3：完成分片上传，指定访问权限为公共读。
	cmur, err := bucket.CompleteMultipartUpload(imur, parts, objectAcl)
	if !e.Ec(err) {
		return err
	}
	fmt.Println("cmur:", cmur)
	return nil
}

// 取消分片上传
func CancelMultipartUpload(bucketName, objectKeyName string) error {
	// 获取存储空间。
	bucket, err := oss.Aliyun.Bucket(bucketName)
	if !e.Ec(err) {
		return err
	}

	// 初始化一个分片上传事件。
	imur, err := bucket.InitiateMultipartUpload(objectKeyName)
	if !e.Ec(err) {
		return err
	}
	// 取消分片上传。
	err = bucket.AbortMultipartUpload(imur)
	if !e.Ec(err) {
		return err
	}
	return nil
}
