package consumer

import (
	"asidikfauzi/xyz-multifinance-api/internal/handler/consumer/dto"
	"asidikfauzi/xyz-multifinance-api/internal/pkg/constant"
	"asidikfauzi/xyz-multifinance-api/internal/pkg/response"
	"asidikfauzi/xyz-multifinance-api/internal/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type ConsumersController struct {
	consumersService ConsumersService
}

func NewConsumersController(cs ConsumersService) *ConsumersController {
	return &ConsumersController{
		consumersService: cs,
	}
}

func (cc *ConsumersController) FindAll(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists {
		response.Error(c, http.StatusUnauthorized, constant.TokenInvalid.Error(), nil)
		return
	}

	if role != constant.ADMIN {
		response.Error(c, http.StatusForbidden, constant.AccessDenied.Error(), nil)
		return
	}

	var query dto.QueryConsumer
	if err := c.ShouldBindQuery(&query); err != nil {
		response.Error(c, http.StatusBadRequest, constant.InvalidQueryParameters.Error(), err.Error())
		return
	}

	res, code, err := cc.consumersService.FindAll(query)
	if err != nil {
		response.Error(c, code, err.Error(), nil)
		return
	}

	response.SuccessPaginate(c, code, "successfully get all consumers", res.Data, res.Page)
}

func (cc *ConsumersController) FindById(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists {
		response.Error(c, http.StatusUnauthorized, constant.TokenInvalid.Error(), nil)
		return
	}

	if role != constant.ADMIN && role != constant.USER {
		response.Error(c, http.StatusForbidden, constant.AccessDenied.Error(), nil)
		return
	}

	consumerID, exists := c.Get("consumer_id")
	if !exists {
		response.Error(c, http.StatusUnauthorized, constant.TokenInvalid.Error(), nil)
		return
	}

	id := c.Param("id")
	idParse, err := uuid.Parse(id)
	if err != nil || (role == constant.USER && consumerID.(uuid.UUID) != idParse) {
		response.Error(c, http.StatusNotFound, constant.ConsumerNotFound.Error(), nil)
		return
	}

	res, code, err := cc.consumersService.FindById(idParse)
	if err != nil {
		response.Error(c, code, err.Error(), nil)
		return
	}

	response.Success(c, code, "successfully get consumer by id", res)
}

func (cc *ConsumersController) Create(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		response.Error(c, http.StatusUnauthorized, constant.TokenInvalid.Error(), nil)
		return
	}

	role, exists := c.Get("role")
	if !exists {
		response.Error(c, http.StatusUnauthorized, constant.TokenInvalid.Error(), nil)
		return
	}

	if role != constant.USER {
		response.Error(c, http.StatusForbidden, constant.AccessDenied.Error(), nil)
		return
	}

	var req dto.CreateConsumerInput
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, constant.InvalidJsonPayload.Error(), err.Error())
		return
	}

	req.UserID = userID.(uuid.UUID)

	validate := utils.FormatValidationError(req)
	if len(validate) > 0 {
		response.Error(c, http.StatusUnprocessableEntity, constant.UnprocessableEntity.Error(), validate)
		return
	}

	data, code, err := cc.consumersService.Create(req)
	if err != nil {
		response.Error(c, code, err.Error(), nil)
		return
	}

	response.Success(c, 200, "successfully created consumer", data)
}

func (cc *ConsumersController) Update(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists {
		response.Error(c, http.StatusUnauthorized, constant.TokenInvalid.Error(), nil)
		return
	}

	if role != constant.USER {
		response.Error(c, http.StatusForbidden, constant.AccessDenied.Error(), nil)
		return
	}

	id := c.Param("id")
	idParse, err := uuid.Parse(id)
	if err != nil {
		response.Error(c, http.StatusNotFound, constant.ConsumerNotFound.Error(), nil)
		return
	}

	var req dto.UpdateConsumerInput
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, constant.InvalidJsonPayload.Error(), err.Error())
		return
	}

	validate := utils.FormatValidationError(req)
	if len(validate) > 0 {
		response.Error(c, http.StatusUnprocessableEntity, constant.UnprocessableEntity.Error(), validate)
		return
	}

	data, code, err := cc.consumersService.Update(idParse, req)
	if err != nil {
		response.Error(c, code, err.Error(), nil)
		return
	}

	response.Success(c, 200, "successfully updated consumer", data)
}
