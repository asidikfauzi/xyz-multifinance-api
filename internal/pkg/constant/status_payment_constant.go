package constant

type StatusPayment string

const (
	PENDING StatusPayment = "PENDING"
	SUCCESS StatusPayment = "SUCCESS"
	FAILED  StatusPayment = "FAILED"
)
