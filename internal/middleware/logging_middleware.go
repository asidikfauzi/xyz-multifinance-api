package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"time"
)

type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	rw.body.Write(b)
	return rw.ResponseWriter.Write(b)
}

func LoggingMiddleware() gin.HandlerFunc {
	file, err := os.OpenFile("assets/log/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	logger := log.New(file, "", log.LstdFlags)

	return func(c *gin.Context) {
		start := time.Now()

		var requestBody bytes.Buffer
		if c.Request.Body != nil {
			bodyBytes, _ := io.ReadAll(c.Request.Body)
			requestBody.Write(bodyBytes)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		rw := &responseWriter{ResponseWriter: c.Writer, body: bytes.NewBufferString("")}
		c.Writer = rw

		c.Next()

		duration := time.Since(start)

		logEntry := fmt.Sprintf("[%s] %s %s %d %v\nRequest: %s\nResponse: %s",
			c.Request.Method,
			c.Request.RequestURI,
			c.ClientIP(),
			c.Writer.Status(),
			duration,
			requestBody.String(),
			rw.body.String(),
		)

		logger.Println(logEntry)
	}
}
