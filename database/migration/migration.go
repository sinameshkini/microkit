package migration

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Model struct {
	Name  string
	Model interface{}
	Seeds interface{}
}

func DoMigration(db *gorm.DB, models []*Model) (err error) {
	forceDrop := viper.GetBool("migration.drop_all")
	for _, m := range models {
		if err = DoTableMigration(db, m.Model, m.Seeds, forceDrop, m.Name); err != nil {
			logrus.Error(err)
		}
	}

	return err
}

func DoTableMigration(db *gorm.DB, table, seeds interface{}, forceDrop bool, name string) (err error) {
	var (
		migrate = viper.GetBool("migration.models." + name + ".migrate")
		drop    = viper.GetBool("migration.models."+name+".drop") || forceDrop
		seed    = viper.GetBool("migration.models." + name + ".seed")
	)

	if !migrate {
		return nil
	}

	if drop {
		if err = db.Migrator().DropTable(table); err != nil {
			return err
		}

		logrus.Info(name + " table drop success")
	}

	if err = db.AutoMigrate(table); err != nil {
		logrus.Error("table"+name+":", err)
		return
	}

	logrus.Info(name + " table migrate success")

	if seed {
		var c int64
		if err = db.Find(table).Count(&c).Error; err != nil {
			return err
		}

		if c > 0 {
			logrus.Warningln(name + " table record exist")
		} else {
			if err = db.Model(table).Create(seeds).Error; err != nil {
				return err
			}

			logrus.Info(name + " table seed success")
		}
	}

	return
}
