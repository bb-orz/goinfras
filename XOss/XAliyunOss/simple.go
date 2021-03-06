package XAliyunOss

import (
	aliOss "github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io/ioutil"
	"strings"
)

type CommonAliyunOss struct {
	client *aliOss.Client
}

func (c *CommonAliyunOss) UploadString(objectKeyName, objectValue string) error {
	return c.UploadStringToBucket(defaultBucket, objectKeyName, objectValue)
}

// 上传普通数据
func (c *CommonAliyunOss) UploadStringToBucket(bucketName, objectKeyName, objectValue string) error {
	// 获取存储空间
	bucket, err := c.client.Bucket(bucketName)
	if err != nil {
		return err
	}

	// 指定存储类型为标准存储，缺省也为标准存储。
	storageType := aliOss.ObjectStorageClass(aliOss.StorageStandard)
	// 指定存储类型为归档存储。
	// storageType := oss.ObjectStorageClass(oss.StorageArchive)

	// 指定访问权限为公共读，缺省为继承bucket的权限。
	objectAcl := aliOss.ObjectACL(aliOss.ACLPublicRead)

	// 上传字符串。
	return bucket.PutObject(objectKeyName, strings.NewReader(objectValue), storageType, objectAcl)
}

func (c *CommonAliyunOss) AppendUpload(objectKeyName string, appendContents ...string) error {
	return c.AppendUploadToBucket(defaultBucket, objectKeyName, appendContents...)
}

// 追加上传
func (c *CommonAliyunOss) AppendUploadToBucket(bucketName, objectKeyName string, appendContents ...string) error {
	// 获取存储空间。
	bucket, err := c.client.Bucket(bucketName)
	if err != nil {
		return err
	}

	var nextPos int64 = 0
	for _, content := range appendContents {
		// 第一次追加的位置是0，返回值为下一次追加的位置。后续追加的位置是追加前文件的长度。
		nextPos, err = bucket.AppendObject(objectKeyName, strings.NewReader(content), nextPos)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *CommonAliyunOss) Uploadfile(objectKeyName, localFilePath string) error {
	return c.UploadfileToBucket(defaultBucket, objectKeyName, localFilePath)
}

// 上传普通文件
func (c *CommonAliyunOss) UploadfileToBucket(bucketName, objectKeyName, localFilePath string) error {

	// 获取存储空间。
	bucket, err := c.client.Bucket(bucketName)
	if err != nil {
		return err
	}

	// 上传本地文件。
	return bucket.PutObjectFromFile(objectKeyName, localFilePath)
}

func (c *CommonAliyunOss) StreamDownload(objectKeyName string) ([]byte, error) {
	return c.StreamDownloadFromBucket(defaultBucket, objectKeyName)
}

// 流下载
func (c *CommonAliyunOss) StreamDownloadFromBucket(bucketName, objectKeyName string) ([]byte, error) {
	// 获取存储空间。
	bucket, err := c.client.Bucket(bucketName)
	if err != nil {
		return nil, err
	}

	// 下载文件到流。
	body, err := bucket.GetObject(objectKeyName)
	if err != nil {
		return nil, err
	}
	// 数据读取完成后，获取的流必须关闭，否则会造成连接泄漏，导致请求无连接可用，程序无法正常工作。
	defer func() {
		body.Close()
	}()

	data, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *CommonAliyunOss) RangeDownload(objectKeyName string, start, end int64) ([]byte, error) {
	return c.RangeDownloadFromBucket(defaultBucket, objectKeyName, start, end)
}

// 仅需要文件中的部分数据，您可以使用范围下载
func (c *CommonAliyunOss) RangeDownloadFromBucket(bucketName, objectKeyName string, start, end int64) ([]byte, error) {
	// 获取存储空间。
	bucket, err := c.client.Bucket(bucketName)
	if err != nil {
		return nil, err
	}

	// 如获取15~35字节范围内的数据，包含15和35，共21个字节的数据。
	// 如果指定的范围无效（比如开始或结束位置的指定值为负数，或指定值大于文件大小），则下载整个文件。
	body, err := bucket.GetObject(objectKeyName, aliOss.Range(start, end))
	if err != nil {
		return nil, err
	} // 数据读取完成后，获取的流必须关闭，否则会造成连接泄漏，导致请求无连接可用，程序无法正常工作。
	defer func() {
		body.Close()
	}()

	data, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *CommonAliyunOss) DownLoadFile(objectKeyName, dstFilePath string) error {
	return c.DownLoadFileFromBucket(defaultBucket, objectKeyName, dstFilePath)
}

// 下载文件到本地
func (c *CommonAliyunOss) DownLoadFileFromBucket(bucketName, objectKeyName, dstFilePath string) error {
	// 获取存储空间。
	bucket, err := c.client.Bucket(bucketName)
	if err != nil {
		return err
	}

	// 下载文件到本地文件。
	err = bucket.GetObjectToFile(objectKeyName, dstFilePath)
	if err != nil {
		return err
	}

	return nil
}

func (c *CommonAliyunOss) CompressDownload(objectKeyName, dstFilePath string) error {
	return c.CompressDownloadFromBucket(defaultBucket, objectKeyName, dstFilePath)
}

// 文件压缩下载
func (c *CommonAliyunOss) CompressDownloadFromBucket(bucketName, objectKeyName, dstFilePath string) error {
	// 获取存储空间。
	bucket, err := c.client.Bucket(bucketName)
	if err != nil {
		return err
	}

	// 文件压缩下载。
	err = bucket.GetObjectToFile(objectKeyName, dstFilePath, aliOss.AcceptEncoding("gzip"))
	if err != nil {
		return err
	}

	return nil
}

func (c *CommonAliyunOss) LimitConditionDownload(objectKeyName, dstFilePath string, options ...aliOss.Option) error {
	return c.LimitConditionDownloadFromBucket(defaultBucket, objectKeyName, dstFilePath, options...)

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
func (c *CommonAliyunOss) LimitConditionDownloadFromBucket(bucketName, objectKeyName, dstFilePath string, options ...aliOss.Option) error {
	// 获取存储空间。
	bucket, err := c.client.Bucket(bucketName)
	if err != nil {
		return err
	}

	// 限定条件不满足，不下载文件。
	return bucket.GetObjectToFile(objectKeyName, dstFilePath, options...)
}
