package jwt

import (
	"GoWebScaffold/infras"
	"github.com/tietang/props/kvs"
)

var tku *tokenUtils

func TokenUtils() *tokenUtils {
	infras.Check(tku)
	return tku
}

type JWTStarter struct {
	infras.BaseStarter
	cfg *jwtConfig
}

func (s *JWTStarter) Init(sctx *infras.StarterContext) {
	configs := sctx.Configs()
	define := jwtConfig{}
	err := kvs.Unmarshal(configs, &define, "Jwt")
	infras.FailHandler(err)
	s.cfg = &define
}

func (s *JWTStarter) Setup(sctx *infras.StarterContext) {
	tku = NewTokenUtils([]byte(s.cfg.privateKey), s.cfg.expSeconds)
}

func (s *JWTStarter) Start(sctx *infras.StarterContext) {
}

func (s *JWTStarter) Stop(sctx *infras.StarterContext) {
}
