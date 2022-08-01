package models

// ParamSignUp 注册请求参数
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"` //binding:"required"指当该结构体绑定了请求参数，请求参数的这些字段必须不能为空，否则报错
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"` //binding:"eqfield=Password"用于判断两个字段的值是否相同
}

// ParamSignIn 登录请求参数
type ParamSignIn struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
