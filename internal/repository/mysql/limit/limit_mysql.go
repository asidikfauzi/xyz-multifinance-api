package limit

import "asidikfauzi/xyz-multifinance-api/internal/model"

type LimitsMySQL interface {
	ApprovalConsumer(model.Consumers, model.Limits) (model.Limits, error)
}
