package main

import (
	"os"

	"github.com/infrakit-io/talos-docker-bootstrap/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}
