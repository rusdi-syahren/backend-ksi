package repository

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/rusdi-syahren/backend-ksi/src/auth/v1/domain"
	"github.com/rusdi-syahren/backend-ksi/src/auth/v1/dto"
	"github.com/rusdi-syahren/backend-ksi/src/shared"

	// "github.com/jinzhu/gorm"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDbMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	sqlDB, mock, err := sqlmock.New()
	mockDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	return mockDB, mock, err
}

func TestNewAuthRepositoryGorm(t *testing.T) {

	mockDB, _, _ := NewDbMock()

	tests := []struct {
		name string
		args *gorm.DB
		want *AuthRepositoryGorm
	}{
		{
			name: "connection test",
			args: mockDB,
			want: &AuthRepositoryGorm{mockDB},
		},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthRepositoryGorm(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthRepositoryGorm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthRepositoryGorm_LoginByPhone(t *testing.T) {

	type args struct {
		params *dto.LoginByPhoneRequest
	}
	tests := []struct {
		name       string
		args       args
		beforeTest func(sqlmock.Sqlmock)
		want       shared.Output
		wantErr    bool
	}{
		{
			name: "fail login byPhone",
			args: args{
				params: &dto.LoginByPhoneRequest{DeviceType: "mobile", DeviceCode: "123", MobilePhone: "6281584832993", Password: "pohodeui70"},
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.
					ExpectQuery(regexp.QuoteMeta(`SELECT * FROM user_management where phone = $1`)).
					WithArgs("6281584832993").
					WillReturnError(errors.New("data not found"))
			},
			want:    shared.Output{Result: nil, Error: errors.New("data not found")},
			wantErr: true,
		},

		{
			name: "success login byPhone",
			args: args{
				params: &dto.LoginByPhoneRequest{DeviceType: "mobile", DeviceCode: "123", MobilePhone: "6281584832993", Password: "pohodeui70"},
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.
					ExpectQuery(regexp.QuoteMeta(`SELECT * FROM user_management where phone = $1`)).
					WithArgs("6281584832993").
					WillReturnRows(sqlmock.NewRows([]string{"mobilePhone", "secUserSignUpId", "expiredDatetime", "smsRateData"}).
						AddRow("xxxx", "xxxxx", "xxxx", "xxxxx"))
			},
			want: shared.Output{Result: domain.SignUpByPhone{}, Error: nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := NewDbMock()

			u := &AuthRepositoryGorm{
				db: mockDB,
			}

			if tt.beforeTest != nil {
				tt.beforeTest(mockSQL)
			}

			got := u.LoginByPhone(tt.args.params)
			if (got.Error != nil) != tt.wantErr {
				t.Errorf("AuthRepositoryGorm.LoginByPhone() error = %v, wantErr %v", got.Error, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthRepositoryGorm.LoginByPhone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthRepositoryGorm_LoginByPhoneOtp(t *testing.T) {
	type args struct {
		params *dto.LoginByPhoneOtpRequest
	}
	tests := []struct {
		name       string
		args       args
		beforeTest func(sqlmock.Sqlmock)
		want       shared.Output
		wantErr    bool
	}{
		{
			name: "fail login by otp",
			args: args{
				params: &dto.LoginByPhoneOtpRequest{SecPatientSignInOtpId: "", Otp: ""},
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.
					ExpectQuery(regexp.QuoteMeta(`SELECT * FROM user_management
					where secPatientSignInOtpId = $1`)).
					WithArgs("xxxx").
					WillReturnError(errors.New("data not found"))
			},
			want:    shared.Output{Result: nil, Error: errors.New("data not found")},
			wantErr: true,
		},

		{
			name: "success login by otp",
			args: args{
				params: &dto.LoginByPhoneOtpRequest{SecPatientSignInOtpId: "xxxx", Otp: "xxxx"},
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.
					ExpectQuery(regexp.QuoteMeta(`SELECT * FROM user_management
					where secPatientSignInOtpId = $1`)).
					WithArgs("xxxx").
					WillReturnRows(sqlmock.NewRows([]string{"mobilePhone", "secUserSignUpId", "expiredDatetime", "smsRateData"}).
						AddRow("xxxx", "xxxx", "xxxx", "xxxx"))
			},
			want: shared.Output{Result: domain.SignUpByPhone{}, Error: nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := NewDbMock()

			u := &AuthRepositoryGorm{
				db: mockDB,
			}

			if tt.beforeTest != nil {
				tt.beforeTest(mockSQL)
			}

			got := u.LoginByPhoneOtp(tt.args.params)
			if (got.Error != nil) != tt.wantErr {
				t.Errorf("AuthRepositoryGorm.LoginByPhoneOtp() error = %v, wantErr %v", got.Error, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthRepositoryGorm.LoginByPhoneOtp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthRepositoryGorm_CheckUserDelete(t *testing.T) {
	type args struct {
		params *dto.LoginByPhoneRequest
	}
	tests := []struct {
		name       string
		args       args
		beforeTest func(sqlmock.Sqlmock)
		want       int
		wantErr    bool
	}{
		// {
		// 	name: "fail check user delete",
		// 	args: args{
		// 		params: &dto.LoginByPhoneRequest{DeviceType: "mobile", DeviceCode: "123", MobilePhone: "6281584832993", Password: "pohodeui70"},
		// 	},
		// 	beforeTest: func(mockSQL sqlmock.Sqlmock) {
		// 		mockSQL.
		// 			ExpectQuery(regexp.QuoteMeta(`SELECT count(*) total FROM security.sec_users
		// 			where mobile_phone = $1 and is_active = false and is_deleted = true`)).
		// 			WithArgs("xxxx").
		// 			WillReturnRows(sqlmock.NewRows([]string{"mobilePhone"}).
		// 				AddRow(false))
		// 	},
		// 	want:    shared.Output{Result: false, Error: nil},
		// 	wantErr: false,
		// },

		// {
		// 	name: "success check user delete",
		// 	args: args{
		// 		params: &dto.LoginByPhoneRequest{DeviceType: "mobile", DeviceCode: "123", MobilePhone: "6281584832993", Password: "pohodeui70"},
		// 	},
		// 	beforeTest: func(mockSQL sqlmock.Sqlmock) {
		// 		mockSQL.
		// 			ExpectQuery(regexp.QuoteMeta(`SELECT count(*) total FROM security.sec_users
		// 			where mobile_phone = $1 and is_active = false and is_deleted = true`)).
		// 			WithArgs("6281584832993").
		// 			WillReturnRows(sqlmock.NewRows([]string{"Total"}).
		// 				AddRow(0))
		// 	},
		// 	want:    shared.Output{Result: false, Error: nil},
		// 	wantErr: false,
		// },
		{
			name: "success check user delete",
			args: args{
				params: &dto.LoginByPhoneRequest{DeviceType: "mobile", DeviceCode: "123", MobilePhone: "6281584832993", Password: "pohodeui70"},
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.
					ExpectQuery(regexp.QuoteMeta(`SELECT count(*) total FROM security.sec_users
					where mobile_phone = $1 and is_active = false and is_deleted = true`)).
					WithArgs("6281584832993").
					WillReturnRows(sqlmock.NewRows([]string{"Total"}).
						AddRow(10))
			},
			want:    10,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := NewDbMock()

			u := &AuthRepositoryGorm{
				db: mockDB,
			}

			if tt.beforeTest != nil {
				tt.beforeTest(mockSQL)
			}

			got := u.CheckUserDelete(tt.args.params)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthRepositoryGorm.CheckUserDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthRepositoryGorm_CheckUserExist(t *testing.T) {
	type args struct {
		params *dto.LoginByPhoneRequest
	}
	tests := []struct {
		name       string
		args       args
		beforeTest func(sqlmock.Sqlmock)
		want       shared.Output
		wantErr    bool
	}{
		{
			name: "fail check user exist 1",
			args: args{
				params: &dto.LoginByPhoneRequest{DeviceType: "mobile", DeviceCode: "123", MobilePhone: "6281584832993", Password: "pohodeui70"},
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.
					ExpectQuery(regexp.QuoteMeta(`SELECT  * FROM security.sec_users
					where mobile_phone = $1 and user_type_code='patient' and is_active = true and is_deleted = false`)).
					WithArgs("6281584832993").
					WillReturnError(errors.New("data not found"))
			},
			want:    shared.Output{Result: nil, Error: errors.New("data not found")},
			wantErr: true,
		},

		{
			name: "fail check user exist 2",
			args: args{
				params: &dto.LoginByPhoneRequest{DeviceType: "mobile", DeviceCode: "123", MobilePhone: "6281584832993", Password: "pohodeui70"},
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.
					ExpectQuery(regexp.QuoteMeta(`SELECT  * FROM security.sec_users
					where mobile_phone = $1 and user_type_code='patient' and is_active = true and is_deleted = false`)).
					WithArgs("6281584832993").
					WillReturnRows(sqlmock.NewRows([]string{"SecUserId"}).
						AddRow(""))
			},
			wantErr: true,
			want:    shared.Output{Result: nil, Error: errors.New("data not found")},
		},

		{
			name: "success check user exist",
			args: args{
				params: &dto.LoginByPhoneRequest{DeviceType: "mobile", DeviceCode: "123", MobilePhone: "6281584832993", Password: "pohodeui70"},
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.
					ExpectQuery(regexp.QuoteMeta(`SELECT  * FROM security.sec_users
					where mobile_phone = $1 and user_type_code='patient' and is_active = true and is_deleted = false`)).
					WithArgs("6281584832993").
					WillReturnRows(sqlmock.NewRows([]string{"SecUserId"}).
						AddRow("6281584832993"))
			},
			wantErr: false,
			want:    shared.Output{Result: domain.SecUsers{SecUserId: "6281584832993"}, Error: nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := NewDbMock()

			u := &AuthRepositoryGorm{
				db: mockDB,
			}

			if tt.beforeTest != nil {
				tt.beforeTest(mockSQL)
			}

			got := u.CheckUserExist(tt.args.params)
			if (got.Error != nil) != tt.wantErr {
				t.Errorf("AuthRepositoryGorm.CheckUserExist() error = %v, wantErr %v", got.Error, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthRepositoryGorm.CheckUserExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthRepositoryGorm_GetSecUserByID(t *testing.T) {
	type args struct {
		params string
	}
	tests := []struct {
		name       string
		args       args
		beforeTest func(sqlmock.Sqlmock)
		want       shared.Output
		wantErr    bool
	}{
		{
			name: "fail getsecuser by id 1",
			args: args{
				params: "secxxxxxx",
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.
					ExpectQuery(regexp.QuoteMeta(`SELECT  * FROM security.sec_users
					where sec_user_id = $1`)).
					WithArgs("secxxxxxx").
					WillReturnError(errors.New("data not found"))
			},
			want:    shared.Output{Result: nil, Error: errors.New("data not found")},
			wantErr: true,
		},

		{
			name: "fail getsecuser by id 2",
			args: args{
				params: "secxxxxxx",
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.
					ExpectQuery(regexp.QuoteMeta(`SELECT  * FROM security.sec_users
					where sec_user_id = $1`)).
					WithArgs("secxxxxxx").
					WillReturnRows(sqlmock.NewRows([]string{"SecUserId"}).
						AddRow(""))
			},
			wantErr: true,
			want:    shared.Output{Result: nil, Error: errors.New("data not found")},
		},

		{
			name: "success check user exist",
			args: args{
				params: "secxxxxxx",
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.
					ExpectQuery(regexp.QuoteMeta(`SELECT  * FROM security.sec_users
					where sec_user_id = $1`)).
					WithArgs("secxxxxxx").
					WillReturnRows(sqlmock.NewRows([]string{"SecUserId"}).
						AddRow("secxxxxxx"))
			},
			wantErr: false,
			want:    shared.Output{Result: domain.SecUsers{SecUserId: "secxxxxxx"}, Error: nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := NewDbMock()

			u := &AuthRepositoryGorm{
				db: mockDB,
			}

			if tt.beforeTest != nil {
				tt.beforeTest(mockSQL)
			}

			got := u.GetSecUserByID(tt.args.params)
			if (got.Error != nil) != tt.wantErr {
				t.Errorf("AuthRepositoryGorm.GetSecUserByID() error = %v, wantErr %v", got.Error, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthRepositoryGorm.GetSecUserByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthRepositoryGorm_LoadActiveSecPatient(t *testing.T) {
	type args struct {
		params string
	}
	tests := []struct {
		name       string
		args       args
		beforeTest func(sqlmock.Sqlmock)
		want       shared.Output
		wantErr    bool
	}{
		{
			name: "fail load active user",
			args: args{
				params: "secxxxxxx",
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.
					ExpectQuery(regexp.QuoteMeta(`SELECT  * FROM security.sec_patient_sign_in_otps
					where sec_user_id = $1  and is_active = true and is_deleted = false order by created_on desc limit 1`)).
					WithArgs("secxxxxxx").
					WillReturnError(errors.New("data not found"))
			},
			want:    shared.Output{Result: nil, Error: errors.New("data not found")},
			wantErr: true,
		},

		{
			name: "success load active user",
			args: args{
				params: "secxxxxxx",
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.
					ExpectQuery(regexp.QuoteMeta(`SELECT  * FROM security.sec_patient_sign_in_otps
					where sec_user_id = $1  and is_active = true and is_deleted = false order by created_on desc limit 1`)).
					WithArgs("secxxxxxx").
					WillReturnRows(sqlmock.NewRows([]string{"SecPatientSignInOtpID"}).
						AddRow("secxxxxxx"))
			},
			wantErr: false,
			want:    shared.Output{Result: domain.SecPatientSignInOtp{SecPatientSignInOtpID: "secxxxxxx"}, Error: nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := NewDbMock()

			u := &AuthRepositoryGorm{
				db: mockDB,
			}

			if tt.beforeTest != nil {
				tt.beforeTest(mockSQL)
			}

			got := u.LoadActiveSecPatient(tt.args.params)
			if (got.Error != nil) != tt.wantErr {
				t.Errorf("AuthRepositoryGorm.LoadActiveSecPatient() error = %v, wantErr %v", got.Error, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthRepositoryGorm.LoadActiveSecPatient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthRepositoryGorm_CountSms(t *testing.T) {
	type args struct {
		params string
	}
	tests := []struct {
		name       string
		args       args
		beforeTest func(sqlmock.Sqlmock)
		want       int
		wantErr    bool
	}{
		{
			name: "fail check user delete",
			args: args{
				params: "xxxx",
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.
					ExpectQuery(regexp.QuoteMeta(`SELECT count(*) total FROM sms.sms_logs
					where mobile_phone = $1 and  created_on > $2`)).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows([]string{"Total"}).
						AddRow(10))
			},
			want:    10,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := NewDbMock()

			u := &AuthRepositoryGorm{
				db: mockDB,
			}

			if tt.beforeTest != nil {
				tt.beforeTest(mockSQL)
			}

			got := u.CountSms(tt.args.params)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthRepositoryGorm.CountSms() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthRepositoryGorm_GetSmsLog(t *testing.T) {
	type args struct {
		params string
	}
	tests := []struct {
		name       string
		args       args
		beforeTest func(sqlmock.Sqlmock)
		want       shared.Output
		wantErr    bool
	}{

		{
			name: "success get sms log",
			args: args{
				params: "xxx",
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.
					ExpectQuery(regexp.QuoteMeta(`SELECT * FROM sms.sms_logs
					where mobile_phone = $1  order by created_on desc limit 1`)).
					WithArgs(sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows([]string{"SmsLogID"}).
						AddRow("xxxx"))
			},
			want: shared.Output{Result: domain.SmsLog{SmsLogID: "xxxx"}, Error: nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := NewDbMock()

			u := &AuthRepositoryGorm{
				db: mockDB,
			}

			if tt.beforeTest != nil {
				tt.beforeTest(mockSQL)
			}

			got := u.GetSmsLog(tt.args.params)
			if (got.Error != nil) != tt.wantErr {
				t.Errorf("AuthRepositoryGorm.GetSmsLog() error = %v, wantErr %v", got.Error, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthRepositoryGorm.GetSmsLog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthRepositoryGorm_SavePatientOtpSignIn(t *testing.T) {
	type args struct {
		params *domain.SecPatientSignInOtp
	}
	tests := []struct {
		name       string
		args       args
		beforeTest func(sqlmock.Sqlmock)
		want       shared.Output
		wantErr    bool
	}{
		{
			name: "failed save sms log message",
			args: args{
				params: &domain.SecPatientSignInOtp{},
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.ExpectBegin()
				mockSQL.
					ExpectExec(regexp.QuoteMeta(`INSERT INTO "security"."sec_patient_sign_in_otps" ("sec_patient_sign_in_otp_id","sec_user_id","mobile_phone","device_type_code","device_code","otp","expired_datetime","retry_counter","is_active","is_deleted","created_by","created_on","updated_by","updated_on") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14)`)).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnError(fmt.Errorf("failed insert"))

				mockSQL.ExpectRollback()
			},
			wantErr: true,
			want:    shared.Output{Result: &domain.SecPatientSignInOtp{}, Error: fmt.Errorf("failed insert")},
		},

		{
			name: "success save sms log message",
			args: args{
				params: &domain.SecPatientSignInOtp{},
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.ExpectBegin()
				mockSQL.
					ExpectExec(regexp.QuoteMeta(`INSERT INTO "security"."sec_patient_sign_in_otps" ("sec_patient_sign_in_otp_id","sec_user_id","mobile_phone","device_type_code","device_code","otp","expired_datetime","retry_counter","is_active","is_deleted","created_by","created_on","updated_by","updated_on") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14)`)).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mockSQL.ExpectCommit()
			},
			wantErr: false,
			want:    shared.Output{Result: &domain.SecPatientSignInOtp{}, Error: nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := NewDbMock()

			u := &AuthRepositoryGorm{
				db: mockDB,
			}

			if tt.beforeTest != nil {
				tt.beforeTest(mockSQL)
			}

			got := u.SavePatientOtpSignIn(tt.args.params)
			if (got.Error != nil) != tt.wantErr {
				t.Errorf("AuthRepositoryGorm.SavePatientOtpSignIn() error = %v, wantErr %v", got.Error, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthRepositoryGorm.SavePatientOtpSignIn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthRepositoryGorm_UpdatePatientOtpSignIn(t *testing.T) {
	type args struct {
		params *domain.SecPatientSignInOtp
	}
	tests := []struct {
		name       string
		args       args
		beforeTest func(sqlmock.Sqlmock)
		want       shared.Output
		wantErr    bool
	}{

		{
			name: "success load active user",
			args: args{
				params: &domain.SecPatientSignInOtp{},
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.
					ExpectQuery(regexp.QuoteMeta(`update security.sec_patient_sign_in_otps set retry_counter = retry_counter + 1 , created_on = $1  where sec_patient_sign_in_otp_id = $2`)).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows([]string{"SecPatientSignInOtpID"}).
						AddRow("secxxxxxx"))
			},
			wantErr: false,
			want:    shared.Output{Result: &domain.SecPatientSignInOtp{SecPatientSignInOtpID: "secxxxxxx"}, Error: nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := NewDbMock()

			u := &AuthRepositoryGorm{
				db: mockDB,
			}

			if tt.beforeTest != nil {
				tt.beforeTest(mockSQL)
			}

			got := u.UpdatePatientOtpSignIn(tt.args.params)
			if (got.Error != nil) != tt.wantErr {
				t.Errorf("AuthRepositoryGorm.UpdatePatientOtpSignIn() error = %v, wantErr %v", got.Error, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthRepositoryGorm.UpdatePatientOtpSignIn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthRepositoryGorm_SaveSmsLogs(t *testing.T) {
	type args struct {
		params *domain.SmsLog
	}
	tests := []struct {
		name       string
		args       args
		beforeTest func(sqlmock.Sqlmock)
		want       shared.Output
		wantErr    bool
	}{
		{
			name: "failed save sms log message",
			args: args{
				params: &domain.SmsLog{},
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.ExpectBegin()
				mockSQL.
					ExpectExec(regexp.QuoteMeta(`INSERT INTO "sms"."sms_logs" ("sms_log_id","sms_reff_id","sms_type_code","mobile_phone","sms_content","sending_count","sms_status","created_by","created_on","updated_by","updated_on") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`)).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnError(fmt.Errorf("failed insert"))

				mockSQL.ExpectRollback()
			},
			wantErr: true,
			want:    shared.Output{Result: &domain.SmsLog{}, Error: fmt.Errorf("failed insert")},
		},

		{
			name: "success save sms log message",
			args: args{
				params: &domain.SmsLog{},
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.ExpectBegin()
				mockSQL.
					ExpectExec(regexp.QuoteMeta(`INSERT INTO "sms"."sms_logs" ("sms_log_id","sms_reff_id","sms_type_code","mobile_phone","sms_content","sending_count","sms_status","created_by","created_on","updated_by","updated_on") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`)).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mockSQL.ExpectCommit()
			},
			wantErr: false,
			want:    shared.Output{Result: &domain.SmsLog{}, Error: nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := NewDbMock()

			u := &AuthRepositoryGorm{
				db: mockDB,
			}

			if tt.beforeTest != nil {
				tt.beforeTest(mockSQL)
			}

			got := u.SaveSmsLogs(tt.args.params)
			if (got.Error != nil) != tt.wantErr {
				t.Errorf("AuthRepositoryGorm.SaveSmsLogs() error = %v, wantErr %v", got.Error, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthRepositoryGorm.SaveSmsLogs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthRepositoryGorm_SaveSmsLogMessages(t *testing.T) {
	type args struct {
		params *domain.SmsLogMessage
	}
	tests := []struct {
		name       string
		args       args
		beforeTest func(sqlmock.Sqlmock)
		want       shared.Output
		wantErr    bool
	}{
		{
			name: "failed save sms log message",
			args: args{
				params: &domain.SmsLogMessage{},
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.ExpectBegin()
				mockSQL.
					ExpectExec(regexp.QuoteMeta(`INSERT INTO "sms"."sms_log_messages" ("sms_log_message_id","sms_log_id","message_rrn","req_res_type_code","response_code","response_message","xml_message","created_by","created_on","updated_by","updated_on") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`)).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnError(fmt.Errorf("failed insert"))

				mockSQL.ExpectRollback()
			},
			wantErr: true,
			want:    shared.Output{Result: &domain.SmsLogMessage{}, Error: fmt.Errorf("failed insert")},
		},

		{
			name: "success save sms log message",
			args: args{
				params: &domain.SmsLogMessage{},
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.ExpectBegin()
				mockSQL.
					ExpectExec(regexp.QuoteMeta(`INSERT INTO "sms"."sms_log_messages" ("sms_log_message_id","sms_log_id","message_rrn","req_res_type_code","response_code","response_message","xml_message","created_by","created_on","updated_by","updated_on") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`)).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mockSQL.ExpectCommit()
			},
			wantErr: false,
			want:    shared.Output{Result: &domain.SmsLogMessage{}, Error: nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := NewDbMock()

			u := &AuthRepositoryGorm{
				db: mockDB,
			}

			if tt.beforeTest != nil {
				tt.beforeTest(mockSQL)
			}

			got := u.SaveSmsLogMessages(tt.args.params)
			if (got.Error != nil) != tt.wantErr {
				t.Errorf("AuthRepositoryGorm.SaveSmsLogMessages() error = %v, wantErr %v", got.Error, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthRepositoryGorm.SaveSmsLogMessages() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthRepositoryGorm_FindBySecPatientSignInOtp(t *testing.T) {
	type args struct {
		params *dto.LoginByPhoneOtpRequest
	}
	tests := []struct {
		name       string
		args       args
		beforeTest func(sqlmock.Sqlmock)
		want       shared.Output
		wantErr    bool
	}{

		{
			name: "success FindBySecPatientSignInOtp",
			args: args{
				params: &dto.LoginByPhoneOtpRequest{SecPatientSignInOtpId: "xxxx", Otp: "xxxx"},
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.
					ExpectQuery(regexp.QuoteMeta(`SELECT * FROM security.sec_patient_sign_in_otps
					where sec_patient_sign_in_otp_id = $1 and is_active = true and is_deleted = false`)).
					WithArgs(sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows([]string{"SecUserID"}).
						AddRow("xxxx"))
			},
			want: shared.Output{Result: domain.SecPatientSignInOtp{SecUserID: "xxxx"}, Error: nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := NewDbMock()

			u := &AuthRepositoryGorm{
				db: mockDB,
			}

			if tt.beforeTest != nil {
				tt.beforeTest(mockSQL)
			}

			got := u.FindBySecPatientSignInOtp(tt.args.params)
			if (got.Error != nil) != tt.wantErr {
				t.Errorf("AuthRepositoryGorm.FindBySecPatientSignInOtp() error = %v, wantErr %v", got.Error, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthRepositoryGorm.FindBySecPatientSignInOtp() = %v, want %v", got, tt.want)
			}
		})
	}
}
