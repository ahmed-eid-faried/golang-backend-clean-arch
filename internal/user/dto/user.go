package dto

import (
	"time"
)

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RegisterReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
}

type RegisterRes struct {
	User User `json:"user"`
}

type LoginReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
}

type LoginRes struct {
	User         User   `json:"user"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenReq struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type RefreshTokenRes struct {
	AccessToken string `json:"access_token"`
}

type UpdateUserReq struct {
	Password    string `json:"password" validate:"required,password"`
	NewPassword string `json:"new_password" validate:"required,password"`
}

type UpdateUserRes struct {
	Message string `json:"message"`
}

//***************************************************************************\\
//***************************************************************************\\

type VerifyPhoneNumberRequest struct {
	PhoneNumber           string `json:"phone_number"`
	VerifyCodePhoneNumber string `json:"verify_code_phone_number"`
}
type ResendVerifyPhoneNumberRequest struct {
	PhoneNumber string `json:"phone_number"`
}

// ***************************************************************************\\
type VerifyEmailRequest struct {
	Email           string `json:"email"`
	VerifyCodeEmail string `json:"verify_code_email"`
}
type ResendVerifyEmailRequest struct {
	Email string `json:"email"`
}

// ***************************************************************************\\
type VerifyResponse struct {
	Message string `json:"message"`
}

//***************************************************************************\\
//***************************************************************************\\
