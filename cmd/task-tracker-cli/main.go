package main

import (
	"fmt"

	"github.com/natnael-alemayehu/task-tracker-cli/internal/server"
)

func main() {
	start := server.ReadCommand()
	fmt.Printf("%v", start)
}
