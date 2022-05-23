package shared

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidator(t *testing.T) {
	var err error

	validInputString := "Ned Stark"
	invalidInputString := ""

	err = ValidateEmptyInput(validInputString)
	assert.NoError(t, err, err)

	err = ValidateEmptyInput(invalidInputString)
	assert.Error(t, err, err)

	validNumberInput := "081728828999"
	invalidNumberInput := "981ahb"
	emptyNumberInput := ""

	err = ValidateNumberOnlyInput(validNumberInput)
	assert.NoError(t, err, err)

	err = ValidateNumberOnlyInput(invalidNumberInput)
	assert.Error(t, err, err)

	err = ValidateNumberOnlyInputAllowEmpty(validNumberInput)
	assert.NoError(t, err, err)

	err = ValidateNumberOnlyInputAllowEmpty(emptyNumberInput)
	assert.NoError(t, err, err)

	err = ValidateNumberOnlyInputAllowEmpty(invalidNumberInput)
	assert.Error(t, err, err)

	validPhoneNumber := "02129872828"
	validPhoneNumberWithExt := "02129872828 12"
	invalidPhoneNumber := "0218293938111111928272222222"

	err = ValidatePhoneNumberMaxInput(validPhoneNumber)
	assert.NoError(t, err, err)

	err = ValidatePhoneNumberMaxInput(validPhoneNumberWithExt)
	assert.NoError(t, err, err)

	err = ValidatePhoneNumberMaxInput(invalidPhoneNumber)
	assert.Error(t, err, err)

	validMobileNumber := "081283187099"
	invalidMobileNumber := "0218293938111111928272"

	err = ValidateMobileNumberMaxInput(validMobileNumber)
	assert.NoError(t, err, err)

	err = ValidateMobileNumberMaxInput(invalidMobileNumber)
	assert.Error(t, err, err)

	validAlphabeticalInput := "Wuriyanto"
	invalidAlphabeticalInput := "981ahb"

	err = ValidateAlphabeticalOnlyInput(validAlphabeticalInput)
	assert.NoError(t, err, err)

	err = ValidateAlphabeticalOnlyInput(invalidAlphabeticalInput)
	assert.Error(t, err, err)

	validAlphabeticalEmptyInput := "Wuriyanto"

	err = ValidateAlphabeticalOnlyInputAllowEmpty(validAlphabeticalInput)
	assert.NoError(t, err, err)

	err = ValidateAlphabeticalOnlyInputAllowEmpty(validAlphabeticalEmptyInput)
	assert.NoError(t, err, err)

	err = ValidateAlphabeticalOnlyInputAllowEmpty(invalidAlphabeticalInput)
	assert.Error(t, err, err)

	validEmail := "wuriyanto.musobar@gmail.com"
	invalidEmail := "wuriyanto.musobar@gmailcom"

	err = ValidateEmail(validEmail)
	assert.NoError(t, err, err)

	err = ValidateEmail(invalidEmail)
	assert.Error(t, err, err)

	validGenderInput := "M"
	invalidGenderInput := "mf"

	err = ValidateGenderInput(validGenderInput)
	assert.NoError(t, err, err)

	err = ValidateGenderInput(invalidGenderInput)
	assert.Error(t, err, err)

	validAlphaNumericInput := `wuriyanto'007 / & , bla. - () :\`
	invalidAlphaNumericInput := "wuriyanto007*"

	err = ValidateAlphaNumericInput(validAlphaNumericInput)
	assert.NoError(t, err, err)

	err = ValidateAlphaNumericInput(invalidAlphaNumericInput)
	assert.Error(t, err, err)

	validAlphaNumericInputAllowEmpty := `wuriyanto'007 / & , bla. - () :\`
	emptyAlphaNumericInputAllowEmpty := ""
	invalidAlphaNumericInputAllowEmpty := "aahaj 89*="

	err = ValidateAlphaNumericInputAllowEmpty(validAlphaNumericInputAllowEmpty)
	assert.NoError(t, err, err)

	err = ValidateAlphaNumericInputAllowEmpty(emptyAlphaNumericInputAllowEmpty)
	assert.NoError(t, err, err)

	err = ValidateAlphaNumericInputAllowEmpty(invalidAlphaNumericInputAllowEmpty)
	assert.Error(t, err, err)

	b1 := ValidateBooleanInput("true")
	assert.Equal(t, true, b1)

	b2 := ValidateBooleanInput("false")
	assert.Equal(t, false, b2)

	b3 := ValidateBooleanInput("")
	assert.Equal(t, false, b3)

	validPass := IsValidPass("Blink182!")
	assert.True(t, validPass)

	validPassErr := IsValidPass("blink182!")
	assert.False(t, validPassErr)

	boolTrue := ValidateAlphanumericWithSpace("oke sip1", false)
	assert.True(t, boolTrue)

	boolFalse := ValidateAlphanumericWithSpace("okesip1", true)
	assert.False(t, boolFalse)

	boolTrue = ValidateAlphanumeric("okesip", false)
	assert.True(t, boolTrue)

	boolFalse = ValidateAlphanumeric("1FgH^*", false)
	assert.False(t, boolFalse)

	boolTrue = ValidateAlphabet("huFtBanGeT")
	assert.True(t, boolTrue)

	boolFalse = ValidateAlphabet("1FgH^*")
	assert.False(t, boolFalse)

	boolFalse = ValidateAlphabetWithSpace("huFtBanGeT*")
	assert.False(t, boolFalse)

	boolTrue = ValidateAlphabetWithSpace("huFt BanGeT")
	assert.True(t, boolTrue)

	boolFalse = ValidateNumeric("1.0.1")
	assert.False(t, boolFalse)

	boolTrue = ValidateNumeric("0123456789")
	assert.True(t, boolTrue)

	boolFalse = ValidateLatinOnly("스칼 k4nj1 k0r34")
	assert.False(t, boolFalse)

	boolTrue = ValidateLatinOnly("oke 123 ~!@#")
	assert.True(t, boolTrue)
}

func TestTextInputValidator(t *testing.T) {

	validInputString := "Ned Stark"
	invalidInputString := ""

	err := ValidateEmptyInput(validInputString)
	assert.NoError(t, err, err)

	err = ValidateEmptyInput(invalidInputString)
	assert.Error(t, err, err)

	shortInputString := "Game of Thrones"
	tooLongInputString := `Let's say we require an item from our drop down list, but instead we get a value fabricated by hackers
	Let's say we require an item from our drop down list, but instead we get a value fabricated by hackers
	Let's say we require an item from our drop down list, but instead we get a value fabricated by hackers
	Let's say we require an item from our drop down list, but instead we get a value fabricated by hackers
	Let's say we require an item from our drop down list, but instead we get a value fabricated by hackers
	Let's say we require an item from our drop down list, but instead we get a value fabricated by hackers
	Let's say we require an item from our drop down list, but instead we get a value fabricated by hackers`

	err = ValidateMaxInput(shortInputString, 250)
	assert.NoError(t, err, err)

	err = ValidateMaxInput(tooLongInputString, 250)
	assert.Error(t, err, err)

	validNumberInput := "081728828999"
	invalidNumberInput := "981ahb"

	err = ValidateNumberOnlyInput(validNumberInput)
	assert.NoError(t, err, err)

	err = ValidateNumberOnlyInput(invalidNumberInput)
	assert.Error(t, err, err)

	validPhoneNumber := "02129872828"
	validPhoneNumberWithExt := "02129872828 12"
	invalidPhoneNumber := "02182939381111119282722222222"

	err = ValidatePhoneNumberMaxInput(validPhoneNumber)
	assert.NoError(t, err, err)

	err = ValidatePhoneNumberMaxInput(validPhoneNumberWithExt)
	assert.NoError(t, err, err)

	err = ValidatePhoneNumberMaxInput(invalidPhoneNumber)
	assert.Error(t, err, err)

	validMobileNumber := "081283187099"
	invalidMobileNumber := "0218293938111111928272"

	err = ValidateMobileNumberMaxInput(validMobileNumber)
	assert.NoError(t, err, err)

	err = ValidateMobileNumberMaxInput(invalidMobileNumber)
	assert.Error(t, err, err)

	validAlphabeticalInput := "Wuriyanto"
	invalidAlphabeticalInput := "981ahb"

	err = ValidateAlphabeticalOnlyInput(validAlphabeticalInput)
	assert.NoError(t, err, err)

	err = ValidateAlphabeticalOnlyInput(invalidAlphabeticalInput)
	assert.Error(t, err, err)

	validGenderInput := "M"
	invalidGenderInput := "mf"

	err = ValidateGenderInput(validGenderInput)
	assert.NoError(t, err, err)

	err = ValidateGenderInput(invalidGenderInput)
	assert.Error(t, err, err)

	validAlphaNumericInput := `wuriyanto'007 / & , bla. - () :\`
	invalidAlphaNumericInput := "aahaj 893*"

	err = ValidateAlphaNumericInput(validAlphaNumericInput)
	assert.NoError(t, err, err)

	err = ValidateAlphaNumericInput(invalidAlphaNumericInput)
	assert.Error(t, err, err)

	validAlphaNumericInputAllowEmpty := `wuriyanto'007 / & , bla. - () :\`
	emptyAlphaNumericInputAllowEmpty := ""
	invalidAlphaNumericInputAllowEmpty := "aahaj 893*"

	err = ValidateAlphaNumericInputAllowEmpty(validAlphaNumericInputAllowEmpty)
	assert.NoError(t, err, err)

	err = ValidateAlphaNumericInputAllowEmpty(emptyAlphaNumericInputAllowEmpty)
	assert.NoError(t, err, err)

	err = ValidateAlphaNumericInputAllowEmpty(invalidAlphaNumericInputAllowEmpty)
	assert.Error(t, err, err)
}
