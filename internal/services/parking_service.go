package services

import (
	"context"
	"errors"
	"net/http"
	"sps-backend/internal/clients"
	"sps-backend/internal/config"
	"sps-backend/internal/domain"
	"time"
)

type ParkingServices struct {
	parkingClient *clients.SPSClient
	httpClient    *http.Client
}

func NewParkingService(parkingClient *clients.SPSClient, c *config.Config) *ParkingServices {
	// apiProxy, error := url.Parse(c.ProxyURL)
	// if error != nil {
	// 	return &ParkingServices{}
	// }
	return &ParkingServices{
		parkingClient: parkingClient,
		httpClient: &http.Client{
			Timeout: 60 * time.Second,
			// Transport: &http.Transport{
			// 	Proxy: http.ProxyURL(apiProxy),
			// },
		},
	}
}

func (s *ParkingServices) ParkingInq(ctx context.Context, req *domain.ParkingInquiryRequest) (*domain.ParkingInquiryResponse, error) {

	response, err := s.parkingClient.PostRequestSearchLocation(req, s.httpClient)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return response, nil
}
