package http_service

import (
	"github.com/gin-gonic/gin"
	"github.com/mzmico/mz"
	"github.com/mzmico/toolkit/state"
)

var (
	service  *mz.HttpService
	handlers []func(engine *gin.Engine)
)

func Default(options ...mz.ServiceOption) (*mz.HttpService, error) {

	if service == nil {
		var (
			err error
		)

		service, err = mz.NewHttpService(options...)

		if err != nil {
			return nil, err
		}

		for _, handler := range handlers {
			handler(service.Engine())
		}

		return service, nil

	} else {
		panic("http_service already call Default")
	}
}

func GetService() *mz.HttpService {

	if service == nil {
		panic("service not init, use http_service.Default() to init.")
	}

	return service
}

func engine() *gin.Engine {
	return GetService().Engine()
}

type Handler func(state *state.HttpState)

func httpHandler(handler Handler) func(context *gin.Context) {

	return func(context *gin.Context) {

		state := state.NewHttpState(service, context)

		if state.GetLastError() != nil {

			state.Error(
				1000,
				state.GetLastError(),
			)
			return
		}

		handler(state)
	}
}

func GET(relativePath string, handler Handler) {

	handlers = append(handlers, func(engine *gin.Engine) {
		engine.GET(relativePath, httpHandler(handler))
	})

}

func POST(relativePath string, handler Handler) {
	handlers = append(handlers, func(engine *gin.Engine) {
		engine.POST(relativePath, httpHandler(handler))
	})

}

func DELETE(relativePath string, handler Handler) {
	handlers = append(handlers, func(engine *gin.Engine) {
		engine.DELETE(relativePath, httpHandler(handler))
	})

}

func PUT(relativePath string, handler Handler) {
	handlers = append(handlers, func(engine *gin.Engine) {
		engine.PUT(relativePath, httpHandler(handler))
	})

}
