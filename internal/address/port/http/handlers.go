package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quangdangfit/gocommon/logger"

	"main/internal/address/dto"
	"main/internal/address/service"
	"main/pkg/config"
	"main/pkg/redis"
	"main/pkg/response"
	"main/pkg/utils"
)

// Address
// address
type AddressHandler struct {
	cache   redis.IRedis
	service service.IAddressService
}

func NewAddressHandler(
	cache redis.IRedis,
	service service.IAddressService,
) *AddressHandler {
	return &AddressHandler{
		cache:   cache,
		service: service,
	}
}

// GetAddressByID godoc
//
//	@Summary	Get Address by id
//	@Tags		Address
//	@Produce	json
//	@Param		id	path	string	true	"Address ID"
//	@Router		/address/{id} [get]
func (p *AddressHandler) GetAddressByID(c *gin.Context) {
	var res dto.Address
	cacheKey := c.Request.URL.RequestURI()
	err := p.cache.Get(cacheKey, &res)
	if err == nil {
		response.JSON(c, http.StatusOK, res)
		return
	}

	AddressId := c.Param("id")
	Address, err := p.service.GetAddressByID(c, AddressId)
	if err != nil {
		logger.Error("Failed to get Address detail: ", err)
		response.Error(c, http.StatusNotFound, err, "Not found")
		return
	}

	utils.Copy(&res, &Address)
	response.JSON(c, http.StatusOK, res)
	_ = p.cache.SetWithExpiration(cacheKey, res, config.AddressCachingTime.Abs())
}

// ListAddress godoc
//
//	@Summary	Get list Address
//	@Tags		Address
//	@Produce	json
//	@Success	200	{object}	dto.ListAddressRes
//	@Router		/address [get]
func (p *AddressHandler) ListAddresses(c *gin.Context) {
	var req dto.ListAddressReq
	if err := c.ShouldBindQuery(&req); err != nil {
		logger.Error("Failed to parse request query: ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	var res dto.ListAddressRes
	cacheKey := c.Request.URL.RequestURI()
	err := p.cache.Get(cacheKey, &res)
	if err == nil {
		response.JSON(c, http.StatusOK, res)
		return
	}

	Addresses, pagination, err := p.service.ListAddresses(c, &req)
	if err != nil {
		logger.Error("Failed to get list Address: ", err)
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}

	utils.Copy(&res.Addresses, &Addresses)
	res.Pagination = pagination
	response.JSON(c, http.StatusOK, res)
	_ = p.cache.SetWithExpiration(cacheKey, res, config.AddressCachingTime)
}

// CreateAddress godoc
//
//	@Summary	create Address
//	@Tags		Address
//	@Produce	json
//	@Security	ApiKeyAuth
//	@Param		_	body	dto.CreateAddressReq	true	"Body"
//	@Router		/address [post]
func (p *AddressHandler) CreateAddress(c *gin.Context) {
	var req dto.CreateAddressReq
	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		logger.Error("Failed to get body", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	Address, err := p.service.Create(c, &req)
	if err != nil {
		logger.Error("Failed to create Address", err.Error())
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}

	var res dto.Address
	utils.Copy(&res, &Address)
	response.JSON(c, http.StatusOK, res)
	_ = p.cache.RemovePattern("*Address*")
}

// UpdateAddress godoc
//
//	@Summary	Update Address
//	@Tags		Address
//	@Produce	json
//	@Security	ApiKeyAuth
//	@Param		id	path	string					true	"Address ID"
//	@Param		_	body	dto.UpdateAddressReq	true	"Body"
//	@Router		/address/{id} [put]
func (p *AddressHandler) UpdateAddress(c *gin.Context) {
	AddressId := c.Param("id")
	var req dto.UpdateAddressReq
	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		logger.Error("Failed to get body", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	Address, err := p.service.Update(c, AddressId, &req)
	if err != nil {
		logger.Error("Failed to Update Address", err.Error())
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}

	var res dto.Address
	utils.Copy(&res, &Address)
	response.JSON(c, http.StatusOK, res)
	_ = p.cache.RemovePattern("*Address*")
}

// DeleteAddress godoc
//
//	@Summary	Delete Address
//	@Tags		Address
//	@Produce	json
//	@Security	ApiKeyAuth
//	@Param		id	path	string					true	"Address ID"
//	@Param		_	body	dto.DeleteAddressReq	true	"Body"
//	@Router		/address/{id} [Delete]
func (p *AddressHandler) DeleteAddress(c *gin.Context) {
	AddressId := c.Param("id")
	var req dto.DeleteAddressReq
	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		logger.Error("Failed to get body", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	Address, err := p.service.Delete(c, AddressId, &req)
	if err != nil {
		logger.Error("Failed to Delete Address", err.Error())
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}

	var res dto.Address
	utils.Copy(&res, &Address)
	response.JSON(c, http.StatusOK, res)
	_ = p.cache.RemovePattern("*Address*")
}
