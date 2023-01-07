package utils

import (
	"github.com/gin-gonic/gin"
	"time"
)

type errorMessage struct {
	Message string `json:"message"`
}

func DoWithTries(fn func() error, attemtps int, delay time.Duration) (err error) {
	for attemtps > 0 {
		if err = fn(); err != nil {
			time.Sleep(delay)
			attemtps--

			continue
		}

		return nil
	}
	return
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, errorMessage{message})
}
