package usecase

import (
	"errors"
	"fmt"
	"net/http"
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
func (u *AuthUsecaseImpl) SignInByPhonePassword(params *dto.LoginByPhoneRequest) shared.OutputV1 {

	// check deleted user
	checkUserDel := u.AuthRepositoryWrite.CheckUserDelete(params)
	if checkUserDel == true {
		errResponse := shared.ErrorResponse{
			Field:   "",
			Code:    "SECURITY_USERPWD_INVALID",
			Message: "Akun telah dihapus & tidak aktif",
		}
		return shared.OutputV1{
			Error:  errors.New("Akun telah dihapus & tidak aktif"),
			Errors: errResponse,
			Code:   http.StatusUnauthorized}
	}

	// check user exist
	checkUserExist := u.AuthRepositoryWrite.CheckUserExist(params)
	if checkUserExist.Error != nil {
		errResponse := shared.ErrorResponse{
			Field:   "",
			Code:    "SECURITY_USERPWD_INVALID",
			Message: checkUserExist.Error.Error(),
		}
		return shared.OutputV1{
			Error:  checkUserExist.Error,
			Errors: errResponse,
			Code:   http.StatusUnauthorized}
	}

	// authenticate user
	getExistUser := checkUserExist.Result.(domain.SecUsers)
	authenticateUser := shared.AuthenticateUser(&getExistUser, params)
	if authenticateUser != nil {
		errResponse := shared.ErrorResponse{
			Field:   "",
			Code:    "SECURITY_USERPWD_INVALID",
			Message: authenticateUser.Error(),
		}
		return shared.OutputV1{
			Error:  authenticateUser,
			Errors: errResponse,
			Code:   http.StatusUnauthorized}
	}

	// get otp by sec patient id
	loadActiveSecPatient := u.AuthRepositoryWrite.LoadActiveSecPatient(getExistUser.SecUserId)
	if loadActiveSecPatient.Error != nil {
		errResponse := shared.ErrorResponse{
			Field:   "",
			Code:    "SECURITY_USERPWD_INVALID",
			Message: loadActiveSecPatient.Error.Error(),
		}
		return shared.OutputV1{
			Error:  loadActiveSecPatient.Error,
			Errors: errResponse,
			Code:   http.StatusUnauthorized}
	}

	secPatientSignInOtp := loadActiveSecPatient.Result.(domain.SecPatientSignInOtp)
	expiredOtp := secPatientSignInOtp.ExpiredDatetime.Local().Add(-time.Hour * 7)

	if expiredOtp.Before(time.Now().Local()) {
		secPatientSignInOtp.SecPatientSignInOtpID = ""
	}

	// validateSmsRate
	return u.validateSignInSms(params, secPatientSignInOtp, getExistUser)

}

func (u *AuthUsecaseImpl) validateSignInSms(
	params *dto.LoginByPhoneRequest,
	secPatientSignInOtp domain.SecPatientSignInOtp,
	getExistUser domain.SecUsers) shared.OutputV1 {
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

	localDateTimeNow := time.Now().Local()
	var allowSendSms bool
	var errorMessage string
	var allowSendLocalDateTime time.Time
	localDateTimeValidation := localDateTimeNow.Add(time.Second * time.Duration(intervalValidationCorrectionSecond))

	getSmsLog := u.AuthRepositoryWrite.GetSmsLog(params.MobilePhone)
	lastSms := getSmsLog.Result.(domain.SmsLog)
	lastSmsDatetime := lastSms.CreatedOn.Local().Add(-time.Hour * 7)

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
		errorMessage = fmt.Sprintf(external.SmsValidationLimitGroup, params.MobilePhone, allowSendLocalDateTime.Format(external.DateFormatStr))
	}

	allowSendInSecond := 0
	if allowSendLocalDateTime.After(localDateTimeNow) {
		getAllowSendInSecond := allowSendLocalDateTime.Sub(localDateTimeNow)
		allowSendInSecond = int(getAllowSendInSecond.Seconds())
	}

	smsRateData.MobilePhone = params.MobilePhone
	smsRateData.AllowSendSms = allowSendSms
	smsRateData.ErrorMessage = errorMessage
	smsRateData.IntervalBetweenSmsSecond = intervalBetweenSmsSecond
	smsRateData.IntervalGroupSecond = intervalGroupSecond
	smsRateData.SmsCount = countSms
	smsRateData.LastSmsDateTime = lastSmsDatetime.Format(external.DateFormat)
	smsRateData.AllowSendInSecond = allowSendInSecond
	smsRateData.AllowSendLocalDateTime = allowSendLocalDateTime.Format(external.DateFormat)

	var dataOptSignIn domain.SecPatientSignInOtp
	strsignInPatientLiveTimeSecond, _ := os.LookupEnv("SIGNINPATIENTLIVETIMESECOND")
	fmt.Sscan(strsignInPatientLiveTimeSecond, &signInPatientLiveTimeSecond)
	expiredDatetime := time.Now().Add(time.Second * time.Duration(signInPatientLiveTimeSecond)).Local()

	if allowSendSms {
		// check jarak sms
		if lastSmsDatetime.Before(time.Now().Local().Add(-time.Second * time.Duration(durationSameNumberSecond))) {
			otpSignIn := ""
			secPatientSignInOtpID := secPatientSignInOtp.SecPatientSignInOtpID
			if secPatientSignInOtp.SecPatientSignInOtpID == "" {
				otpSignIn = shared.RandomString(4)
				dataOptSignIn.SecPatientSignInOtpID = shared.GenerateUUID()
				dataOptSignIn.SecUserID = getExistUser.SecUserId
				dataOptSignIn.MobilePhone = params.MobilePhone
				dataOptSignIn.DeviceTypeCode = params.DeviceType
				dataOptSignIn.DeviceCode = params.DeviceCode
				dataOptSignIn.Otp = otpSignIn
				dataOptSignIn.ExpiredDatetime = expiredDatetime
				dataOptSignIn.RetryCounter = 0
				dataOptSignIn.IsActive = true
				dataOptSignIn.CreatedBy = getExistUser.SecUserId
				dataOptSignIn.UpdatedBy = getExistUser.SecUserId
				dataOptSignIn.CreatedOn = time.Now().Local()
				dataOptSignIn.UpdatedOn = time.Now().Local()
				savePatientOtp := u.AuthRepositoryWrite.SavePatientOtpSignIn(&dataOptSignIn)
				if savePatientOtp.Error != nil {
					errResponse := shared.ErrorResponse{
						Field:   "",
						Code:    "SECURITY_USERPWD_INVALID",
						Message: savePatientOtp.Error.Error(),
					}
					return shared.OutputV1{Error: savePatientOtp.Error, Result: errResponse, Code: http.StatusBadRequest}

				}
				secPatientSignInOtpID = dataOptSignIn.SecPatientSignInOtpID

			} else {
				dataOptSignIn.SecPatientSignInOtpID = secPatientSignInOtp.SecPatientSignInOtpID
				updatePatientOtp := u.AuthRepositoryWrite.UpdatePatientOtpSignIn(&dataOptSignIn)
				if updatePatientOtp.Error != nil {
					errResponse := shared.ErrorResponse{
						Field:   "",
						Code:    "SECURITY_USERPWD_INVALID",
						Message: updatePatientOtp.Error.Error(),
					}
					return shared.OutputV1{Error: updatePatientOtp.Error, Result: errResponse, Code: http.StatusBadRequest}

				}
				otpSignIn = secPatientSignInOtp.Otp
			}

			// saveSmsLog
			var dataSmsLog domain.SmsLog
			dataSmsLog.SmsStatus = false
			dataSmsLog.MobilePhone = params.MobilePhone
			dataSmsLog.SendingCount = 1
			dataSmsLog.SmsLogID = shared.GenerateUUID()
			dataSmsLog.SmsTypeCode = "userSignIn"
			dataSmsLog.SmsReffID = secPatientSignInOtpID
			dataSmsLog.SmsContent = fmt.Sprintf(external.SmsOtpSignIn, otpSignIn)
			dataSmsLog.CreatedBy = getExistUser.SecUserId
			dataSmsLog.UpdatedBy = getExistUser.SecUserId
			dataSmsLog.CreatedOn = time.Now().Local()
			dataSmsLog.UpdatedOn = time.Now().Local()
			savesmslogs := u.AuthRepositoryWrite.SaveSmsLogs(&dataSmsLog)
			if savesmslogs.Error != nil {
				errResponse := shared.ErrorResponse{
					Field:   "",
					Code:    "SECURITY_USERPWD_INVALID",
					Message: savesmslogs.Error.Error(),
				}
				return shared.OutputV1{Error: savesmslogs.Error, Errors: errResponse, Code: http.StatusBadRequest}

			}

			// send sms by acs service -> if simulation false
			var smsReq shared.AcsSmsRequest
			var smsSimulation bool
			strAcsPartnerID, _ := os.LookupEnv("ACS_PARTNER_ID")
			strAcsPartnerName, _ := os.LookupEnv("ACS_PARTNER_NAME")
			strAcsPassword, _ := os.LookupEnv("ACS_PASSWORD")
			strsmsSiluation, _ := os.LookupEnv("SMS_SIMULATION")
			fmt.Sscan(strsmsSiluation, &smsSimulation)

			smsReq.SmsBc.Request.Datetime = time.Now().Local().Format(external.SmsDateFormat)
			smsReq.SmsBc.Request.DestinationNumber = params.MobilePhone
			smsReq.SmsBc.Request.Rrn = shared.GenerateUUID()
			smsReq.SmsBc.Request.PartnerId = strAcsPartnerID
			smsReq.SmsBc.Request.PartnerName = strAcsPartnerName
			smsReq.SmsBc.Request.Password = strAcsPassword
			smsReq.SmsBc.Request.Message = fmt.Sprintf(external.SmsOtpSignIn, otpSignIn)

			// saveSmsLog request
			var dataSmsLogMessage domain.SmsLogMessage
			dataSmsLogMessage.SmsLogMessageID = shared.GenerateUUID()
			dataSmsLogMessage.SmsLogID = dataSmsLog.SmsLogID
			dataSmsLogMessage.MessageRrn = smsReq.SmsBc.Request.Rrn
			dataSmsLogMessage.ReqResTypeCode = "req"
			dataSmsLogMessage.XmlMessage = shared.CreateSmsXmlRequest(&smsReq)
			dataSmsLogMessage.CreatedBy = "APP-SMS"
			dataSmsLogMessage.UpdatedBy = "APP-SMS"
			dataSmsLogMessage.CreatedOn = time.Now().Local()
			dataSmsLogMessage.UpdatedOn = time.Now().Local()
			saveSmsLogMessReq := u.AuthRepositoryWrite.SaveSmsLogMessages(&dataSmsLogMessage)
			if saveSmsLogMessReq.Error != nil {
				errResponse := shared.ErrorResponse{
					Field:   "",
					Code:    "SECURITY_USERPWD_INVALID",
					Message: saveSmsLogMessReq.Error.Error(),
				}
				return shared.OutputV1{Error: saveSmsLogMessReq.Error, Errors: errResponse, Code: http.StatusBadRequest}

			}

			responeSms := u.External.SendSms(smsReq, smsSimulation)
			getResponseSms := responeSms.Result.(*external.AcsSmsAllResponse)
			if responeSms.Error != nil {
				errResponse := shared.ErrorResponse{
					Field:   "",
					Code:    "SECURITY_USERPWD_INVALID",
					Message: responeSms.Error.Error(),
				}
				return shared.OutputV1{Error: responeSms.Error, Errors: errResponse, Code: http.StatusBadRequest}

			}

			// saveSmsLogMessageRes
			dataSmsLogMessage.SmsLogMessageID = shared.GenerateUUID()
			dataSmsLogMessage.ReqResTypeCode = "resp"
			dataSmsLogMessage.ResponseCode = getResponseSms.Response.Response.Rc
			dataSmsLogMessage.ResponseMessage = getResponseSms.Response.Response.Rm
			dataSmsLogMessage.XmlMessage = shared.CreateSmsXmlResponse(&smsReq)
			saveSmsLogMessRes := u.AuthRepositoryWrite.SaveSmsLogMessages(&dataSmsLogMessage)
			if saveSmsLogMessRes.Error != nil {
				errResponse := shared.ErrorResponse{
					Field:   "",
					Code:    "SECURITY_USERPWD_INVALID",
					Message: saveSmsLogMessRes.Error.Error(),
				}
				return shared.OutputV1{Error: saveSmsLogMessRes.Error, Errors: errResponse, Code: http.StatusBadRequest}

			}

		} else {

			allowAfter20Sec := lastSmsDatetime.Add(time.Second * time.Duration(intervalBetweenSmsSecond))
			insecLeft := allowAfter20Sec.Sub(time.Now().Local())
			leftSecondOf60 := int(insecLeft.Seconds())
			smsRateData.AllowSendInSecond = leftSecondOf60
			smsRateData.AllowSendSms = false
			smsRateData.AllowSendLocalDateTime = allowAfter20Sec.Format(external.DateFormat)
			errorMessage := fmt.Sprintf(external.SmsValidationLimit, params.MobilePhone, lastSmsDatetime.Format(external.DateFormatStr), strdurationSameNumberSecond)
			smsRateData.ErrorMessage = errorMessage
			patientLoginPasswordResp.SecPatientSignInOtpId = ""
		}

		// send sms otp
		// insert to sms log & sms log messages
		patientLoginPasswordResp.SecPatientSignInOtpId = dataOptSignIn.SecPatientSignInOtpID
		patientLoginPasswordResp.ExpiredDatetime = expiredDatetime.Format(external.DateFormat)

	} else {

		patientLoginPasswordResp.SecPatientSignInOtpId = ""
		patientLoginPasswordResp.ExpiredDatetime = ""
	}

	patientLoginPasswordResp.SmsRateData = smsRateData

	if patientLoginPasswordResp.SecPatientSignInOtpId == "" {
		errResponse := shared.ErrorResponse{
			Field:   "",
			Code:    "SECURITY_USERPWD_INVALID",
			Message: patientLoginPasswordResp.SmsRateData.ErrorMessage,
		}

		return shared.OutputV1{
			Error:  errors.New(patientLoginPasswordResp.SmsRateData.ErrorMessage),
			Errors: errResponse,
			Result: patientLoginPasswordResp,
			Code:   http.StatusUnauthorized}
	}

	return shared.OutputV1{Result: patientLoginPasswordResp}
}

// SignUpByPhone function
func (u *AuthUsecaseImpl) SignInByPhoneOtp(params *dto.LoginByPhoneOtpRequest) shared.Output {

	signUp := u.AuthRepositoryWrite.LoginByPhoneOtp(params)
	if signUp.Error != nil {
		return shared.Output{Error: signUp.Error}
	}

	response := signUp.Result.(domain.SignUpByPhone)

	return shared.Output{Result: response}
}
