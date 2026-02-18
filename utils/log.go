package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func LogActivity(ctx *gin.Context) {
	location, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		location = time.UTC
	}

	now := time.Now().In(location)
	year := now.Format("2006")
	month := now.Format("Jan")
	fullDate := now.Format("2006-01-02")

	logDir := filepath.Join("./logs", year, month, "access")
	logFile := filepath.Join(logDir, fullDate+".txt")

	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		fmt.Println("Error creating access log directory:", err)
		return
	}

	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening access log file:", err)
		return
	}
	defer file.Close()

	method := ctx.Request.Method
	path := ctx.Request.URL.Path
	ip := ctx.ClientIP()
	status := ctx.Writer.Status()

	var params string

	if len(ctx.Request.URL.Query()) > 0 {
		params += "Query: " + ctx.Request.URL.Query().Encode() + " "
	}
	if method == "POST" || method == "PUT" || method == "PATCH" {
		if body, exists := ctx.Get("requestBody"); exists {
			if bodyStr, ok := body.(string); ok && bodyStr != "" {
				params += fmt.Sprintf("Body: %s ", bodyStr)
			}
		}
	}

	logMessage := fmt.Sprintf("[%s] %s - Status: %d Method: %s Path: %s",
		now.Format("2006-01-02 15:04:05"), ip, status, method, path)

	if len(params) > 0 {
		logMessage += fmt.Sprintf(" %s", params)
	}

	logger := log.New(file, "", 0)
	logger.Println(logMessage)
}

type LogInterfacer interface {
	LogInfo() string
}

func LogError(ctx *gin.Context, errorMsg string, errorDetails ...interface{}) {
	location, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		location = time.UTC
	}

	now := time.Now().In(location)
	year := now.Format("2006")
	month := now.Format("Jan")
	fullDate := now.Format("2006-01-02")

	logDir := filepath.Join("./logs", year, month, "error")
	logFile := filepath.Join(logDir, fullDate+".txt")

	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		fmt.Println("Error creating error log directory:", err)
		return
	}

	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening error log file:", err)
		return
	}
	defer file.Close()

	method := ctx.Request.Method
	path := ctx.Request.URL.Path
	ip := ctx.ClientIP()
	status := ctx.Writer.Status()

	requestBody := ""
	if body, exists := ctx.Get("requestBody"); exists {
		if bodyStr, ok := body.(string); ok && bodyStr != "" {
			requestBody = fmt.Sprintf(" - Body: %s", bodyStr)
		}
	}

	params := ""
	if len(ctx.Request.URL.Query()) > 0 {
		params = fmt.Sprintf(" - Query: %s", ctx.Request.URL.Query().Encode())
	}

	logMessage := fmt.Sprintf("[%s] ERROR - %s - Status: %d Method: %s Path: %s - Error: %s%s%s",
		now.Format("2006-01-02 15:04:05"), ip, status, method, path, errorMsg, params, requestBody)

	if len(errorDetails) > 0 {
		detailsStr := ""
		for _, detail := range errorDetails {
			if loggable, ok := detail.(LogInterfacer); ok {
				detailsStr += fmt.Sprintf("%s, ", loggable.LogInfo())
			} else {
				detailsStr += fmt.Sprintf("%+v, ", detail)
			}
		}
		if len(detailsStr) > 2 {
			detailsStr = detailsStr[:len(detailsStr)-2]
		}
		logMessage += fmt.Sprintf(" - Details: [%s]", detailsStr)
	}

	logger := log.New(file, "", 0)
	logger.Println(logMessage)
}
