package main

import (
	"go2ban/api/REST"
	"go2ban/api/gRPC"
	"go2ban/cmd/commandLine"
	"go2ban/cmd/fakeSocks"
	"go2ban/cmd/firewall"
	"go2ban/cmd/localService"
	"go2ban/pkg/config"
	"go2ban/pkg/logger"
	"go2ban/pkg/proFile"
)

func main() {
	pprofEnd := proFile.Start("cpu & mem & mutex & block")
	config.Load()
	logger.Start()
	commandLine.Run()
	firewall.WorkerStart(config.Get().Flags.RunAsDaemon)
	fakeSocks.Listen(config.Get().FakeSocksPorts)
	REST.Start(config.Get().Flags.RunAsDaemon) //TODO validator check port
	gRPC.Start(config.Get().Flags.RunAsDaemon)
	localService.WorkerStart(config.Get().Services, pprofEnd)

}
