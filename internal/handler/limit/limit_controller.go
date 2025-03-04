package limit

import (
	"asidikfauzi/xyz-multifinance-api/internal/handler/limit/dto"
	"asidikfauzi/xyz-multifinance-api/internal/pkg/constant"
	"asidikfauzi/xyz-multifinance-api/internal/pkg/response"
	"asidikfauzi/xyz-multifinance-api/internal/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type LimitsController struct {
	limitsService LimitsService
}

func NewLimitsController(cs LimitsService) *LimitsController {
	return &LimitsController{
		limitsService: cs,
	}
}

func (cc *LimitsController) ApprovalConsumer(c *gin.Context) {
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

	if role != constant.ADMIN {
		response.Error(c, http.StatusForbidden, constant.AccessDenied.Error(), nil)
		return
	}

	id := c.Param("id")
	idParse, err := uuid.Parse(id)
	if err != nil {
		response.Error(c, http.StatusNotFound, constant.ConsumerNotFound.Error(), nil)
		return
	}

	var req dto.ApprovalLimitInput
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, constant.InvalidJsonPayload.Error(), err.Error())
		return
	}

	req.CreatedBy = userID.(uuid.UUID)

	validate := utils.FormatValidationError(req)
	if len(validate) > 0 {
		response.Error(c, http.StatusUnprocessableEntity, constant.UnprocessableEntity.Error(), validate)
		return
	}

	data, code, err := cc.limitsService.ApprovalConsumer(idParse, req)
	if err != nil {
		response.Error(c, code, err.Error(), nil)
		return
	}

	response.Success(c, code, "successfully approval limit", data)
}
