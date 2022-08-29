package delivery

import (
	"reflect"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/rusdi-syahren/backend-ksi/src/auth/v1/usecase"
)

func TestNewEchoHandler(t *testing.T) {
	type args struct {
		AuthUsecase usecase.AuthUsecase
	}
	tests := []struct {
		name string
		args args
		want *EchoHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEchoHandler(tt.args.AuthUsecase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEchoHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEchoHandler_Mount(t *testing.T) {
	type fields struct {
		AuthUsecase usecase.AuthUsecase
	}
	type args struct {
		group *echo.Group
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &EchoHandler{
				AuthUsecase: tt.fields.AuthUsecase,
			}
			h.Mount(tt.args.group)
		})
	}
}

func TestEchoHandler_SignInByPhonePassword(t *testing.T) {
	type fields struct {
		AuthUsecase usecase.AuthUsecase
	}
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &EchoHandler{
				AuthUsecase: tt.fields.AuthUsecase,
			}
			if err := h.SignInByPhonePassword(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("EchoHandler.SignInByPhonePassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEchoHandler_SignInByPhoneOtp(t *testing.T) {
	type fields struct {
		AuthUsecase usecase.AuthUsecase
	}
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &EchoHandler{
				AuthUsecase: tt.fields.AuthUsecase,
			}
			if err := h.SignInByPhoneOtp(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("EchoHandler.SignInByPhoneOtp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
