package repository

type Repository interface {
	HelloWorld() string
    Hello(input string) string
}
