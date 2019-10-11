package store

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"ntci/ci-build/k8s/dataBus"
)

/*
The backend store.

K8s build server will store all builds in postgres db.
*/

var pb *PGBus

type PGBus struct {
	db *sql.DB
}

func PG() *PGBus {
	return pb
}

func PGInit(bus *dataBus.DataBus) {
	pb = new(PGBus)
	connStr := ""

	if bus.Postgres.Passwd == "" {
		connStr = fmt.Sprintf("postgres://%s@%s:%d/%s?sslmode=disable", bus.Postgres.User, bus.Postgres.Addr, bus.Postgres.Port, bus.Postgres.Name)
	} else {
		connStr = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", bus.Postgres.User, bus.Postgres.Passwd, bus.Postgres.Addr, bus.Postgres.Port, bus.Postgres.Name)
	}

	logrus.Debugf("Postgres Connstr: %s", connStr)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logrus.Fatalf("Connect Postgres Error: %s", err.Error())
	}

	pb.db = db

	if err := pb.db.Ping(); err != nil {
		logrus.Fatalf("Ping Postgres Error: %s", err.Error())
	}

	logrus.Info("Connect Postgres Success.")
}
