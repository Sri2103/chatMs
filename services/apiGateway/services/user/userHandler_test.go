package userHandler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	userMock "github.com/sri2103/chat_me/mocks/user"
	userpb "github.com/sri2103/chat_me/protos/user"
	"github.com/sri2103/chat_me/services/apiGateway/config"
)

func newEchoContex(e *echo.Echo, req *http.Request, res http.ResponseWriter) echo.Context {
	return e.NewContext(req, res)
}
func Test_userHandler_Register(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(""))
	res := httptest.NewRecorder()
	ctx := newEchoContex(e, req, res)
	mockInterface := userMock.NewMockUserServiceClient(t)
	mockInterface.EXPECT().UpdateUserDetails(req.Context(), &userpb.UpdateUserRequest{}).Return(&userpb.UpdateUserResponse{}, nil)
	h := New(&config.Config{
		UserClientService: mockInterface,
	})
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		cfg *config.Config
		// Named input parameters for target function.
		c       echo.Context
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "no-data test",
			c:       ctx,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := h.Register(tt.c)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Register() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Register() succeeded unexpectedly")
			}
		})
	}
}
