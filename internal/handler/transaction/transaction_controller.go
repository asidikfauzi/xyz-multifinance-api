package transaction

import (
	"asidikfauzi/xyz-multifinance-api/internal/handler/transaction/dto"
	"asidikfauzi/xyz-multifinance-api/internal/pkg/constant"
	"asidikfauzi/xyz-multifinance-api/internal/pkg/response"
	"asidikfauzi/xyz-multifinance-api/internal/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type TransactionsController struct {
	transactionService TransactionsService
}

func NewTransactionsController(ts TransactionsService) *TransactionsController {
	return &TransactionsController{
		transactionService: ts,
	}
}

func (cc *TransactionsController) FindAll(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists {
		response.Error(c, http.StatusUnauthorized, constant.TokenInvalid.Error(), nil)
		return
	}

	if role != constant.ADMIN && role != constant.USER {
		response.Error(c, http.StatusForbidden, constant.AccessDenied.Error(), nil)
		return
	}

	var query dto.QueryTransaction
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

	res, code, err := cc.transactionService.FindAll(query)
	if err != nil {
		response.Error(c, code, err.Error(), nil)
		return
	}

	response.SuccessPaginate(c, code, "successfully get all consumers", res.Data, res.Page)
}

func (cc *TransactionsController) Transactions(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists {
		response.Error(c, http.StatusUnauthorized, constant.TokenInvalid.Error(), nil)
		return
	}

	if role != constant.USER {
		response.Error(c, http.StatusForbidden, constant.AccessDenied.Error(), nil)
		return
	}

	var req dto.TransactionInput
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, constant.InvalidJsonPayload.Error(), err.Error())
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		response.Error(c, http.StatusUnauthorized, constant.TokenInvalid.Error(), nil)
		return
	}

	consumerID, exists := c.Get("consumer_id")
	if !exists {
		response.Error(c, http.StatusUnauthorized, constant.TokenInvalid.Error(), nil)
		return
	}

	if consumerID == uuid.Nil {
		response.Error(c, http.StatusNotFound, constant.ConsumerNotFound.Error(), nil)
	}

	req.ConsumerID = consumerID.(uuid.UUID)
	req.CreatedBy = userID.(uuid.UUID)

	validate := utils.FormatValidationError(req)
	if len(validate) > 0 {
		response.Error(c, http.StatusUnprocessableEntity, constant.UnprocessableEntity.Error(), validate)
		return
	}

	data, code, err := cc.transactionService.Transaction(req)
	if err != nil {
		response.Error(c, code, err.Error(), nil)
		return
	}

	response.Success(c, code, "transaction successfully", data)
}
