package mysql

import (
	"herrz-backend-base/comm"
	"herrz-backend-base/models"
)

var sqlUtil SqlUtil

func SearchOne() (int, error) {
	sql := "select count(1) from t_users where id > ? and username != ?;"
	var cnt int
	name := "hh"
	num := 1
	err := sqlUtil.FetchOne(db, sql, &cnt, num, name)
	if err != nil {
		return 0, err
	}
	return cnt, nil
}

func SearchAll() (interface{}, error) {
	sql := "select tu.id, tu.username from t_users tu limit ? offset ?;"
	var res []*models.ParamUser
	limit, offset := 10, 0
	err := sqlUtil.FetchAll(db, sql, &res, limit, offset)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func SearchAllWithPost(reqData *models.ParamPage) (interface{}, error) {
	sql := "select tu.id, tu.username from t_users tu limit ? offset ?;"
	var res []*models.ParamUser
	err := sqlUtil.FetchAll(db, sql, &res, reqData.Limit, reqData.Offset)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func UpdateName(reqData *models.ParamUpdateName) (err error) {
	sql := "update t_users set username = ? where id = ?;"
	res, err := sqlUtil.Exec(db, sql, reqData.NewName, reqData.UserId)
	if err != nil {
		return
	}
	// 根据实际需要
	lastInsertId, _ := res.LastInsertId()
	affectedRows, _ := res.RowsAffected()
	comm.Logger.Debug().Msgf("lastInsertId: %v, affectedRows: %v ", lastInsertId, affectedRows)
	return
}

func InsertOne(reqData *models.ParamAddUser) (err error) {
	sql := "insert into t_users (username, password) values (?, ?);"
	res, err := sqlUtil.Exec(db, sql, reqData.UserName, reqData.Password)
	if err != nil {
		return
	}
	// 根据实际需要
	lastInsertId, _ := res.LastInsertId()
	affectedRows, _ := res.RowsAffected()
	comm.Logger.Debug().Msgf("lastInsertId: %v, affectedRows: %v ", lastInsertId, affectedRows)
	return
}
