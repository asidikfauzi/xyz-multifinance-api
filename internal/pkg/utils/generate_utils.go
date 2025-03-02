package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateContractNumber() string {
	today := time.Now().Format("20060102")

	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	randomStr := make([]byte, 6)
	for i := range randomStr {
		randomStr[i] = charset[rand.Intn(len(charset))]
	}

	contractNumber := fmt.Sprintf("CN-%s-%s", today, string(randomStr))

	return contractNumber
}
