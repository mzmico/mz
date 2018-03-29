package http_service

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mzmico/mz"
	"github.com/mzmico/protobuf"
	"github.com/mzmico/toolkit/state"
)

var (
	service *mz.HttpService
)

func Default(options ...mz.ServiceOption) *mz.HttpService {

	if service == nil {
		var (
			err error
		)

		service, err = mz.NewHttpService(options...)

		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}

		return service

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
				int(protobuf.State_ArgumentError),
				protobuf.State_ArgumentError.String(),
				state.GetLastError(),
			)

			return
		}

		handler(state)
	}
}

func GET(relativePath string, handler Handler) {
	engine().GET(relativePath, httpHandler(handler))
}

func POST(relativePath string, handler Handler) {
	engine().POST(relativePath, httpHandler(handler))
}

func DELETE(relativePath string, handler Handler) {
	engine().DELETE(relativePath, httpHandler(handler))
}

func PUT(relativePath string, handler Handler) {
	engine().DELETE(relativePath, httpHandler(handler))
}
