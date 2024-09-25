package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"herrz-backend-base/comm"
	"herrz-backend-base/dao/mysql"
	"herrz-backend-base/models"
)

func SearchOne(ctx context.Context, c *app.RequestContext) {
	data, err := mysql.SearchOne()
	if err != nil {
		ResponseErrWithMsg(c, CodeServerBusy, err.Error())
		return
	}
	ResponseSuccess(c, data)
}

func SearchAll(ctx context.Context, c *app.RequestContext) {
	data, err := mysql.SearchAll()
	if err != nil {
		ResponseErrWithMsg(c, CodeServerBusy, err.Error())
		return
	}
	ResponseSuccess(c, data)
}

func SearchAllWithPost(ctx context.Context, c *app.RequestContext) {
	reqData := new(models.ParamPage)
	if err := c.Bind(reqData); err != nil {
		comm.Logger.Error().Msg(err.Error())
		ResponseErrWithMsg(c, CodeInvalidParams, err.Error())
		return
	}

	data, err := mysql.SearchAllWithPost(reqData)
	if err != nil {
		ResponseErrWithMsg(c, CodeServerBusy, err.Error())
		return
	}
	ResponseSuccess(c, data)
}

func Update(ctx context.Context, c *app.RequestContext) {
	updaterId := c.Request.Header.Get("user_id")
	updaterName := c.Request.Header.Get("username")
	reqData := new(models.ParamUpdateName)
	if err := c.Bind(reqData); err != nil {
		comm.Logger.Error().Msg(err.Error())
		ResponseErrWithMsg(c, CodeInvalidParams, err.Error())
		return
	}
	err := mysql.UpdateName(reqData)
	if err != nil {
		ResponseErrWithMsg(c, CodeServerBusy, err.Error())
		return
	}
	comm.Logger.Info().Msgf("updaterId: %v, updaterName: %v", updaterId, updaterName)
	ResponseSuccess(c, "update successfully!")
}

func Insert(ctx context.Context, c *app.RequestContext) {
	var addUser models.ParamAddUser
	addUser.Password = string(c.FormValue("password"))
	addUser.UserName = string(c.FormValue("username"))
	err := mysql.InsertOne(&addUser)
	if err != nil {
		ResponseErrWithMsg(c, CodeServerBusy, err.Error())
		return
	}
	ResponseSuccess(c, "insert one successfully!")
}
