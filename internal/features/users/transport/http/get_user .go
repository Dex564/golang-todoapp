package users_transport_http

import (
	"net/http"

	core_logger "github.com/Dex564/golang-todoapp/internal/core/logger"
	core_http_request "github.com/Dex564/golang-todoapp/internal/core/transport/http/request"
	core_http_response "github.com/Dex564/golang-todoapp/internal/core/transport/http/response"
)

type GetUserResponce UserDTOResponse

func (h *UsersHTTPHandler) GetUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responceHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responceHandler.ErrorResponse(
			err,
			"failed to get userID path value",
		)
		return
	}

	user, err := h.usersService.GetUser(ctx, userID)
	if err != nil {
		responceHandler.ErrorResponse(
			err,
			"failed to get user",
		)
		return
	}

	response := GetUserResponce(userDTOFromDomain(user))

	responceHandler.JsonResponse(response, http.StatusOK)
}
