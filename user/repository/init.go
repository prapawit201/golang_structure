package repository

//go:generate /Users/babe/go/bin/mockgen -source=init.go -destination=mocks/repository.go

type UserRepository interface {
	GetAll() error
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}
