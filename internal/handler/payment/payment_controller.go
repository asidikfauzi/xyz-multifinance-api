package payment

import (
	"asidikfauzi/xyz-multifinance-api/internal/handler/payment/dto"
	"asidikfauzi/xyz-multifinance-api/internal/pkg/constant"
	"asidikfauzi/xyz-multifinance-api/internal/pkg/response"
	"asidikfauzi/xyz-multifinance-api/internal/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type PaymentsController struct {
	paymentService PaymentsService
}

func NewPaymentsController(ps PaymentsService) *PaymentsController {
	return &PaymentsController{
		paymentService: ps,
	}
}

func (cc *PaymentsController) FindAll(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists {
		response.Error(c, http.StatusUnauthorized, constant.TokenInvalid.Error(), nil)
		return
	}

	if role != constant.ADMIN && role != constant.USER {
		response.Error(c, http.StatusForbidden, constant.AccessDenied.Error(), nil)
		return
	}

	var query dto.QueryPayment
	if err := c.ShouldBindQuery(&query); err != nil {
		response.Error(c, http.StatusBadRequest, constant.InvalidQueryParameters.Error(), err.Error())
		return
	}

	consumerId, exists := c.Get("consumer_id")
	if !exists {
		response.Error(c, http.StatusUnauthorized, constant.TokenInvalid.Error(), nil)
		return
	}

	if role == constant.USER {
		query.ConsumerId = consumerId.(uuid.UUID).String()
	}

	res, code, err := cc.paymentService.FindAll(query)
	if err != nil {
		response.Error(c, code, err.Error(), nil)
		return
	}

	response.SuccessPaginate(c, code, "successfully get all payments", res.Data, res.Page)
}

func (cc *PaymentsController) Pay(c *gin.Context) {
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

	var req dto.PaymentInput
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, constant.InvalidJsonPayload.Error(), err.Error())
		return
	}

	userId, exists := c.Get("user_id")
	if !exists {
		response.Error(c, http.StatusUnauthorized, constant.TokenInvalid.Error(), nil)
		return
	}

	consumerId, exists := c.Get("consumer_id")
	if !exists {
		response.Error(c, http.StatusUnauthorized, constant.TokenInvalid.Error(), nil)
		return
	}

	req.ConsumerId = consumerId.(uuid.UUID)
	req.UpdatedBy = userId.(uuid.UUID)

	validate := utils.FormatValidationError(req)
	if len(validate) > 0 {
		response.Error(c, http.StatusUnprocessableEntity, constant.UnprocessableEntity.Error(), validate)
		return
	}

	data, code, err := cc.paymentService.Pay(idParse, req)
	if err != nil {
		response.Error(c, code, err.Error(), nil)
		return
	}

	response.Success(c, code, "successful installment payment", data)
}
