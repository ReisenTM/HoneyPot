package config

import "fmt"

type Config struct {
	DB        DB       `yaml:"db"`
	Logger    Logger   `yaml:"logger"`
	Redis     Redis    `yaml:"redis"`
	System    System   `yaml:"system"`
	Jwt       Jwt      `yaml:"jwt"`
	WhiteList []string `yaml:"whiteList"`
	VsNet     VsNet    `yaml:"vsNet"`
	//MQ        MQ       `yaml:"mq"`
}

// DB 数据库设置
type DB struct {
	DbName          string `yaml:"db_name"`
	Host            string `yaml:"host"`
	Port            int    `yaml:"port"`
	User            string `yaml:"user"`
	Password        string `yaml:"password"`
	MaxIdleConns    int    `yaml:"maxIdleConns"`
	MaxOpenConns    int    `yaml:"maxOpenConns"`
	ConnMaxLifetime int    `yaml:"connMaxLifetime"`
}

func (cfg DB) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
	)
}

// Logger 日志输出格式设置
type Logger struct {
	Format  string `yaml:"format"`
	Level   string `yaml:"level"`
	AppName string `yaml:"appName"`
}

// Redis  配置
type Redis struct {
	Addr     string
	Password string
	DB       int
}

type System struct {
	WebAddr string `yaml:"webAddr"`
	//GrpcAddr string `yaml:"grpcAddr"`
	Mode string `yaml:"mode"`
}

type Jwt struct {
	Expires int    `yaml:"expires"` // 单位为秒
	Issuer  string `yaml:"issuer"`
	Secret  string `yaml:"secret"`
}
type VsNet struct {
	Name   string `yaml:"name" json:"name"`
	Prefix string `yaml:"prefix" json:"prefix"`
	Net    string `yaml:"net" json:"net"`
}
