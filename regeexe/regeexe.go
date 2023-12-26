package regeexe

import (
	"fmt"
	"reflect"
	"regexp"
)

/*
Split in the regeexe package splits a block of text by anonymizing any data that matches the regular expression.
Parameters:
	  - block: The block of text to split.
	  - anonym: The regular expression to split the block by. ([^:]+) > will split range to find : or ([^ ]+) > will split range to find " "
	  - expectResponse: The expected response object to unmarshal.
Returns:
	  - response: The block of text with the anonymized data replaced with the empty string.
	  - err: Any error that occurred during processing.

“
func main() {
	responseObj := ExampleResponse{}
	err := regeexe.Split("email:password:codeunix", "([^:]+):([^:]+):([^:]+)", &responseObj)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Email: %s\nPassword: %s\nUniqCodes:", responseObj.Email, responseObj.Password, responseObj.UniqCodes)

	responseObj1 := ExampleResponse{}
	err1 := regeexe.Split("email password codeunix", "([^ ]+) ([^ ]+)", &responseObj)
	if err1 != nil {
		log.Fatal(err)
	}

	fmt.Printf("Email: %s\nPassword: %s\n", responseObj1.Email, responseObj1.Password)
}
“
*/

// Split splits a block of text by anonymizing any data that matches the regular expression.
// The anonymized data is replaced with the empty string. If an expected response is provided,
// it is unmarshaled from the matched substrings.
func Split(block string, anonym string, expectResponse interface{}) error {
	// re is a regular expression that matches the anonym string.
	re := regexp.MustCompile(anonym)
	// response is the block of text with the anonymized data replaced with the empty string.
	matches := re.FindStringSubmatch(block)

	// Extract the matched substrings, excluding the full match at index 0
	matchedSubstrings := matches[1:]

	convertArrayToObject(matchedSubstrings, expectResponse)

	// Return the response and any error.
	return nil
}

// convertArrayToObject converts an array to an object.
func convertArrayToObject(array []string, exampleResponseInstance interface{}) error {
	// Check if the array is large enough to fill the struct fields
	if len(array) != reflect.ValueOf(exampleResponseInstance).Elem().NumField() {
		// Return an error if the array doesn't have the correct number of elements
		return fmt.Errorf("incorrect number of elements in the array to create the object")
	}

	// Iterate over struct fields and assign values from the array
	for i := 0; i < len(array); i++ {
		value := array[i]

		// Set the struct field value using reflection
		reflect.ValueOf(exampleResponseInstance).Elem().Field(i).SetString(value)
	}

	return nil
}
