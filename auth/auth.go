package auth

import (
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// generate an jwt token

var secretKey = []byte("your-secret-key")

func GenerateJWT(username string) (string, error) {

	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	}

	// generate  atoken
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	// sign it with a secret key
	return token.SignedString(secretKey)

}

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "authorization header missing", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			// check the signing of the jwt
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}

			return secretKey, nil

		})

		if err != nil || !token.Valid {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		// access claims

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			username := claims["username"].(string)
			r.Header.Set("X-User", username)
		}

		next.ServeHTTP(w, r)
	})
}
