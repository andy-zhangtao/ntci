package store

import (
	_sql "database/sql"
	"errors"

	"github.com/sirupsen/logrus"
)

func (p *PGBus) GetCommonEnv() (env map[string]string, err error) {
	sql := "SELECT key, value FROM priavte_data WHERE owner=$1 AND name=$2"
	var rows *_sql.Rows

	rows, err = p.db.Query(sql, "super", "super")
	if err != nil {
		return
	}

	env = make(map[string]string)

	for rows.Next() {
		var key, value string
		err = rows.Scan(&key, &value)
		if err != nil {
			logrus.Error(err)
		} else {
			env[key] = value
		}
	}

	return
}
func (p *PGBus) GetBuildByID(user, name string, id int) (b Build, err error) {
	sql := "SELECT * FROM build WHERE owner=$1 AND name=$2 AND id=$3"
	logrus.Debugf("Select SQL: %s . $1=%s $2=%s $3=%d ", sql, user, name, id)
	var rows *_sql.Rows

	rows, err = p.db.Query(sql, user, name, id)
	if err != nil {
		return
	}

	if rows.Next() {
		err = rows.Scan(&b.Name, &b.Id, &b.Branch, &b.Git, &b.Timestamp, &b.Status, &b.User, &b.Sha, &b.Message, &b.Language, &b.Lanversion)
		if err != nil {
			return
		}

	} else {
		return b, errors.New("Not Found This Build ")
	}

	return
}

func (p *PGBus) GetBuild(user, name string) (bs []Build, err error) {
	sql := ""
	var rows *_sql.Rows

	if name == "" {
		sql = "SELECT * FROM build where owner=$1 order by timestamp desc  LIMIT 30"
		logrus.Debugf("Select SQL: %s . $1= %s", sql, user)
		rows, err = p.db.Query(sql, user)
		if err != nil {
			return
		}
	} else {
		sql = "SELECT * FROM build where owner=$1 AND name=$2 order by timestamp desc LIMIT 30"
		logrus.Debugf("Select SQL: %s . $1= %s . $2= %s", sql, user, name)
		rows, err = p.db.Query(sql, user, name)
		if err != nil {
			return
		}
	}

	for rows.Next() {

		b := Build{}
		err = rows.Scan(&b.Name, &b.Id, &b.Branch, &b.Git, &b.Timestamp, &b.Status, &b.User, &b.Sha, &b.Message)
		if err == nil {
			bs = append(bs, b)
		} else {
			logrus.Error(err)
		}

	}

	return
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

	sql := "INSERT INTO build (name,id,branch,git,timestamp,status,owner, sha, message) VALUES ($1, $2, $3, $4, $5, 0, $6, $7, $8)"
	logrus.Debugf("Insert New ID SQL: %s ", sql)

	_, err = p.db.Exec(sql, b.Name, id, b.Branch, b.Git, b.Timestamp, b.User, b.Sha, b.Message)
	return id, err
}

// UpdateBuildLanguage
// Language can not empty. Lan version default is latest
func (p *PGBus) UpdateBuildLanguage(language, lanver string, id int, name, user string) (err error) {
	b := Build{
		Name: name,
		Id:   id,
		User: user,
	}

	return p.updateLanguage(language, lanver, b)
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
func (p *PGBus) UpdataBuildStatus(status int32, id int, name, user string) (err error) {

	//i, _ := strconv.Atoi(id)
	b := Build{
		Name: name,
		Id:   id,
		User: user,
	}

	return p.updateBuild(status, b)
}

/*
getNextId

Get specify build ID. If there is no build record, create a new one.
*/
func (p *PGBus) getNextId(b Build) (id int, err error) {
	query := "SELECT ID FROM id WHERE name=$1 AND owner=$2"
	logrus.Debugf("GetNextID SQL: %s , name: %s, owner:%s", query, b.Name, b.User)
	rows, err := p.db.Query(query, b.Name, b.User)
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
	sql := "INSERT INTO id(owner,name,id) VALUES($1,$2,1)"
	logrus.Debugf("Insert New ID SQL: %s ", sql)
	_, err := p.db.Exec(sql, b.User, b.Name)
	return err
}

func (p *PGBus) addBuildId(b Build) error {
	sql := "UPDATE id set id=id+1 WHERE name=$1 AND owner=$2"
	logrus.Debugf("UPDATE ID SQL: %s ", sql)
	_, err := p.db.Exec(sql, b.Name, b.User)
	return err
}

func (p *PGBus) updateBuild(status int32, b Build) error {
	sql := "UPDATE build SET status=$1 WHERE name=$2 and id=$3 and owner=$4"
	logrus.Debugf("UPDATE Build Status SQL: %s ", sql)
	_, err := p.db.Exec(sql, status, b.Name, b.Id, b.User)
	return err
}

func (p *PGBus) updateLanguage(lan, lanv string, b Build) error {
	sql := "UPDATE build SET language=$1, langver=$2 WHERE name=$3 and id=$4 and owner=$5"
	logrus.Debugf("UPDATE Build Language SQL: %s ", sql)
	_, err := p.db.Exec(sql, lan, lanv, b.Name, b.Id, b.User)
	return err
}
