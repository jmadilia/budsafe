package model

// Business entity that holds licenses
type Business struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Type        BusinessType `json:"type"`
	Description *string      `json:"description,omitempty"`
	Licenses    []*License   `json:"licenses,omitempty"`
	Locations   []*Location  `json:"locations,omitempty"`
	OwnerID     string       `json:"ownerId" db:"owner_id"`
	CreatedAt   string       `json:"createdAt" db:"created_at"`
	UpdatedAt   *string      `json:"updatedAt,omitempty" db:"updated_at"`
}