package keys

import "github.com/lai0xn/orka/internal/domain"

type Repo interface {
	Create(apiKey *domain.ApiKey) error
	GetByKey(key string) (*domain.ApiKey, error)
	GetByUserID(userID uint) ([]domain.ApiKey, error)
	DeleteByKey(key string) error
}
