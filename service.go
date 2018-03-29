package mz

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mzmico/toolkit/balance"
	"github.com/mzmico/toolkit/errors"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type IService interface {
	Name() string
	Address() string
	Run() error
	Rpc(target string) *grpc.ClientConn

	setName(name string)
	setAddress(address string)
}

type Service struct {
	listen  string
	name    string
	balance *balance.DNSBalance
}

func (m *Service) Rpc(target string) *grpc.ClientConn {
	return m.balance.GetConn(target)
}

func (m *Service) setName(name string) {
	m.name = name
}

func (m *Service) setAddress(address string) {
	m.listen = address
}

func WithAddress(addr string) ServiceOption {
	return func(service IService) {
		service.setAddress(addr)
	}
}

func (m *Service) init() error {

	viper.SetConfigFile("settings.toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil {
		errors.By(err)
	}

	if len(m.name) == 0 {
		m.name = viper.GetString("service.name")

		fmt.Println(viper.Get("name"))

		if len(m.name) == 0 {
			m.name = strings.TrimSuffix(filepath.Base(os.Args[0]), ".exe")
		}
	}

	if len(m.listen) == 0 {
		m.listen = viper.GetString("service.address")

		if len(m.listen) == 0 {
			m.listen = "0.0.0.0:80"
		}
	}

	return nil
}

type ServiceOption func(service IService)

func (m *Service) Name() string {
	return m.name
}

func (m *Service) Address() string {
	return m.listen
}
