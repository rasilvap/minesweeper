package model

import "time"

type Config struct {
	Env      string       `json:"-"`
	Version  string       `json:"string"`
	Server   ServerConfig `json:"server"`
	Database DbConfig     `json:"db"`
}

type DbConfig struct {
	Server          string        `json:"server"`
	Port            int           `json:"port"`
	User            string        `json:"user"`
	Password        string        `json:"password"`
	Database        string        `json:"database"`
	MaxOpenConn     int           `json:"maxOpenConnections"`
	MaxIdleConn     int           `json:"maxIdleConnections"`
	ConnMaxLifeTime time.Duration `json:"connectionMaxLifetime"`
}

// Config representation
type ServerConfig struct {
	Port int `json:"port"`
}
