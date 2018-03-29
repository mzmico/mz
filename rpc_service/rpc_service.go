package rpc_service

import (
	"fmt"
	"os"

	"github.com/mzmico/mz"
)

var (
	service *mz.RpcService
)

func Default(options ...mz.ServiceOption) *mz.RpcService {

	if service == nil {
		var (
			err error
		)

		service, err = mz.NewRpcService(options...)

		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}

		return service

	} else {
		panic("rpc_service already call Default")
	}
}

func GetService() *mz.RpcService {

	if service == nil {
		panic("service not init, use rpc_service.Default() to init.")
	}

	return service
}
