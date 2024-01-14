package util

import (
	"fmt"
	"log"
)

var logger = log.Default()

func LogErrorf(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	logger.Printf("[Error]: %s\n", msg)
}

func LogInfo(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	logger.Printf("[Info]: %s\n", msg)
}

func LogFatalf(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	logger.Fatalf("[Fatal]: %s\n", msg)
}
