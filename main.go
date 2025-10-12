package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"q-game-app/repository/mysql"
	"q-game-app/service/userservice"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", healthCheckHandler)
	mux.HandleFunc("/users/register-user", registerUserHandler)

	log.Println("Starting server on :8080...")
	server := http.Server{Addr: ":8080", Handler: mux}
	log.Fatal(server.ListenAndServe())
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "server is available")
}

func registerUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	data, dErr := io.ReadAll(r.Body)
	if dErr != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%s"}`, dErr.Error()), http.StatusBadRequest)
		return
	}

	var uReq userservice.RegisterRequest
	err := json.Unmarshal(data, &uReq)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusBadRequest)
		return
	}

	mysqlRepo := mysql.New()
	userSvc := userservice.New(mysqlRepo)
	_, sErr := userSvc.Register(uReq)
	if sErr != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%s"}`, sErr.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"success": "success"}`))
}
