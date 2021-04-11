package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/liuhongdi/digv31/pkg/md5"
	"math/rand"
	"strconv"
	"time"
)
//token的过期时长
//const TokenExpireDuration = time.Hour * 2
const TokenExpireDuration = time.Minute * 5
//secret,签名时使用
var MySecret = []byte("thisislaoliusecret")

//用来生成token的struct
type MyClaims struct {
	UserToken string `json:"usertoken"`
	jwt.StandardClaims
}
//通过username得到origintoken
func GetOriginToken(username string) (string) {
	now := time.Now().Unix()
	rand := rand.Intn(8999)+1000
	dest := username+strconv.FormatInt(now, 10)+strconv.Itoa(rand)
	return md5.MD5(dest)
}

//创建token
func GenToken(userToken string) (string, error) {
	c := MyClaims{
		userToken, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "my-project",                               // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}

// 解析token
func ParseToken(tokenString string) (*MyClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		fmt.Println("jwt ok")
		fmt.Println(claims.UserToken)
		return claims, nil
	}
	return nil, errors.New("invalid token")
}