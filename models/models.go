package models

import "gorm.io/gorm"

type Payouts struct {
	gorm.Model
	Amount int
	Sender string
	Receiver string
	CRN      int64  `gorm:"default:nextval('crn_seq')"`
}