package config

import (
	"fmt"
	"time"
)

// Log will introduce the logging concept into the EvoGuard applications output
func Log(format string, args ...any) {
	stamp := time.Now().Format("2006-01-02 3:4:5")
	fmt.Printf("\x1b[48;5;111;38;5;16;1m %s \x1b[0m %s\r\n", stamp, fmt.Sprintf(format, args...))
}
