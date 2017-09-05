package postgres

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/gorm"
	"sync"
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
)

var (
	once sync.Once
)

func GetPostgresConnection() (db *gorm.DB, err error) {
	once.Do(func() {
		db, err = gorm.Open("postgres", "user=root dbname=gamebotdb sslmode=disable password=root")
		if err != nil {
			loger.Log.Errorf("Error during connection to Postgres has occurred %s", err.Error())
		}else {
			loger.Log.Info("connected to db")
		}
	})
	return db, err
}
