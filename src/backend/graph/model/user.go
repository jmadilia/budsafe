package model

// User account with authentication and permissions
type User struct {
	ID         	string      `json:"id"`
	Email      	string      `json:"email"`
	FirstName  	*string     `json:"firstName,omitempty" db:"first_name"`
	LastName   	*string     `json:"lastName,omitempty" db:"last_name"`
	Role       	UserRole    `json:"role"`
	FirebaseUID *string     `json:"firebaseUid" db:"firebase_uid"`
	Businesses 	[]*Business `json:"businesses,omitempty"`
	CreatedAt  	string      `json:"createdAt" db:"created_at"`
	UpdatedAt  	*string     `json:"updatedAt,omitempty" db:"updated_at"`
}