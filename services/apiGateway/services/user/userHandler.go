package userHandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	userpb "github.com/sri2103/chat_me/protos/user"
	"github.com/sri2103/chat_me/services/apiGateway/config"
	"go.uber.org/zap"
)

type userHandler struct {
	logger      *zap.Logger
	userService userpb.UserServiceClient
}

func New(cfg *config.Config) *userHandler {
	return &userHandler{
		logger:      cfg.Log,
		userService: cfg.UserClientService,
	}
}

type registerBody struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	UserId   string `json:"user_id"`
}

func (h *userHandler) Register(c echo.Context) error {
	var b registerBody
	err := c.Bind(&b)
	if err != nil {
		h.logger.Error("body error in register", zap.Error(err))
		return c.String(http.StatusBadRequest, "bad Request")
	}
	var gr userpb.UpdateUserRequest
	gr.Email = b.Email
	gr.UserId = b.UserId
	gr.Username = b.Username
	resp, err := h.userService.UpdateUserDetails(c.Request().Context(), &gr)
	if err != nil {
		h.logger.Error("Error from the user service")
		return c.String(http.StatusInternalServerError, "Error from client user")
	}
	if !resp.Success {
		h.logger.Error("error from the internal of grpc")
		return c.String(http.StatusInternalServerError)
	}
	return c.String(http.StatusCreated, "user crested")

}
