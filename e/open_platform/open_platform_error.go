package open_platform

import "fmt"

// PlatformError is a trivial implementation of error.
type PlatformError struct {
	message 	string
	code		int32
}

//NewPlatformError returns an error that formats as the given text.
func NewPlatformError(code int32, text string) error {
	return &PlatformError{
		code: code,
		message: text,
	}
}

//NewPlatformErrorForEnum returns an error that formats as the given enum.
func NewPlatformErrorForEnum(code PlatFormErrorEnum) error {
	return &PlatformError{
		code: int32(code),
		message: code.String(),
	}
}

func (e PlatformError) Error() string {
	return fmt.Sprintf("code : %d message : %s", e.code, e.message)
}