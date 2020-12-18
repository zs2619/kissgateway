package main

import (
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/zs2619/kissgateway/common"
	"github.com/zs2619/kissgateway/gateway"
	"github.com/zs2619/kissnet-go"
)

var gAcceptor kissnet.IAcceptor

func main() {
	conf, err := common.LoadWebConfig("assets/config.json")
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Info("common.LoadWebConfig error")
		return
	}

	port := 0
	portStr := os.Getenv("GATEWAY_CONFIGPATH")
	if len(portStr) == 0 {
		port = conf.ServerPort
	} else {
		port, err = strconv.Atoi(portStr)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"error": err,
			}).Info("os.Getenv(GATEWAY_CONFIGPATH) error")
			return
		}
	}

	event := kissnet.NewNetEvent()
	logrus.Info("acceptor start")
	gAcceptor, err := kissnet.AcceptorFactory(
		conf.ServerType,
		port,
		gateway.GWClientCB,
	)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("AcceptorFactory error")
		return
	}

	gAcceptor.Run()
	event.EventLoop()
	gAcceptor.Close()
	logrus.Info("acceptor end")
}
