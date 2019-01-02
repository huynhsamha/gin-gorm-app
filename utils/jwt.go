package utils

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(DefaultGetEnv("JWT_SECRET_KEY", ""))

// JWTCustomClaims :
type JWTCustomClaims struct {
	jwt.StandardClaims
	Payload interface{} `json:"payload"`
}

// JWT : JWT utils
type JWT struct{}

// GenerateToken : JWT method
func (JWT) GenerateToken(data interface{}) (string, error) {
	claims := JWTCustomClaims{
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(7 * time.Second).Unix(),        // only valid in 5s after valid
			NotBefore: time.Now().Add(100 * time.Millisecond).Unix(), // just valid after 100ms
		},
		Payload: data,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(jwtSecret)

	return tokenString, err
}

// ParseToken : JWT method
func (JWT) ParseToken(tokenString string) (interface{}, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return jwtSecret, nil
	})

	if claims, ok := token.Claims.(*JWTCustomClaims); ok && token.Valid {
		return claims.Payload, nil
	}
	return nil, err
}
