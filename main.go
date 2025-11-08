package main

import (
	"goruntime_explore/middleware"
	"log"
	"net/http"

	"github.com/shopspring/decimal"
)

type contextKey string

const userIDKey contextKey = "userID"

type Stats struct {
	Goroutines int
	Memory     uint64
}

var balance = decimal.NewFromFloat(1000.00)

func main() {
	// Create main router
	mux := http.NewServeMux()

	// Public routes (no auth required)
	publicRouter := http.NewServeMux()
	publicRouter.HandleFunc("GET /health", healthHandler)
	publicRouter.HandleFunc("GET /monsters", listMonstersHandler)
	publicRouter.HandleFunc("GET /monsters/{id}", getMonsterHandler)

	// Admin routes (auth required)
	adminRouter := http.NewServeMux()
	adminRouter.HandleFunc("POST /monsters", createMonsterHandler)
	adminRouter.HandleFunc("PUT /monsters/{id}", updateMonsterHandler)
	adminRouter.HandleFunc("PATCH /monsters/{id}", patchMonsterHandler)
	adminRouter.HandleFunc("DELETE /monsters/{id}", deleteMonsterHandler)
	// Apply middleware to admin routes
	adminWithAuth := middleware.EnsureAuthenticated(adminRouter)

	// Mount routers with versioning
	mux.Handle("/v1/", http.StripPrefix("/v1", publicRouter))
	mux.Handle("/v1/admin/", http.StripPrefix("/v1/admin", adminWithAuth))

	// Create middleware stack
	stack := middleware.CreateStack(
		middleware.Logging,
	)

	// Wrap entire router with middleware
	server := &http.Server{
		Addr:    ":8080",
		Handler: stack(mux),
	}

	log.Println("Server starting on :8080")
	log.Fatal(server.ListenAndServe())

}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func listMonstersHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	newBalance := balance.Add(decimal.NewFromFloat(100))
	log.Println(newBalance) // 1100

	log.Println(balance) // 1000

	w.Write([]byte("List of monsters"))
}

func getMonsterHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Monster ID: " + id))
}

func createMonsterHandler(w http.ResponseWriter, r *http.Request) {
	// Get userID from context (set by middleware)
	userID, ok := r.Context().Value(userIDKey).(string)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Monster created by user: " + userID))
}

func updateMonsterHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	userID, _ := r.Context().Value(userIDKey).(string)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Monster " + id + " updated by user: " + userID))
}

func patchMonsterHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	userID, _ := r.Context().Value(userIDKey).(string)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Monster " + id + " patched by user: " + userID))
}

func deleteMonsterHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	userID, _ := r.Context().Value(userIDKey).(string)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Monster " + id + " deleted by user: " + userID))
}
