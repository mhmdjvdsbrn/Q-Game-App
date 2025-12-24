package main

import (
	"fmt"
	"log"
	"net/http"
	"q-game-app/config"
	"q-game-app/delivery/httpserver"
	"time"

	"q-game-app/repository/mysql"
	"q-game-app/service/authservice"
	"q-game-app/service/userservice"
)

const (
	jwtSecret = "my_super_secret_key_123"

	AccessTokenExpireDuration  = time.Hour * 24
	RefreshTokenExpireDuration = time.Hour * 24 * 7

	AccessTokenSubject  = "access"
	RefreshTokenSubject = "refresh"

	DBUserName = "myuser"
	DBUserPass = "mypassword"
	DBHost     = "localhost"
	DBName     = "mydb"
	DBPort     = 3308
)

func main() {
	cfg := config.Config{
		HttpServer: config.HttpServer{Port: 8080},
		Auth: authservice.Config{
			SignKey:               jwtSecret,
			AccessExpirationTime:  AccessTokenExpireDuration,
			RefreshExpirationTime: RefreshTokenExpireDuration,
			AccessSubject:         AccessTokenSubject,
			RefreshSubject:        RefreshTokenSubject,
		},
		Mysql: mysql.Config{
			Username: DBUserName,
			Password: DBUserPass,
			Port:     DBPort,
			DBName:   DBName,
			Host:     DBHost,
		},
	}

	authSvc, userSvc := setupServices(cfg)

	server := httpserver.New(cfg, authSvc, userSvc)
	server.Serve()

	log.Println("Starting server on :8080...")
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "server is available")
}

//func userLoginHandler(w http.ResponseWriter, r *http.Request) {
//	if r.Method != http.MethodPost {
//		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
//		return
//	}
//
//	data, err := io.ReadAll(r.Body)
//	if err != nil {
//		http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusBadRequest)
//		return
//	}
//
//	var lReq userservice.LoginRequest
//	if err := json.Unmarshal(data, &lReq); err != nil {
//		http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusBadRequest)
//		return
//	}
//
//	mysqlRepo := mysql.New()
//	authSvc := authservice.NewService(jwtSecret, "access", "refresh", time.Hour*1, time.Hour*24)
//	userSvc := userservice.New(mysqlRepo, authSvc)
//
//	tokens, sErr := userSvc.Login(lReq)
//	if sErr != nil {
//		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, sErr.Error()), http.StatusInternalServerError)
//		return
//	}
//
//	result, err := json.Marshal(tokens)
//	if err != nil {
//		http.Error(w, `{"error":"failed to marshal response"}`, http.StatusInternalServerError)
//		return
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusOK)
//	w.Write(result)
//}
//
//func userProfileHandler(w http.ResponseWriter, r *http.Request) {
//	if r.Method != http.MethodGet {
//		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
//		return
//	}
//
//	userIDStr := r.URL.Query().Get("user_id")
//	if userIDStr == "" {
//		http.Error(w, `{"error":"missing user_id"}`, http.StatusBadRequest)
//		return
//	}
//
//	userID, err := strconv.Atoi(userIDStr)
//	if err != nil {
//		http.Error(w, `{"error":"invalid user_id"}`, http.StatusBadRequest)
//		return
//	}
//
//	pReq := userservice.ProfileRequest{UserID: uint(userID)}
//
//	mysqlRepo := mysql.New()
//	authSvc := authservice.NewService(jwtSecret, "access", "refresh", time.Hour*1, time.Hour*24)
//	userSvc := userservice.New(mysqlRepo, authSvc)
//
//	resp, svcErr := userSvc.Profile(pReq)
//	if svcErr != nil {
//		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, svcErr.Error()), http.StatusInternalServerError)
//		return
//	}
//
//	jsonData, _ := json.Marshal(resp)
//	w.Header().Set("Content-Type", "application/json")
//	w.Write(jsonData)
//}

func setupServices(cfg config.Config) (*authservice.Service, *userservice.Service) {
	// auth service pointer
	authSvc := authservice.New(cfg.Auth)

	// mysql repo pointer
	mysqlRepo := mysql.New(cfg.Mysql)

	userSvc := userservice.New(mysqlRepo, authSvc)

	return authSvc, userSvc
}
