package services

import (
	"errors"
	"portfolio-backend/internal/dto"
	"portfolio-backend/internal/models"
	"portfolio-backend/internal/repositories"
	"portfolio-backend/internal/utils"
)

// AuthService menangani business logic untuk autentikasi.
type AuthService struct {
	userRepo  *repositories.UserRepository
	jwtSecret string
}

// NewAuthService membuat instance baru AuthService.
func NewAuthService(userRepo *repositories.UserRepository, jwtSecret string) *AuthService {
	return &AuthService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

// Login memvalidasi credentials dan mengembalikan JWT token.
func (s *AuthService) Login(req dto.LoginRequest) (string, error) {
	user, err := s.userRepo.FindByUsername(req.Username)
	if err != nil {
		return "", errors.New("username atau password salah")
	}

	if !utils.CheckPassword(user.Password, req.Password) {
		return "", errors.New("username atau password salah")
	}

	token, err := utils.GenerateToken(user.ID, s.jwtSecret)
	if err != nil {
		return "", errors.New("gagal membuat token")
	}

	return token, nil
}

// Register membuat user baru.
func (s *AuthService) Register(req dto.RegisterRequest) (*models.User, error) {
	// Cek apakah username sudah dipakai
	existing, _ := s.userRepo.FindByUsername(req.Username)
	if existing != nil {
		return nil, errors.New("username sudah digunakan")
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("gagal memproses password")
	}

	user := &models.User{
		Username: req.Username,
		Password: hashedPassword,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, errors.New("gagal membuat user")
	}

	return user, nil
}
