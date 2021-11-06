package util

func ValidateErrorStatus(err error) {
	if err != nil {
		panic(err)
	}
}
