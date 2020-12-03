package jwt

import (
	"GoWebScaffold/infras"
)

var tku ITokenUtils

func JWTComponent() ITokenUtils {
	infras.Check(tku)
	return tku
}

func SetComponent(t ITokenUtils) {
	tku = t
}
