package errorhandler

import (
	"errors"
	"strings"

	"gitlab.com/tixia-backend/go-boilerplate/helpers/constants/rpcstd"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type stdRespError struct {
	StatusCode string `json:"status_code"`
	Message    string `json:"message"`
}

// ErrorHandler handles errors and returns an appropriate response.
func FiberErrHandler(ctx *fiber.Ctx, err error) error {

	// Check if the error is a validation error
	if validationErr, ok := err.(validator.ValidationErrors); ok {
		for _, val := range validationErr {
			switch strings.ToLower(val.Tag()) {
			case "required":
				return ctx.Status(fiber.StatusBadRequest).JSON(stdRespError{
					StatusCode: rpcstd.INVALID_ARGUMENT,
					Message:    strings.ToLower(val.Field()) + " is required",
				})
			}
		}
	}

	// check std erro
	var stdError *ErrorStd
	ok := errors.As(err, &stdError)
	if ok {
		return ctx.Status(stdError.HttpStatusCode).JSON(stdRespError{
			StatusCode: stdError.ErrorCode(),
			Message:    strings.ToLower(stdError.Error()),
		})
	}

	return nil
}
