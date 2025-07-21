package domain

type ParkingInquiryRequest struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Radius    int     `json:"radius"`
	Category  string  `json:"category"`
}
