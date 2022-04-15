package config

import (
	"io/fs"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server ServerConfig `yaml:"server"`
	Db     DbConfig     `yaml:"db"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

type DbConfig struct {
	DBname   string `yaml:"dbname"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
}

func (c *Config) SetDefault() {
	c.Server.Port = "8080"
	c.Db.DBname = "expenditure"
	c.Db.Port = "5432"
	c.Db.Username = "postgres"
}

func LoadConfig() Config {
	var conf Config
	conf.SetDefault()
	data, err := ioutil.ReadFile("config.yaml")
	if os.IsNotExist(err) {

		if data, err := yaml.Marshal(conf); err == nil {
			_ = ioutil.WriteFile("config.yaml", data, fs.ModePerm)
		}
	} else {
		_ = yaml.Unmarshal(data, &conf)
	}
	return conf
}
