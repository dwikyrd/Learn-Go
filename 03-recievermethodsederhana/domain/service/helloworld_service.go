package service

import "fmt"

func (s service) HelloWorldService() string {
	return s.repository.HelloWorld()
}

func (s service) PrintHelloWorldService() {
	fmt.Println(s.repository.HelloWorld())
}

func (s service) PrintHello(input string) {
	fmt.Println(s.repository.Hello(input))
}
