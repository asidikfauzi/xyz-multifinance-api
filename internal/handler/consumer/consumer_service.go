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
	"math"
	"net/http"
	"time"
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
			ID:              c.ID,
			Email:           c.User.Email,
			NIK:             utils.FormatDefaultString(c.NIK, ""),
			FullName:        c.FullName,
			LegalName:       c.LegalName,
			Phone:           c.Phone,
			PlaceOfBirth:    c.PlaceOfBirth,
			DateOfBirth:     c.DateOfBirth,
			Salary:          c.Salary,
			KtpImage:        c.KTPImage,
			SelfieImage:     c.SelfieImage,
			IsVerified:      c.IsVerified,
			RejectionReason: c.RejectionReason,
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
		Email:           consumerData.User.Email,
		NIK:             utils.FormatDefaultString(consumerData.NIK, ""),
		FullName:        consumerData.FullName,
		LegalName:       consumerData.LegalName,
		Phone:           consumerData.Phone,
		PlaceOfBirth:    consumerData.PlaceOfBirth,
		DateOfBirth:     consumerData.DateOfBirth,
		Salary:          consumerData.Salary,
		KtpImage:        consumerData.KTPImage,
		SelfieImage:     consumerData.SelfieImage,
		IsVerified:      consumerData.IsVerified,
		RejectionReason: consumerData.RejectionReason,
		CreatedAt:       utils.FormatTime(consumerData.CreatedAt),
		UpdatedAt:       utils.FormatTime(consumerData.UpdatedAt),
	}

	if len(consumerData.Limits) > 0 {
		res.Limit = &dto2.LimitResponse{
			ID:             consumerData.Limits[0].ID,
			LimitAvailable: consumerData.Limits[0].LimitAvailable,
			CreatedAt:      utils.FormatTime(consumerData.Limits[0].CreatedAt),
			UpdatedAt:      utils.FormatTime(consumerData.Limits[0].UpdatedAt),
		}
	}

	return res, http.StatusOK, nil
}

func (c *consumerService) Update(id uuid.UUID, input dto.UpdateConsumerInput) (res dto.ConsumerResponse, code int, err error) {
	checkConsumer, err := c.consumerMySQL.FindById(id)
	if err != nil {
		if errors.Is(err, constant.ConsumerNotFound) {
			return res, http.StatusNotFound, constant.ConsumerNotFound
		}
		return res, http.StatusInternalServerError, err
	}

	if checkConsumer.NIK != nil && *checkConsumer.NIK != input.NIK {
		_, err = c.consumerMySQL.FindByNIK(input.NIK)
		if err == nil {
			return res, http.StatusConflict, constant.NIKConsumerAlreadyExists
		}
	}

	consumerData := model.Consumers{
		ID:           id,
		NIK:          &input.NIK,
		FullName:     input.FullName,
		LegalName:    input.LegalName,
		Phone:        input.Phone,
		PlaceOfBirth: input.PlaceOfBirth,
		DateOfBirth:  input.DateOfBirth,
		Salary:       math.Round(input.Salary*100) / 100,
		KTPImage:     input.KtpImage,
		SelfieImage:  input.SelfieImage,
		UpdatedAt:    time.Now(),
	}

	editConsumer, err := c.consumerMySQL.Update(consumerData)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	res = dto.ConsumerResponse{
		ID:           editConsumer.ID,
		Email:        editConsumer.User.Email,
		NIK:          utils.FormatDefaultString(editConsumer.NIK, ""),
		FullName:     editConsumer.FullName,
		LegalName:    editConsumer.LegalName,
		Phone:        editConsumer.Phone,
		PlaceOfBirth: editConsumer.PlaceOfBirth,
		DateOfBirth:  editConsumer.DateOfBirth,
		Salary:       editConsumer.Salary,
		KtpImage:     editConsumer.KTPImage,
		SelfieImage:  editConsumer.SelfieImage,
		UpdatedAt:    utils.FormatTime(editConsumer.UpdatedAt),
	}

	return res, http.StatusOK, nil
}
