package auth_transport_http

import (
	"net/http"

	core_logger "github.com/Dex564/golang-todoapp/internal/core/logger"
	core_http_response "github.com/Dex564/golang-todoapp/internal/core/transport/http/response"
)

func (h *AuthHTTPHandler) Login(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

}
