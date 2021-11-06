package util

import "os"

func ValidateFile(fileToValidate string) {
	_, err := os.Stat(fileToValidate)
	ValidateErrorStatus(err)
}
