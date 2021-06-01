package misc

import "regexp"

var (
	// EmailValidation (const)
	EmailValidation = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

// IsEmailValid checks if the given email is valid using regex only.
//
// Identical to the function found on this page: https://golangcode.com/validate-an-email-address/.
func IsEmailValid(email string) bool {
	if len(email) < 3 && len(email) > 254 {
		return false
	}
	return EmailValidation.MatchString(email)
}
