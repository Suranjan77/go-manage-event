package domain

import "gorm.io/gorm"

type LoginDetails struct {
	gorm.Model
	Email           string
	TotalLogins     uint64
	LastIpUsed      string
	LastBrowserUsed string
}
