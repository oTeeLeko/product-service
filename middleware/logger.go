package middleware

import (
	"bytes"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/oTeeLeko/product-service/utils"
)

// AccessLogger สำหรับ log การเข้าถึง
func AccessLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// เก็บ request body ไว้สำหรับ error logging
		var bodyBytes []byte
		if ctx.Request.Body != nil {
			bodyBytes, _ = io.ReadAll(ctx.Request.Body)
			ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			ctx.Set("requestBody", string(bodyBytes))
		}

		ctx.Next()
		// Log ทุกครั้งไม่ว่า LogLevel จะเป็นอะไร
		utils.LogActivity(ctx)
	}
}

// ErrorLogger สำหรับ log errors
func ErrorLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		// ถ้า status >= 400 ให้ log error
		if ctx.Writer.Status() >= 400 {
			errorMsg := "HTTP Error"
			// ลองดึง error จาก gin context ก่อน
			if errors := ctx.Errors; len(errors) > 0 {
				errorMsg = errors.String()
			} else {
				// ถ้าไม่มี ให้ใช้ "Unknown Error" แทน
				errorMsg = fmt.Sprintf("HTTP %d Error", ctx.Writer.Status())
			}
			utils.LogError(ctx, errorMsg)
		}
	}
}
