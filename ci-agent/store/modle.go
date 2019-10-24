package store

import "time"

type Build struct {
	Name      string
	Branch    string
	Status    int
	Id        int
	Git       string
	Timestamp time.Time
	Image     string
	Token     string
	Addr      string
	User      string
	Sha       string
	Message   string
}

const (
	BuildReady  = 0
	BuildFailed = -2
	BuildEnv    = 2
	Building    = 3
)
