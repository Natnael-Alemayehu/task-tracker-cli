package main

import (
	"fmt"

	"github.com/natnael-alemayehu/task-tracker-cli/internal/server"
)

func main() {
	ans := server.ReadCommand()
	fmt.Printf("The ans: %v", ans)
}
