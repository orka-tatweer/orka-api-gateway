package repo

import (
	"errors"

	"github.com/lai0xn/orka/internal/domain"
	"gorm.io/gorm"
)

type authRepo struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) *authRepo {
	return &authRepo{db: db}
}

func (r *authRepo) CreateUser(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *authRepo) FindUser(email string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}
