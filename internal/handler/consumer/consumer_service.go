package consumer

import (
	"asidikfauzi/xyz-multifinance-api/internal/handler/consumer/dto"
	dto2 "asidikfauzi/xyz-multifinance-api/internal/handler/limit/dto"
	"asidikfauzi/xyz-multifinance-api/internal/model"
	"asidikfauzi/xyz-multifinance-api/internal/pkg/constant"
	"asidikfauzi/xyz-multifinance-api/internal/pkg/response"
	"asidikfauzi/xyz-multifinance-api/internal/pkg/utils"
	"asidikfauzi/xyz-multifinance-api/internal/repository/mysql/consumer"
	"errors"
	"github.com/google/uuid"
	"net/http"
)

type consumerService struct {
	consumerMySQL consumer.ConsumersMySQL
}

func NewConsumersService(cm consumer.ConsumersMySQL) ConsumersService {
	return &consumerService{
		consumerMySQL: cm,
	}
}

func normalizeCategoryQuery(q dto.QueryConsumer) dto.QueryConsumer {
	if q.Page == 0 {
		q.Page = 1
	}
	if q.Limit == 0 {
		q.Limit = 10
	}
	if q.Paginate == "" || (q.Paginate != "false" && q.Paginate != "true") {
		q.Paginate = "true"
	}
	return q
}

func (c *consumerService) FindAll(q dto.QueryConsumer) (res dto.ConsumersResponseWithPage, code int, err error) {
	q = normalizeCategoryQuery(q)

	consumers, totalItems, err := c.consumerMySQL.FindAll(q)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	for _, c := range consumers {
		res.Data = append(res.Data, dto.ConsumerResponse{
			ID:           c.ID,
			NIK:          c.NIK,
			FullName:     c.FullName,
			LegalName:    c.LegalName,
			Phone:        c.Phone,
			PlaceOfBirth: c.PlaceOfBirth,
			DateOfBirth:  c.DateOfBirth,
			Salary:       c.Salary,
			KTPImage:     c.KTPImage,
			SelfieImage:  c.SelfieImage,
			IsVerified:   c.IsVerified,
		})
	}

	res.Page = response.PaginationResponse{
		TotalItems:   totalItems,
		ItemCount:    len(res.Data),
		ItemsPerPage: q.Limit,
		CurrentPage:  q.Page,
	}

	return res, http.StatusOK, nil
}

func (c *consumerService) FindById(id uuid.UUID) (res dto.ConsumerResponse, code int, err error) {
	consumerData, err := c.consumerMySQL.FindById(id)
	if err != nil {
		if errors.Is(err, constant.ConsumerNotFound) {
			return res, http.StatusNotFound, constant.ConsumerNotFound
		}
		return res, http.StatusInternalServerError, err
	}

	res = dto.ConsumerResponse{
		ID:              consumerData.ID,
		NIK:             consumerData.NIK,
		FullName:        consumerData.FullName,
		LegalName:       consumerData.LegalName,
		Phone:           consumerData.Phone,
		PlaceOfBirth:    consumerData.PlaceOfBirth,
		DateOfBirth:     consumerData.DateOfBirth,
		Salary:          consumerData.Salary,
		KTPImage:        consumerData.KTPImage,
		SelfieImage:     consumerData.SelfieImage,
		IsVerified:      consumerData.IsVerified,
		RejectionReason: consumerData.RejectionReason,
		CreatedAt:       utils.FormatTime(consumerData.CreatedAt),
		UpdatedAt:       utils.FormatTime(consumerData.UpdatedAt),
	}

	if len(consumerData.Limits) > 0 {
		res.Limit = &dto2.LimitResponse{
			ID:             consumerData.Limits[0].ID,
			Tenor:          consumerData.Limits[0].Tenor,
			LimitAvailable: consumerData.Limits[0].LimitAvailable,
			CreatedAt:      utils.FormatTime(consumerData.Limits[0].CreatedAt),
			UpdatedAt:      utils.FormatTime(consumerData.Limits[0].UpdatedAt),
		}
	}

	return res, http.StatusOK, nil
}

func (c *consumerService) Create(input dto.CreateConsumerInput) (res dto.ConsumerResponse, code int, err error) {
	_, err = c.consumerMySQL.FindByNIK(input.NIK)
	if err == nil {
		return res, http.StatusConflict, constant.NIKConsumerAlreadyExists
	}

	consumerData := model.Consumers{
		ID:           uuid.New(),
		NIK:          input.NIK,
		FullName:     input.FullName,
		LegalName:    input.LegalName,
		Phone:        input.Phone,
		PlaceOfBirth: input.PlaceOfBirth,
		DateOfBirth:  input.DateOfBirth,
		Salary:       input.Salary,
		KTPImage:     input.KTPImage,
		SelfieImage:  input.SelfieImage,
		UserID:       input.UserID,
	}

	newConsumer, err := c.consumerMySQL.Create(consumerData)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	res = dto.ConsumerResponse{
		ID:           newConsumer.ID,
		NIK:          newConsumer.NIK,
		FullName:     newConsumer.FullName,
		LegalName:    newConsumer.LegalName,
		Phone:        newConsumer.Phone,
		PlaceOfBirth: newConsumer.PlaceOfBirth,
		DateOfBirth:  newConsumer.DateOfBirth,
		Salary:       newConsumer.Salary,
		KTPImage:     newConsumer.KTPImage,
		SelfieImage:  newConsumer.SelfieImage,
		CreatedAt:    utils.FormatTime(newConsumer.CreatedAt),
	}

	return res, http.StatusCreated, nil
}

func (c *consumerService) Update(id uuid.UUID, input dto.UpdateConsumerInput) (res dto.ConsumerResponse, code int, err error) {
	checkConsumer, err := c.consumerMySQL.FindById(id)
	if err != nil {
		if errors.Is(err, constant.ConsumerNotFound) {
			return res, http.StatusNotFound, constant.ConsumerNotFound
		}
		return res, http.StatusInternalServerError, err
	}

	if checkConsumer.NIK != input.NIK {
		_, err = c.consumerMySQL.FindByNIK(input.NIK)
		if err == nil {
			return res, http.StatusConflict, constant.NIKConsumerAlreadyExists
		}
	}

	consumerData := model.Consumers{
		ID:           id,
		NIK:          input.NIK,
		FullName:     input.FullName,
		LegalName:    input.LegalName,
		Phone:        input.Phone,
		PlaceOfBirth: input.PlaceOfBirth,
		DateOfBirth:  input.DateOfBirth,
		Salary:       input.Salary,
		KTPImage:     input.KTPImage,
		SelfieImage:  input.SelfieImage,
	}

	editConsumer, err := c.consumerMySQL.Update(consumerData)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	res = dto.ConsumerResponse{
		ID:           editConsumer.ID,
		NIK:          editConsumer.NIK,
		FullName:     editConsumer.FullName,
		LegalName:    editConsumer.LegalName,
		Phone:        editConsumer.Phone,
		PlaceOfBirth: editConsumer.PlaceOfBirth,
		DateOfBirth:  editConsumer.DateOfBirth,
		Salary:       editConsumer.Salary,
		KTPImage:     editConsumer.KTPImage,
		SelfieImage:  editConsumer.SelfieImage,
		CreatedAt:    utils.FormatTime(editConsumer.CreatedAt),
		UpdatedAt:    utils.FormatTime(editConsumer.UpdatedAt),
	}

	return res, http.StatusOK, nil
}
