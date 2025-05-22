package handler

import (
	"context"
	"net/http"
	"wire-gorm-server/model"
	"wire-gorm-server/service"

	"github.com/go-chi/chi/v5"
	"github.com/swaggest/rest/nethttp"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)

type UserHandler struct {
	service *service.UserService
}


func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}


func (h *UserHandler) RegisterRoutes(r chi.Router) {
	r.Method(http.MethodGet, "/users", nethttp.NewHandler(h.GetUsers()))
}

// GetUsers returns a usecase.Interactor for GET /users
func (h *UserHandler) GetUsers() usecase.Interactor {
	// Output port
	type Response struct {
		Users []model.User `json:"users" required:"true"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, _ struct{}, out *Response) error {
		users, err := h.service.GetAllUsers()
		if err != nil {
			return status.Wrap(err, status.Internal)
		}
		out.Users = users
		return nil
	})

	u.SetTitle("Get Users Route")
	u.SetDescription("Returns all users. You can get from this route all user data like user name,id etc")
	u.SetTags("user")

	return u
}
