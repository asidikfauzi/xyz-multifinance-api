package limit

import (
	"asidikfauzi/xyz-multifinance-api/internal/handler/limit/dto"
	"asidikfauzi/xyz-multifinance-api/internal/model"
	"asidikfauzi/xyz-multifinance-api/internal/pkg/constant"
	"asidikfauzi/xyz-multifinance-api/internal/pkg/utils"
	"asidikfauzi/xyz-multifinance-api/internal/repository/mysql/consumer"
	"asidikfauzi/xyz-multifinance-api/internal/repository/mysql/limit"
	"errors"
	"github.com/google/uuid"
	"net/http"
)

type limitService struct {
	limitMySQL    limit.LimitsMySQL
	consumerMySQL consumer.ConsumersMySQL
}

func NewLimitsService(lm limit.LimitsMySQL, cm consumer.ConsumersMySQL) LimitsService {
	return &limitService{
		limitMySQL:    lm,
		consumerMySQL: cm,
	}
}

func (c *limitService) ApprovalConsumer(id uuid.UUID, input dto.ApprovalLimitInput) (res dto.LimitResponse, code int, err error) {
	consumerData, err := c.consumerMySQL.FindById(id)
	if err != nil {
		if errors.Is(err, constant.ConsumerNotFound) {
			return res, http.StatusNotFound, constant.ConsumerNotFound
		}
		return res, http.StatusInternalServerError, err
	}

	consumerData.IsVerified = input.IsVerified

	if !input.IsVerified && input.RejectionReason != "" {
		consumerData.RejectionReason = input.RejectionReason

		_, err := c.consumerMySQL.Update(consumerData)
		if err != nil {
			return res, http.StatusInternalServerError, err
		}

		return res, http.StatusOK, nil
	}

	limitData := model.Limits{
		ID:             uuid.New(),
		Tenor:          input.Tenor,
		LimitAvailable: input.LimitAvailable,
		ConsumerID:     consumerData.ID,
		CreatedBy:      input.CreatedBy,
	}

	consumerData.RejectionReason = ""

	newLimit, err := c.limitMySQL.ApprovalConsumer(consumerData, limitData)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	res = dto.LimitResponse{
		ID:             newLimit.ID,
		Tenor:          newLimit.Tenor,
		LimitAvailable: newLimit.LimitAvailable,
		CreatedAt:      utils.FormatTime(newLimit.CreatedAt),
	}

	return res, http.StatusOK, nil
}
