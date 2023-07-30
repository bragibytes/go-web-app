package controllers

import (
	"os"
	"time"

	"github.com/dedpidgon/go-web-app/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Custom claims for JWT
type customClaims struct {
	UserID string
	jwt.StandardClaims
}

type token_controller struct {
	key    string
	header string
}

func new_token_controller() *token_controller {
	x := &token_controller{}
	x.key = os.Getenv("TOKEN_KEY")
	x.header = "Authorization"
	return x
}

// Middleware to verify JWT and extract user information
func (tc *token_controller) token_wall() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader(tc.header)

		if tokenString == "" {
			response.Unauthorized(c, "no token")
			return
		}

		token, err := tc.check(tokenString)
		if err != nil {
			response.Unauthorized(c, "invalid token")
			return
		}

		if claims, ok := token.Claims.(*customClaims); ok && token.Valid {
			// Set user information in the request context
			c.Set("userID", claims.UserID)
			c.Next()
		} else {
			response.Unauthorized(c, "invalid token")
		}
	}
}

func (tc *token_controller) get_user_id(tokenString string) (string, error) {
	token, err := tc.check(tokenString)
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*customClaims); ok && token.Valid {
		return claims.UserID, nil
	} else {
		return "", err
	}
}

func (tc *token_controller) check(str string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(str, &customClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(tc.key), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (tc *token_controller) make_verification_token(id primitive.ObjectID) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := customClaims{
		UserID: id.Hex(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(tc.key))
	return tokenString, err
}
