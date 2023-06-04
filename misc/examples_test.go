package misc

import "fmt"

// Check if the email:
//
//	test@example.com
//
// Is valid and then check whether some normal text is valid.
func ExampleIsEmailValid() {
	// Valid example
	e := "test@example.com"
	if IsEmailValid(e) {
		fmt.Println(e + " is a valid email")
	}

	// Invalid example
	if !IsEmailValid("just text") {
		fmt.Println("not a valid email")
	}
	// Output:
	// test@example.com is a valid email
	// not a valid email
}

// Compare some constraints.Ordered values!
func ExampleCompare() {
	fmt.Printf("Compare(1, 2) = %s\n", Compare(1, 2))
	fmt.Printf("Compare(1.23, 1.23) = %s\n", Compare(1.23, 1.23))
	fmt.Printf("Compare(\"world\", \"hello\") = %s\n", Compare("world", "hello"))
	// Output:
	// Compare(1, 2) = Less
	// Compare(1.23, 1.23) = Equal
	// Compare("world", "hello") = Greater
}
