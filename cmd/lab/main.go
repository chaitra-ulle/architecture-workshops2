package main

import (
	"fmt"
	"os"

	"github.com/infobloxopen/architecture-workshops2/pkg/api"
	"github.com/infobloxopen/architecture-workshops2/pkg/dep"
	"github.com/infobloxopen/architecture-workshops2/pkg/worker"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: lab <api|worker|dep>")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "api":
		api.Run()
	case "worker":
		worker.Run()
	case "dep":
		dep.Run()
	default:
		fmt.Fprintf(os.Stderr, "Unknown mode: %s\n", os.Args[1])
		os.Exit(1)
	}
}
