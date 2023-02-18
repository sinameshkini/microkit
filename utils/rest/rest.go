package rest

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"github.com/go-resty/resty/v2"
	"github.com/sinameshkini/microkit/helpers"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var pool []*RestyClientInfo

func Add(info *RestyClientInfo) (client *resty.Client, err error) {
	if info == nil || info.Name == "" {
		return nil, helpers.ErrInvalidRestClientInfo
	}

	if _, err = url.Parse(info.BaseURL); err != nil {
		return nil, err
	}

	client = resty.New().
		SetBaseURL(info.BaseURL).
		SetAuthToken(info.Token).
		SetHeader("Content-Type", info.ContentType).
		SetDebug(info.Debug)

	info.Client = client

	pool = append(pool, info)

	return client, nil
}

func Get(name string) *resty.Client {
	for _, c := range pool {
		if name == c.Name {
			return c.Client
		}
	}

	return nil
}

type RestyClientInfo struct {
	Name        string        `json:"name,omitempty" mapstructure:"name"`
	BaseURL     string        `json:"base_url,omitempty" mapstructure:"base_url"`
	Token       string        `json:"token,omitempty" mapstructure:"token"`
	ContentType string        `json:"content_type,omitempty" mapstructure:"content_type"`
	Debug       bool          `json:"debug,omitempty" mapstructure:"debug"`
	Client      *resty.Client `json:"-" gorm:"-"`
}

func (c RestyClientInfo) Value() (value driver.Value, err error) {
	var b []byte
	if b, err = json.Marshal(c); err != nil {
		return
	}

	return string(b), nil
}

func (c *RestyClientInfo) Scan(value interface{}) (err error) {
	if value == nil {
		*c = RestyClientInfo{}
		return nil
	}

	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := RestyClientInfo{}
	if err = json.Unmarshal(bytes, &result); err != nil {
		return
	}

	*c = result

	return nil
}

// GormDBDataType gorm db data type
func (RestyClientInfo) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "sqlite":
		return "JSON"
	case "mysql":
		return "JSON"
	case "postgres":
		return "JSONB"
	}
	return ""
}
