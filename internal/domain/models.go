package domain

import (
	"time"

	"github.com/google/uuid"
)

// Institution represents an organization
type Institution struct {
	InstID        uuid.UUID `json:"instId"`
	InstCode      string    `json:"instCode"`
	InstName      string    `json:"instName"`
	InstAddress   string    `json:"instAddress,omitempty"`
	InstType      string    `json:"instType,omitempty"`
	RegisteredAt  string    `json:"registeredAt"`  // ISO8601 formatted string
	LastUpdatedAt string    `json:"lastUpdatedAt"` // ISO8601 formatted string
}

// InstitutionCreds represents credentials for an institution
type InstitutionCreds struct {
	ID            int       `json:"id"`
	InstID        uuid.UUID `json:"instId"`
	ClientKey     string    `json:"clientKey"`
	ClientSecret  string    `json:"clientSecret"`
	RegisteredAt  string    `json:"registeredAt"`
	LastUpdatedAt string    `json:"lastUpdatedAt"`
}

// Member represents a member of an institution
type Member struct {
	MemberID              uuid.UUID `json:"memberId"`
	NIK                   string    `json:"nik"`
	MemberName            string    `json:"memberName"`
	MemberOf              uuid.UUID `json:"memberOf"`
	BeneficiaryIdentifier string    `json:"beneficiaryIdentifier,omitempty"`
	RegisteredAt          string    `json:"registeredAt"`
	LastUpdatedAt         string    `json:"lastUpdatedAt"`
}

// Vehicle represents a vehicle
type Vehicle struct {
	VehicleID     uuid.UUID `json:"vehicleId"`
	Type          string    `json:"type"`
	Brand         string    `json:"brand"`
	RegisteredAt  string    `json:"registeredAt"`
	LastUpdatedAt string    `json:"lastUpdatedAt"`
}

// VehicleOwn represents ownership between members and vehicles
type VehicleOwn struct {
	MemberID      uuid.UUID `json:"memberId"`
	VehicleID     uuid.UUID `json:"vehicleId"`
	OwnerType     string    `json:"ownerType"`
	LastUpdatedAt string    `json:"lastUpdatedAt"`
}

// VehicleRegistration represents vehicle documents
type VehicleRegistration struct {
	ID            uuid.UUID `json:"id"`
	VehicleID     uuid.UUID `json:"vehicleId"`
	FrontViewURL  string    `json:"frontViewUrl,omitempty"`
	RearViewURL   string    `json:"rearViewUrl,omitempty"`
	SideViewURL   string    `json:"sideViewUrl,omitempty"`
	RegisteredAt  string    `json:"registeredAt"`
	LastUpdatedAt string    `json:"lastUpdatedAt"`
}

type UserSession struct {
	Email     string    `json:"email"`
	UserID    string    `json:"user_id"`
	ExpiresAt time.Time `json:"expires_at"`
	Token     string    `json:"token"`
	IPAddress string    `json:"ip_address"`
	Longitude float64   `json:"longitude"`
	Latitude  float64   `json:"latitude"`
	Location  string    `json:"location"`
}
