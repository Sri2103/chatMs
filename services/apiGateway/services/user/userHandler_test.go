package userHandler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	userMock "github.com/sri2103/chat_me/mocks/user"
	"github.com/sri2103/chat_me/services/apiGateway/config"
	httpServer "github.com/sri2103/chat_me/services/apiGateway/servers/http"
	"github.com/stretchr/testify/assert"
)

func newEchoContex(e *echo.Echo, req *http.Request, res http.ResponseWriter) echo.Context {
	return e.NewContext(req, res)
}
func Test_userHandler_Register(t *testing.T) {
	e := echo.New()
	e = httpServer.SetCustomValidator(e)

	t.Run("doing with wrong data", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(""))
		res := httptest.NewRecorder()
		ctx := newEchoContex(e, req, res)
		mockInterface := userMock.NewMockUserServiceClient(t)
		h := New(&config.Config{
			UserClientService: mockInterface,
		})
		err := h.Register(ctx)
		assert.Error(t, err)
	})

	t.Run("run a test server", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(""))
		res := httptest.NewRecorder()
		mockInterface := userMock.NewMockUserServiceClient(t)
		h := New(&config.Config{
			UserClientService: mockInterface,
		})
		e.POST("/", h.Register)
		e.ServeHTTP(res, req)
		assert.Equal(t, http.StatusBadRequest, res.Code)

	})

	t.Run("user data adding", func(t *testing.T) {

	})

}
