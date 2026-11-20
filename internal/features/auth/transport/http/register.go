package auth_transport_http

import (
	"net/http"

	core_logger "github.com/Dex564/golang-todoapp/internal/core/logger"
	core_http_request "github.com/Dex564/golang-todoapp/internal/core/transport/http/request"
	core_http_response "github.com/Dex564/golang-todoapp/internal/core/transport/http/response"
)

type RegisterUserRequest struct {
	Username string `json:"username"`
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterUserResponse UserDTOResponse

func (h *AuthHTTPHandler) Register(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	var request RegisterUserRequest

	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")
		return
	}

	userDomain, err := h.authService.Register(ctx, request.Username, request.Login, request.Email, request.Password)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to register user")
		return
	}

	response := RegisterUserResponse(userDTOFromDomain(userDomain))

	responseHandler.JsonResponse(response, http.StatusCreated)
}
