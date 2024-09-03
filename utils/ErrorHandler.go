package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ErrorLogger.Printf("Panic occurred: %v", err)
				ErrorLogger.Printf("Stack trace: %s", debug.Stack())
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
				c.Abort()
			}
		}()

		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			switch {
			case errors.Is(err, ErrEmailAlreadyExists):
				c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			case errors.Is(err, ErrInvalidCredentials):
				c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			default:
				ErrorLogger.Printf("Unexpected error: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			}
			c.Abort()
		}
	}
}
