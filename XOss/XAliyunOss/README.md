# XAliyunOss Starter

> 基于 github.com/aliyun/aliyun-oss-go-sdk 包

### AliyunOss Documentation

> Documentation https://help.aliyun.com/product/31815.html?spm=a2c4g.750001.list.18.4afe7b13ZutdDu


> Example: https://github.com/aliyun/aliyun-oss-go-sdk/tree/master/sample

> Crypto Example： https://github.com/aliyun/aliyun-oss-go-sdk/blob/master/sample_crypto/sample_crypto.go

### XAliyunOss Starter Usage
```
goinfras.RegisterStarter(XAliyunOss.NewStarter())

```

### XAliyunOss Config Setting

```
AccessKeySecret string // 开发者AccessKeySecret
ConnTimeout     int    // 请求超时时间，包括连接超时、Socket读写超时，单位秒,默认连接超时30秒，读写超时60秒
RWTimeout       int    // 读写超时设置
EnableMD5       bool   // 是否开启MD5校验。推荐使用CRC校验，CRC的效率高于MD5
EnableCRC       bool   // 是否开启CRC校验
AuthProxy       string // 带账号密码的代理服务器
Proxy           string // 代理服务器，如http://8.8.8.8:3128
AccessKeyId     string //
BucketName      string // 存储库名
Endpoint        string // 机房节点
UseCname        bool   // 是否使用自定义域名CNAME
SecurityToken   string // 临时用户的SecurityToken
```

### XAliyunOss Usage

// 实例使用请查看run_test.go
```

// 简单通用的上传下载实例
XAliyunOss.XCommonOss()

// 断点上传下载实例
XAliyunOss.XBreakPointOss()

// 分块上传下载实例
XAliyun.XMultipartOss()

// 含进度的上传下载实例
XAliyunOss.XProgressOss()

```