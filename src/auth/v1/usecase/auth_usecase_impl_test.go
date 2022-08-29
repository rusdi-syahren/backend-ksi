package usecase

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/rusdi-syahren/backend-ksi/src/auth/v1/domain"
	"github.com/rusdi-syahren/backend-ksi/src/auth/v1/dto"
	"github.com/rusdi-syahren/backend-ksi/src/auth/v1/repository"
	mockDB "github.com/rusdi-syahren/backend-ksi/src/auth/v1/repository/mocks"
	"github.com/rusdi-syahren/backend-ksi/src/shared"
	"github.com/rusdi-syahren/backend-ksi/src/shared/external"
	"github.com/stretchr/testify/mock"
)

func TestNewAuthUsecaseImpl(t *testing.T) {
	type args struct {
		AuthRepositoryWrite mockDB.AuthRepository
		External            *external.SmsAcs
	}
	tests := []struct {
		name string
		args args
		want *AuthUsecaseImpl
	}{
		{
			name: "testing TestNewAuthUsecaseImpl",
			args: args{AuthRepositoryWrite: mockDB.AuthRepository{}, External: &external.SmsAcs{}},
			want: &AuthUsecaseImpl{AuthRepositoryWrite: &mockDB.AuthRepository{}, External: &external.SmsAcs{}},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthUsecaseImpl(&tt.args.AuthRepositoryWrite, tt.args.External); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthUsecaseImpl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthUsecaseImpl_SignInByPhonePassword(t *testing.T) {

	tests := []struct {
		name                 string
		checkUserDelete      int
		checkUserExist       interface{}
		authenticateUser     error
		loadActiveSecPatient interface{}
		countSms             int
		getSmsLog            interface{}
		savePatientOtpSignIn interface{}
		args                 *dto.LoginByPhoneRequest
		want                 shared.OutputV1
	}{
		{
			name: "Testing SignInByPhonePassword - CheckUserDelete",

			args: &dto.LoginByPhoneRequest{},
			want: shared.OutputV1{Errors: shared.ErrorResponse{
				Field:   "",
				Code:    "SECURITY_USERPWD_INVALID",
				Message: "Akun telah dihapus & tidak aktif",
			}, Error: errors.New("Akun telah dihapus & tidak aktif"), Code: 401},
			checkUserDelete: 10,
		},
		{
			name: "Testing SignInByPhonePassword - checkUserExist",
			args: &dto.LoginByPhoneRequest{},
			want: shared.OutputV1{Errors: shared.ErrorResponse{
				Field:   "",
				Code:    "SECURITY_USERPWD_INVALID",
				Message: "ooops",
			}, Error: errors.New("ooops"), Code: 401},
			checkUserDelete: 0,
			checkUserExist:  shared.Output{Error: errors.New("ooops")},
		},

		{
			name: "Testing SignInByPhonePassword - AuthenticateUser",
			args: &dto.LoginByPhoneRequest{},
			want: shared.OutputV1{Errors: shared.ErrorResponse{
				Field:   "",
				Code:    "SECURITY_USERPWD_INVALID",
				Message: "Username atau password salah",
			}, Error: errors.New("Username atau password salah"), Code: 401},
			checkUserDelete:  0,
			checkUserExist:   shared.Output{Result: domain.SecUsers{}},
			authenticateUser: errors.New("Username atau password salah"),
		},
		{
			name: "Testing SignInByPhonePassword - LoadActiveSecPatient",
			args: &dto.LoginByPhoneRequest{Password: "pohodeui70"},
			want: shared.OutputV1{Errors: shared.ErrorResponse{
				Field:   "",
				Code:    "SECURITY_USERPWD_INVALID",
				Message: "Username atau password salah",
			}, Error: errors.New("Username atau password salah"), Code: 401},
			checkUserDelete:      0,
			checkUserExist:       shared.Output{Result: domain.SecUsers{SecUserId: "xxx", IsDeleted: false, IsActive: true, AccountExpired: false, CredentialsExpired: false, Password: "$2a$10$OLU5XvlNqy/S5qlFEcF9cuIKFnIjC/co1B3gm2T619SShFDy6tM16"}},
			authenticateUser:     nil,
			loadActiveSecPatient: shared.Output{Error: errors.New("Username atau password salah")},
		},
		{
			name: "Testing SignInByPhonePassword - LoadActiveSecPatient",
			args: &dto.LoginByPhoneRequest{Password: "pohodeui70"},
			want: shared.OutputV1{Result: shared.ErrorResponse{
				Field:   "",
				Code:    "SECURITY_USERPWD_INVALID",
				Message: "Username atau password salah",
			}, Errors: nil, Error: errors.New("Username atau password salah"), Code: 400},
			checkUserDelete:      0,
			checkUserExist:       shared.Output{Result: domain.SecUsers{SecUserId: "xxx", IsDeleted: false, IsActive: true, AccountExpired: false, CredentialsExpired: false, Password: "$2a$10$OLU5XvlNqy/S5qlFEcF9cuIKFnIjC/co1B3gm2T619SShFDy6tM16"}},
			authenticateUser:     nil,
			loadActiveSecPatient: shared.Output{Result: domain.SecPatientSignInOtp{ExpiredDatetime: time.Now().Local()}},
			countSms:             0,
			getSmsLog:            shared.Output{Result: domain.SmsLog{}},
			savePatientOtpSignIn: shared.Output{Result: nil, Error: errors.New("Username atau password salah")},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockArticleRepo := new(mockDB.AuthRepository)
			u := &AuthUsecaseImpl{AuthRepositoryWrite: mockArticleRepo, External: &external.SmsAcs{}}
			mockArticleRepo.On("CheckUserDelete", tt.args).Return(tt.checkUserDelete)
			mockArticleRepo.On("CheckUserExist", tt.args).Return(tt.checkUserExist)
			mockArticleRepo.On("AuthenticateUser", tt.checkUserExist, tt.args).Return(tt.checkUserExist)
			mockArticleRepo.On("LoadActiveSecPatient", mock.Anything).Return(tt.loadActiveSecPatient)
			mockArticleRepo.On("CountSms", mock.Anything).Return(tt.countSms)
			mockArticleRepo.On("GetSmsLog", mock.Anything).Return(tt.getSmsLog)
			mockArticleRepo.On("SavePatientOtpSignIn", mock.Anything).Return(tt.savePatientOtpSignIn)

			if got := u.SignInByPhonePassword(tt.args); !reflect.DeepEqual(got, tt.want) {
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
