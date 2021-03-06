package XAliyunOss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func XClient() *oss.Client {
	return aliyunOssClient
}

// 资源组件闭包执行
func XFClient(f func(c *oss.Client) error) error {
	return f(aliyunOssClient)
}

// 通用实例
func XCommonOss() *CommonAliyunOss {
	common := new(CommonAliyunOss)
	common.client = XClient()
	return common
}

// 断点操作实例
func XBreakPointOss() *BreakPointOss {
	bp := new(BreakPointOss)
	bp.client = XClient()
	return bp
}

// 分片上传操作实例
func XMultipartOss() *MultipartOss {
	mp := new(MultipartOss)
	mp.client = XClient()
	return mp
}

// 含进度的上传下载操作实例
func XProgressOss() *ProgressOss {
	p := new(ProgressOss)
	p.client = XClient()
	return p
}
