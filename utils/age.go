package utils

import (
	"strings"
	"time"
)

func CalculateAge(dob string) int {

	if len(dob) >= 10 {
		dob = dob[:10]
	}

	birthDate, err := time.Parse("2006-01-02", strings.TrimSpace(dob))
	if err != nil {
		return 0
	}

	now := time.Now()

	age := now.Year() - birthDate.Year()

	if now.YearDay() < birthDate.YearDay() {
		age--
	}

	return age
}
