package domain

import "time"

type UserProfile struct {
    ID          uint      `gorm:"primaryKey" json:"id"`
    UserID      uint      `gorm:"uniqueIndex" json:"user_id"`
    FirstName   string    `json:"first_name,omitempty"`
    LastName    string    `json:"last_name,omitempty"`
    Bio         string    `json:"bio,omitempty"`
    Location    string    `json:"location,omitempty"`
    ProfilePic  string    `json:"profile_pic,omitempty"`
    UpdatedAt   time.Time `json:"updated_at"`
}