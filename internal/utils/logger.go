package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Logger struct {
	basePath string // Stores path without date/extension
	mu       sync.Mutex
}

// NewLogger creates a new logger instance
func NewLogger(basePath string) *Logger {
	dir := filepath.Dir(basePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		panic(fmt.Sprintf("failed to create log directory: %v", err))
	}
	return &Logger{basePath: basePath}
}

// Log writes an entry to the current day's log file
func (l *Logger) Log(operation, trxID, status, data string) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	logEntry := fmt.Sprintf("[%s] [%s] [%s] [%s] [%s]\n",
		now.Format("2006-01-02 15:04:05"),
		operation,
		trxID,
		status,
		data,
	)

	// Generate filename with date
	today := now.Format("2006-01-02")
	logPath := fmt.Sprintf("%s_%s.log", l.basePath, today)

	file, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}
	defer file.Close()

	if _, err := file.WriteString(logEntry); err != nil {
		return fmt.Errorf("failed to write log entry: %w", err)
	}

	return nil
}

// Helper function to create loggers with consistent naming
func createLogger(baseDir, fileName string) *Logger {
	return NewLogger(filepath.Join(baseDir, fileName))
}

// Pre-configured loggers with your exact requested filenames
func GetTransferLogger() *Logger {
	return createLogger("sps/log/transfer", "log_file_transfer")
}

func GetPaymentVALogger() *Logger {
	return createLogger("sps/log/payment/va", "log_file_payment_va")
}

func GetPaymentQRISLogger() *Logger {
	return createLogger("sps/log/payment/qris", "log_file_payment_qris")
}

func GetUserAuthLogger() *Logger {
	return createLogger("sps/log/user_auth", "log_file_user_auth")
}
