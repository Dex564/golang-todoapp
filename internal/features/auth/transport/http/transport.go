package auth_transport_http

import (
	"context"

	"github.com/Dex564/golang-todoapp/internal/core/domain"
	core_http_server "github.com/Dex564/golang-todoapp/internal/core/transport/http/server"
)

type AuthHTTPHandler struct {
	authService AuthService
}

type AuthService interface {
	Register(
		ctx context.Context,
		username string,
		login string,
		email string,
		password string,
	) (domain.User, error)
}

func NewAuthHTTPHandler(authService AuthService) *AuthHTTPHandler {
	return &AuthHTTPHandler{authService: authService}
}

func (h *AuthHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{}
}
