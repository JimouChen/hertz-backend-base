package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"herrz-backend-base/comm"
)

var db *sqlx.DB

// InitMysqlCfg init mysql db
func InitMysqlCfg() (err error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		comm.CfgLoader.GetString("mysql.user"),
		comm.Decode(
			comm.CfgLoader.GetString("mysql.password"),
			comm.CfgLoader.GetString("key.aes_32_code"),
		),
		comm.CfgLoader.GetString("mysql.host"),
		comm.CfgLoader.GetInt("mysql.port"),
		comm.CfgLoader.GetString("mysql.dbname"),
	)
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		comm.MysqlLogger.Error().Msgf("connect DB failed, err:%v\n", err.Error())
		return
	}
	db.SetMaxOpenConns(comm.CfgLoader.GetInt("mysql.max_open_connections"))
	db.SetMaxIdleConns(comm.CfgLoader.GetInt("mysql.max_idle_connections"))
	comm.MysqlLogger.Info().Msg("connect DB successfully!")
	return
}

func Close() {
	_ = db.Close()
}
