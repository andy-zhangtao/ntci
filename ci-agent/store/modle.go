package store

import "time"

type Build struct {
	Name       string
	Branch     string
	Status     int
	Id         int
	Git        string
	Timestamp  time.Time
	Image      string
	Token      string
	Addr       string
	User       string
	Sha        string
	Message    string
	Language   string
	Lanversion string
	Namespace  string
}

const (
	BuildReady    = 0
	BuildEnv      = 2
	Building      = 3
	BuildSuccess  = 4
	BuildFailed   = -4
	ProcessFailed = -5
)
