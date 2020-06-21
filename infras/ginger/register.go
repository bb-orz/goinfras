package ginger

/*
 注册服务所提供的api，由具体项目实现
*/

// 每个模块服务应该实现的接口
type IApiModule interface {
	SetRoutes() // 模块服务应该实现的方法，各模块启动器设置相应路由
}

// API模块注册器
type ApiModuleRegister struct {
	modules []IApiModule
}

// 注册API模块
func (register *ApiModuleRegister) Register(module IApiModule) {
	register.modules = append(register.modules, module)
}

// 初始化API模块注册器
var apiRegister = new(ApiModuleRegister)

// 注册WEB API初始化对象
func RegisterApiModule(module IApiModule) {
	apiRegister.Register(module)
}

// 获取注册的web api初始化对象
func GetApiModules() []IApiModule {
	return apiRegister.modules
}
