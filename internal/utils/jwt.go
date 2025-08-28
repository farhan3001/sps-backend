package utils

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"sps-backend/internal/config"
	"sps-backend/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func MakeJWTSession(ipAddress string, timestamp string, jwtSecret string) (string, time.Time, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":        time.Now().Add(time.Minute * 10).Unix(),
		"ip_address": ipAddress,
		"timestamp":  timestamp,
	})

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", time.Now().Add(time.Minute * 10), err
	}

	return tokenString, time.Now().Add(time.Minute * 10), nil
}

func AuthenticateSession(cfg *config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			Error(ctx, http.StatusUnauthorized, ErrUnauthorized)
			ctx.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			Error(ctx, http.StatusUnauthorized, ErrUnauthorized)
			ctx.Abort()
			return
		}

		_, err := ValidateSession(tokenString, cfg.JWTSecret)
		if err != nil {
			Error(ctx, http.StatusUnauthorized, gin.H{"message": err.Error()})
			ctx.Abort()
			return
		}

		// Set user info in context
		ctx.Next()
	}
}

func ValidateSession(tokenString string, jwtSecret string) (*domain.UserSession, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token signature")
		} else {
			return []byte(jwtSecret), nil
		}
	})
	if err != nil {
		return nil, fmt.Errorf("token validation failed: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Validate expiration
	exp, ok := claims["exp"].(float64)
	if !ok {
		return nil, fmt.Errorf("invalid expiration time")
	}
	if time.Now().After(time.Unix(int64(exp), 0)) {
		return nil, fmt.Errorf("token expired")
	}

	// Populate UserSession with all available claims
	return &domain.UserSession{
		Timestamp: getStringClaim(claims, "timestamp"),
		IPAddress: getStringClaim(claims, "ip_address"),
		ExpiresAt: time.Unix(int64(exp), 0),
		Token:     tokenString,
	}, nil
}

// Helper functions for safe claim extraction
func getStringClaim(claims jwt.MapClaims, key string) string {
	if val, ok := claims[key].(string); ok {
		return val
	}
	return ""
}
