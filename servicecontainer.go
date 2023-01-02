package main

import (
	"crypto/tls"
	"sync"

	"github.com/metalcorpe/payshield-rest-api/controllers"
	"github.com/metalcorpe/payshield-rest-api/engine"
	"github.com/metalcorpe/payshield-rest-api/misc"
	"github.com/metalcorpe/payshield-rest-api/services"
	"go.uber.org/zap"
)

type IServiceContainer interface {
	InjectHsmController() controllers.HsmController
}

type kernel struct {
	log  *zap.Logger
	conf misc.Config
}

func (k *kernel) InjectHsmController() controllers.HsmController {
	k.log.Debug("Injecting Dependencies")
	cert, _ := tls.LoadX509KeyPair(k.conf.Hsm.ClientCert, k.conf.Hsm.ClientKey)
	config := tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}
	tcpConfig := engine.TcpConfig{Host: k.conf.Hsm.Ip, Port: k.conf.Hsm.Port, TlsConfig: &config, MaxIdleConns: 2, MaxOpenConn: 64}
	connectionPool, _ := engine.CreateTcpConnPool(&tcpConfig)
	hsmRepository := &engine.HsmRepository{IConnectionHandler: connectionPool}
	hsmService := &services.HsmService{IHsmRepository: hsmRepository}
	hsmController := controllers.HsmController{IHsmService: hsmService}
	return hsmController
}

var (
	k             *kernel
	containerOnce sync.Once
)

func ServiceContainer(log *zap.Logger, conf misc.Config) IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{log: log, conf: conf}
		})
	}
	return k
}