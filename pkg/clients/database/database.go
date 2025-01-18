package database

import "gorm.io/gorm"

type Database interface {
	DB() *gorm.DB
	Close() error
}
