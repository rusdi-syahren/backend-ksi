package shared

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"html/template"
	"io"
	"math/rand"
	"net/smtp"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/labstack/echo"
	"github.com/maddevsio/fcm"

	humanize "github.com/dustin/go-humanize"
)

// SendNotification func
func SendNotification(data map[string]string, deviceID string) error {
	var errorFcm error
	c := fcm.NewFCM("AAAATGi3SSI:APA91bFgez6UvJGirsVA-2qssWCedbbBK3_gfjJGmPH7ajn1cbv9h2O9-WRsMTYfEOqDhaR_qEko9Y2359tq2swaXSUT3OX5KW3Yqu_GMaiB3LQ9vro5jE60v4JpVBXkjZyvNDO5z4XJ")
	responseFcm, err := c.Send(fcm.Message{
		To:               deviceID,
		Data:             data,
		ContentAvailable: true,
		Priority:         fcm.PriorityHigh,
	})
	if err != nil {
		errorFcm = err
	}
	fmt.Printf("response from firebase notification : %+v\n", responseFcm)

	return errorFcm
}

// StringInSlice function for checking whether string in slice
// str string searched string
// list []string slice
func StringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

// HasPassword func
func HasPassword(password string) string {
	var sha = sha1.New()
	sha.Write([]byte(password))
	var encrypted = sha.Sum(nil)
	var encryptedString = fmt.Sprintf("%x", encrypted)

	return encryptedString
}

// CheckPassword func
func Check(password string, curentPassword string) bool {
	result := false
	// fmt.Printf(curentPassword)
	// fmt.Printf(HasPassword(password))
	if curentPassword == HasPassword(password) {
		result = true
	}
	return result
}

var emailAuth smtp.Auth

func SendMail(to []string, cc []string, subject, message string) error {
	emailHost := "smtp.mailgun.org"
	emailFrom := "postmaster@gasrem.id"
	emailPassword := "7ce77b97de80d0a0d649ae3ab448204d-a09d6718-3beb11ca"
	emailPort := 587

	body := "From: " + emailFrom + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Cc: " + strings.Join(cc, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	auth := smtp.PlainAuth("", emailFrom, emailPassword, emailHost)
	smtpAddr := fmt.Sprintf("%s:%d", emailHost, emailPort)

	err := smtp.SendMail(smtpAddr, auth, emailFrom, append(to, cc...), []byte(body))
	if err != nil {
		return err
	}

	return nil
}

// SendEmailSMTP func
func SendEmailSMTP(to []string, data interface{}, template string) (bool, error) {
	emailHost := "smtp.mailgun.org"
	emailFrom := "postmaster@gasrem.id"
	emailPassword := "7ce77b97de80d0a0d649ae3ab448204d-a09d6718-3beb11ca"
	emailPort := 465

	emailAuth = smtp.PlainAuth("", emailFrom, emailPassword, emailHost)

	emailBody, err := parseTemplate(template, data)
	if err != nil {
		return false, errors.New("unable to parse email template")
	}

	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + "Test Email" + "!\n"
	msg := []byte(subject + mime + "\n" + emailBody)
	addr := fmt.Sprintf("%s:%s", emailHost, emailPort)

	if err := smtp.SendMail(addr, emailAuth, emailFrom, to, msg); err != nil {
		return false, err
	}
	return true, nil
}

// parseTemplate func
func parseTemplate(templateFileName string, data interface{}) (string, error) {
	templatePath, err := filepath.Abs(fmt.Sprintf("gomail/email_templates/%s", templateFileName))
	if err != nil {
		return "", errors.New("invalid template name")
	}
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}
	body := buf.String()
	return body, nil
}

// ParseTemplatePdf func
func ParseTemplatePdf(templateFileName string, data interface{}) (string, error) {

	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	templatePath := dir + "/static/templates/" + templateFileName

	// t, err := template.ParseFiles(templatePath)
	// fmap := template.FuncMap{
	// 	"FormatRupiah": FormatRupiah,
	// }
	// t, err := template.New(templatePath).Funcs(fmap).ParseFiles(templatePath)
	t, err := template.ParseFiles(templatePath)

	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}
	body := buf.String()
	return body, nil
}

//FormatRupiah func
func FormatRupiah(amount float64) string {
	humanizeValue := humanize.CommafWithDigits(amount, 0)
	stringValue := strings.Replace(humanizeValue, ",", ".", -1)
	return "Rp " + stringValue
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%d/%d", day, month, year)
}

// RandomString func
func RandomString(n int) string {
	var letters = []rune("0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

// Upload func
func Upload(c echo.Context, fileName string, path string) (string, error) {
	randBytes := make([]byte, 16)
	rand.Read(randBytes)
	var err error

	file, err := c.FormFile(fileName)
	if err != nil {
		return "", err
	}

	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	filename := fmt.Sprintf("%s%s", hex.EncodeToString(randBytes), filepath.Ext(file.Filename))

	defaultPath := "/static/profiles/"
	if path != "" {
		defaultPath = path
	}
	fileLocation := filepath.Join(dir, defaultPath, filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer targetFile.Close()

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	if _, err := io.Copy(targetFile, src); err != nil {
		return "", err
	}

	return defaultPath + filename, nil
}

// UploadBase64 func
func UploadBase64(fileName string, content string) error {
	decode, err := base64.StdEncoding.DecodeString(content)
	file, err := os.Create(fileName)
	defer file.Close()
	_, err = file.Write(decode)
	return err
}

// StatusJualBeli func
func StatusJualBeli(status int) string {
	strStatus := ""
	if status == 5 {
		strStatus = "upnpaid"
	} else if status == 6 {
		strStatus = "paid"
	} else if status == 7 {
		strStatus = "process"
	} else if status == 8 {
		strStatus = "ship"
	} else if status == 9 {
		strStatus = "delivered"
	}
	return strStatus
}

// StatusInvoice func
func StatusInvoice(status int) string {
	strStatus := ""
	if status == 1 {
		strStatus = "unpaid"
	} else if status == 2 {
		strStatus = "paid"
	} else if status == 0 {
		strStatus = "pending"

	}
	return strStatus
}

// PyamentType func
func PyamentType(paymentType int) string {
	strpaymentType := ""
	if paymentType == 1 {
		strpaymentType = "gopay"
	} else if paymentType == 2 {
		strpaymentType = "CC"
	}
	return strpaymentType
}

// ItemExists func
func ItemExists(arrayType interface{}, item interface{}) bool {
	arr := reflect.ValueOf(arrayType)

	if arr.Kind() != reflect.Array {
		panic("Invalid data-type")
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}

	return false
}
