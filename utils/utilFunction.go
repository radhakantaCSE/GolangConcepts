package utils

// Not exposed to outside of this package
func greetMessage() string {
	return "Hello learner"
}

// Exposed to outside of package
func SendMessage() string {
	return "Sending message..."
}
