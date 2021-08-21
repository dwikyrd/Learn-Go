package main

import (
	"fmt"
	repository "recievermethodsederhana/domain/repository/mysql"
	"recievermethodsederhana/domain/service"
)

func main() {
	repo := repository.NewRepository()
	service := service.NewService(repo)

    service.HelloWorldService()
    fmt.Println(service.HelloWorldService())
    service.PrintHello("abda")
}
