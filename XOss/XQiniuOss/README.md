# XQiniuOss Starter

> 基于 https://github.com/qiniu/api.v7 包

### QiniuOss Documentation

> Documentation https://developer.qiniu.com/kodo/sdk/1238/go

> Example https://github.com/qiniu/api.v7/tree/master/examples



### XQiniuOss Starter Usage
```
goinfras.RegisterStarter(XQiniuOss.NewStarter())

```

### XQiniuOss Config Setting

```
AccessKey        string // 开发者key
SecretKey        string // 开发者secret
Bucket           string // 存储库名
UseHTTPS         bool   // 是否使用https域名
UseCdnDomains    bool   // 上传是否使用CDN上传加速
UpTokenExpires   int    // 上传凭证有效期
CallbackURL      string // 上传回调地址
CallbackBodyType string // 上传回调信息格式
EndUser          string // 唯一宿主标识
FsizeMin         int    // 限定上传文件大小最小值，单位Byte。
FsizeMax         int    // 限定上传文件大小最大值，单位Byte。超过限制上传文件大小的最大值会被判为上传失败，返回 413 状态码。
MimeLimit        string // 限定上传类型

```

### XQiniuOss Usage

1 、 客户端上传下载
```
var upToken string
upToken = XQiniuOss.XClient().SimpleUpload("bucket")
Println("Client Upload Token:", upToken)

upToken := XQiniuOss.XClient().OverwriteUpload("bucket", "keyToOverwrite")
Println("Client Overwrite Upload Token:", upToken)

callbackUploadToken := XQiniuOss.XClient().CallbackUpload("bucket")
Println("Client Callback Upload Token:", callbackUploadToken)

```


2、 服务端断点上传
```
putRet, err := XQiniuOss.XClient().BreakPointUpload("bucket", "fileKey", "localFilePath", "recordDir")
So(err, ShouldBeNil)
Println("BreakPointUpload Key:", putRet.Key)
Println("BreakPointUpload PersistentID:", putRet.PersistentID)
Println("BreakPointUpload Hash:", putRet.Hash)

```

3、服务端表单上传
```
// 服务器表单上传
putRet1, err := XQiniuOss.XClient().FormUploadWithLocalFile("bucket", "localFilePath", "fileKey")
So(err, ShouldBeNil)
Println("FormUploadWithLocalFile Key:", putRet1.Key)
Println("FormUploadWithLocalFile PersistentID:", putRet1.PersistentID)
Println("FormUploadWithLocalFile Hash:", putRet1.Hash)

// 服务器字节数组上传
var data []byte
putRet2, err := XQiniuOss.XClient().FormUploadWithByteSlice("bucket", "fileKey", data)
So(err, ShouldBeNil)
Println("FormUploadWithByteSlice Key:", putRet2.Key)
Println("FormUploadWithByteSlice PersistentID:", putRet2.PersistentID)
Println("FormUploadWithByteSlice Hash:", putRet2.Hash)
```

4、服务端分块上传

```
putRet, err := XQiniuOss.XClient().MultipartUpload("bucket", "localFilePath", "fileKey")
So(err, ShouldBeNil)
Println("MultipartUpload Key:", putRet.Key)
Println("MultipartUpload PersistentID:", putRet.PersistentID)
Println("MultipartUpload Hash:", putRet.Hash)

```