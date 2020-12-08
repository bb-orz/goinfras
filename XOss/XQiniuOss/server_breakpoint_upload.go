package XQiniuOss

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/qiniu/api.v7/v7/storage"
	"golang.org/x/net/context"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

func md5Hex(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

type ProgressRecord struct {
	Progresses []storage.BlkputRet `json:"progresses"`
}

/*
@param bucket string 指定上传的bucket
@param fileKey string 文件唯一key
@param localFilePath string 本地文件路径
@param recordKey string 指定的进度文件保存目录，实际情况下，请确保该目录存在，而且只用于记录进度文件

*/
func (client *QnClient) BreakPointUpload(bucket, fileKey, localFilePath, recordDir string) (storage.PutRet, error) {
	var err error
	ret := storage.PutRet{}

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

	// 必须仔细选择一个能标志上传唯一性的 recordKey 用来记录上传进度
	// 我们这里采用 md5(bucket+key+local_path+local_file_last_modified)+".progress" 作为记录上传进度的文件名
	fileInfo, statErr := os.Stat(localFilePath)
	if statErr != nil {
		return ret, statErr
	}
	fileSize := fileInfo.Size()
	fileLmd := fileInfo.ModTime().UnixNano()
	recordKey := md5Hex(fmt.Sprintf("%s:%s:%s:%d", bucket, fileKey, localFilePath, fileLmd)) + ".progress"
	err = os.MkdirAll(recordDir, 0755)
	if err != nil {
		return ret, err
	}
	recordPath := filepath.Join(recordDir, recordKey)
	progressRecord := ProgressRecord{}
	// 尝试从旧的进度文件中读取进度
	recordFp, openErr := os.Open(recordPath)
	if openErr == nil {
		progressBytes, readErr := ioutil.ReadAll(recordFp)
		if readErr == nil {
			mErr := json.Unmarshal(progressBytes, &progressRecord)
			if mErr == nil {
				// 检查context 是否过期，避免701错误
				for _, item := range progressRecord.Progresses {
					if storage.IsContextExpired(item) {
						fmt.Println(item.ExpiredAt)
						progressRecord.Progresses = make([]storage.BlkputRet, storage.BlockCount(fileSize))
						break
					}
				}
			}
		}
		recordFp.Close()
	}
	if len(progressRecord.Progresses) == 0 {
		progressRecord.Progresses = make([]storage.BlkputRet, storage.BlockCount(fileSize))
	}
	resumeUploader := storage.NewResumeUploader(&cfg)
	progressLock := sync.RWMutex{}
	putExtra := storage.RputExtra{
		Progresses: progressRecord.Progresses,
		Notify: func(blkIdx int, blkSize int, ret *storage.BlkputRet) {
			progressLock.Lock()
			progressLock.Unlock()
			// 将进度序列化，然后写入文件
			progressRecord.Progresses[blkIdx] = *ret
			progressBytes, _ := json.Marshal(progressRecord)
			fmt.Println("write progress file", blkIdx, recordPath)
			wErr := ioutil.WriteFile(recordPath, progressBytes, 0644)
			if wErr != nil {
				fmt.Println("write progress file error,", wErr)
			}
		},
	}
	err = resumeUploader.PutFile(context.Background(), &ret, upToken, fileKey, localFilePath, &putExtra)
	if err != nil {
		return ret, err
	}
	// 上传成功之后，删除这个进度文件
	os.Remove(recordPath)
	return ret, nil
}
