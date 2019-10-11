package store

import (
	"github.com/sirupsen/logrus"
)

/*
AddNewBuild

Insert new build record.
*/
func (p *PGBus) AddNewBuild(b Build) (err error) {
	id, err := p.getNextId(b)
	if err != nil {
		return err
	}

	sql := "INSERT INTO build ('name','id','branch','git','timestamp') VALUES ($1, $2, $3, $4, $5)"
	logrus.Debugf("Insert New ID SQL: %s ", sql)

	_, err = p.db.Exec(sql, b.Name, id, b.Branch, b.Git, b.Timestamp)
	return
}

/*
getNextId

Get specify build ID. If there is no build record, create a new one.
*/
func (p *PGBus) getNextId(b Build) (id int, err error) {
	query := "SELECT ID FROM id WHERE name=$1"
	logrus.Debugf("GetNextID SQL: %s , name: %s", query, b.Name)
	rows, err := p.db.Query(query, b.Name)
	if err != nil {
		return 0, err
	}

	if rows.Next() {
		rows.Scan(&id)
		return id, p.addBuildId(b)
	} else {

		return 0, p.createNewId(b)
	}
}

func (p *PGBus) createNewId(b Build) error {
	sql := "INSERT INTO id('name','id') VALUES($1,1)"
	logrus.Debugf("Insert New ID SQL: %s ", sql)
	_, err := p.db.Exec(sql, b.Name)
	return err
}

func (p *PGBus) addBuildId(b Build) error {
	sql := "UPDATE id set id=id+1 WHERE name=$1"
	logrus.Debugf("UPDATE ID SQL: %s ", sql)
	_, err := p.db.Exec(sql, b.Name)
	return err
}
