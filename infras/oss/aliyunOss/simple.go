package aliyunOss

import (
	aliOss "github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io/ioutil"
	"strings"
)

// 上传普通数据
func UploadString(bucketName, objectKeyName, objectValue string) error {

	// 获取存储空间
	bucket, err := oss.Aliyun.Bucket(bucketName)
	if !e.Ec(err) {
		return err
	}

	// 指定存储类型为标准存储，缺省也为标准存储。
	storageType := aliOss.ObjectStorageClass(aliOss.StorageStandard)
	// 指定存储类型为归档存储。
	// storageType := oss.ObjectStorageClass(oss.StorageArchive)

	// 指定访问权限为公共读，缺省为继承bucket的权限。
	objectAcl := aliOss.ObjectACL(aliOss.ACLPublicRead)

	// 上传字符串。
	err = bucket.PutObject(objectKeyName, strings.NewReader(objectValue), storageType, objectAcl)
	if !e.Ec(err) {
		return err
	}

	return nil
}

// 追加上传
func AppendUpload(bucketName, objectKeyName string, appendContents ...string) error {
	// 获取存储空间。
	bucket, err := oss.Aliyun.Bucket(bucketName)
	if !e.Ec(err) {
		return err
	}

	var nextPos int64 = 0
	for _, content := range appendContents {
		// 第一次追加的位置是0，返回值为下一次追加的位置。后续追加的位置是追加前文件的长度。
		nextPos, err = bucket.AppendObject(objectKeyName, strings.NewReader(content), nextPos)
		if !e.Ec(err) {
			return err
		}
	}

	return nil
}

// 上传普通文件
func Uploadfile(bucketName, objectKeyName, localFilePath string) error {

	// 获取存储空间。
	bucket, err := oss.Aliyun.Bucket(bucketName)
	if !e.Ec(err) {
		return err
	}

	// 上传本地文件。
	err = bucket.PutObjectFromFile(objectKeyName, localFilePath)
	if !e.Ec(err) {
		return err
	}
	return nil
}

// 流下载
func StreamDownload(bucketName, objectKeyName string) ([]byte, error) {
	// 获取存储空间。
	bucket, err := oss.Aliyun.Bucket(bucketName)
	if !e.Ec(err) {
		return nil, err
	}

	// 下载文件到流。
	body, err := bucket.GetObject(objectKeyName)
	if !e.Ec(err) {
		return nil, err
	}
	// 数据读取完成后，获取的流必须关闭，否则会造成连接泄漏，导致请求无连接可用，程序无法正常工作。
	defer body.Close()

	data, err := ioutil.ReadAll(body)
	if !e.Ec(err) {
		return nil, err
	}
	return data, nil
}

// 仅需要文件中的部分数据，您可以使用范围下载
func RangeDownload(bucketName, objectKeyName string, start, end int64) ([]byte, error) {
	// 获取存储空间。
	bucket, err := oss.Aliyun.Bucket(bucketName)
	if !e.Ec(err) {
		return nil, err
	}

	// 如获取15~35字节范围内的数据，包含15和35，共21个字节的数据。
	// 如果指定的范围无效（比如开始或结束位置的指定值为负数，或指定值大于文件大小），则下载整个文件。
	body, err := bucket.GetObject(objectKeyName, aliOss.Range(start, end))
	if !e.Ec(err) {
		return nil, err
	}
	// 数据读取完成后，获取的流必须关闭，否则会造成连接泄漏，导致请求无连接可用，程序无法正常工作。
	defer body.Close()

	data, err := ioutil.ReadAll(body)
	if !e.Ec(err) {
		return nil, err
	}

	return data, nil
}

// 下载文件到本地
func DownLoadFile(bucketName, objectKeyName, dstFilePath string) error {
	// 获取存储空间。
	bucket, err := oss.Aliyun.Bucket(bucketName)
	if !e.Ec(err) {
		return err
	}

	// 下载文件到本地文件。
	err = bucket.GetObjectToFile(objectKeyName, dstFilePath)
	if !e.Ec(err) {
		return err
	}

	return nil
}

// 文件压缩下载
func CompressDownload(bucketName, objectKeyName, dstFilePath string) error {
	// 获取存储空间。
	bucket, err := oss.Aliyun.Bucket(bucketName)
	if !e.Ec(err) {
		return err
	}

	// 文件压缩下载。
	err = bucket.GetObjectToFile(objectKeyName, dstFilePath, aliOss.AcceptEncoding("gzip"))
	if !e.Ec(err) {
		return err
	}

	return nil
}

// 限定条件下载
/*
传入options条件参数：

参数						描述																						如何设置
IfModifiedSince			如果指定的时间早于实际修改时间，则正常传输文件，否则返回错误（304 Not modified）。					oss.IfModifiedSince
IfUnmodifiedSince		如果指定的时间等于或者晚于文件实际修改时间，则正常传输文件，否则返回错误（412 Precondition failed）。	oss.IfUnmodifiedSince
IfMatch					如果指定的ETag和OSS文件的ETag匹配，则正常传输文件，否则返回错误（412 Precondition failed）。		oss.IfMatch
IfNoneMatch				如果指定的ETag和OSS文件的ETag不匹配，则正常传输文件，否则返回错误（304 Not modified）。				oss.IfNoneMatch
*/
func LimitConditionDownload(bucketName, objectKeyName, dstFilePath string, options ...aliOss.Option) error {
	// 获取存储空间。
	bucket, err := oss.Aliyun.Bucket(bucketName)
	if !e.Ec(err) {
		return err
	}

	// 限定条件不满足，不下载文件。
	err = bucket.GetObjectToFile(objectKeyName, dstFilePath, options...)
	if !e.Ec(err) {
		return err
	}
	return nil
}
