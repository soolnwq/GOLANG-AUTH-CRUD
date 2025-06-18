package errs

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AppError struct {
	Status  int                   `json:"status"`
	Message string                `json:"message"`
	Errors  []ValidationErrorItem `json:"errors,omitempty"`
}

type ValidationErrorItem struct {
	Field string `json:"field"`
	Error string `json:"message"`
}

func (e *AppError) Error() string { return e.Message }

func NewAppError(message string, status int) *AppError {
	return &AppError{Status: status, Message: message}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{Status: 404, Message: message}
}

func NewUnexpectedError(message string) *AppError {
	return &AppError{Status: 500, Message: message}
}

func NewValidationErrorItem(errors []ValidationErrorItem) *AppError {
	return &AppError{Status: 422, Message: "validation error", Errors: errors}
}

func NewBadRequestError(message string) *AppError {
	return &AppError{Status: 400, Message: message}
}

func HandleError(c *fiber.Ctx, err error) error {
	switch e := err.(type) {
	case *AppError:
		return c.Status(e.Status).JSON(e)
	default:
		return c.Status(500).JSON(NewUnexpectedError(err.Error()))
	}
}

func validationMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fe.Field() + " is required"
	case "email":
		return "invalid email format"
	case "min":
		return fe.Field() + " must be at least " + fe.Param()
	case "max":
		return fe.Field() + " must be at most " + fe.Param()
	default:
		return fe.Error()
	}
}

func ParseValidationErrors(err error) error {
	var valErrs validator.ValidationErrors
	if errors.As(err, &valErrs) {
		details := make([]ValidationErrorItem, 0)
		for _, fe := range valErrs {
			details = append(details, ValidationErrorItem{
				Field: fe.Field(),
				Error: validationMessage(fe),
			})
		}
		return NewValidationErrorItem(details)
	}
	return NewUnexpectedError(err.Error())
}
