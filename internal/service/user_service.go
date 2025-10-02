package service

import (
	"errors"
	"user-service/internal/domain"
	"user-service/internal/repository"
	"user-service/internal/utils"
)

type UserService struct {
    UserRepo    *repository.UserRepository
    ProfileRepo *repository.ProfileRepository
}

func NewUserService(u *repository.UserRepository, p *repository.ProfileRepository) *UserService {
    return &UserService{UserRepo: u, ProfileRepo: p}
}

func (s *UserService) Register(user *domain.User) error {
    if user.Name == "" || user.Email == "" || user.Password == "" {
		return errors.New("name, email, and password are required")
	}
    hashed, _ := utils.HashPassword(user.Password)
    user.Password = hashed
    if err := s.UserRepo.CreateUser(user); err != nil {
		return err
	}
	// create empty profile
	return s.UserRepo.CreateEmptyProfile(user.ID)
}
func (s *UserService) GetUserByName(name string) (*domain.User, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}
	return s.UserRepo.GetUserByName(name)
}
func (s *UserService) Login(email, password string) (*domain.User, error) {
    
    user, err := s.UserRepo.GetByEmail(email)
    if err != nil {
        return nil, err
    }
    if !utils.CheckPasswordHash(password, user.Password) {
        return nil, err
    }
    return user, nil
}


func (s *UserService) UpdateProfile(userID uint, profile *domain.UserProfile) error {
	if userID == 0 {
		return errors.New("invalid user ID")
	}
	return s.ProfileRepo.UpdateProfile(userID, profile)
}