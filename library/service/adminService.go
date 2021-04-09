package service

import (
	"gfast/library/response"
	"gfast/library/utils"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

var (
	AdminMultiLogin      bool  //是否允许后台管理员多端登陆
	AdminPageNum         = 20  //后台分页长度
	NotCheckAuthAdminIds []int //无需验证权限的用户id
)

//AdminLogin 后台用户登陆验证
func AdminLogin(r *ghttp.Request) (string, interface{}) {

	data := r.GetFormMapStrStr()
	rules := map[string]string{
		"idValueC": "required",
		"username": "required",
		"password": "required",
	}
	msgs := map[string]interface{}{
		"idValueC": "请输入验证码",
		"username": "账号不能为空",
		"password": "密码不能为空",
	}

	if e := gvalid.CheckMap(data, rules, msgs); e != nil {
		response.JsonExit(r, response.ErrorCode, e.String())
	}
	//判断验证码是否正确
	if !VerifyString(data["idKeyC"], data["idValueC"]) {
		response.JsonExit(r, response.ErrorCode, "验证码输入错误")
	}
	password := utils.EncryptCBC(data["password"], utils.AdminCbcPublicKey)
	var keys string
	if AdminMultiLogin {
		keys = data["username"] + password + gmd5.MustEncryptString(utils.GetClientIp(r))
	} else {
		keys = data["username"] + password
	}
	ip := utils.GetClientIp(r)
	userAgent := r.Header.Get("User-Agent")
	if err, user := signIn(data["username"], password, r); err != nil {
		go loginLog(0, data["username"], ip, userAgent, err.Error(), "系统后台")
		response.JsonExit(r, response.ErrorCode, err.Error())
	} else {
		//判断是否后台用户
		if user.IsAdmin != 1 {
			response.JsonExit(r, response.ErrorCode, "抱歉!此用户不属于后台管理员!")
		}
		r.SetParam("userInfo", user)
		go loginLog(1, data["username"], ip, userAgent, "登录成功", "系统后台")
		return keys, user
	}
	return keys, nil
}
