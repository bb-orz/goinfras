# XOAuth Starter

###  Documentation

> Documentation 

> QQ OAuth2 接口文档 ：https://wiki.connect.qq.com/准备工作_oauth2-0

> 微博 OAuth2 接口文档 ：https://open.weibo.com/wiki/Connect/login

> 微信 OAuth2 接口文档 ：https://developers.weixin.qq.com/doc/oplatform/Mobile_App/WeChat_Login/Development_Guide.html


### XOAuth Starter Usage
```
goinfras.RegisterStarter(XOAuth.NewStarter())

```

### XOAuth Config Setting

```
WechatSignSwitch      bool   // 微信三方登录开关
WechatOAuth2AppKey    string // 微信开发者appkey
WechatOAuth2AppSecret string // 微信开发者secret
WeiboSignSwitch       bool   // 微博三方登录开关
WeiboOAuth2AppKey     string // 微博开发者appkey
WeiboOAuth2AppSecret  string // 微博开发者secret
QQSignSwitch          bool   // qq三方登录开关
QQOAuth2AppKey        string // qq开发者appkey
QQOAuth2AppSecret     string // qq开发者secret
```

### XOAuth Usage

1、qq 登录
```
// TODO 前端先获取预授权码 qqprecode ...

var qqprecode string
qqOAuthResult := XOAuth.XQQOAuthManager().Authorize(qqprecode)
Println("qqOAuthResult", qqOAuthResult)
// 获取信息后进入登录或注册流程...

```

2、  微博登录 
```
// TODO 前端先获取预授权码 weiboprecode...
var weiboprecode string
weiboOAuthResult := XOAuth.XWeiboOAuthManager().Authorize(weiboprecode)
Println("weiboOAuthResult", weiboOAuthResult)
// 获取信息后进入登录或注册流程...

```

3、微信登录
```
// TODO 前端先获取预授权码 wechatprecode ...
var wechatprecode string
wechatOAuthResult := XOAuth.XWechatOAuthManager().Authorize(wechatprecode)
Println("wechatOAuthResult", wechatOAuthResult)
// 获取信息后进入登录或注册流程...

```