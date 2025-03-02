package limit

import (
	"asidikfauzi/xyz-multifinance-api/internal/handler/limit/dto"
	"github.com/google/uuid"
)

type LimitsService interface {
	ApprovalConsumer(uuid.UUID, dto.ApprovalLimitInput) (dto.LimitResponse, int, error)
}
