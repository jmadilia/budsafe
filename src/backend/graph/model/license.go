package model

// License issued to a business by a regulatory authority
type License struct {
	ID                  string                `json:"id"`
	BusinessID          string                `json:"businessId" db:"business_id"`
	Business            *Business             `json:"business"`
	LocationID          *string               `json:"locationId,omitempty" db:"location_id"`
	Location            *Location             `json:"location,omitempty"`
	LicenseNumber       string                `json:"licenseNumber" db:"license_number"`
	LicenseType         LicenseType           `json:"licenseType" db:"type"`
	JurisdictionID      string                `json:"jurisdictionId" db:"jurisdiction_id"`
	Jurisdiction        *Jurisdiction         `json:"jurisdiction"`
	IssuedDate          string                `json:"issuedDate" db:"issued_date"`
	ExpirationDate      string                `json:"expirationDate" db:"expiration_date"`
	Status              LicenseStatus         `json:"status"`
	RenewalDate					*string               `json:"renewalDate,omitempty" db:"renewal_date"`
	RenewalRequirements []*RenewalRequirement `json:"renewalRequirements,omitempty"`
	ComplianceChecks    []*ComplianceCheck    `json:"complianceChecks,omitempty"`
	Documents           []*Document           `json:"documents,omitempty"`
	FeeAmount						*float64              `json:"feeAmount,omitempty" db:"fee_amount"`
	Notes               *string               `json:"notes,omitempty"`
	CreatedAt           string                `json:"createdAt" db:"created_at"`
	UpdatedAt           *string               `json:"updatedAt,omitempty" db:"updated_at"`
}

type LicenseFilter struct {
	BusinessID     *string        `json:"businessId,omitempty"`
	JurisdictionID *string        `json:"jurisdictionId,omitempty"`
	LicenseType    *LicenseType   `json:"licenseType,omitempty"`
	Status         *LicenseStatus `json:"status,omitempty"`
	ExpiringBefore *string        `json:"expiringBefore,omitempty"`
}