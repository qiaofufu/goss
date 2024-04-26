package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var hmacSampleSecret = []byte("feaf-=-f#adwacasfa")

type Claims struct {
	ID int64 `json:"id"`
	jwt.RegisteredClaims
}

// GenToken 生成 Token
func GenToken(id int64) (string, error) {
	c := Claims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "goss",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Hour * 12)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                          // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                          // 生效时间
		}}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	tokenStr, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func GetUserIDFromPayload(token string) (int64, error) {
	parser := jwt.NewParser()
	var claims Claims
	_, _, err := parser.ParseUnverified(token, &claims)
	return claims.ID, err
}

func VerifyToken(token string) error {
	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return hmacSampleSecret, nil
	})
	return err
}
