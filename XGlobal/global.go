package XGlobal

func init() {
	_g = make(map[string]interface{})
}

var _g Global

type Global map[string]interface{}

func (g Global) Set(k string, v interface{}) {
	g[k] = v
}

func (g Global) Get(k string) interface{} {
	if v, ok := g[k]; ok {
		return v
	}
	return nil
}
