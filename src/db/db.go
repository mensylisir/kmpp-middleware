package db

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/mensylisir/kmpp-middleware/src/util/encrypt"
)

var DB *gorm.DB

const phaseName = "db"

type InitDBPhase struct {
	Host         string
	Port         int
	Name         string
	User         string
	Password     string
	MaxOpenConns int
	MaxIdleConns int
}

func (i *InitDBPhase) Init() error {
	aesPasswd, er1 := encrypt.StringEncrypt(i.Password)
	if er1 != nil {
		return er1
	}
	p, err := encrypt.StringDecrypt(aesPasswd)
	if err != nil {
		return err
	}
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Asia%%2FShanghai",
		i.User,
		p,
		i.Host,
		i.Port,
		i.Name)
	db, err := gorm.Open("mysql", url)
	if err != nil {
		return err
	}

	gorm.DefaultTableNameHandler = func(DB *gorm.DB, defaultTableName string) string {
		return "rdev_" + defaultTableName
	}
	db.SingularTable(true)
	db.DB().SetMaxOpenConns(i.MaxOpenConns)
	db.DB().SetMaxIdleConns(i.MaxIdleConns)
	db.DB().SetConnMaxLifetime(time.Hour)
	DB = db
	DB.LogMode(false)
	return nil
}

func (i *InitDBPhase) PhaseName() string {
	return phaseName
}
