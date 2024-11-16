package models

//定义请求的参数结果体

type ParamSignUp struct {
	Username   string `json:"username" binding:"required" `
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"rePassword" binding:"required,eqfield=Password"`
}
