package userHandler

import (
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/sri2103/chat_me/services/apiGateway/config"
)

func Test_userHandler_Register(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		cfg *config.Config
		// Named input parameters for target function.
		c       echo.Context
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(tt.cfg)
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
