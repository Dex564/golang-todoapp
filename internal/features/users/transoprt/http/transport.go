package users_transport_http

import (
	"net/http"

	core_http_server "github.com/Dex564/golang-todoapp/internal/core/transport/http/server"
)

type UsersHTTPHandler struct {
	usersService UsersService
}

type UsersService interface {
}

func NewUsersHTTPHandler(userService UsersService) *UsersHTTPHandler {
	return &UsersHTTPHandler{
		usersService: userService,
	}
}

func (h *UsersHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: h.CreateUser,
		},
	}
}
