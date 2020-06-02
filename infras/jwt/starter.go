package jwt

import (
	"GoWebScaffold/infras"
	"github.com/tietang/props/kvs"
)

var tokenService *TokenService

func TokenUtil() *TokenService {
	infras.Check(tokenService)
	return tokenService
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
	tokenService = NewTokenService([]byte(s.cfg.privateKey))
}

func (s *JWTStarter) Start(sctx *infras.StarterContext) {
}

func (s *JWTStarter) Stop(sctx *infras.StarterContext) {
}
