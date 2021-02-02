package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type appError struct {
	error
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e appError) Error() string {
	return fmt.Sprintf("异常code:%d,错误原因：%s", e.Code, e.Message)
}

func JSONAppErrorReporter() gin.HandlerFunc {
	return jsonAppErrorReporterT(gin.ErrorTypeAny)
}

func jsonAppErrorReporterT(errType gin.ErrorType) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		detectedErrors := c.Errors.ByType(errType)
		if len(detectedErrors) > 0 {
			err := detectedErrors[0].Err
			var parsedError *appError
			switch err.(type) {
			case *appError:
				parsedError = err.(*appError)
			default:
				parsedError = &appError{
					Code:    http.StatusInternalServerError,
					Message: "Internal Server Error",
				}
			}
			c.IndentedJSON(parsedError.Code, parsedError)
			c.Abort()
			return
		}

	}
}
