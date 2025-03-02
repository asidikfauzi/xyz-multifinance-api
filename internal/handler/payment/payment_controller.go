package payment

import (
	"asidikfauzi/xyz-multifinance-api/internal/handler/payment/dto"
	"asidikfauzi/xyz-multifinance-api/internal/pkg/constant"
	"asidikfauzi/xyz-multifinance-api/internal/pkg/response"
	"asidikfauzi/xyz-multifinance-api/internal/pkg/utils"
	"github.com/gin-gonic/gin"
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

func (cc *PaymentsController) Create(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists {
		response.Error(c, http.StatusUnauthorized, constant.TokenInvalid.Error(), nil)
		return
	}

	if role != constant.USER {
		response.Error(c, http.StatusForbidden, constant.AccessDenied.Error(), nil)
		return
	}

	var req dto.PaymentInput
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, constant.InvalidJsonPayload.Error(), err.Error())
		return
	}

	validate := utils.FormatValidationError(req)
	if len(validate) > 0 {
		response.Error(c, http.StatusUnprocessableEntity, constant.UnprocessableEntity.Error(), validate)
		return
	}

	data, code, err := cc.paymentService.Create(req)
	if err != nil {
		response.Error(c, code, err.Error(), nil)
		return
	}

	response.Success(c, code, "successful installment payment", data)
}
