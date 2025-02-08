package repo

import (
	"github.com/lai0xn/orka/internal/domain"
	"gorm.io/gorm"
)

type apiKeyRepo struct {
	db *gorm.DB
}

func NewKeyRepository(db *gorm.DB) *apiKeyRepo {
	return &apiKeyRepo{db: db}
}

func (r *apiKeyRepo) Create(apiKey *domain.ApiKey) error {
	return r.db.Create(apiKey).Error
}

func (r *apiKeyRepo) GetByKey(key string) (*domain.ApiKey, error) {
	var apiKey domain.ApiKey
	if err := r.db.Where("key = ?", key).First(&apiKey).Error; err != nil {
		return nil, err
	}
	return &apiKey, nil
}

func (r *apiKeyRepo) GetByUserID(userID uint) ([]domain.ApiKey, error) {
	var apiKeys []domain.ApiKey
	if err := r.db.Where("user_id = ?", userID).Find(&apiKeys).Error; err != nil {
		return nil, err
	}
	return apiKeys, nil
}

func (r *apiKeyRepo) DeleteByKey(key string) error {
	return r.db.Where("key = ?", key).Delete(&domain.ApiKey{}).Error
}
