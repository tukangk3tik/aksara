package security

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tukangk3tik/aksara/utils"
	"go.uber.org/zap"
)

func TraceMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := c.GetHeader("X-Trace-Id")
		if traceID == "" {
			traceID = uuid.New().String()
		}

		logWithTrace := utils.GlobalLogger.With(
			zap.String("trace_id", traceID),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
		)
		ctxWithLogger := utils.WithContext(c.Request.Context(), logWithTrace)

		// Replace Gin context with updated request context
		c.Request = c.Request.WithContext(ctxWithLogger)

		c.Set("trace_id", traceID)
		c.Writer.Header().Set("x-trace-Id", traceID)
		start := time.Now()

		c.Next()

		elapsed := time.Since(start)
		log := utils.FromContext(c.Request.Context())
		log.Info("request processed", zap.Float64("elapsed_ms", float64(elapsed.Nanoseconds())/1000000))
	}
}
