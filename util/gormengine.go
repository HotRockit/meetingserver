package util

import (
	"github.com/jinzhu/gorm"
	"log"
	"meetingserver/model"
)

type GormEngine struct {
	*gorm.DB
}

var GlobalGormEngine *GormEngine = nil

func InitDataBase(cfg *Config) (*GormEngine,error) {
	dataBaseConfig := cfg.DataBaseConfig
	db, err := gorm.Open(dataBaseConfig.Driver,dataBaseConfig.User+":"+dataBaseConfig.Password+
		"@tcp("+dataBaseConfig.Host+":"+dataBaseConfig.Port+")"+"/"+dataBaseConfig.DbName+
		"?charset="+dataBaseConfig.Charset+"&parseTime=True&loc=Local")
	if err != nil {
		log.Print(err.Error())
		panic(err)
		return nil, err
	}

	gormEngine := new(GormEngine)
	gormEngine.DB = db
	GlobalGormEngine = gormEngine


	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Meeting{})
	db.AutoMigrate(&model.MeetingRoom{})
	return gormEngine,nil
}
