package models

type ParamPage struct {
	Limit  int32 `json:"limit" binding:"required"`
	Offset int32 `json:"offset" binding:"required"`
}

type ParamUpdateName struct {
	UserId  int32  `json:"user_id" binding:"required"`
	NewName string `json:"new_name" binding:"required"`
}

type ParamAddUser struct {
	Password string `json:"password" binding:"required"`
	UserName string `json:"username" binding:"required"`
}
