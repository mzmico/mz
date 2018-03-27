package mz

import (
	"github.com/gin-gonic/gin"
	"github.com/mzmico/toolkit/errors"
)

type HttpService struct {
	Service

	engine *gin.Engine
}

func (m *HttpService) Engine() *gin.Engine {
	return m.engine
}

func (m *HttpService) Run() error {

	err := m.engine.Run(m.listen)

	if err != nil {
		return errors.By(err)
	}
	return nil
}

func NewHttpService(opts ...ServiceOption) *HttpService {

	service := &HttpService{
		engine: gin.New(),
	}
	service.init()

	for _, opt := range opts {
		opt(service)
	}

	return service
}
