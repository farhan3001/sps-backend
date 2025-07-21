package domain

type ParkingInquiryResponse struct {
	ResponseStatus  string         `json:"responseStatus"`
	ResponseCode    string         `json:"responseCode"`
	ResponseMessage string         `json:"responseMessage"`
	Data            []LocationData `json:"data"`
}

type LocationData struct {
	LocationCode string       `json:"location_code"`
	LocationName string       `json:"location_name"`
	Address      string       `json:"address"`
	Coordinate   Coordinate   `json:"coordinate"`
	Category     string       `json:"category"`
	ParkingLot   []ParkingLot `json:"parking_lot"`
}

type Coordinate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type ParkingLot struct {
	TotalTraffic   int `json:"TOTAL_TRAFFIC"`
	CarUsedLot     int `json:"CAR_USED_LOT"`
	MotorUsedLot   int `json:"MOTOR_USED_LOT"`
	CarAvailable   int `json:"CAR_AVAILABLE"`
	MotorAvailable int `json:"MOTOR_AVAILABLE"`
}
