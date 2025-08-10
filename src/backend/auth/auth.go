package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
)

// A private key for the context to avoid collisions
type contextKey string

const userContextKey = contextKey("user")

// AuthClient holds the initialized Firebase Auth client
type AuthClient struct {
	Client *auth.Client
}

// User holds the essential information from the verified token
type User struct {
	UID    string
	Email	 string
}

func Init(ctx context.Context) (*AuthClient, error) {
	// IMPORTANT: Use environment variables for your service account credentials
	// In production (like Cloud Run), this can be automatically inferred.
	// For local development, set GOOGLE_APPLICATION_CREDENTIALS env var.
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("error getting Firebase app: %w", err)
	}

	client, err := app.Auth(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting Firebase Auth client: %w", err)
	}

	return &AuthClient{Client: client}, nil
}

// Middleware is the HTTP middleware for authenticating requests
func (ac *AuthClient) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. Get the token from the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			// Allow unauthenticated requests to proceed.
			// Your resolvers will then decide if they need an authenticated user.
			next.ServeHTTP(w, r)
			return
		}

		// 2. Validate the token format ("Bearer <token>")
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
			http.Error(w, "Authorization header format must be Bearer {token}", http.StatusUnauthorized)
			return
		}
		idToken := tokenParts[1]

		// 3. Verify the token using our reusable function
		user, err := ac.VerifyToken(r.Context(), idToken)
		if err != nil {
			log.Printf("Error verifying token: %v", err)
			http.Error(w, "Invalid authentication token", http.StatusUnauthorized)
			return
		}

		// 4. Add the user info to the request context
		ctxWithUser := context.WithValue(r.Context(), userContextKey, user)
		rWithUser := r.WithContext(ctxWithUser)

		// 5. Call the next handler in the chain
		next.ServeHTTP(w, rWithUser)
	})
}

// VerifyToken is the reusable function that verifies a Firebase ID token.
// It returns a User struct with the UID and email on success.
func (ac *AuthClient) VerifyToken(ctx context.Context, idToken string) (*User, error) {
	token, err := ac.Client.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, fmt.Errorf("could not verify token: %w", err)
	}

	return &User{
		UID:   token.UID,
		Email: token.Claims["email"].(string),
	}, nil
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *User {
	raw, _ := ctx.Value(userContextKey).(*User)
	return raw
}

// NewContext returns a new context that carries the provided user.
func NewContext(ctx context.Context, user *User) context.Context {
	return context.WithValue(ctx, userContextKey, user)
}