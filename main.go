package main

import (
	"herrz-backend-base/comm"
	"herrz-backend-base/dao/mysql"
	"herrz-backend-base/routes"
)

func main() {
	if err := comm.InitViperCfg(); err != nil {
		return
	}
	if err := comm.InitLogger(); err != nil {
		return
	}
	log := comm.Logger

	if err := mysql.InitMysqlCfg(); err != nil {
		log.Error().Msgf("init mysql failed: %s!", err.Error())
		return
	}
	defer mysql.Close()

	hz := routes.Init()

	// 启动服务器
	hz.Spin()
}
