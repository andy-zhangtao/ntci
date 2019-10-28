package store

import "github.com/sirupsen/logrus"

func (p *PGBus) AddNtci(user, name, branch, ntci string) (err error) {
	_n, err := p.GetNtci(user, name, branch)
	if err != nil {
		return err
	}

	if _n != "" {
		return nil
	}

	sql := "INSERT INTO ntci (owner, name, branch, ntci) VALUES($1, $2, $3, $4)"
	logrus.Infof("Insert ntci SQL: %s", sql)

	_, err = p.db.Exec(sql, user, name, branch, ntci)
	return
}

func (p *PGBus) GetNtci(user, name, branch string) (ntci string, err error) {
	sql := "SELECT ntci FROM ntci WHERE owner=$1 AND name=$2 AND branch=$3"

	logrus.Infof("Select SQL: %s . $1=%s, $2=%s, $2=%s", sql, user, name, branch)

	rows, err := p.db.Query(sql, user, name, branch)
	if err != nil {
		return ntci, err
	}

	if rows.Next() {
		err = rows.Scan(&ntci)
		if err != nil {
			return ntci, err
		}
	}

	return
}
