package log

import (
	"fmt"
	"time"
)

// Info TODO: documentation
func Info(msg string) {
	log("INFO", msg)
}

// Warn TODO: documentation
func Warn(msg string) {
	log("WARN", msg)
}

// Error TODO: documentation
func Error(err error) error {
	log("ERROR", err.Error())
	return err
}

func log(level string, msg string) {
	if level == "" {
		return
	}
	if msg == "" {
		return
	}
	fmt.Println("[" + level + "] " + time.Now().Format("[2006-01-02] [15:04:05] ") + msg)
}
