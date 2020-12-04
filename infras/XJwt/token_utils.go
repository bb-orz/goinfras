package XJwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type ITokenUtils interface {
	Encode(user UserClaim) (string, error)
	Decode(tokenString string) (*CustomerClaim, error)
}

// JWT中携带的用户个人信息
type UserClaim struct {
	Id     string `json:id`
	Name   string `json:name`
	Avatar string `json:avatar`
}

// 聚合jwt内部实现的Claims
type CustomerClaim struct {
	UserClaim
	*jwt.StandardClaims
}

// 实现token服务
type tokenUtils struct {
	privateKey []byte    // 编解码私钥，在生产环境中，该私钥请使用生成器生成，并妥善保管，此处使用简单字符串。
	expTime    time.Time // 超时秒数
}

func NewTokenUtils(privateKey []byte, expSeconds int) *tokenUtils {
	ts := new(tokenUtils)
	ts.privateKey = privateKey
	ts.expTime = time.Now().Add(time.Second * time.Duration(expSeconds))
	return ts
}

// 传入用户信息编码成token
func (tks *tokenUtils) encode(user UserClaim) (string, error) {
	// privateKey, _ := base64.URLEncoding.DecodeString(string(privateKey))

	// 设置Claim
	customer := CustomerClaim{user, &jwt.StandardClaims{ExpiresAt: tks.expTime.Unix()}}

	// 生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customer)

	return token.SignedString(tks.privateKey)

}

// token字符串解码成用户信息
func (tks *tokenUtils) decode(tokenString string) (*CustomerClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomerClaim{}, func(token *jwt.Token) (interface{}, error) {
		// return base64.URLEncoding.DecodeString(string(privateKey))
		return tks.privateKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claim, ok := token.Claims.(*CustomerClaim); ok && token.Valid {
		return claim, nil
	} else {
		return nil, err
	}
}

func (tks *tokenUtils) Encode(user UserClaim) (string, error) {
	return tks.encode(user)
}

func (tks *tokenUtils) Decode(tokenString string) (*CustomerClaim, error) {
	return tks.decode(tokenString)
}