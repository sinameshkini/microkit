package connection

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func newMssql(dbConn *DBConnection) (db *gorm.DB, err error) {
	var newLogger logger.Interface
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		dbConn.User,
		dbConn.Pass,
		dbConn.Host,
		dbConn.Port,
		dbConn.DBName,
	)

	if dbConn.Debug {
		newLogger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: time.Second, // Slow SQL threshold
				LogLevel:      logger.Info, // Log level
				Colorful:      true,        // Disable color
				// IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			},
		)
	}

	return gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
}
