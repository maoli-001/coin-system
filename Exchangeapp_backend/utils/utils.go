package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// 对用户密码进行加密
// 传入string类型密码，返回加密后的密码和错误信息
func HashPassword(pwd string) (string, error) {
	//bcrypt恒诚加密密码，加密强度12
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), 12)
	return string(hash), err
}

// 生成JWT令牌
// 传入用户名，返回JWT令牌和错误信息
func GenerateJWT(username string) (string, error) {
	//使用HS256算法创建JWTtoken
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		//存入用户名
		"username": username,
		//设置过期时间
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})
	//使用密钥“secret”进行签名
	signedToken, err := token.SignedString([]byte("secret"))
	return "Bearer " + signedToken, err
}

// 验证密码的正确性
func CheckPassword(password string, hash string) bool {
	//用bcrypt内置函数验证密码是否正确
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// 解析JWT令牌，传入token,返回用户名或错误
func ParseJWT(tokenString string) (string, error) {
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	//JWT官方解析函数
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//检查签名是不是我们预期的HS256，防止伪造
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte("secret"), nil
	})

	if err != nil {
		return "", err
	}

	//拿到token信息
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//拿出同户名，返回给中间件
		username, ok := claims["username"].(string)
		if !ok {
			return "", errors.New("username claim is not a string")
		}
		return username, nil
	}

	//token无效返回
	return "", err
}
