package repository

import (
    "user-service/internal/domain"
    "gorm.io/gorm"
)

type ProfileRepository struct {
    db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) *ProfileRepository {
    return &ProfileRepository{db: db}

}
/*
ID          uint      `gorm:"primaryKey" json:"id"`
    UserID      uint      `gorm:"uniqueIndex" json:"user_id"`
    FirstName   string    `json:"first_name,omitempty"`
    LastName    string    `json:"last_name,omitempty"`
    Bio         string    `json:"bio,omitempty"`
    Location    string    `json:"location,omitempty"`
    ProfilePic  string    `json:"profile_pic,omitempty"`
    UpdatedAt   time.Time `json:"updated_at"`
*/
func (r *UserRepository) CreateEmptyProfile(userID uint) error {
	profile := domain.UserProfile{
		UserID:    userID,
		FirstName:"",
		LastName: "",
		Bio:       "",
		Location:  "",
		ProfilePic: "",
	}
	return r.db.Create(&profile).Error
}

func (r *ProfileRepository) Create(profile *domain.UserProfile) error {
    return r.db.Create(profile).Error
}

func (r *ProfileRepository) GetByUserID(userID uint) (*domain.UserProfile, error) {
    var Profile domain.UserProfile
    err := r.db.Where("user_id = ?", userID).First(&Profile).Error
    return &Profile, err
}
func (r *ProfileRepository) UpdateProfile(userID uint, profile *domain.UserProfile) error {
	var existing domain.UserProfile
	if err := r.db.Where("user_id = ?", userID).First(&existing).Error; err != nil {
		return err
	}
	existing.FirstName=profile.FirstName
	existing.LastName=profile.LastName
	existing.Bio=profile.Bio
	existing.Location=profile.Location
	existing.ProfilePic=profile.ProfilePic


	return r.db.Save(&existing).Error
}
func (r *ProfileRepository) Update(profile *domain.UserProfile) error {
    return r.db.Save(profile).Error
}