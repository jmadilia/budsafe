package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"budsafe/backend/auth"
	"budsafe/backend/graph"
	"budsafe/backend/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Load environment variables from root .env.local
	// IMPORTANT: Make sure your .env.local contains the GOOGLE_APPLICATION_CREDENTIALS variable
	// pointing to your service account JSON file for local development.
	// e.g., GOOGLE_APPLICATION_CREDENTIALS=../path/to/your/serviceAccountKey.json
	rootDir := filepath.Join("../..", ".env.local")
	if err := godotenv.Load(rootDir); err != nil {
		log.Printf("Warning: Could not load .env.local file: %v", err)
	}

	// Get database connection info
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}
	
	// Connect to database
	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Test database connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	log.Println("Successfully connected to PostgreSQL database!")

	// Initialize Firebase Auth client
	authClient, err := auth.Init(context.Background())
	if err != nil {
		log.Fatalf("Could not initialize Firebasee Auth client: %v", err)
	}
	log.Println("Successfully initialized Firebase Auth client!")

	// Create GraphQL server with database connection
	resolver := &graph.Resolver{DB: db}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))


	// --- CORS Middleware ---
	corsMiddleware := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*") // Change to your frontend URL in production
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
			h.ServeHTTP(w, r)
		})
	}

	// GraphQL playground
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", corsMiddleware(authClient.Middleware(srv)))

	// Health check endpoints
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if err := db.Ping(); err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte("Database connection failed"))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Backend and database are healthy"))
	})

	// Get the port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the server
	log.Printf("GraphQL server starting at http://localhost:%s", port)
	log.Printf("GraphQL playground at http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
