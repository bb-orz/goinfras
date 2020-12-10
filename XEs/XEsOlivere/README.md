# ElasticSearch 客户端

> 基于 github.com/olivere/elastic/v7 包

### ElasticSearch Documentation
> Documentation https://www.elastic.co/guide/cn/elasticsearch/guide/current/index.html


### XEsOlivere Starter Usage
```
// ...

goinfras.Register(XEsOlivere.NewStarter())

// ...
```

### XEsOlivere Config Setting
```
URL         string // 服务地址
	Username    string // 鉴权用户名
	Password    string // 鉴权密码
	Sniff       bool   // 启用或禁用嗅探器。
	Healthcheck bool   // 启用连接健康检查
	Infolog     string // Info级别日志记录文件路径
	Errorlog    string // Error级别日志记录文件路径
	Tracelog    string // Trace级别日志记录文件路径

```


### XEsOlivere Usage
```
// 查找文档是否存在
XEsOlivere.XCommon().IsDocExists(1,"xxxxxxxx")

// 获取
XEsOlivere.XCommon().GetDoc(1,"xxxxxxxx")

// 新增文档
XEsOlivere.XCommon().AddDoc(1,"xxxxxxxx")

// 更新文档
XEsOlivere.XCommon().UpdateDoc(1,"xxxxxxxx")

// 删除文档
XEsOlivere.XCommon().DeleteDoc(1,"xxxxxxxx")
```