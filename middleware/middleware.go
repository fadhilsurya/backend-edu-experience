package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func GenerateJWT(id uint) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = id
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY_TOKEN")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Missing token"})
			c.Abort()
			return
		}

		checkTokenBlacklist := IsTokenRevoked(tokenString)
		if checkTokenBlacklist {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Token Already Revoked"})
			c.Abort()
			return

		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("SECRET_KEY_TOKEN")), nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			c.Abort()
			return
		}

		expirationTime := int64(claims["exp"].(float64))
		if time.Unix(expirationTime, 0).Before(time.Now()) {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Token has expired"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func GetUserIDFromToken(tokenString string) (int, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY_TOKEN")), nil
	})

	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("invalid claims")
	}

	userIDInt, ok := claims["user_id"].(float64)

	if !ok {
		return 0, fmt.Errorf("user ID not found or not a valid type")
	}

	return int(userIDInt), nil
}

var tokenBlacklist = make(map[string]bool)

func RevokeToken(token string) {
	tokenBlacklist[token] = true
}

func IsTokenRevoked(token string) bool {
	return tokenBlacklist[token]
}
