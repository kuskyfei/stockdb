package main

import (
	"fmt"
	"os"
	"time"
)

const (
	logInfo    = "[ INFO  ]"
	logSuccess = "[SUCCESS]"
	logError   = "[ ERROR ]"
	logFatal   = "[ FATAL ]"
)

func log(level string, msgs ...interface{}) {
	fmt.Printf(
		"[%s] %9s %s\n",
		time.Now().Format("01-02 15:04:05"),
		level,
		fmt.Sprint(msgs...),
	)
	if level == logFatal {
		os.Exit(1)
	}
}
