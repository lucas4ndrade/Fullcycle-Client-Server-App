package database

import "gorm.io/gorm"

type Bid struct {
	gorm.Model
	Value string
}
