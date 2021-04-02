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
	MaxOpenConn     int           `json:"-"`
	MaxIdleConn     int           `json:"-"`
	ConnMaxLifeTime time.Duration `json:"-"`
}

// Config representation
type ServerConfig struct {
	Port    int    `json:"port"`
}
