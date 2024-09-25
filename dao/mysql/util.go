package mysql

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"herrz-backend-base/comm"
)

type SqlUtil struct{}

func (s *SqlUtil) NewSession() *sqlx.Tx {
	session, err := db.Beginx()
	if err != nil {
		comm.MysqlLogger.Error().Msgf("new session failed: %s", err.Error())
		return nil
	}
	return session
}

func (s *SqlUtil) FetchOne(db *sqlx.DB, sql string, dest interface{}, args ...interface{}) error {
	return db.Get(dest, sql, args...)
}

func (s *SqlUtil) FetchAll(db *sqlx.DB, sql string, dest interface{}, args ...interface{}) error {
	return db.Select(dest, sql, args...)
}

func (s *SqlUtil) Exec(db *sqlx.DB, sql string, args ...interface{}) (res sql.Result, err error) {
	session, err := db.Begin()
	if err != nil {
		comm.MysqlLogger.Error().Msgf("create session failed!", err.Error())
		_ = session.Rollback()
		return
	}
	res, err = session.Exec(sql, args...)
	if err != nil {
		_ = session.Rollback()
		return
	}
	if err = session.Commit(); err != nil {
		return
	}
	return
}
