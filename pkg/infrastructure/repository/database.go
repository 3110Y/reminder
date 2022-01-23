package repository

import "gorm.io/gorm"

type database interface {
	Connection() (*gorm.DB, error)
}
