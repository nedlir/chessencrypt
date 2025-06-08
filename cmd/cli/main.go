package main

import (
	"fmt"
	"os"

	"github.com/nedlir/chessencrypt/cli"
)

func main() {
	if err := cli.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
