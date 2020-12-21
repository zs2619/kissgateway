package gateway

import (
	"github.com/sirupsen/logrus"
	"github.com/zs2619/kissnet-go"
)

type UserClinet struct {
	Conn         kissnet.IConnection
	UserClientID string
}

func (this *UserClinet) Init() {
	logrus.WithFields(logrus.Fields{
		"userID": this.UserClientID,
	}).Info("UserChat Init")
}

func (this *UserClinet) Dispose() {
	//退出channel
	logrus.Info("UserChat Dispose" + this.UserClientID)
}
