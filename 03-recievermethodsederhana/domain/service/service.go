package service

import (
	"recievermethodsederhana/domain/repository"
)

type service struct {
	repository repository.Repository
}

type serviceOrkestra interface {
    HelloWorldService() string
    PrintHelloWorldService()
    PrintHello(input string)
}

func NewService(repo repository.Repository) serviceOrkestra {
	return service{
        repository: repo,
    }
}
