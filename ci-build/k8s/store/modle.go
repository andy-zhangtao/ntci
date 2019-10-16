package store

import "time"

type Build struct {
	Name      string
	Branch    string
	Id        int
	Git       string
	Timestamp time.Time
	Image     string
	Token     string
	Addr      string
	User      string
}
