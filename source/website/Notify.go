package website

import (
	"EvoGuard/source/config"
	"time"
)

/*
	Widget.go is how we only print to the terminal when the actual
	HTTPs server has been started without any errors spawning.
*/

// NewNotify will run for 5 seconds maximum, if the function reached 5 seconds it will print a success message, else it will end the routine
func NewNotify(cancel chan error, serverConfig *ServerConfig) {
	timeout := time.NewTicker(1 * time.Second)

	select {
	/* If function reaches here it has been acitve for over 5 seconds */
	case <-timeout.C:
		config.Log("\x1b[38;5;10;1mSuccessfully\x1b[0m sent signal to start server \"\x1b[0;38;5;13;3;1m%s:%d\x1b[0m\"", serverConfig.Address, serverConfig.Port)
		timeout.Stop()
		return

	/* Error prevented function from spawning success message */
	case <-cancel:
		timeout.Stop()
		return
	}
}
