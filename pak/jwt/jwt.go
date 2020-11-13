package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
	jwt.StandardClaims
}

var jwtSecret []byte

//创建token
func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)
	clims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, clims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err

}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if clims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return clims, nil
		}
	}
	return nil, err
}
