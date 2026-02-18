package middleware

import (
	"bytes"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/oTeeLeko/product-service/utils"
)

func AccessLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var bodyBytes []byte
		if ctx.Request.Body != nil {
			bodyBytes, _ = io.ReadAll(ctx.Request.Body)
			ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			ctx.Set("requestBody", string(bodyBytes))
		}

		ctx.Next()
		utils.LogActivity(ctx)
	}
}

func ErrorLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if ctx.Writer.Status() >= 400 {
			errorMsg := "HTTP Error"
			if errors := ctx.Errors; len(errors) > 0 {
				errorMsg = errors.String()
			} else {
				errorMsg = fmt.Sprintf("HTTP %d Error", ctx.Writer.Status())
			}
			utils.LogError(ctx, errorMsg)
		}
	}
}
