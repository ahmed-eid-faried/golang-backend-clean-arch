package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"main/internal/user/model"
	"main/pkg/dbs"
)

//go:generate mockery --name=IUserRepository
type IUserRepository interface {
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	FindByEmailAndVerifyCode(ctx context.Context, email, verifyCode string) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) error
}

type UserRepo struct {
	db dbs.IDatabase
}

// repo repository.IUserRepository
func NewUserRepository(db dbs.IDatabase) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(ctx context.Context, user *model.User) error {
	return r.db.Create(ctx, user)
}

func (r *UserRepo) Update(ctx context.Context, user *model.User) error {
	return r.db.Update(ctx, user)
}

func (r *UserRepo) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	if err := r.db.FindById(ctx, id, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepo) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	query := dbs.NewQuery("email = ?", email)
	if err := r.db.FindOne(ctx, &user, dbs.WithQuery(query)); err != nil {
		return nil, err
	}

	return &user, nil
}
func (r *UserRepo) FindByEmailAndVerifyCode(ctx context.Context, email, verifyCode string) (*model.User, error) {
	var user model.User
	query := "SELECT id, email, verify_code, approve FROM users WHERE email = ? AND verify_code = ?"

	result := r.db.GetDB().WithContext(ctx).Raw(query, email, verifyCode).Scan(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepo) UpdateUser(ctx context.Context, user *model.User) error {
	query := "UPDATE users SET approve = ? WHERE email = ?"
	err := r.db.Exec(ctx, query, user.Approve, user.Email)
	return err
}
