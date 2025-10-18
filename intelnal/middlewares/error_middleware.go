package middlewares

import (
	"awesomeProject1/intelnal/apperror"
	"awesomeProject1/intelnal/dtos/dto_responses"

	"errors"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			var appErr *apperror.AppError
			if errors.As(err, &appErr) {

				apiError := &dto_responses.APIError{
					Code:    appErr.Code,
					Message: appErr.Error(),
				}
				dto_responses.ErrorResponse(c, appErr.Code, appErr.Error(), apiError)
				c.Abort()
				return
			}
			apiError := &dto_responses.APIError{
				Code:    500,
				Message: "An internal server error occurred",
			}
			dto_responses.ErrorResponse(c, 500, err.Error(), apiError)
		}
	}
}
