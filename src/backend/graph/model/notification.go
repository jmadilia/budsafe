package model

// Notification for upcoming deadlines or compliance issues
type Notification struct {
	ID                string           `json:"id"`
	UserID            string           `json:"userId" db:"user_id"`
	User              *User            `json:"user"`
	Title             string           `json:"title"`
	Message           string           `json:"message"`
	Type              NotificationType `json:"type"`
	IsRead            bool             `json:"isRead" db:"is_read"`
	RelatedEntityID   *string          `json:"relatedEntityId,omitempty" db:"related_entity_id"`
	RelatedEntityType *string          `json:"relatedEntityType,omitempty" db:"related_entity_type"`
	CreatedAt         string           `json:"createdAt" db:"created_at"`
	UpdatedAt         string           `json:"updatedAt" db:"updated_at"`
}