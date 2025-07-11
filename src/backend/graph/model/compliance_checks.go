package model

// Compliance check for a ComplianceCheck
type ComplianceCheck struct {
	ID													string  					`json:"id"`
	LicenseID 									string   					`json:"licenseId" db:"license_id"`
	ComplianceCheckLicense   		*License 					`json:"license"`
	Title 											string 						`json:"title" db:"check_type"`
	DueDate 										string 						`json:"dueDate" db:"next_check_date"`
	CheckedAt 									*string          	`json:"checkedAt,omitempty" db:"checked_at"`
	Status    									ComplianceStatus 	`json:"status"`
	UserID											*string 					`json:"userId" db:"checked_by_id"`
	ComplianceCheckUser 				*User 						`json:"assignedTo,omitempty"`
	Notes     									*string 					`json:"notes,omitempty"`
	CreatedAt 									string  					`json:"createdAt" db:"created_at"`
	UpdatedAt 									*string 					`json:"updatedAt,omitempty" db:"updated_at"`
}