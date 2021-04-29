package method

import (
	"bishe/backend/util"

	"github.com/gin-gonic/gin"
)

func UserInfo(c *gin.Context) {
	var req UserInfoRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(200, util.BuildError(util.PARAMERROR, util.ErrMap[util.PARAMERROR]))
		return
	}
	resp := UserInfoResponse{
		Code: 0,
		Msg:  "",
	}
	userId, err := GetId(c)
	if err != nil {
		c.JSON(200, util.BuildError(util.NOTLOGIN, err.Error()))
		return
	}
	userType, err := GetWho(c)
	if err != nil {
		c.JSON(200, util.BuildError(util.NOTLOGIN, err.Error()))
		return
	}
	// 从数据库拿到用户名称

	resp.UserId = userId
	resp.UserType = userType
	resp.Name = "wzy"
	c.JSON(200, resp)
}

type UserInfoRequest struct {
}

type UserInfoResponse struct {
	Name     string `json:"name"`
	UserId   int32  `json:"user_id"`
	UserType int32  `json:"user_type"`

	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}