package response

type ExtraError struct {
	AdditionalProperties string `json:"additionalProperties"`
}

type ValidationError struct {
	Message string      `json:"message"`
	Errors  *ExtraError `json:"errors"`
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
