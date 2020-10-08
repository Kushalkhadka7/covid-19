package util

import "fmt"

// Error logs error to stderr.
func Error(message string, err error) error {
	newMessage := message + "%v"
	return fmt.Errorf(newMessage, err)
}
