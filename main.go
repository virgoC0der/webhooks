package main

import (
	"os"
	"os/signal"
	"syscall"

	. "github.com/virgoC0der/go-base/logging"
	"go.uber.org/zap"

	"webhooks/services"
)

func main() {
	InitLog()
	go services.WeatherHook()

	c := make(chan os.Signal, 1)
	defer close(c)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	s := <-c
	Logger.Info("Got signal: ", zap.Any("signal", s))
	Logger.Info("Exiting...")
}
