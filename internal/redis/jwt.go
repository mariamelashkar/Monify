package redis

import (
	"os"
	"time"
	"github.com/golang-jwt/jwt"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateToken(userID string) (string, error) {
    claims := jwt.MapClaims{
        "userId": userID,
        "exp":    time.Now().Add(time.Minute * 30).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtSecret)
    return tokenString, err
}
