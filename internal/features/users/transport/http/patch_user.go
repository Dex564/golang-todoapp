package users_transport_http

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Dex564/golang-todoapp/internal/core/domain"
	core_logger "github.com/Dex564/golang-todoapp/internal/core/logger"
	core_http_request "github.com/Dex564/golang-todoapp/internal/core/transport/http/request"
	core_http_response "github.com/Dex564/golang-todoapp/internal/core/transport/http/response"
	core_http_types "github.com/Dex564/golang-todoapp/internal/core/transport/http/types"
)

type PatchUserRequest struct {
	Username    core_http_types.Nullable[string] `json:"username" swaggertype:"string" example:"Максим Максимович"`
	PhoneNumber core_http_types.Nullable[string] `json:"phone_number" swaggertype:"string" example:"+71112223344"`
}

func (r *PatchUserRequest) Validate() error {
	if r.Username.Set {
		if r.Username.Value == nil {
			return fmt.Errorf("`Username` can't be NUll")
		}

		usernameLen := len([]rune(*r.Username.Value))
		if usernameLen < 3 || usernameLen > 100 {
			return fmt.Errorf("`Username` must be betweem 3 and 100 symbols")
		}
	}

	if r.PhoneNumber.Set {
		if r.PhoneNumber.Value != nil {
			phoneNumberLen := len([]rune(*r.PhoneNumber.Value))
			if phoneNumberLen < 10 || phoneNumberLen > 15 {
				return fmt.Errorf("`PhoneNumber` must be betweem 10 and 15 symbols")
			}

			if !strings.HasPrefix(*r.PhoneNumber.Value, "+") {
				return fmt.Errorf("`PhoneNumber` must startswith `+` symbol")
			}
		}
	}

	return nil
}

type PatchUserResponse UserDTOResponse

// PatchUser    godoc
// @Summary     Изменение пользователя
// @Description Изменение информации об уже существующем пользователе
// @Description ### Логика обновления полей (Three-state logic):
// @Description 1. **Поле не передано**: поле игнорируется, значение в БД не меняется
// @Description 2. **Явно передано значение**: `"phone_number": "+71112223344"` - устанавливается новое поле
// @Description 3. **Передан null**: `"phone_number: null"` - очищает поле в БД (set to NULL)
// @Description Ограничения: `username` не может быть выставлено как null
// @Tags        users
// @Accept      json
// @Produce     json
// @Param       id path int true "ID изменяемого пользователя"
// @Param       request body PatchUserRequest true "PatchUser тело запроса"
// @Success     200 {object} PatchUserResponse "Успешно изменённый пользователь"
// @Failure     400 {object} core_http_response.ErrorResponse "Bad Request"
// @Failure     404 {object} core_http_response.ErrorResponse "User not found"
// @Failure     409 {object} core_http_response.ErrorResponse "Conflict"
// @Failure     500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router      /users/{id} [patch]
func (h *UsersHTTPHandler) PatchUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get userID path value")
		return
	}

	var request PatchUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")
		return
	}

	userPatch := userPatchFromRequest(request)

	userDomain, err := h.usersService.PatchUser(ctx, userID, userPatch)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to patch user",
		)
		return
	}

	response := PatchUserResponse(userDTOFromDomain(userDomain))

	responseHandler.JsonResponse(response, http.StatusOK)
}

func userPatchFromRequest(request PatchUserRequest) domain.UserPatch {
	return domain.NewUserPatch(
		request.Username.ToDomain(),
		request.PhoneNumber.ToDomain(),
	)
}
