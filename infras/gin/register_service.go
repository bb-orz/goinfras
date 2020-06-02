package gin

/*
 注册服务所提供的api，由具体项目实现
*/

// API初始化接口
type Initializer interface {
	//用于对象实例化后的初始化操作
	Init()
}

//初始化注册器
type InitializeRegister struct {
	Initializers []Initializer
}

//注册一个初始化对象
func (i *InitializeRegister) Register(ai Initializer) {
	i.Initializers = append(i.Initializers, ai)
}

var apiInitializerRegister *InitializeRegister = new(InitializeRegister)

//注册WEB API初始化对象

func RegisterApi(ai Initializer) {
	apiInitializerRegister.Register(ai)
}

//获取注册的web api初始化对象
func GetApiInitializers() []Initializer {
	return apiInitializerRegister.Initializers
}
