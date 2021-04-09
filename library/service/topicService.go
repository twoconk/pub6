package service

import (
	"gfast/app/model/admin/user_online"
	"gfast/app/model/topic/users"
	"gfast/app/model/topic/userstrace"
	"gfast/library/response"
	"gfast/library/utils"

	"database/sql"
	"errors"

	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"github.com/mssola/user_agent"

	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

// 用户登录，成功返回用户信息，否则返回nil
func signInTopicUser(username, password string, r *ghttp.Request) (error, *users.Entity) {
	userInfo, err := users.Model.Where("name=? and passwd=?", username, password).One()
	if err != nil && err != sql.ErrNoRows {
		return err, nil 
	} 
	if userInfo == nil {
		return errors.New("账号或密码错误"), nil
	}
	//判断用户状态
	if userInfo.Status == 0 || userInfo.Status == 10 {
		return errors.New("用户已被冻结或者被删除"), nil
	}
	returnData := *userInfo

	//更新登陆时间及ip
	usertrace, err := userstrace.Model.Where("uid=?", userInfo.Id).One()

	if usertrace == nil || err != nil {
		usertrace2 := userstrace.Entity{
			Uid:   userInfo.Id,
			Ext:   gconv.String(gtime.Timestamp()),
			Type:  "1", //类型(0:注册1::登录2:退出3:修改4:删除)
			Ip:    utils.GetClientIp(r),
			Ctime: gconv.String(gtime.Timestamp()),
		}
		userstrace.Model.Save(usertrace2)
	} else {
		usertrace.Ctime = gconv.String(gtime.Timestamp())
		usertrace.Ip = utils.GetClientIp(r)
		usertrace.Ext = gconv.String(gtime.Timestamp())
		userstrace.Model.Save(usertrace)
	}

	return nil, &returnData
}

// 登录返回方法
func UserLoginAfter(r *ghttp.Request, respData gtoken.Resp) {
	if !respData.Success() {
		r.Response.WriteJson(respData)
	} else {
		token := respData.GetString("token")
		uuid := respData.GetString("uuid")
		var userInfo *users.Entity
		r.GetParamVar("userInfo").Struct(&userInfo)
		//保存用户在线状态token到数据库
		userAgent := r.Header.Get("User-Agent")
		ua := user_agent.New(userAgent)
		os := ua.OS()
		explorer, _ := ua.Browser()
		entity := user_online.Entity{
			Uuid:       uuid,
			Token:      token,
			CreateTime: gconv.Uint64(gtime.Timestamp()),
			UserName:   userInfo.Name,
			Ip:         utils.GetClientIp(r),
			Explorer:   explorer,
			Os:         os,
		}
		user_online.Model.Save(entity)
		r.Response.WriteJson(gtoken.Succ(g.Map{
			"token": token,
			"uid": userInfo.Id,
			"joinTime":userInfo.Ctime,
			"ext": userInfo.Ext,
			"action": "",
		}))
	}
}

//UserLogin 前端用户登陆验证
func UserLogin(r *ghttp.Request) (string, interface{}) {

	data := r.GetFormMapStrStr()
	rules := map[string]string{
		"idValueC": "required",
		"username": "required",
		"passwd":   "required",
	}
	msgs := map[string]interface{}{
		"idValueC": "请输入验证码",
		"username": "账号不能为空",
		"passwd":   "密码不能为空",
	}

	if e := gvalid.CheckMap(data, rules, msgs); e != nil {
		response.JsonExit(r, response.ErrorCode, e.String())
	}
	//判断验证码是否正确
	if !VerifyString(data["idKeyC"], data["idValueC"]) {
		response.JsonExit(r, response.ErrorCode, "验证码输入错误")
	}
	password := utils.EncryptCBC(data["passwd"], utils.AdminCbcPublicKey)
	var keys = data["username"] + password + gmd5.MustEncryptString(utils.GetClientIp(r))

	ip := utils.GetClientIp(r)
	userAgent := r.Header.Get("User-Agent")
	if err, user := signInTopicUser(data["username"], password, r); err != nil {
		go loginLog(0, data["username"], ip, userAgent, err.Error(), "系统后台")
		response.JsonExit(r, response.ErrorCode, err.Error())
	} else {
		r.SetParam("userInfo", user)
		go loginLog(1, data["username"], ip, userAgent, "登录成功", "系统后台")
		return keys, user
	}
	return keys, nil
}
