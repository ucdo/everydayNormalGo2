package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"runtime"
)

var db *gorm.DB

type Config struct {
	User     string `json:"user"`
	PassWord string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
}

func getConfig() (*Config, error) {
	var conf *Config

	// you cant write file path like ./config
	// can only use runtime to find the file
	// the enter point is which the main func run
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return nil, errors.New("file not exist")
	}

	path := filepath.Join(filepath.Dir(filename), "config.json")

	file, err := os.ReadFile(path)
	if err != nil {
		return conf, err
	}

	err = json.Unmarshal(file, &conf)
	if err != nil {
		return conf, err
	}

	return conf, nil
}

func Conn() (*gorm.DB, error) {
	config, err := getConfig()
	if err != nil {
		panic(err)
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai", config.Host, config.User, config.PassWord, config.Database)
	db, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		return db, err
	}

	return db, nil
}
