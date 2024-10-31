package redis

import (
	"errors"
	"fmt"
	"time"
)

func StoreToken(token string) error {
	claims, err := ParseJWT(token)
	if err != nil {
		fmt.Println("Error parsing JWT:", err)
		return err
	}

	userID, ok := claims["userId"].(string)
	if !ok {
		fmt.Println("Error: userID not found in token claims")
		return errors.New("invalid token: userID not found")
	}

	// Debug statements
	fmt.Println("Storing token for userID:", userID)
	fmt.Println("Token:", token)

	err = Rdb.SetEX(Ctx, "token:"+userID, token, 30*time.Minute).Err()
	if err != nil {
		fmt.Println("Error storing token in Redis:", err)
	}
	fmt.Println("token stored sucessfully")
	return err
}
