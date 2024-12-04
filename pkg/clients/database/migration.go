package database

import (
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

func Drop(db *gorm.DB, tables []interface{}) (err error) {
	for _, t := range tables {
		if err = db.Migrator().DropTable(t); err != nil {
			log.Error(err.Error())
		}
	}

	return nil
}

func Migrate(db *gorm.DB, tables []interface{}) (err error) {
	if err = db.AutoMigrate(tables...); err != nil {
		return err
	}

	return nil
}
