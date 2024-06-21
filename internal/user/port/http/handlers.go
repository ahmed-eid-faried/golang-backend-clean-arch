package http

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/quangdangfit/gocommon/logger"

	"main/internal/user/dto"
	"main/internal/user/service"
	"main/pkg/config"
	"main/pkg/redis"
	"main/pkg/response"
	"main/pkg/utils"
)

type UserHandler struct {
	cache   redis.IRedis
	service service.IUserService
}

func NewUserHandler(service service.IUserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

// Login , Register, GetMe, RefreshToken, UpdateUser, VerfiyCode FOR PHONE

// Login godoc
//
//	@Summary	Login
//	@Tags		users
//	@Produce	json
//	@Param		_	body		dto.LoginReq	true	"Body"
//	@Success	200	{object}	dto.LoginRes
//	@Router		/auth/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var req dto.LoginReq
	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		logger.Error("Failed to get body ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	user, accessToken, refreshToken, err := h.service.Login(c, &req)
	if err != nil {
		logger.Error("Failed to login ", err)
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}

	var res dto.LoginRes
	utils.Copy(&res.User, &user)
	res.AccessToken = accessToken
	res.RefreshToken = refreshToken
	response.JSON(c, http.StatusOK, res)
}

// Register godoc
//
//	@Summary	Register new user
//	@Tags		users
//	@Produce	json
//	@Param		_	body		dto.RegisterReq	true	"Body"
//	@Success	200	{object}	dto.RegisterRes
//	@Router		/auth/register [post]
func (h *UserHandler) Register(c *gin.Context) {
	var req dto.RegisterReq
	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		logger.Error("Failed to get body", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	user, err := h.service.Register(c, &req)
	if err != nil {
		logger.Error(err.Error())
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}

	var res dto.RegisterRes
	utils.Copy(&res.User, &user)
	response.JSON(c, http.StatusOK, res)
}

// GetMe godoc
//
//	@Summary	get my profile
//	@Tags		users
//	@Security	ApiKeyAuth
//	@Produce	json
//	@Success	200	{object}	dto.User
//	@Router		/auth/me [get]
func (h *UserHandler) GetMe(c *gin.Context) {
	userID := c.GetString("userId")
	if userID == "" {
		response.Error(c, http.StatusUnauthorized, errors.New("unauthorized"), "Unauthorized")
		return
	}

	user, err := h.service.GetUserByID(c, userID)
	if err != nil {
		logger.Error(err.Error())
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}

	var res dto.User
	utils.Copy(&res, &user)
	response.JSON(c, http.StatusOK, res)
}

// GetMe godoc
//
//	@Summary	get my profile
//	@Tags		users
//	@Security	ApiKeyAuth
//	@Produce	json
//	@Param		_	body		dto.RefreshTokenReq	true	"Body"
//	@Success	200	{object}	dto.RefreshTokenRes
//	@Router		/auth/refresh-token [get]
func (h *UserHandler) RefreshToken(c *gin.Context) {
	userID := c.GetString("userId")
	if userID == "" {
		response.Error(c, http.StatusUnauthorized, errors.New("unauthorized"), "Unauthorized")
		return
	}

	accessToken, err := h.service.RefreshToken(c, userID)
	if err != nil {
		logger.Error("Failed to refresh token", err)
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}

	res := dto.RefreshTokenRes{
		AccessToken: accessToken,
	}
	response.JSON(c, http.StatusOK, res)
}

// UpdateUser godoc
//
//	@Summary	changes the password
//	@Tags		users
//	@Security	ApiKeyAuth
//	@Produce	json
//	@Param		_	body	dto.UpdateUserReq	true	"Body"
//	@Success	200	{object}	dto.UpdateUserRes
//	@Router		/auth/update-user [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	var req dto.UpdateUserReq
	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		logger.Error("Failed to get body", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	userID := c.GetString("userId")
	err := h.service.UpdateUser(c, userID, &req)
	if err != nil {
		logger.Error(err.Error())
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}
	response.JSON(c, http.StatusOK, nil)
}

// VerfiyCodeEmail godoc
//
//	@Summary	Verfiy Code for Email
//	@Tags		users
//	@Security	ApiKeyAuth
//	@Produce	json
//	@Param		_	body	dto.VerifyEmailRequest	true	"Body"
//	@Success	200	{object}	dto.VerifyResponse
//	@Router		/auth/verfiy-code-email [put]
func (h *UserHandler) VerfiyCodeEmail(c *gin.Context) {
	var req dto.VerifyEmailRequest
	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		logger.Error("Failed to get body", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	resp, err := h.service.VerifyEmail(c, req)
	if err != nil {
		logger.Error(err.Error())
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}
	response.JSON(c, http.StatusOK, resp)
}

// VerfiyCodePhoneNumber godoc
//
//	@Summary	Verfiy Code for PhoneNumber
//	@Tags		users
//	@Security	ApiKeyAuth
//	@Produce	json
//	@Param		_	body	dto.VerifyPhoneNumberRequest	true	"Body"
//	@Success	200	{object}	dto.VerifyResponse
//	@Router		/auth/verfiy-code-phone-number [put]
func (h *UserHandler) VerfiyCodePhoneNumber(c *gin.Context) {
	var req dto.VerifyPhoneNumberRequest
	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		logger.Error("Failed to get body", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	resp, err := h.service.VerifyPhoneNumber(c, req)
	if err != nil {
		logger.Error(err.Error())
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}
	response.JSON(c, http.StatusOK, resp)
}

// VerfiyCodePhoneNumber godoc
//
//	@Summary	Verfiy Code for PhoneNumber
//	@Tags		users
//	@Security	ApiKeyAuth
//	@Produce	json
//	@Param		_	body	dto.ResendVerifyPhoneNumberRequest	true	"Body"
//	@Success	200	{object}	dto.VerifyResponse
//	@Router		/auth/resend-verfiy-code-phone-number [put]
func (h *UserHandler) VerfiyCodePhoneNumberResend(c *gin.Context) {
	var req dto.ResendVerifyPhoneNumberRequest
	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		logger.Error("Failed to get body", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	resp, err := h.service.ResendVerfiyCodePhone(c, req)
	if err != nil {
		logger.Error(err.Error())
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}
	response.JSON(c, http.StatusOK, resp)
}

// VerfiyCodeEmail godoc
//
//	@Summary	Verfiy Code for Email
//	@Tags		users
//	@Security	ApiKeyAuth
//	@Produce	json
//	@Param		_	body	dto.ResendVerifyEmailRequest	true	"Body"
//	@Success	200	{object}	dto.VerifyResponse
//	@Router		/auth/resend-verfiy-code-email [put]
func (h *UserHandler) VerfiyCodeEmailResend(c *gin.Context) {
	var req dto.ResendVerifyEmailRequest
	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		logger.Error("Failed to get body", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	resp, err := h.service.ResendVerfiyCodeEmail(c, req)
	if err != nil {
		logger.Error(err.Error())
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}
	response.JSON(c, http.StatusOK, resp)
}

// ListUsers godoc
//
//		@Summary	Get list Users
//		@Tags		Users
//		@Produce	json
//	 @Param id_user header string true "id_user"
//	@Param		name	path	string					true	"name"
//	@Param		page	path	string					true	"page"
//	@Param		limit	path	string					true	"limit"
//
// @Success	200	{object}	dto.ListUsersRes
// @Router		/auth/users  [get]
func (p *UserHandler) ListUsers(c *gin.Context) {
	Name := c.Param("name")
	IDUser := c.Param("id_user")
	PageStr := c.Param("page")
	LimitStr := c.Param("limit")

	Page, err1 := strconv.ParseInt(PageStr, 10, 64)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, utils.HTTPError{Code: http.StatusBadRequest, Message: "Invalid page number"})
		return
	}

	Limit, err2 := strconv.ParseInt(LimitStr, 10, 64)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, utils.HTTPError{Code: http.StatusBadRequest, Message: "Invalid limit number"})
		return
	}

	// var req dto.ListUsersReq
	req := dto.ListUsersReq{
		Name:   Name,
		IDUser: IDUser,
		Page:   Page,
		Limit:  Limit,
	}
	if err := c.ShouldBindQuery(&req); err != nil {
		logger.Error("Failed to parse request query: ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	var res dto.ListUsersRes
	cacheKey := c.Request.URL.RequestURI()
	err := p.cache.Get(cacheKey, &res)
	if err == nil {
		response.JSON(c, http.StatusOK, res)
		return
	}

	Users, pagination, err := p.service.ListUsers(c, req)
	if err != nil {
		logger.Error("Failed to get list Users: ", err)
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}

	utils.Copy(&res.Users, &Users)
	res.Pagination = pagination
	response.JSON(c, http.StatusOK, res)
	_ = p.cache.SetWithExpiration(cacheKey, res, config.UsersCachingTime)
}

// DeleteUser godoc
//
//	@Summary	Delete User
//	@Tags		User
//	@Produce	json
//	@Security	ApiKeyAuth
//	@Param		_	body	dto.DeleteUserReq	true	"Body"
//	@Success	200	{object}	dto.User
//	@Router		/auth/{id} [Delete]
func (p *UserHandler) DeleteUser(c *gin.Context) {
	//	@Param		id	path	string					true	"User ID"
	// UserId := c.Param("id")
	var req dto.DeleteUserReq
	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		logger.Error("Failed to get body", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	User, err := p.service.Delete(c, req.ID, &req)
	if err != nil {
		logger.Error("Failed to Delete User", err.Error())
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}

	var res dto.User
	utils.Copy(&res, &User)
	response.JSON(c, http.StatusOK, res)
	_ = p.cache.RemovePattern("*User*")
}
