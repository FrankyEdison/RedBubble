package controller

import (
	"RedBubble/common/response"
	"RedBubble/common/responseCode"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// 实际上就是先初始化gin路由，然后写死参数、请求该路由，判断返回结果是否符合预期

func TestAddPostHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/api/post/addPost"
	r.POST(url, AddPostHandler)

	body := `{
		"category_id" : 1,
		"title" : "叱咤乐坛生力军",
		"content" : "mc张天赋与chantal姚焯菲合唱"
	}`

	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body))) //创建请求对象
	w := httptest.NewRecorder()                                                    //创建响应对象
	r.ServeHTTP(w, req)

	// 判断响应的内容是不是按预期返回了需要登录的错误
	assert.Equal(t, 200, w.Code) //判断w的响应码是否为200
	// 将w的响应体反序列化到ResponseData 然后判断字段与预期是否一致
	res := new(response.ResponseData)
	if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
		t.Fatalf("json.Unmarshal w.Body failed, err:%v\n", err)
	}
	assert.Equal(t, res.Code, responseCode.CodeNeedLogin)
	assert.Equal(t, res.Code.Msg(), responseCode.CodeNeedLogin.Msg())
}
