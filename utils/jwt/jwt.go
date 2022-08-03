package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const TokenExpireDuration = time.Hour * 1 //设置超时时间1h

var mySecret = []byte("RedBubble是一个社交网络平台") //密钥，不能泄漏

// GenerateToken 生成JWT
func GenerateToken(userID int64, username string) (string, error) {

	/*
		  payload：jwt提供的7个标准字段
			iss(issuer)			    签发者，其值为大小写敏感的字符串或Uri
			sub(subject)	        主题,用于鉴别一个用户
			exp(expiration time)	过期时间
			aud(audience)			受众
			iat(issued at)			签发时间
			nbf(not before)			生效时间
			jti(JWT ID)				编号
	*/

	// 1、自定义payload内容，并生成token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  userID,                                     //用户id
		"username": username,                                   //用户名
		"exp":      time.Now().Add(TokenExpireDuration).Unix(), //过期时间
		"iss":      "RedBubble",                                //签发人
	})

	// 2、使用密钥签名并获得完整的编码后的jwt字符串
	return token.SignedString(mySecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (claims jwt.MapClaims, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				// That's not even a token
				return nil, err
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// Token is either expired or not active yet
				return nil, errors.New("令牌已过期或尚未激活")
			} else {
				// Couldn't handle this token
				return nil, errors.New("无法处理此令牌")
			}
		} else {
			// Couldn't handle this token
			return nil, errors.New("无法处理此令牌")
		}
	}
	if !token.Valid {
		return nil, errors.New("令牌无效")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("令牌解析claims失败")
	}

	return claims, nil
}
