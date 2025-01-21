package userHandler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	userpb "github.com/sri2103/chat_me/protos/user"
	"github.com/sri2103/chat_me/services/apiGateway/config"
	"go.uber.org/zap"
)

type userHandler struct {
	logger      *zap.Logger
	userService userpb.UserServiceClient
	authService userpb.AuthServiceClient
}

func New(cfg *config.Config) *userHandler {
	return &userHandler{
		logger:      cfg.Log,
		userService: cfg.UserClientService,
		authService: cfg.AuthClientService,
	}
}

type registerBody struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (h *userHandler) Register(c echo.Context) error {
	var b registerBody
	err := c.Bind(&b)
	if err != nil {
		h.logger.Error("body error in register", zap.Error(err))
		return c.String(http.StatusBadRequest, "bad Request")
	}
	if err := c.Validate(b); err != nil {
		return err
	}
	var gr userpb.RegisterUserRequest
	gr.Email = b.Email
	gr.Username = b.Username
	gr.Password = b.Password

	resp, err := h.userService.RegisterUserDetails(c.Request().Context(), &gr)
	if err != nil {
		h.logger.Error("Error from the user service")
		return c.String(http.StatusInternalServerError, "Error from client user")
	}
	if !resp.Success {
		h.logger.Error("error from the internal of grpc")
		return c.String(http.StatusInternalServerError, "registration success")
	}
	return c.String(http.StatusCreated, "user crested")

}

type GetUserResponse struct {
	UserId   string `json:"user_id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func (h *userHandler) GetUserDetails(c echo.Context) error {
	userId := c.QueryParams().Get("userId")
	var gr userpb.GetUserRequest
	gr.UserId = userId
	resp, err := h.userService.GetUserDetails(c.Request().Context(), &gr)

	if err != nil {
		return c.String(http.StatusInternalServerError, "error from the userCLient")
	}

	data := &GetUserResponse{
		UserId:   resp.GetUserId(),
		Email:    resp.GetEmail(),
		Username: resp.GetUsername(),
	}

	return c.JSON(http.StatusOK, data)

}

type LoginRequest struct {
	Email    string `join:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token  string `json:"token"`
	AuthID string `json:"auth_id"`
}

func (h *userHandler) LoginUser(c echo.Context) error {
	var b LoginRequest
	err := c.Bind(&b)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = c.Validate(b)
	if err != nil {
		return err
	}

	AuthReq := &userpb.AuthenticateRequest{
		Username: b.Email,
		Password: b.Password,
	}

	resp, err := h.userService.AuthenticateUser(c.Request().Context(), AuthReq)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error in remote request")
	}

	if !resp.Authenticated {
		return echo.NewHTTPError(http.StatusUnauthorized, "user authentication failed")
	}

	token, err := CreateToken(&TokenPayload{
		Name:   resp.GetUsername(),
		UserId: resp.GetUserId(),
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error creating token")
	}

	authResp, err := h.authService.CreateAuth(c.Request().Context(), &userpb.CreateAuthRequest{
		UserId: resp.GetUserId(),
		AuthId: uuid.NewString(),
		Token:  *token,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error adding an auth entry")
	}

	return c.JSON(http.StatusOK, LoginResponse{
		Token:  *token,
		AuthID: authResp.AuthId,
	})
}

func (h *userHandler) Logout(c echo.Context) error {
	ctx := c.Request().Context()
	authId := ctx.Value("authId").(string)
	resp, err := h.authService.DestroyAuth(ctx, &userpb.DestroyAuthRequest{
		AuthId: authId,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error deleting auth session")
	}
	if !resp.Success {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error deleting auth")
	}
	return c.String(http.StatusOK, "logged out")
}
