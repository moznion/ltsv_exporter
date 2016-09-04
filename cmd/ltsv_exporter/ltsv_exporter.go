package main

import (
	"os"

	"github.com/moznion/ltsv_exporter"
)

func main() {
	ltsvExporter.Run(os.Args[1:])
}
