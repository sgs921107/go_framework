/*************************************************************************
> File Name: jwt.go
> Author: sgs921107
> Mail: 757513128@gmail.com
> Created Time: 2025-01-01 14:34:22 星期三
> Content: This is a desc
*************************************************************************/

package token

import (
	"errors"
	"sync"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/sgs921107/go_framework/common"
)

var (
	once        sync.Once
	jwtInstance = &JWT{}
)

// 生成基础的声明
func BaseClaims() jwt.MapClaims {
	return jwt.MapClaims{
		"iss": "goframework.com",
		"sub": "web",
	}
}

// 获取签名的key
func KeyFunc(*jwt.Token) (interface{}, error) {
	return []byte(common.Setting.Jwt.Key), nil
}

type JWT struct {
	SignMethod jwt.SigningMethod
	Type       string
	// token的有效期
	Validity time.Duration
	Key      []byte
}

// 生成token
func (j *JWT) GenWithClaims(claims jwt.MapClaims) (string, error) {
	claims["exp"] = time.Now().Add(j.Validity).Unix()
	obj := jwt.NewWithClaims(j.SignMethod, claims)
	return obj.SignedString(j.Key)
}

// 解析token
func (j *JWT) Parse(token string, options ...jwt.ParserOption) (jwt.Claims, error) {
	tokenObj, err := jwt.ParseWithClaims(token, BaseClaims(), KeyFunc, options...)
	if err != nil {
		return nil, err
	}
	return tokenObj.Claims, nil
}

// Valid 验证token是否合法
func (j *JWT) Valid(token string, options ...jwt.ParserOption) error {
	_, err := jwt.ParseWithClaims(token, BaseClaims(), KeyFunc, options...)
	return err
}

// AllowRefresh 是否可以刷新token  token过期一定时间内可以刷新token
func (j *JWT) AllowRefresh(token string, options ...jwt.ParserOption) (bool, error) {
	claims, err := j.Parse(token, options...)
	if err != nil {
		return false, err
	}
	exp, err := claims.GetExpirationTime()
	if err != nil {
		return false, err
	}
	if exp.Add(j.Validity).Before(time.Now()) {
		return false, errors.New("token已超出最大刷新时间")
	}
	return true, nil
}

// NewJWT 获取jwt单例
func NewJWT() *JWT {
	once.Do(func() {
		jwtInstance.SignMethod = jwt.SigningMethodHS256
		jwtInstance.Type = "JWT"
		jwtInstance.Validity = time.Hour * 6
		jwtInstance.Key = []byte(common.Setting.Jwt.Key)
	})
	return jwtInstance
}
