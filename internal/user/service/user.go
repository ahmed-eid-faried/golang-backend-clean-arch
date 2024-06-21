package service

import (
	"context"
	"errors"

	"github.com/quangdangfit/gocommon/logger"
	"github.com/quangdangfit/gocommon/validation"
	"golang.org/x/crypto/bcrypt"

	"main/internal/user/dto"
	"main/internal/user/model"
	"main/internal/user/repository"
	"main/pkg/jtoken"
	"main/pkg/paging"
	"main/pkg/utils"
)

//go:generate mockery --name=IUserService
type IUserService interface {
	Login(ctx context.Context, req *dto.LoginReq) (*model.User, string, string, error)
	Register(ctx context.Context, req *dto.RegisterReq) (*model.User, error)
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	RefreshToken(ctx context.Context, userID string) (string, error)
	VerifyEmail(ctx context.Context, request dto.VerifyEmailRequest) (dto.VerifyResponse, error)
	VerifyPhoneNumber(ctx context.Context, request dto.VerifyPhoneNumberRequest) (dto.VerifyResponse, error)
	ResendVerfiyCodePhone(ctx context.Context, request dto.ResendVerifyPhoneNumberRequest) (dto.VerifyResponse, error)
	ResendVerfiyCodeEmail(ctx context.Context, request dto.ResendVerifyEmailRequest) (dto.VerifyResponse, error)
	ListUsers(ctx context.Context, request dto.ListUsersReq) ([]*model.User, *paging.Pagination, error)
	UpdateUser(ctx context.Context, id string, req *dto.UpdateUserReq) error
	Delete(ctx context.Context, id string, req *dto.DeleteUserReq) (*model.User, error)
}

type UserService struct {
	validator validation.Validation
	repo      repository.IUserRepository
}

func NewUserService(
	validator validation.Validation,
	repo repository.IUserRepository) *UserService {
	return &UserService{
		validator: validator,
		repo:      repo,
	}
}

func (s *UserService) Login(ctx context.Context, req *dto.LoginReq) (*model.User, string, string, error) {
	if err := s.validator.ValidateStruct(req); err != nil {
		return nil, "", "", err
	}

	user, err := s.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		logger.Errorf("Login.GetUserByEmail fail, email: %s, error: %s", req.Email, err)
		return nil, "", "", err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, "", "", errors.New("wrong password")
	}

	tokenData := map[string]interface{}{
		"id":    user.ID,
		"email": user.Email,
		"role":  user.Role,
	}
	accessToken := jtoken.GenerateAccessToken(tokenData)
	refreshToken := jtoken.GenerateRefreshToken(tokenData)
	return user, accessToken, refreshToken, nil
}

func (s *UserService) Register(ctx context.Context, req *dto.RegisterReq) (*model.User, error) {
	if err := s.validator.ValidateStruct(req); err != nil {
		return nil, err
	}

	var user model.User
	utils.Copy(&user, &req)
	user.BeforeUpdate()
	err := s.repo.Create(ctx, &user)
	if err != nil {
		logger.Errorf("Register.Create fail, email: %s, error: %s", req.Email, err)
		return nil, err
	}
	return &user, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		logger.Errorf("GetUserByID fail, id: %s, error: %s", id, err)
		return nil, err
	}

	return user, nil
}

func (s *UserService) RefreshToken(ctx context.Context, userID string) (string, error) {
	user, err := s.repo.GetUserByID(ctx, userID)
	if err != nil {
		logger.Errorf("RefreshToken.GetUserByID fail, id: %s, error: %s", userID, err)
		return "", err
	}

	tokenData := map[string]interface{}{
		"id":    user.ID,
		"email": user.Email,
		"role":  user.Role,
	}
	accessToken := jtoken.GenerateAccessToken(tokenData)
	return accessToken, nil
}

func (s *UserService) UpdateUser(ctx context.Context, id string, req *dto.UpdateUserReq) error {
	if err := s.validator.ValidateStruct(req); err != nil {
		return err
	}
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		logger.Errorf("UpdateUser.GetUserByID fail, id: %s, error: %s", id, err)
		return err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return errors.New("wrong password")
	}

	user.Password = utils.HashAndSalt([]byte(req.NewPassword))
	err = s.repo.Update(ctx, user)
	if err != nil {
		logger.Errorf("UpdateUser.Update fail, id: %s, error: %s", id, err)
		return err
	}

	return nil
}

func (s *UserService) VerifyEmail(ctx context.Context, request dto.VerifyEmailRequest) (dto.VerifyResponse, error) {
	user, err := s.repo.FindByEmailAndVerifyCode(ctx, request.Email, request.VerifyCodeEmail)
	if err != nil {
		return dto.VerifyResponse{Message: "Verification failed"}, err
	}

	if user == nil {
		return dto.VerifyResponse{Message: "Verify code not correct"}, errors.New("verify code not correct")
	}

	user.ApproveEmail = true
	if err := s.repo.Update(ctx, user); err != nil {
		return dto.VerifyResponse{Message: "Failed to update user"}, err
	}

	return dto.VerifyResponse{Message: "Verification successful"}, nil
}

func (s *UserService) VerifyPhoneNumber(ctx context.Context, request dto.VerifyPhoneNumberRequest) (dto.VerifyResponse, error) {
	user, err := s.repo.FindByPhoneAndVerifyCode(ctx, request.PhoneNumber, request.VerifyCodePhoneNumber)
	if err != nil {
		return dto.VerifyResponse{Message: "Verification failed"}, err
	}

	if user == nil {
		return dto.VerifyResponse{Message: "Verify code not correct"}, errors.New("verify code not correct")
	}

	user.ApprovePhoneNumber = true
	if err := s.repo.Update(ctx, user); err != nil {
		return dto.VerifyResponse{Message: "Failed to update user"}, err
	}

	return dto.VerifyResponse{Message: "Verification successful"}, nil
}

func (s *UserService) ResendVerfiyCodePhone(ctx context.Context, request dto.ResendVerifyPhoneNumberRequest) (dto.VerifyResponse, error) {

	user, err := s.repo.FindByPhone(ctx, request.PhoneNumber)
	if err != nil {
		return dto.VerifyResponse{Message: "Resend Verification failed"}, err
	}

	if user == nil {
		return dto.VerifyResponse{Message: "Resend Verify code not correct"}, errors.New("verify code not correct")
	}
	// sendVerificationCodes sends the verification codes to the user's phone and email
	user.BeforeUpdateVerificationPhone()
	if err := s.repo.Update(ctx, user); err != nil {
		return dto.VerifyResponse{Message: "Failed to Resend user"}, err
	}

	return dto.VerifyResponse{Message: "Resend Verify code is successful"}, nil
}

func (s *UserService) ResendVerfiyCodeEmail(ctx context.Context, request dto.ResendVerifyEmailRequest) (dto.VerifyResponse, error) {

	user, err := s.repo.FindByEmail(ctx, request.Email)
	if err != nil {
		return dto.VerifyResponse{Message: "Resend Verification failed"}, err
	}

	if user == nil {
		return dto.VerifyResponse{Message: "Resend Verify code not correct"}, errors.New("verify code not correct")
	}
	// sendVerificationCodes sends the verification codes to the user's phone and email
	user.BeforeUpdateVerificationEmail()
	if err := s.repo.Update(ctx, user); err != nil {
		return dto.VerifyResponse{Message: "Failed to Resend Resend"}, err
	}

	return dto.VerifyResponse{Message: "Resend Verify code is successful"}, nil
}

func (p *UserService) ListUsers(ctx context.Context, req dto.ListUsersReq) ([]*model.User, *paging.Pagination, error) {
	Userss, pagination, err := p.repo.ListUsers(ctx, req)
	if err != nil {
		return nil, nil, err
	}

	return Userss, pagination, nil
}

func (p *UserService) Delete(ctx context.Context, id string, req *dto.DeleteUserReq) (*model.User, error) {
	if err := p.validator.ValidateStruct(req); err != nil {
		return nil, err
	}

	User, err := p.repo.GetUserByID(ctx, id)
	if err != nil {
		logger.Errorf("Delete.GetUserByID fail, id: %s, error: %s", id, err)
		return nil, err
	}

	utils.Copy(User, req)
	err = p.repo.Delete(ctx, User)
	if err != nil {
		logger.Errorf("Delete fail, id: %s, error: %s", id, err)
		return nil, err
	}

	return User, nil
}
