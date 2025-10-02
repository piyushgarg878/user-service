package repository

import (
    "user-service/internal/domain"
    "gorm.io/gorm"
)

type UserRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *domain.User) error {
    return r.db.Create(user).Error
}

func (r *UserRepository) GetByEmail(email string) (*domain.User, error) {
    var user domain.User
    err := r.db.Where("email = ?", email).First(&user).Error
    return &user, err
}

func (r *UserRepository) GetByID(id uint) (*domain.User, error) {
    var user domain.User
    err := r.db.First(&user, id).Error
    return &user, err
}
func (r *UserRepository) GetUserByName(name string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("name=?",name).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}