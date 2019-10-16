package store

import (
	_sql "database/sql"
	"strconv"

	"github.com/sirupsen/logrus"
)

func (p *PGBus) GetBuild(user, name string) (bs []Build, err error) {
	sql := ""
	var rows *_sql.Rows

	if name == "" {
		sql = "SELECT * FROM build where owner=$1"
		logrus.Debugf("Select SQL: %s . $1= %s", sql, user)
		rows, err = p.db.Query(sql, user)
		if err != nil {
			return
		}
	} else {
		sql = "SELECT * FROM build where owner=$1 AND name=$2"
		logrus.Debugf("Select SQL: %s . $1= %s . $2= %s", sql, user, name)
		rows, err = p.db.Query(sql, user, name)
		if err != nil {
			return
		}
	}

	for {
		if rows.Next() {
			var b Build
			if rows.Scan(&b) == nil {
				bs = append(bs, b)
			}
		} else {
			return
		}
	}
}

/*
AddNewBuild

Insert new build record.
*/
func (p *PGBus) AddNewBuild(b Build) (id int, err error) {
	id, err = p.getNextId(b)
	if err != nil {
		return 0, err
	}

	sql := "INSERT INTO build (name,id,branch,git,timestamp,status,owner) VALUES ($1, $2, $3, $4, $5,0,$6)"
	logrus.Debugf("Insert New ID SQL: %s ", sql)

	_, err = p.db.Exec(sql, b.Name, id, b.Branch, b.Git, b.Timestamp, b.User)
	return id, err
}

/*
UpdataBuildStatus

Update specify job status.
// 0 - Ready
// 1 - Git clone success
//-1 - Git clone failed
// 2 - Ntci parse success
//-2 - Ntci parse failed
// 3 - Building
// 4 - Build success
//-4 - Build failed
*/
func (p *PGBus) UpdataBuildStatus(status int32, name, id, user string) (err error) {

	i, _ := strconv.Atoi(id)
	b := Build{
		Name: name,
		Id:   i,
		User: user,
	}

	return p.updateBuild(status, b)
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
	sql := "INSERT INTO id(name,id) VALUES($1,1)"
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

func (p *PGBus) updateBuild(status int32, b Build) error {
	sql := "UPDATE build SET status=$1 WHERE name=$2 and id=$3 and owner=$4"
	logrus.Debugf("UPDATE Build Status SQL: %s ", sql)
	_, err := p.db.Exec(sql, status, b.Name, b.Id, b.User)
	return err
}
