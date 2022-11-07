package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"strings"
	"wTools/enum/code"
	"wTools/global"
	"wTools/response"
)

type CustomClaim struct {
	Name     string `json:"name"`
	ExpireAt int64  `json:"expireAt"`
	Code     string `json:"code"`
	jwt.StandardClaims
}

type JWT struct {
	SigningKey []byte
}

func JETAuth(c *gin.Context) {
	authorization := getAuthorization(c)

	if authorization == "" {
		response.Result(code.Unauthorized, "未登录或非法访问", nil, c)
		c.Abort()
		return
	}

	claims, err := NewJWT().ParseToken(authorization)
	if err != nil {
		response.Result(code.Unauthorized, "登陆授权校验失败", nil, c)
		c.Abort()
	}
	userInfo := claims.Code
	fmt.Println(userInfo)

}

// getAuthorization
//  @Description: 获取 token 正文
//  @param c
//  @return string
//  @author	wujingfeng 2022-07-01 15:17:05
func getAuthorization(c *gin.Context) string {
	authorization := c.Request.Header.Get("Authorization")
	if authorization == "" {
		authorization = c.Request.Header.Get("authorization")
	}

	authorization = strings.TrimLeft(authorization, "Bearer")
	authorization = strings.TrimSpace(authorization)

	return authorization
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.Config.JWT.SigningKey),
	}
}

func (j *JWT) ParseToken(tokenString string) (CustomClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, CustomClaim{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return CustomClaim{}, err
	}

	if token == nil {
		if claims, ok := token.Claims.(CustomClaim); ok && token.Valid {
			return claims, nil
		}
	}
	return CustomClaim{}, err
}
