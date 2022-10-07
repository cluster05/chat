package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type configuration struct {
	app      `json:"app"`
	server   `json:"server"`
	database `json:"database"`
}

type app struct {
	Env                        string   `json:"env"`
	JWTSecret                  string   `json:"jwtSecret"`
	TokenExpireDuration        int64    `json:"tokenExpireDuration"`
	RefreshTokenExpireDuration int64    `json:"refreshTokenExpireDuration"`
	AllowOrigin                []string `json:"allowOrigin"`
}

type server struct {
	RunMode      string        `json:"runMode"`
	Port         string        `json:"port"`
	ReadTimeout  time.Duration `json:"readTimeout"`
	WriteTimeout time.Duration `json:"writeTimeout"`
}

type DB struct {
	DSN        string `json:"dns"`
	DBName     string `json:"dbName"`
	DBUsername string `json:"dbUsername"`
	DBPassword string `json:"dbPassword"`
	DBPort     string `json:"dbPort"`
}

type database struct {
	Mongo DB `json:"mongo"`
}

var AppConfig app
var ServerConfig server
var DatabaseConfig database

func Setup(filename string) error {

	var config configuration

	file := fmt.Sprintf("config_%v.json", filename)

	env, err := ioutil.ReadFile(file)

	if err != nil {
		return err
	}

	if err = json.Unmarshal(env, &config); err != nil {
		return err
	}

	AppConfig = config.app
	ServerConfig = config.server
	DatabaseConfig = config.database

	return nil
}
