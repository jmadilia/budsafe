package graph_test

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"testing"

	"budsafe/backend/auth"
	"budsafe/backend/graph"
	"budsafe/backend/graph/model"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupTestDB remains the same
func setupTestDB(t *testing.T) *sqlx.DB {
	rootDir := filepath.Join("../../..", ".env.local")
	if err := godotenv.Load(rootDir); err != nil {
		log.Printf("Warning: Could not load .env.local file: %v", err)
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is required for testing")
	}
	db, err := sqlx.Connect("postgres", dbURL)
	require.NoError(t, err, "Failed to connect to database for testing")
	return db
}

func TestMutationResolver_CreateUser(t *testing.T) {
	// --- 1. SETUP ---
	db := setupTestDB(t)
	// The resolver now directly uses the database connection.
	resolver := &graph.Resolver{DB: db}
	mutationResolver := resolver.Mutation()

	// --- 2. SIMULATE AUTHENTICATION ---
	fakeAuthUser := &auth.User{
		UID:   "test-firebase-uid-cleanup-123",
		Email: "cleanup.test@example.com",
	}
	ctx := auth.NewContext(context.Background(), fakeAuthUser)

	// --- 3. DEFINE TEST INPUT ---
	input := model.CreateUserInput{
		Email: 	 fakeAuthUser.Email,
		FirstName: "Jane",
		LastName:  "Doe",
		Role:      model.UserRoleAdmin,
		FirebaseUID: fakeAuthUser.UID,
	}

	// --- 4. EXECUTE THE FUNCTION ---
	createdUser, err := mutationResolver.CreateUser(ctx, input)

	// --- 5. SCHEDULE CLEANUP ---
	// We use `defer` to ensure this cleanup code runs at the end of the test,
	// even if one of the assertions below fails.
	if createdUser != nil {
		defer func() {
			_, delErr := db.Exec("DELETE FROM users WHERE id = $1", createdUser.ID)
			require.NoError(t, delErr, "Cleanup failed: Could not delete test user.")
			log.Printf("Cleaned up user with ID: %s", createdUser.ID)
		}()
	}

	// --- 6. ASSERT THE RESULTS ---
	require.NoError(t, err)
	require.NotNil(t, createdUser)

	assert.Equal(t, input.FirstName, *createdUser.FirstName)
	assert.Equal(t, input.LastName, *createdUser.LastName)
	assert.Equal(t, fakeAuthUser.Email, createdUser.Email)
	assert.Equal(t, fakeAuthUser.UID, *createdUser.FirebaseUID)

	// --- 7. VERIFY DATABASE STATE (Optional but good) ---
	// We can still check that the user exists in the DB before the cleanup runs.
	var dbUser model.User
	err = db.Get(&dbUser, "SELECT * FROM users WHERE id = $1", createdUser.ID)
	require.NoError(t, err, "User should be findable in the database before cleanup")
	
	// Check that the pointer is not nil and then dereference it
	require.NotNil(t, dbUser.FirstName, "FirstName should not be nil")
	assert.Equal(t, "Jane", *dbUser.FirstName)
}