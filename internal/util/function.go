package util

import (
	"net/mail"
	"time"
)

func TimeOverlap(f1, t1, f2, t2 time.Time) bool {
	return !(t1.Before(f2) || f1.After(t2))
}

func ValidateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
