package store

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

/*
The backend store.

K8s build server will store all builds in postgres db.
*/

var pb *PGBus

type PGBus struct {
	db *sql.DB
}

func PG(addr string) *PGBus {
	if pb == nil {
		pgInit(addr)
		return pb
	}

	return pb
}

func pgInit(addr string) {

	pb = new(PGBus)

	logrus.Debugf("Postgres Connstr: %s", addr)

	db, err := sql.Open("postgres", addr)
	if err != nil {
		logrus.Fatalf("Connect Postgres Error: %s", err.Error())
	}

	pb.db = db

	if err := pb.db.Ping(); err != nil {
		logrus.Fatalf("Ping Postgres Error: %s", err.Error())
	}

	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(30)
	db.SetConnMaxLifetime(5 * time.Minute)
	logrus.Info("Connect Postgres Success.")
}
