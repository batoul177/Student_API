package GORM_model

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	FirstName string
	LastName  string
}
