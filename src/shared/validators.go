package shared

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

var (
	// alphanumericComaPointDashSpaceRegexp regex for validating alphanumeric, coma, point and space
	alphanumericComaPointDashSpaceRegexp = regexp.MustCompile("^ *([a-zA-Z0-9,.-] ?)+ *$")
	// nonEnglishCharacter regex for validating non  English character
	nonEnglishCharacter = regexp.MustCompile("[^\x00-\x7F]+")
)

// ValidateID function
func ValidateID(input string) error {
	if !bson.IsObjectIdHex(input) {
		return errors.New("invalid object id format")
	}

	return nil
}

// ValidateAlphanumericWithComaDashPointSpace function for validating string to alpha numeric
func ValidateAlphanumericWithComaDashPointSpace(str string) error {
	if alphanumericComaPointDashSpaceRegexp.MatchString(str) {
		return nil
	}
	return errors.New("alphanumeric, comma, dash, point and one space between word is allowed")
}

// ValidateRangeName function
func ValidateRangeName(input string) error {
	if len(input) > 30 {
		return errors.New("name length cannot greater than 30 character")
	}

	return nil
}

// ValidateAlphaNumericInput function for validating alpha numeric
func ValidateAlphaNumericInput(input string) error {
	check := regexp.MustCompile(`^[A-Za-z0-9.,\\;\\_\\&\\/\-\(\\)\\:\\'\\\ ]+$`).MatchString
	if !check(input) {
		err := errors.New("this field value is in bad format")
		return err
	}
	return nil
}

// ValidateAlphaNumericInputAllowEmpty function for validating alpha numeric and allow empty
func ValidateAlphaNumericInputAllowEmpty(input string) error {
	check := regexp.MustCompile(`^[A-Za-z0-9.,\\;\\_\\&\\/\-\(\\)\\:\\'\\\ ]*$`).MatchString
	if !check(input) {
		err := errors.New("this field value is in bad format")
		return err
	}
	return nil
}

// ValidateGenderInput function for validating gender
func ValidateGenderInput(input string) error {
	check := regexp.MustCompile("^[MF]+$").MatchString
	if len(input) > 1 {
		err := errors.New("this field value should be M or F")
		return err
	}
	if !check(input) {
		err := errors.New("this field value is in bad format")
		return err
	}
	return nil
}

// ValidateBooleanInput function for validating boolean input
func ValidateBooleanInput(input string) bool {
	b, e := strconv.ParseBool(input)
	if e != nil {
		return false
	}
	return b
}

// ValidateEmptyInput function for validating empty input
func ValidateEmptyInput(input string) error {
	if len(input) == 0 {
		err := errors.New("this field cannot be empty")
		return err
	}
	return nil
}

// ValidateNonEnglishCharacter function for validating string to non english
func ValidateNonEnglishCharacter(str string) error {
	if nonEnglishCharacter.MatchString(str) {
		return errors.New("character latin only is allowed")
	}

	return nil
}

// StringArrayReplace function for replacing whether string in array
// str string searched string
// list []string array
func StringArrayReplace(str string, listFind, listReplace []string) string {
	for i, v := range listFind {
		if strings.Contains(str, v) {
			str = strings.Replace(str, v, listReplace[i], -1)
		}
	}
	return str
}

// ValidateMaxInput function for validating maximum input
func ValidateMaxInput(input string, limit int) error {
	if len(input) > limit {
		err := errors.New(" value is too long")
		return err
	}

	return nil
}

// ValidatePhoneNumberMaxInput function for validating phone number
func ValidatePhoneNumberMaxInput(input string) error {
	check := regexp.MustCompile("^[0-9 ]*$").MatchString
	if len(input) > 27 {
		err := errors.New("this field value is too long")
		return err
	}
	if !check(input) {
		err := errors.New("this field value should be Number")
		return err
	}
	return nil
}

// ValidateExtPhoneNumberMaxInput function for validating phone number extension
func ValidateExtPhoneNumberMaxInput(input string) error {
	check := regexp.MustCompile("^[0-9 ]*$").MatchString
	if len(input) > 10 {
		err := errors.New("this field value is too long")
		return err
	}
	if !check(input) {
		err := errors.New("this field value should be Number")
		return err
	}
	return nil
}

// ValidateMobileNumberMaxInput function for validating mobile phone number
func ValidateMobileNumberMaxInput(input string) error {
	check := regexp.MustCompile("^[0-9]+$").MatchString
	if len(input) > 13 {
		err := errors.New("this field value is too long")
		return err
	}
	if !check(input) {
		err := errors.New("this field value should be Number")
		return err
	}
	check = regexp.MustCompile(`^[0]+8[1235789]+[0-9]`).MatchString
	if !check(input) {
		err := errors.New("this input mobile number is in bad format")
		return err
	}
	return nil
}

// ValidateNumberOnlyInput function for validating number only
func ValidateNumberOnlyInput(input string) error {
	check := regexp.MustCompile("^[0-9]+$").MatchString
	if !check(input) {
		err := errors.New("this field value should be number")
		return err
	}
	return nil
}

// ValidateNumberOnlyInputAllowEmpty function for validating number only and allow empty
func ValidateNumberOnlyInputAllowEmpty(input string) error {
	check := regexp.MustCompile("^[0-9]*$").MatchString
	if !check(input) {
		err := errors.New("this field value should be number")
		return err
	}
	return nil
}

// ValidateAlphabeticalOnlyInput function for validating alphabet only
func ValidateAlphabeticalOnlyInput(input string) error {
	check := regexp.MustCompile("^[A-Za-z ]+$").MatchString
	if !check(input) {
		err := errors.New("this field value should be alphabetical")
		return err
	}
	return nil
}

// ValidateAlphabeticalOnlyInputAllowEmpty function for validating
// alphabet only and allow empty
func ValidateAlphabeticalOnlyInputAllowEmpty(input string) error {
	check := regexp.MustCompile("^[A-Za-z]*$").MatchString
	if !check(input) {
		err := errors.New("this field value should be alphabetical")
		return err
	}
	return nil
}

// ValidateEmail function for validating email
func ValidateEmail(input string) error {
	check := regexp.MustCompile(`^[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,4}$`).MatchString
	if !check(input) {
		err := errors.New("email address is invalid")
		return err
	}
	return nil
}

// IsValidPass func for check valid password
func IsValidPass(pass string) bool {
	if len(pass) == 0 || len(pass) < 6 {
		return false
	}

	var uppercase, lowercase, num int
	for _, r := range pass {
		if r >= 65 && r <= 90 { //code ascii for [A-Z]
			uppercase = +1
		} else if r >= 97 && r <= 122 { //code ascii for [a-z]
			lowercase = +1
		} else if r >= 48 && r <= 57 { //code ascii for [0-9]
			num = +1
		}
	}
	return uppercase >= 1 && lowercase >= 1 && num >= 1
}

// ValidateNumeric func for check valid numeric
func ValidateNumeric(str string) bool {
	var num, symbol int
	for _, r := range str {
		if r >= 48 && r <= 57 { //code ascii for [0-9]
			num = +1
		} else {
			symbol = +1
		}
	}

	if symbol > 0 {
		return false
	}

	return num >= 1
}

// ValidateAlphabet func for check alphabet
func ValidateAlphabet(str string) bool {
	var uppercase, lowercase, symbol int
	for _, r := range str {
		if r >= 65 && r <= 90 { //code ascii for [A-Z]
			uppercase = +1
		} else if r >= 97 && r <= 122 { //code ascii for [a-z]
			lowercase = +1
		} else { //except alphabet
			symbol = +1
		}
	}

	if symbol > 0 {
		return false
	}
	return uppercase >= 1 || lowercase >= 1
}

// ValidateAlphabetWithSpace func for check alphabet with space
func ValidateAlphabetWithSpace(str string) bool {
	var uppercase, lowercase, space, symbol int
	for _, r := range str {
		if r >= 65 && r <= 90 { //code ascii for [A-Z]
			uppercase = +1
		} else if r >= 97 && r <= 122 { //code ascii for [a-z]
			lowercase = +1
		} else if r == 32 { //code ascii for space
			space = +1
		} else { //except alphabet
			symbol = +1
		}
	}

	if symbol > 0 {
		return false
	}
	return uppercase >= 1 || lowercase >= 1 || space >= 1
}

// ValidateAlphanumeric func for check valid alphanumeric
func ValidateAlphanumeric(str string, must bool) bool {
	var uppercase, lowercase, num, symbol int
	for _, r := range str {
		if r >= 65 && r <= 90 { //code ascii for [A-Z]
			uppercase = +1
		} else if r >= 97 && r <= 122 { //code ascii for [a-z]
			lowercase = +1
		} else if r >= 48 && r <= 57 { //code ascii for [0-9]
			num = +1
		} else {
			symbol = +1
		}
	}

	if symbol > 0 {
		return false
	}

	if must { //must alphanumeric
		return uppercase >= 1 && lowercase >= 1 && num >= 1
	}

	return uppercase >= 1 || lowercase >= 1 || num >= 1
}

// ValidateAlphanumericWithSpace func for validating string to alpha numeric with space
func ValidateAlphanumericWithSpace(str string, must bool) bool {
	var uppercase, lowercase, num, space, symbol int
	for _, r := range str {
		if r >= 65 && r <= 90 { //code ascii for [A-Z]
			uppercase = +1
		} else if r >= 97 && r <= 122 { //code ascii for [a-z]
			lowercase = +1
		} else if r >= 48 && r <= 57 { //code ascii for [0-9]
			num = +1
		} else if r == 32 { //code ascii for space
			space = +1
		} else {
			symbol = +1
		}
	}

	if symbol > 0 {
		return false
	}

	if must { //must alphanumeric
		return uppercase >= 1 && lowercase >= 1 && num >= 1 && space >= 1
	}

	return (uppercase >= 1 || lowercase >= 1 || num >= 1) || space >= 1
}

// ValidateLatinOnly func for check valid latin only
func ValidateLatinOnly(str string) bool {
	var uppercase, lowercase, num, allowed, symbol int
	for _, r := range str {
		if r >= 65 && r <= 90 { //code ascii for [A-Z]
			uppercase = +1
		} else if r >= 97 && r <= 122 { //code ascii for [a-z]
			lowercase = +1
		} else if r >= 48 && r <= 57 { //code ascii for [0-9]
			num = +1
		} else if r >= 32 && r <= 47 || r >= 58 && r <= 64 || r >= 91 && r <= 96 || r >= 123 && r <= 126 {
			allowed = +1 //code ascii for [space, coma, ., !, ", #, $, %, &, ', (, ), *, +, -, /, :, ;, <, =, >, ?, @, [, \, ], ^, _, `, {, |, }, ~]
		} else {
			symbol = +1
		}
	}

	if symbol > 0 {
		return false
	}

	return uppercase >= 1 || lowercase >= 1 || num >= 1 || allowed >= 0
}

// ValidateHTML function for validating HTML
func ValidateHTML(src string) string {
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	return re.ReplaceAllString(src, "")
}
