package httpnative

import (
	"log"
	"net/http"

	"github.com/shopspring/decimal"
)

type contextKey string

var balance = decimal.NewFromFloat(1000.00)

const userIDKey contextKey = "userID"

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
