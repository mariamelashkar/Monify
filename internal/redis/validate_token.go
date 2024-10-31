package redis

import (
	"errors"
	"github.com/go-redis/redis/v8"
"github.com/golang-jwt/jwt/v4"
 	"fmt"
)

func ValidateToken(tokenStr string) (jwt.MapClaims, error) {
    claims, err := ParseJWT(tokenStr)
    if err != nil {
        if ve, ok := err.(*jwt.ValidationError); ok {
            if ve.Errors&jwt.ValidationErrorExpired != 0 {
                if userID, ok := claims["userId"].(string); ok {
                    Rdb.Del(Ctx, "token:"+userID)
                }
                return nil, errors.New("token is expired")
            }
        }
        return nil, err
    }

    userID, ok := claims["userId"].(string)
    if !ok {
        return nil, errors.New("user ID not found in token claims")
    }

    storedToken, err := Rdb.Get(Ctx, "token:"+userID).Result()
    if err == redis.Nil || storedToken != tokenStr {
        fmt.Println("invalid token")
        return nil, errors.New("invalid token")
    } else if err != nil {
        return nil, err
    }

    return claims, nil
}
