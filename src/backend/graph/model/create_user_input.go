package model

type CreateUserInput struct {
	FirebaseUID string   `json:"firebaseUid" db:"firebase_uid"`
	Email       string   `json:"email"`
	FirstName   string   `json:"firstName" db:"first_name"`
	LastName    string   `json:"lastName" db:"last_name"`
	Role        UserRole `json:"role"`
}