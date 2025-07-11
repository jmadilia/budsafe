package graph

import (
	"budsafe/backend/graph/model"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

// Helper function to scan a user row from database
func scanUser(rows *sql.Rows) (*model.User, error) {
	var user model.User
	var createdAt, updatedAt time.Time
	var firstName, lastName sql.NullString
	
	err := rows.Scan(
		&user.ID,
		&user.Email,
		&firstName,
		&lastName,
		&user.Role,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return nil, err
	}
	
	if firstName.Valid {
		user.FirstName = &firstName.String
	}
	if lastName.Valid {
		user.LastName = &lastName.String
	}
	
	user.CreatedAt = createdAt.Format(time.RFC3339)
	if !updatedAt.IsZero() {
		updatedAtStr := updatedAt.Format(time.RFC3339)
		user.UpdatedAt = &updatedAtStr
	}
	
	return &user, nil
}

// Helper function to scan a business row from database
func scanBusiness(rows *sql.Rows) (*model.Business, error) {
	var business model.Business
	var createdAt, updatedAt time.Time
	var description sql.NullString
	
	err := rows.Scan(
		&business.ID,
		&business.Name,
		&business.Type,
		&description,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return nil, err
	}
	
	if description.Valid {
		business.Description = &description.String
	}
	
	business.CreatedAt = createdAt.Format(time.RFC3339)
	if !updatedAt.IsZero() {
		updatedAtStr := updatedAt.Format(time.RFC3339)
		business.UpdatedAt = &updatedAtStr
	}
	
	return &business, nil
}

// Helper function to scan a license row from database
func scanLicense(rows *sql.Rows) (*model.License, error) {
	var license model.License
	var locationID sql.NullString
	var notes sql.NullString
	var renewalDate sql.NullString
	var feeAmount sql.NullFloat64
	var issuedDate, expirationDate, createdAt, updatedAt sql.NullString

	err := rows.Scan(
		&license.ID,
		&license.BusinessID,
		&license.JurisdictionID,
		&locationID,
		&license.LicenseNumber,
		&license.LicenseType,
		&license.Status,
		&issuedDate,
		&expirationDate,
		&renewalDate,
		&feeAmount,
		&notes,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return nil, err
	}

	if locationID.Valid {
		license.LocationID = &locationID.String
	}
	if notes.Valid {
		license.Notes = &notes.String
	}
	if renewalDate.Valid {
		license.RenewalDate = &renewalDate.String
	}
	if feeAmount.Valid {
		license.FeeAmount = &feeAmount.Float64
	}
	if issuedDate.Valid {
		license.IssuedDate = issuedDate.String
	}
	if expirationDate.Valid {
		license.ExpirationDate = expirationDate.String
	}
	if createdAt.Valid {
		license.CreatedAt = createdAt.String
	}
	if updatedAt.Valid {
		license.UpdatedAt = &updatedAt.String
	}

	return &license, nil
}

// Helper function to scan a compliance check row from database
func scanComplianceCheck(rows *sql.Rows) (*model.ComplianceCheck, error) {
	var check model.ComplianceCheck
	var createdAt, updatedAt, checkedAt, dueDate time.Time
	var notes sql.NullString
	var checkedByID sql.NullString
	
	err := rows.Scan(
		&check.ID,
		&check.LicenseID,
		&check.Title,
		&check.Status,
		&checkedAt,
		&dueDate,
		&notes,
		&checkedByID,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return nil, err
	}
	
	if notes.Valid {
		check.Notes = &notes.String
	}
	
	check.DueDate = dueDate.Format(time.RFC3339)
	check.CreatedAt = createdAt.Format(time.RFC3339)
	if !updatedAt.IsZero() {
		updatedAtStr := updatedAt.Format(time.RFC3339)
		check.UpdatedAt = &updatedAtStr
	}
	
	return &check, nil
}

// Helper function to parse Postgres array string (e.g., {A,B,C}) to []string
func parsePostgresArray(arrayStr string) []string {
	trimmed := strings.Trim(arrayStr, "{}")
	if trimmed == "" {
		return []string{}
	}
	parts := strings.Split(trimmed, ",")
	for i, p := range parts {
		parts[i] = strings.TrimSpace(p)
	}
	return parts
}

// scanner interface for Scan compatibility
// Both *sql.Row and *sql.Rows implement this

type scanner interface {
	Scan(dest ...any) error
}

// Helper function to scan a jurisdiction row from database (works for both *sql.Row and *sql.Rows)
func scanJurisdiction(s scanner) (*model.Jurisdiction, error) {
	var jurisdiction model.Jurisdiction
	var licenseTypesStr sql.NullString
	var createdAt, updatedAt sql.NullString

	err := s.Scan(
		&jurisdiction.ID,
		&jurisdiction.Name,
		&jurisdiction.Type,
		&jurisdiction.Country,
		&jurisdiction.RegulatoryBody,
		&jurisdiction.RegulatoryWebsite,
		&licenseTypesStr,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get jurisdiction: %v", err)
	}

	if licenseTypesStr.Valid {
		jurisdiction.LicenseTypes = parsePostgresArray(licenseTypesStr.String)
	}
	if createdAt.Valid {
		jurisdiction.CreatedAt = createdAt.String
	}
	if updatedAt.Valid {
		jurisdiction.UpdatedAt = &updatedAt.String
	}

	return &jurisdiction, nil
}


// buildUpdateQuery dynamically constructs an SQL UPDATE statement.
// It takes a map of column names to their new values.
// It only includes non-nil values in the SET clause.
func buildUpdateQuery(table string, id string, updates map[string]interface{}) (string, []interface{}) {
	var setClauses []string
	args := []interface{}{}
	argIndex := 1

	for col, val := range updates {
		// Check for pointer types and dereference to check for nil
		switch v := val.(type) {
		case *string:
			if v != nil {
				setClauses = append(setClauses, fmt.Sprintf("%s = $%d", col, argIndex))
				args = append(args, *v)
				argIndex++
			}
		case *bool:
			if v != nil {
				setClauses = append(setClauses, fmt.Sprintf("%s = $%d", col, argIndex))
				args = append(args, *v)
				argIndex++
			}
		case *time.Time:
			if v != nil {
				setClauses = append(setClauses, fmt.Sprintf("%s = $%d", col, argIndex))
				args = append(args, *v)
				argIndex++
			}
		case *float64:
			if v != nil {
				setClauses = append(setClauses, fmt.Sprintf("%s = $%d", col, argIndex))
				args = append(args, *v)
				argIndex++
			}
		// Add other types as needed, e.g., model enums
		case *model.LicenseStatus:
			if v != nil {
				setClauses = append(setClauses, fmt.Sprintf("%s = $%d", col, argIndex))
				args = append(args, *v)
				argIndex++
			}
		case *model.ComplianceStatus:
			if v != nil {
				setClauses = append(setClauses, fmt.Sprintf("%s = $%d", col, argIndex))
				args = append(args, *v)
				argIndex++
			}
		}
	}

	if len(setClauses) == 0 {
		return "", nil // No update to perform
	}

	// Always update the updated_at timestamp
	setClauses = append(setClauses, fmt.Sprintf("updated_at = $%d", argIndex))
	args = append(args, time.Now())
	argIndex++

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d RETURNING *", table, strings.Join(setClauses, ", "), argIndex)
	args = append(args, id)

	return query, args
}