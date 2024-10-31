package redis

import (
	"fmt"
	"errors"
)

func RemoveToken(tokenStr string) error {
	_, err := ValidateToken(tokenStr)
	if err != nil {
		return err
	}
	claims, err := ParseJWT(tokenStr)
    if err != nil {
        fmt.Println("Error parsing JWT:", err)
        return err
    }

    userID, ok := claims["userId"].(string)
    if !ok {
        fmt.Println("Error: userID not found in token claims")
        return errors.New("invalid token: userID not found")
    }
	return Rdb.Del(Ctx, "token:"+userID).Err()
}