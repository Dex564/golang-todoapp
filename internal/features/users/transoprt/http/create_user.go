package users_transport_http

import (
	"encoding/json"
	"fmt"
	"net/http"

	core_logger "github.com/Dex564/golang-todoapp/internal/core/logger"
)

type CreateUserRequest struct {
	Username    string  `json:"username"`
	PhoneNumber *string `json:"phone_number"`
}

type CreateUserResponse struct {
	Id          int     `json:"id"`
	Version     int     `json:"version"`
	Username    string  `json:"username"`
	PhoneNumber *string `json:"phone_number"`
}

func (h *UsersHTTPHandler) CreateUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	log.Debug("invoke CreateUser handler")
	var request CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		fmt.Println("ошибочка")
	}

	rw.WriteHeader(http.StatusOK)
}
