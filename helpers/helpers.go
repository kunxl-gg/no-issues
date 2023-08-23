package helpers

// Check Error
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

// Check Error with Message
func CheckErrorWithMessage(err error, message string) {
	if err != nil {
		panic(message)
	}
}