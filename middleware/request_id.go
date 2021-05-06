package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/pborman/uuid"
)

func RequestID(allow bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestID string

		if allow {
			requestID = c.Request.Header.Get("Request-Id")
		}

		if requestID == "" {
			requestID = uuid.New()
		}

		c.Writer.Header().Set("Request-Id", requestID)
		c.Next()
	}
}
