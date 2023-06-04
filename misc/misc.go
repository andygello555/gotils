// Package misc contains some miscellaneous functions and constants that don't really fit elsewhere.
package misc

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"regexp"
)

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

type Ordered int

const (
	Less    Ordered = -1
	Equal   Ordered = 0
	Greater Ordered = 1
)

func (o Ordered) String() string {
	switch o {
	case Less:
		return "Less"
	case Equal:
		return "Equal"
	case Greater:
		return "Greater"
	default:
		return fmt.Sprintf("%d", o)
	}
}

// Compare compares the two constraints.Ordered values and returns:
//   - Less: a < b
//   - Equal: a == b
//   - Greater: a > b
func Compare[V constraints.Ordered](a, b V) Ordered {
	switch {
	case a < b:
		return Less
	case a == b:
		return Equal
	default:
		return Greater
	}
}
