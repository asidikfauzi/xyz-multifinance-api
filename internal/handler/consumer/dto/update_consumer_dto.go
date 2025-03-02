package dto

type UpdateConsumerInput struct {
	NIK          string  `json:"nik" validate:"required,len=16,numeric"`
	FullName     string  `json:"full_name" validate:"required"`
	LegalName    string  `json:"legal_name" validate:"required"`
	Phone        string  `json:"phone" validate:"required,numeric"`
	PlaceOfBirth string  `json:"place_of_birth" validate:"required"`
	DateOfBirth  string  `json:"date_of_birth" validate:"required,datetime=02-01-2006"`
	Salary       float64 `json:"salary" validate:"required"`
	KtpImage     string  `json:"ktp_image" validate:"required"`
	SelfieImage  string  `json:"selfie_image" validate:"required"`
}
