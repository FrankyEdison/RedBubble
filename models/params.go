package models

/**
form格式是一种“键值对”的数据格式，比如：a=1&b=2&c=3
而json格式则类似这样：{"a":1,"b":2,"c":3}
get请求都用form格式，post请求json和form格式都行
*/
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

// 发布帖子参数
type ParamAddPost struct {
	CategoryId int64  `json:"category_id" binding:"required"`
	Title      string `json:"title" binding:"required"`
	Content    string `json:"content" binding:"required"`
}

// 分页参数
type ByPage struct {
	PageSize   int `form:"page_size" binding:"required"`
	PageNumber int `form:"page_number" binding:"required"`
}
