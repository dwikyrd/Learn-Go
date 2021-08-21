package mysql

import "recievermethodsederhana/domain/repository"

type repositories struct {
}

func NewRepository() repository.Repository {
	return repositories{}
}

func (r repositories) HelloWorld() string {
	return "hello world"
}

func (r repositories) Hello(input string) string {
	return "hello "+input
}