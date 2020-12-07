package database

import (
	"fmt"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // dialect mysql for gorm
	"github.com/time-off-system/shared/config"
	Logger "github.com/time-off-system/shared/log"
)

var (
	once   sync.Once
	logger = Logger.NewLog()
)

type (

	// MysqlInterface is an interface that represent mysql methods in package database
	MysqlInterface interface {
		OpenMysqlConn() (*gorm.DB, error)
	}

	// database is a struct to map given struct
	database struct {
		SharedConfig config.ImmutableConfig
	}
)

func (d *database) OpenMysqlConn() (*gorm.DB, error) {
	fmt.Println("Start open mysql connection...")
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local",
		d.SharedConfig.GetDatabaseUserName(),
		d.SharedConfig.GetDatabasePassword(),
		d.SharedConfig.GetDatabaseHost(),
		d.SharedConfig.GetDatabaseName()))

	if err != nil {
		return nil, err
	}

	db.DB().SetMaxIdleConns(64)
	db.DB().SetMaxIdleConns(128)
	db.DB().SetConnMaxLifetime(2 * time.Second)
	db.LogMode(true)
	return db, nil
}

// NewMysql is an factory that implement of mysql database configuration
func NewMysql(config config.ImmutableConfig) MysqlInterface {
	if config == nil {
		panic("[CONFIG] immutable config is required")
	}

	return &database{config}
}
