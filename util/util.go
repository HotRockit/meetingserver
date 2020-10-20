package util

import (
	"bufio"
	"encoding/json"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)

//Config is a struct to describe the config of app
type Config struct {
	AppName        string         `json:"app_name"`
	AppMode        string         `json:"app_mode"`
	AppHost        string         `json:"app_host"`
	AppPort        string         `json:"app_port"`
	DataBaseConfig DataBaseConfig `json:"database"`
}

//DataBaseConfig is a config of database
type DataBaseConfig struct {
	Driver   string `json:"driver"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DbName   string `json:"db_name"`
	Charset  string `json:"charset"`
}

var _cfg *Config = nil

//ParseConfig is a function to parse the config of this app
func ParseConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Print(err.Error())
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&_cfg)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	return _cfg, nil
}


