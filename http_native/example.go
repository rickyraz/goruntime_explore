package httpnative

import (
	"goruntime_explore/middleware"
	"log"
	"net/http"
)

func RunHTTPNative() {
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
