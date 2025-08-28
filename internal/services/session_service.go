package services

import (
	"context"
	"errors"
	"net/http"
	"sps-backend/internal/config"
	"sps-backend/internal/domain"
	"sps-backend/internal/utils"
	"time"
)

type SessionServices struct {
	httpClient *http.Client
	jwtSecret  string
}

func NewSessionService(c *config.Config, jwtSecret string) *SessionServices {
	return &SessionServices{
		httpClient: &http.Client{
			Timeout: 60 * time.Second,
		},
		jwtSecret: jwtSecret,
	}
}

func (s *SessionServices) GetTokenService(ctx context.Context, ipAddress string, timeStamp string) (*domain.GetTokenResponse, error) {

	if len(ipAddress) == 0 {
		return nil, errors.New("X-IP-Address tidak boleh kosong")
	}

	if len(timeStamp) == 0 {
		return nil, errors.New("X-TIMESTAMP tidak boleh kosong")
	}

	jwtToken, expTime, err := utils.MakeJWTSession(ipAddress, timeStamp, s.jwtSecret)
	if err != nil {
		return nil, err
	}

	jwtResponse := &domain.GetTokenResponse{
		Token:     jwtToken,
		ExpiresAt: expTime,
	}

	return jwtResponse, nil
}
