package model

type ComplianceStatusSummary struct {
	BusinessID        string           `json:"businessId" db:"business_id"`
	CompliantCount    int              `json:"compliantCount" db:"compliant_count"`
	NonCompliantCount int              `json:"nonCompliantCount" db:"non_compliant_count"`
	PendingCount      int              `json:"pendingCount" db:"pending_count"`
	AttentionCount    int              `json:"attentionCount" db:"attention_count"`
	OverallStatus     ComplianceStatus `json:"overallStatus"`
}