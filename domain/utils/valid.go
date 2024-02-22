package utils

import (
	"api/domain/user"
	"time"
)

func Valid(user *user.Users) bool {

	if len(user.Name) == 0 && user.Name == "" {
		return false
	}

	if len(user.Email) == 0 && user.Email == "" {
		return false
	}

	if _, err := time.Parse("2006-01-02", user.Age); err != nil {
		return false
	}

	return true
}
