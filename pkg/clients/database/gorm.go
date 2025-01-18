package database

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type impl struct {
	db *gorm.DB
}

func (i *impl) DB() *gorm.DB {
	return i.db
}

func (i *impl) Close() error {
	db, err := i.db.DB()
	if err != nil {
		return err
	}

	return db.Close()
}

func NewSQLite(path string, debug bool) (Database, error) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{Logger: getLogger(debug)})
	if err != nil {
		return nil, err
	}

	return &impl{db: db}, nil
}

func NewDBWithConf(conf *Config) (Database, error) {
	db, err := gorm.Open(postgres.Open(conf.dsn()), &gorm.Config{Logger: getLogger(conf.Debug)})
	if err != nil {
		return nil, err
	}

	return &impl{db: db}, nil
}

func NewDBWithDsn(dsn string, debug bool) (Database, error) {
	sqlDB, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: sqlDB}), &gorm.Config{Logger: getLogger(debug)})
	if err != nil {
		return nil, err
	}

	return &impl{db: db}, nil
}

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	Schema   string
	Debug    bool
}

func (c *Config) dsn() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran",
		c.Host,
		c.Username,
		c.Password,
		c.Schema,
		c.Port,
	)
}

func getLogger(debug bool) logger.Interface {
	logLevel := logger.Error
	if debug {
		logLevel = logger.Info
	}

	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logLevel,    // Log level
			Colorful:      true,        // Disable color
		},
	)
}
