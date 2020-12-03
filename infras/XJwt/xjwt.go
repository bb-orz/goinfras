package XJwt

var tku ITokenUtils

// 资源组件实例调用
func XTokenUtils() ITokenUtils {
	return tku
}

// 资源组件闭包执行
func XFTokenUtils(f func(t ITokenUtils) error) error {
	return f(tku)
}
