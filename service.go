package mz

import (
	"os"
	"path/filepath"
	"strings"
)

type IService interface {
	Name() string
	Address() string
	Run() error

	setName(name string)
	setAddress(address string)
}

type Service struct {
	listen string
	name   string
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

func (m *Service) init() {
	if len(m.name) == 0 {
		m.name = strings.TrimSuffix(filepath.Base(os.Args[0]), ".exe")
	}
}

type ServiceOption func(service IService)

func (m *Service) Name() string {
	return m.name
}

func (m *Service) Address() string {
	return m.listen
}
