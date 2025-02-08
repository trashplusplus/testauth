package crypto

import (
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("secret")

func CreateToken(username string, userID int, email string) (string, error) {
	//створюємо новий токен
	token := jwt.New(jwt.SigningMethodHS256)

	//встановлюємо поля
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["id"] = userID
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() //термін дії

	//підпис токена з приватним ключом
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	//парсимо та верифікуємо токен
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//підписуємо ключ
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	//перевіряємо валідність токену
	if !token.Valid {
		return nil, fmt.Errorf("token is not valid")
	}

	//витягуємо клейми з токену
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}

func VerifyTokenBearer(tokenString string) (jwt.MapClaims, error) {

	//парсимо та верифікуємо токен
	token, err := jwt.Parse(strings.TrimPrefix(tokenString, "Bearer "), func(token *jwt.Token) (interface{}, error) {
		//підписуємо ключ
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	//перевіряємо що валідний
	if !token.Valid {
		return nil, fmt.Errorf("token is not valid")
	}

	//витягуємо клейми з токену
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
