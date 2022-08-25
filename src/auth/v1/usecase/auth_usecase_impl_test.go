package usecase

import (
	"reflect"
	"testing"

	"github.com/Klinisia/backend-ksi/src/auth/v1/domain"
	"github.com/Klinisia/backend-ksi/src/auth/v1/dto"
	"github.com/Klinisia/backend-ksi/src/auth/v1/repository"
	"github.com/Klinisia/backend-ksi/src/shared"
	"github.com/Klinisia/backend-ksi/src/shared/external"
)

func TestNewAuthUsecaseImpl(t *testing.T) {
	type args struct {
		AuthRepositoryWrite repository.AuthRepository
		External            *external.SmsAcs
	}
	tests := []struct {
		name string
		args args
		want *AuthUsecaseImpl
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthUsecaseImpl(tt.args.AuthRepositoryWrite, tt.args.External); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthUsecaseImpl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthUsecaseImpl_SignInByPhonePassword(t *testing.T) {
	type fields struct {
		AuthRepositoryWrite repository.AuthRepository
		External            *external.SmsAcs
	}
	type args struct {
		params *dto.LoginByPhoneRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   shared.OutputV1
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &AuthUsecaseImpl{
				AuthRepositoryWrite: tt.fields.AuthRepositoryWrite,
				External:            tt.fields.External,
			}
			if got := u.SignInByPhonePassword(tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthUsecaseImpl.SignInByPhonePassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthUsecaseImpl_validateSignInSms(t *testing.T) {
	type fields struct {
		AuthRepositoryWrite repository.AuthRepository
		External            *external.SmsAcs
	}
	type args struct {
		params              *dto.LoginByPhoneRequest
		secPatientSignInOtp domain.SecPatientSignInOtp
		getExistUser        domain.SecUsers
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   shared.OutputV1
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &AuthUsecaseImpl{
				AuthRepositoryWrite: tt.fields.AuthRepositoryWrite,
				External:            tt.fields.External,
			}
			if got := u.validateSignInSms(tt.args.params, tt.args.secPatientSignInOtp, tt.args.getExistUser); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthUsecaseImpl.validateSignInSms() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthUsecaseImpl_SignInByPhoneOtp(t *testing.T) {
	type fields struct {
		AuthRepositoryWrite repository.AuthRepository
		External            *external.SmsAcs
	}
	type args struct {
		params *dto.LoginByPhoneOtpRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   shared.OutputV1
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &AuthUsecaseImpl{
				AuthRepositoryWrite: tt.fields.AuthRepositoryWrite,
				External:            tt.fields.External,
			}
			if got := u.SignInByPhoneOtp(tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthUsecaseImpl.SignInByPhoneOtp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthUsecaseImpl_ValidateOtp(t *testing.T) {
	type fields struct {
		AuthRepositoryWrite repository.AuthRepository
		External            *external.SmsAcs
	}
	type args struct {
		loginOtpReq         *dto.LoginByPhoneOtpRequest
		secPatientSignInOtp *domain.SecPatientSignInOtp
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   shared.OutputV1
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &AuthUsecaseImpl{
				AuthRepositoryWrite: tt.fields.AuthRepositoryWrite,
				External:            tt.fields.External,
			}
			if got := u.ValidateOtp(tt.args.loginOtpReq, tt.args.secPatientSignInOtp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthUsecaseImpl.ValidateOtp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthUsecaseImpl_CreateSuccessLoginResp(t *testing.T) {
	type fields struct {
		AuthRepositoryWrite repository.AuthRepository
		External            *external.SmsAcs
	}
	type args struct {
		secPatientSignInOtp *domain.SecPatientSignInOtp
		secUser             *domain.SecUsers
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   dto.LoginByPhoneOtpResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &AuthUsecaseImpl{
				AuthRepositoryWrite: tt.fields.AuthRepositoryWrite,
				External:            tt.fields.External,
			}
			if got := u.CreateSuccessLoginResp(tt.args.secPatientSignInOtp, tt.args.secUser); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthUsecaseImpl.CreateSuccessLoginResp() = %v, want %v", got, tt.want)
			}
		})
	}
}
