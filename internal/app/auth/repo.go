package auth

import "github.com/lai0xn/orka/internal/domain"

type AuthRepo interface {
	CreateUser(user *domain.User) error
	FindUser(email string) (*domain.User, error)
}
