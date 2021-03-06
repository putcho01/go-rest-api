package usecase

import "github.com/ga-tcho/go-rest-api/domain"

type UserRepository interface {
	FindByID(id int) (domain.User, error)
	Store(domain.User) (domain.User, error)
	Update(domain.User) (domain.User, error)
	DeleteByID(domain.User) error
	FindAll() (domain.Users, error)
}
