package utils

import (
	"asidikfauzi/xyz-multifinance-api/internal/config"
	"asidikfauzi/xyz-multifinance-api/internal/pkg/constant"
	"fmt"
	"strings"
	"time"
)

func FormatFieldName(fieldName string) string {
	var formatted strings.Builder
	runes := []rune(fieldName)

	for i, r := range runes {
		if i > 0 && r >= 'A' && r <= 'Z' {
			formatted.WriteRune(' ')
		}
		formatted.WriteRune(r)
	}

	return strings.ToLower(formatted.String())
}

func FormatTimeWithTimezone(utcTime time.Time) (string, error) {
	timezone := config.Env("APP_TIMEZONE")

	location, err := time.LoadLocation(timezone)
	if err != nil {
		return "", fmt.Errorf("%w %v", constant.FailedToLoadTimeZone, err)
	}

	localTime := utcTime.In(location)
	return localTime.Format("02-01-2006 15:04:05"), nil
}

func FormatTime(t time.Time) *string {
	formattedTime, _ := FormatTimeWithTimezone(t)
	return &formattedTime
}

func FormatDefaultString(str *string, defaultValue string) *string {
	if str != nil {
		return str
	}
	return &defaultValue
}
