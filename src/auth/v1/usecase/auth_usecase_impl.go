package usecase

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/Klinisia/backend-ksi/src/auth/v1/domain"
	"github.com/Klinisia/backend-ksi/src/auth/v1/dto"
	"github.com/Klinisia/backend-ksi/src/auth/v1/repository"
	"github.com/Klinisia/backend-ksi/src/shared"
	"github.com/Klinisia/backend-ksi/src/shared/external"
)

// AuthUsecaseImpl struct
type AuthUsecaseImpl struct {
	AuthRepositoryWrite repository.AuthRepository
	External            *external.SmsAcs
}

// NewAuthUsecaseImpl function
func NewAuthUsecaseImpl(AuthRepositoryWrite repository.AuthRepository, External *external.SmsAcs) *AuthUsecaseImpl {
	return &AuthUsecaseImpl{
		AuthRepositoryWrite: AuthRepositoryWrite,
		External:            External,
	}
}

// SignUpByPhone function
func (u *AuthUsecaseImpl) SignUpByPhone(filter *dto.SignUpByPhoneRequest) shared.Output {

	// assign Auth process
	signUp := u.AuthRepositoryWrite.SignUpByPhone(filter)
	if signUp.Error != nil {
		return shared.Output{Error: signUp.Error}
	}

	response := signUp.Result.(domain.SignUpByPhone)

	return shared.Output{Result: response}
}

// SignUpByPhone function
func (u *AuthUsecaseImpl) LoginByPhonePassword(params *dto.LoginByPhoneRequest) shared.Output {

	// check deleted user
	checkUserDel := u.AuthRepositoryWrite.CheckUserDelete(params)
	if checkUserDel == true {
		return shared.Output{Error: errors.New("Akun telah dihapus & tidak aktif")}
	}

	// check user exist
	checkUserExist := u.AuthRepositoryWrite.CheckUserExist(params)
	if checkUserExist.Error != nil {
		return shared.Output{Error: checkUserExist.Error}
	}

	// authenticate user
	getExistUser := checkUserExist.Result.(domain.SecUsers)
	authenticateUser := shared.AuthenticateUser(&getExistUser, params)

	if authenticateUser != nil {
		return shared.Output{Error: authenticateUser}

	}

	// get otp by sec patient id
	loadActiveSecPatient := u.AuthRepositoryWrite.LoadActiveSecPatient(getExistUser.SecUserId)
	if loadActiveSecPatient.Error != nil {
		return shared.Output{Error: loadActiveSecPatient.Error}

	}

	secPatientSignInOtp := loadActiveSecPatient.Result.(domain.SecPatientSignInOtp)
	expiredOtp := secPatientSignInOtp.ExpiredDatetime

	// check otp expired ?
	if expiredOtp.Before(time.Now().UTC()) {
		secPatientSignInOtp.SecPatientSignInOtpID = ""
	}

	// validateSmsRate
	var smsRateData dto.SmsRateData
	var patientLoginPasswordResp dto.PatientLoginPasswordResp
	countSms := u.AuthRepositoryWrite.CountSms(params.MobilePhone)
	smsRateData.SmsCount = countSms

	var intervalValidationCorrectionSecond, signInPatientLiveTimeSecond, durationSameNumberSecond,
		maxGroupSmsCount, intervalGroupSecond, intervalBetweenSmsSecond int

	strdurationSameNumberSecond, _ := os.LookupEnv("DURATIONSAMENUMBERSECOND")
	fmt.Sscan(strdurationSameNumberSecond, &durationSameNumberSecond)

	strintervalValidationCorrectionSecond, _ := os.LookupEnv("INTERVALVALIDATIONCORRECTIONSECOND")
	fmt.Sscan(strintervalValidationCorrectionSecond, &intervalValidationCorrectionSecond)

	strmaxGroupSmsCount, _ := os.LookupEnv("MAXGROUPSMSCOUNT")
	fmt.Sscan(strmaxGroupSmsCount, &maxGroupSmsCount)

	strintervalGroupSecond, _ := os.LookupEnv("INTERVALGROUPSECOND")
	fmt.Sscan(strintervalGroupSecond, &intervalGroupSecond)

	strintervalBetweenSmsSecond, _ := os.LookupEnv("INTERVALBETWEENSMSSECOND")
	fmt.Sscan(strintervalBetweenSmsSecond, &intervalBetweenSmsSecond)

	localDateTimeNow := time.Now().UTC()
	var allowSendSms bool
	var errorMessage string
	var allowSendLocalDateTime time.Time
	localDateTimeValidation := localDateTimeNow.Add(time.Second * time.Duration(intervalValidationCorrectionSecond))

	getSmsLog := u.AuthRepositoryWrite.GetSmsLog(params.MobilePhone)

	lastSms := getSmsLog.Result.(domain.SmsLog)

	lastSmsDatetime := lastSms.CreatedOn.UTC()

	if countSms >= maxGroupSmsCount {
		allowSendLocalDateTime = lastSmsDatetime.Add(time.Second * time.Duration(intervalGroupSecond))
		allowSendSms = !allowSendLocalDateTime.After(localDateTimeValidation)

	} else {

		if lastSms.SmsLogID != "" {
			allowSendLocalDateTime = localDateTimeNow.Add(time.Second * time.Duration(intervalBetweenSmsSecond))
			allowSendSms = true
		} else {
			allowSendLocalDateTime = lastSmsDatetime.Add(time.Second * time.Duration(intervalBetweenSmsSecond))
			allowSendSms = !allowSendLocalDateTime.After(localDateTimeValidation)
		}
	}

	if allowSendSms {
		errorMessage = ""

	} else {
		errorMessage = "Maaf, nomor " + params.MobilePhone + " dapat mengirim sms lagi pada " + allowSendLocalDateTime.String()
	}

	allowSendInSecond := 0

	if allowSendLocalDateTime.After(localDateTimeNow) {
		getAllowSendInSecond := allowSendLocalDateTime.Sub(localDateTimeNow)
		allowSendInSecond = int(getAllowSendInSecond.Seconds())
	} else {
		allowSendInSecond = 0
	}

	smsRateData.MobilePhone = params.MobilePhone
	smsRateData.AllowSendSms = allowSendSms
	smsRateData.ErrorMessage = errorMessage
	smsRateData.IntervalBetweenSmsSecond = intervalBetweenSmsSecond
	smsRateData.IntervalGroupSecond = intervalGroupSecond
	smsRateData.SmsCount = countSms
	smsRateData.LastSmsDateTime = lastSmsDatetime.Format("2006-01-02T15:04:05.99999")
	smsRateData.AllowSendInSecond = allowSendInSecond
	smsRateData.AllowSendLocalDateTime = allowSendLocalDateTime.Format("2006-01-02T15:04:05.99999")

	var dataOptSignIn domain.SecPatientSignInOtp
	if allowSendSms {
		// check jarak sms

		if lastSmsDatetime.Before(time.Now().UTC().Add(-time.Second * time.Duration(durationSameNumberSecond))) {

			otpSignIn := shared.RandomString(4)
			if secPatientSignInOtp.SecPatientSignInOtpID == "" {
				strsignInPatientLiveTimeSecond, _ := os.LookupEnv("SIGNINPATIENTLIVETIMESECOND")
				fmt.Sscan(strsignInPatientLiveTimeSecond, &signInPatientLiveTimeSecond)
				dataOptSignIn.SecPatientSignInOtpID = shared.GenerateUUID()
				dataOptSignIn.SecUserID = getExistUser.SecUserId
				dataOptSignIn.MobilePhone = params.MobilePhone
				dataOptSignIn.DeviceTypeCode = params.DeviceType
				dataOptSignIn.DeviceCode = params.DeviceCode
				dataOptSignIn.Otp = otpSignIn
				dataOptSignIn.ExpiredDatetime = time.Now().UTC().Add(time.Second * time.Duration(signInPatientLiveTimeSecond))
				dataOptSignIn.RetryCounter = 0
				dataOptSignIn.CreatedBy = getExistUser.SecUserId
				dataOptSignIn.UpdatedBy = getExistUser.SecUserId
				dataOptSignIn.CreatedOn = time.Now().UTC()
				dataOptSignIn.UpdatedOn = time.Now().UTC()
				u.AuthRepositoryWrite.SavePatientOtpSignIn(&dataOptSignIn)
			} else {
				dataOptSignIn.SecPatientSignInOtpID = secPatientSignInOtp.SecPatientSignInOtpID
				u.AuthRepositoryWrite.UpdatePatientOtpSignIn(&dataOptSignIn)
			}

			var smsReq shared.AcsSmsRequest
			strAcsPartnerID, _ := os.LookupEnv("ACS_PARTNER_ID")
			strAcsPartnerName, _ := os.LookupEnv("ACS_PARTNER_NAME")
			strAcsPassword, _ := os.LookupEnv("ACS_PASSWORD")
			smsReq.SmsBc.Request.Datetime = time.Now().UTC().Format("0102150405")
			smsReq.SmsBc.Request.DestinationNumber = params.MobilePhone
			smsReq.SmsBc.Request.Rrn = shared.GenerateUUID()
			smsReq.SmsBc.Request.PartnerId = strAcsPartnerID
			smsReq.SmsBc.Request.PartnerName = strAcsPartnerName
			smsReq.SmsBc.Request.Password = strAcsPassword
			smsReq.SmsBc.Request.Message = "OTP Login pengguna = " + otpSignIn + ". Jangan share OTP ke siapapun"
			// smsPayload := shared.CreateSmsXmlRequest(&smsReq)

			// saveSmsLog
			var dataSmsLog domain.SmsLog
			dataSmsLog.SmsStatus = false
			dataSmsLog.MobilePhone = params.MobilePhone
			dataSmsLog.SendingCount = 1
			dataSmsLog.SmsLogID = shared.GenerateUUID()
			dataSmsLog.SmsTypeCode = "userSignIn"
			dataSmsLog.SmsReffID = secPatientSignInOtp.SecPatientSignInOtpID
			dataSmsLog.SmsContent = "OTP Login = " + otpSignIn + ". Jangan share OTP ke siapapun"
			dataSmsLog.CreatedBy = getExistUser.SecUserId
			dataSmsLog.UpdatedBy = getExistUser.SecUserId
			dataSmsLog.CreatedOn = time.Now().UTC()
			dataSmsLog.UpdatedOn = time.Now().UTC()

			u.AuthRepositoryWrite.SaveSmsLogs(&dataSmsLog)

			// u.External.SendSms(smsPayload)

		} else {

			allowAfter20Sec := lastSmsDatetime.Add(time.Second * time.Duration(intervalBetweenSmsSecond))

			insecLeft := allowAfter20Sec.Sub(time.Now().UTC())
			leftSecondOf60 := int(insecLeft.Seconds())
			smsRateData.AllowSendInSecond = leftSecondOf60
			smsRateData.AllowSendSms = false
			smsRateData.AllowSendLocalDateTime = allowAfter20Sec.Format("2006-01-02T15:04:05.99999")
			errorMessage := "SMS terakhir untuk no hp " + params.MobilePhone + ", jam " + lastSmsDatetime.Format("2006-01-02 15:04:05") + ", minimal jarak " + strdurationSameNumberSecond + " detik"
			smsRateData.ErrorMessage = errorMessage
			patientLoginPasswordResp.SecPatientSignInOtpId = ""
		}

		// send sms otp
		// insert to sms log & sms log messages
		patientLoginPasswordResp.SecPatientSignInOtpId = dataOptSignIn.SecPatientSignInOtpID
	} else {

		patientLoginPasswordResp.SecPatientSignInOtpId = ""
	}

	patientLoginPasswordResp.SmsRateData = smsRateData

	return shared.Output{Result: patientLoginPasswordResp}
}

// SignUpByPhone function
func (u *AuthUsecaseImpl) LoginByPhoneOtp(params *dto.LoginByPhoneOtpRequest) shared.Output {

	signUp := u.AuthRepositoryWrite.LoginByPhoneOtp(params)
	if signUp.Error != nil {
		return shared.Output{Error: signUp.Error}
	}

	response := signUp.Result.(domain.SignUpByPhone)

	return shared.Output{Result: response}
}
