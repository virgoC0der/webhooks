package main

import (
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"webhooks/services"
	. "webhooks/utils/log"
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
