package auth

import (
	"github.com/lai0xn/orka/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo AuthRepo
}

func NewAuthService(repo AuthRepo) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s AuthService) Signup(dto SignupDTO) error {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(dto.Password), 12)
	if err != nil {
		return err
	}
	err = s.repo.CreateUser(&domain.User{
		Email:    dto.Email,
		Username: dto.Username,
		Password: string(hashedPwd),
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) Authenticate(dto LoginDTO) (*domain.User, error) {
	user, err := s.repo.FindUser(dto.Email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password))
	if err != nil {
		return nil, nil
	}

	return user, nil
}
