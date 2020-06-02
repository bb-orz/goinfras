package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type ITokenService interface {
	Encode(user UserClaim) (string, error)
	Decode(tokenString string) (*CustomerClaim, error)
}

// JWT中携带的用户个人信息
type UserClaim struct {
	Id     int64  `json:id`
	Name   string `json:name`
	Avatar string `json:avatar`
}

// 聚合jwt内部实现的Claims
type CustomerClaim struct {
	UserClaim
	*jwt.StandardClaims
}

// 实现token服务
type TokenService struct {
	privateKey []byte // 编解码私钥，在生产环境中，该私钥请使用生成器生成，并妥善保管，此处使用简单字符串。
}

func NewTokenService(privateKey []byte) *TokenService {
	ts := new(TokenService)
	ts.privateKey = privateKey
	return ts
}

// 传入用户信息编码成token
func (tks *TokenService) Encode(user UserClaim) (string, error) {
	// privateKey, _ := base64.URLEncoding.DecodeString(string(privateKey))

	// 设置超时时间
	expTime := time.Now().Add(time.Hour * 24 * 3).Unix()

	// 设置Claim
	customer := CustomerClaim{user, &jwt.StandardClaims{ExpiresAt: expTime}}

	// 生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customer)

	return token.SignedString(tks.privateKey)

}

// token字符串解码成用户信息
func (tks *TokenService) Decode(tokenString string) (*CustomerClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomerClaim{}, func(token *jwt.Token) (interface{}, error) {
		// return base64.URLEncoding.DecodeString(string(privateKey))
		return tks.privateKey, nil
	})

	if err != nil {
		//logger.Error("JWT Decode Wrong",
		//	zap.String("path", "util/jwt.TokenService.Decode"),
		//	zap.String("warming", err.Error()),
		//	zap.String("receive_token", tokenString),
		//)
		return nil, err
	}

	if claim, ok := token.Claims.(*CustomerClaim); ok && token.Valid {
		return claim, nil
	} else {
		return nil, err
	}
}
