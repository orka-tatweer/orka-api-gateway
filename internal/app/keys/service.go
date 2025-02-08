package keys

import "github.com/lai0xn/orka/internal/domain"

type ApiKeyService struct {
	repo Repo
}

func NewApiKeyService(repo Repo) *ApiKeyService {
	return &ApiKeyService{repo: repo}
}

func (s *ApiKeyService) GenerateApiKey(userID uint, key string) (*domain.ApiKey, error) {
	apiKey := &domain.ApiKey{
		Key:    key,
		UserID: userID,
	}
	err := s.repo.Create(apiKey)
	return apiKey, err
}

func (s *ApiKeyService) ValidateApiKey(key string) (*domain.ApiKey, error) {
	return s.repo.GetByKey(key)
}

func (s *ApiKeyService) GetUserApiKeys(userID uint) ([]domain.ApiKey, error) {
	return s.repo.GetByUserID(userID)
}

func (s *ApiKeyService) RevokeApiKey(key string) error {
	return s.repo.DeleteByKey(key)
}
