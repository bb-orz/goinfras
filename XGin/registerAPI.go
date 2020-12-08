package XGin

/*
 注册项目所提供的api，由具体项目实现
*/

// 每个模块服务应该实现的接口
type IApi interface {
	SetRoutes() // 模块服务应该实现的方法，各模块启动器设置相应路由
}

// API模块注册器
type ApiRegister struct {
	apis []IApi
}

// 注册API模块
func (register *ApiRegister) Register(module IApi) {
	register.apis = append(register.apis, module)
}

// 初始化API模块注册器
var apiRegister = new(ApiRegister)

// 注册WEB API初始化对象
func RegisterApi(module IApi) {
	apiRegister.Register(module)
}

// 获取注册的web api初始化对象
func GetApis() []IApi {
	return apiRegister.apis
}
