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
}
