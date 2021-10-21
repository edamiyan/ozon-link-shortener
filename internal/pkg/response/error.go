package response

type Error struct {
	Message string `json:"message"`
}

type ValidationError struct {
	Message string      `json:"message"`
	Errors  *ExtraError `json:"errors"`
}

type ExtraError struct {
	AdditionalProperties string `json:"additionalProperties"`
}

func NewError(message string) *Error {
	return &Error{Message: message}
}

func NewValidationError(additionalProperties string) *ValidationError {
	return &ValidationError{
		Message: "invalid data",
		Errors:  NewExtraError(additionalProperties),
	}
}

func NewExtraError(additionalProperties string) *ExtraError {
	return &ExtraError{AdditionalProperties: additionalProperties}
}
