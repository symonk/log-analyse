package main

import (
	"fmt"
	"log/slog"

	"github.com/symonk/log-analyse/cmd"
	"github.com/symonk/log-analyse/internal/logger"
)

func main() {
	logger.Init()
	slog.Info(fmt.Sprintf("%s loaded. version: %s", name, version))
	cmd.Execute()
}
