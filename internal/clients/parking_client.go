package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sps-backend/internal/domain"
	"sps-backend/internal/utils"
	"time"
)

type SPSClient struct {
	baseURL      string
	clientKey    string
	clientSecret string
}

func NewSPSClient(baseURL, clientKey, clientSecret string) *SPSClient {
	return &SPSClient{
		baseURL:      baseURL,
		clientKey:    clientKey,
		clientSecret: clientSecret,
	}
}

func (c *SPSClient) PostRequestSearchLocation(request *domain.ParkingInquiryRequest, httpClient *http.Client) (*domain.ParkingInquiryResponse, error) {
	// Prepare the request body
	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request body: %v", err)
	}

	_, err = json.Marshal(request)
	if err != nil {
		return nil, err
	}

	// fmt.Println(string(jsonBody))

	// Generate signature
	path := "/v1/partner/get-geolocation"
	signature, timestamp, err := utils.GenerateSignatureForAPIReq(
		c.clientSecret,
		c.clientKey,
		time.Now())

	if err != nil {
		return nil, fmt.Errorf("error generating signature: %v", err)
	}

	// Create the HTTP request
	req, err := http.NewRequest(
		"POST",
		c.baseURL+path,
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("timestamp", timestamp)
	req.Header.Set("signature", signature)
	req.Header.Set("clientKey", c.clientKey)

	// Send the request
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	var response *domain.ParkingInquiryResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("error parsing response: %v", err)
	}

	return response, nil
}
