package models

import "time"

// User model
type User struct {
	ID              uint64     `json:"id"`
	FirstName       string     `json:"first_name"`
	LastName        string     `json:"last_name"`
	ProfileImage    string     `json:"profile_image"`
	Email           string     `json:"email" binding:"required,email"`
	IsEmailVerified bool       `json:"is_email_verified"`
	Bio             string     `json:"bio"`
	PhoneNumber     string     `json:"phone_number"`
	IsPhoneVerified bool       `json:"is_phone_verified"`
	Country         string     `json:"country"`
	State           string     `json:"state"`
	Area            string     `json:"area"`
	City            string     `json:"city"`
	Address         string     `json:"address"`
	PostCode        string     `json:"post_code"`
	BirthDate       *time.Time `json:"birth_date"`
	TosAccepted     bool       `json:"tos_accepted"`
	InvitedByUserID *uint64    `json:"invited_by_user_id"`
	IsActive        bool       `json:"is_active"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at"`
}
