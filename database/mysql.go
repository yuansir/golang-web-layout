package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang-echo-layout/config"
	"sync"
)

var (
	// Mysql DB
	Mysql *gorm.DB
	once  sync.Once
)

// NewMysql new mysql connection
func NewMysql() *gorm.DB {
	once.Do(func() {
		db, err := gorm.Open("mysql", config.Conf.Mysql.Connect)
		if err != nil {
			panic(err)
		}
		Mysql = db
		Mysql.DB().SetMaxIdleConns(config.Conf.Mysql.MaxIdle)
		Mysql.DB().SetMaxOpenConns(config.Conf.Mysql.MaxOpen)
	})
	return Mysql
}
