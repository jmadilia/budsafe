package model

// Jurisdiction (state/country) with specific regulations
type Jurisdiction struct {
	ID                string           `json:"id"`
	Name              string           `json:"name"`
	Type              JurisdictionType `json:"type"`
	Country           string           `json:"country"`
	RegulatoryBody    string           `json:"regulatoryBody" db:"regulatory_body"`
	RegulatoryWebsite *string          `json:"regulatoryWebsite,omitempty" db:"regulatory_website"`
	LicenseTypes      []string         `json:"licenseTypes" db:"license_types"`
	Regulations       []*Regulation    `json:"regulations,omitempty"`
	CreatedAt         string           `json:"createdAt" db:"created_at"`
	UpdatedAt         *string          `json:"updatedAt,omitempty" db:"updated_at"`
}