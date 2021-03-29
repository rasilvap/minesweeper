package model

import "time"

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
