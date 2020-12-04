package XJwt

import (
	"errors"
	"time"
)

type tokenUtilsX struct {
	tokenUtils
	cache *redisCache
}

func NewTokenUtilsX(privateKey []byte, expSeconds int) *tokenUtilsX {
	ts := new(tokenUtilsX)
	ts.privateKey = privateKey
	ts.expTime = time.Now().Add(time.Second * time.Duration(expSeconds))
	ts.cache = NewRedisCache()
	return ts
}

func (tks *tokenUtilsX) Encode(user UserClaim) (string, error) {
	token, err := tks.encode(user)
	if err != nil {
		return "", err
	}

	// 将token缓存到redis
	if user.Id == "" {
		return "", errors.New("user id empty not allowed")
	}
	err = tks.cache.SetToken(user.Id, token)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (tks *tokenUtilsX) Decode(tokenString string) (*CustomerClaim, error) {
	// 如不能解码，直接返回err
	claim, err := tks.decode(tokenString)
	if err != nil {
		return claim, err
	}

	// redis 鉴定缓存数据
	cacheToken, err := tks.cache.GetToken(claim.UserClaim.Id)
	if cacheToken != tokenString {
		return nil, errors.New("Token string is invalid with cache data ")
	}
	return claim, nil
}
