package connection

import (
	"errors"

	"github.com/sinameshkini/microkit/utils/enums"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	ErrConnectionLost    = errors.New("database connection lost")
	ErrInvalidConnection = errors.New("invalid database connection")
	ErrInvalidDriver     = errors.New("invalid driver selected")

	conn *DBConnection
)

func Init(c *DBConnection) (err error) {
	if err = Connect(c); err != nil {
		return
	}

	conn = c

	return nil
}

func GetInstance() (dbConnrction *DBConnection, err error) {
	if conn != nil {
		return conn, nil
	}

	return nil, ErrConnectionLost
}

type DBConnection struct {
	Name   string               `json:"name,omitempty"`
	Driver enums.DatabaseDriver `json:"driver,omitempty"`
	Host   string               `json:"host,omitempty"`
	Port   string               `json:"port,omitempty"`
	User   string               `json:"user,omitempty"`
	Pass   string               `json:"pass,omitempty"`
	DBName string               `json:"db_name,omitempty" mapstructure:"db_name"`
	Debug  bool                 `json:"debug,omitempty"`
	DB     *gorm.DB             `json:"db,omitempty"`
}

func (c *DBConnection) isValid() bool {
	return c != nil &&
		c.Name != ""
}

func Connect(c *DBConnection) (err error) {
	logrus.Infoln(
		"connecting to database with",
		"driver:", c.Driver,
		"name:", c.Name,
		"host:", c.Host,
		"port:", c.Port,
		"user:", c.User,
		"dn_name:", c.DBName,
	)

	if !c.isValid() {
		return ErrInvalidConnection
	}

	switch c.Driver {
	case enums.PostgresSQL:
		if c.DB, err = newPsql(c); err != nil {
			return
		}

	case enums.SQLServer:
		if c.DB, err = newMssql(c); err != nil {
			return
		}

	default:
		return ErrInvalidDriver
	}

	logrus.Infoln("database connected successfully!")

	return
}

// use in unit tests
func ConnectionTest() (testConn *DBConnection) {
	testConn = new(DBConnection)

	if err := viper.UnmarshalKey("database.test", testConn); err != nil {
		logrus.Errorln(err)
		return
	}

	if err := Connect(testConn); err != nil {
		logrus.Errorln(err)
		return
	}

	return
}
