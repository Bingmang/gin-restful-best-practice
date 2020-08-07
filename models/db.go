package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"gin-restful-best-practice/conf"
	"log"
)

var DB *gorm.DB

func init() {
	var err error
	config := conf.Conf()
	log.Printf("Connecting to database %s:%d...", config.DB_HOST, config.DB_PORT)
	DB, err = gorm.Open("postgres", config.DB_CONN_STR)
	if err != nil {
		log.Panicln(err)
	}
	DB.SingularTable(true)
	InitTableRole()
	InitTableUser()
	InitTableModel()
}
