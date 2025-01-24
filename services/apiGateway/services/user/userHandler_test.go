package userHandler

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	userMock "github.com/sri2103/chat_me/mocks/user"
	"github.com/sri2103/chat_me/protos/user"
	"github.com/sri2103/chat_me/services/apiGateway/config"
	"github.com/sri2103/chat_me/services/apiGateway/utils"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func newEchoContext(e *echo.Echo, req *http.Request, res http.ResponseWriter) echo.Context {
	return e.NewContext(req, res)
}
func Test_userHandler_Register(t *testing.T) {
	e := echo.New()
	e.Validator = utils.SetCustomValidator()

	t.Run("doing with wrong data", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(""))
		res := httptest.NewRecorder()
		ctx := newEchoContext(e, req, res)
		mockInterface := userMock.NewMockUserServiceClient(t)
		h := New(&config.Config{
			UserClientService: mockInterface,
		})
		err := h.Register(ctx)
		assert.Error(t, err)
	})

	t.Run("run a test server", func(t *testing.T) {

		data := &registerBody{
			Username: "uname",
			Email:    "uname@email.com",
			Password: "password",
		}
		d, _ := json.Marshal(&data)
		logger, _ := zap.NewDevelopment()

		defer func() {
			err := logger.Sync()
			if err != nil {
				log.Println(err)
			}
		}()

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(d)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		mockInterface := userMock.NewMockUserServiceClient(t)
		mockInterface.EXPECT().RegisterUserDetails(req.Context(), &user.RegisterUserRequest{
			Username: data.Username,
			Email:    data.Email,
			Password: data.Password,
		}).Return(&user.RegisterUserResponse{
			Success: true,
		}, nil)
		h := New(&config.Config{
			UserClientService: mockInterface,
			Log:               logger,
		})
		e.POST("/", h.Register)
		e.ServeHTTP(res, req)
		assert.Equal(t, http.StatusCreated, res.Code)

	})

	t.Run("user data adding", func(t *testing.T) {

	})

}

func Test_Login(t *testing.T) {
	e := echo.New()
	e.Validator = utils.SetCustomValidator()
	t.Run("Data error response here", func(t *testing.T) {
		reqData := &LoginRequest{
			Email:    "email",
			Password: "pwd",
		}

		d, _ := json.Marshal(reqData)

		reqTest := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(d)))

		reqTest.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()

		context := newEchoContext(e, reqTest, res)
		userService := userMock.NewMockUserServiceClient(t)
		userAuthServiceMock := userMock.NewMockAuthServiceClient(t)
		userService.EXPECT().
			AuthenticateUser(reqTest.Context(), &user.AuthenticateRequest{
				Email:    "email",
				Password: "pwd",
			}).
			Return(&user.AuthenticateResponse{
				UserId:        "abcd",
				Username:      "abcd",
				Authenticated: true,
			}, nil)
		userAuthServiceMock.EXPECT().
			CreateAuth(reqTest.Context(), &user.CreateAuthRequest{
				UserId: "abcd",
				AuthId: "abcd",
			}).Return(&user.CreateAuthResponse{
			Id:     "abcd",
			UserId: "abcd",
			AuthId: "abcd",
		}, nil).Once()

		h := New(&config.Config{
			UserClientService: userService,
			AuthClientService: userAuthServiceMock,
		})

		t.Run("handling controller", func(t *testing.T) {

			err := h.LoginUser(context)

			assert.NoError(t, err)
		})

		t.Run("login calling to a server", func(t *testing.T) {
			e.POST("/", h.LoginUser)
			e.ServeHTTP(res, reqTest)
			assert.Equal(t, http.StatusOK, res.Code)
		})
	})

}
