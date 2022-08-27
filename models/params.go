package models

/**
form格式是一种“键值对”的数据格式，比如：a=1&b=2&c=3
而json格式则类似这样：{"a":1,"b":2,"c":3}
get请求都用form格式，post请求json和form格式都行
*/
// binding tag的内容是根据validator库的文档进行添加规则，这些规则是用来校验参数的，比如说哪些参数必须要填，哪些参数只能填1或-1等等，校验参数不通过就会直接报错返回，本项目已配置validator返回中文错误

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
	// 用户id从jwt里拿
	CategoryId int64  `json:"category_id" binding:"required"`
	Title      string `json:"title" binding:"required"`   //帖子标题
	Content    string `json:"content" binding:"required"` //帖子内容
}

// 分页参数
type ByPage struct {
	PageSize   int `form:"page_size" binding:"required"`
	PageNumber int `form:"page_number" binding:"required"`
}

// 点赞/点灭帖子参数
type ParamVotePost struct {
	// 用户id从jwt里拿
	//PostId 	   int64 `json:"post_id,string" binding:"required"`
	PostId     string `json:"post_id" binding:"required"`         //不转成int64了，因为存到redis里也是要string类型
	VoteAction int8   `json:"vote_action" binding:"oneof=-1 0 1"` //点赞（1）、点灭（-1）、取消点赞或点灭（0），oneof规则的意思是该参数只能是这三个数之一，不需要required因为传0会说我没传
}
