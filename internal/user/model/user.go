// package model

// import (
// 	"time"
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//

// 	"github.com/google/uuid"
// 	"gorm.io/gorm"

// 	"main/pkg/utils"
// )

// type UserRole string

// const (
// 	UserRoleAdmin    UserRole = "admin"
// 	UserRoleCustomer UserRole = "customer"
// )

// type User struct {
// 	ID        string     `json:"id" gorm:"unique;not null;index;primary_key"`
// 	CreatedAt time.Time  `json:"created_at"`
// 	UpdatedAt time.Time  `json:"updated_at"`
// 	DeletedAt *time.Time `json:"deleted_at" gorm:"index"`
// 	Email     string     `json:"email" gorm:"unique;not null;index:idx_user_email"`
// 	Password  string     `json:"password"`
// 	Role      UserRole   `json:"role"`
// }

// func (user *User) BeforeCreate(tx *gorm.DB) error {
// 	user.ID = uuid.New().String()
// 	user.Password = utils.HashAndSalt([]byte(user.Password))
// 	if user.Role == "" {
// 		user.Role = UserRoleCustomer
// 	}
// 	return nil
// }

package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"main/pkg/utils"
)

// UserRole represents the role of a user
type UserRole string

// Constants for user roles
const (
	UserRoleAdmin    UserRole = "admin"    // Administrator role
	UserRoleCustomer UserRole = "customer" // Customer role
)

// User represents a user in the system
type User struct {
	ID         string     `json:"id" gorm:"unique;not null;index;primary_key"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at" gorm:"index"`
	Email      string     `json:"email" gorm:"unique;not null;index:idx_user_email"`
	Password   string     `json:"password"`
	Role       UserRole   `json:"role"`
	VerifyCode string     `json:"verify_code"`
	Approve    bool       `json:"approve"`
}

// BeforeCreate is a hook that is called before creating a new user
func (user *User) BeforeCreate(tx *gorm.DB) error {
	// Generate a unique ID for the user
	user.ID = uuid.New().String()

	// Hash and salt the password for security
	user.Password = utils.HashAndSalt([]byte(user.Password))

	// Set the default role to customer if not specified
	if user.Role == "" {
		user.Role = UserRoleCustomer
	}

	return nil
}
