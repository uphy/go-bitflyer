package main

import (
	"fmt"
	"os"

	"github.com/uphy/go-bitflyer/bitflyer/commands"
)

func main() {
	if err := commands.Execute(); err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Command failed. (err=%s)", err))
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
