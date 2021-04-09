package boot

import (
	"gfast/library/service"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/frame/g"
)

var TopicGfToken *gtoken.GfToken

func initTopic() {
	//无需验证权限的用户id
	service.NotCheckAuthAdminIds = g.Cfg().GetInts("adminInfo.notCheckAuthAdminIds")
	//后端分页长度配置
	service.AdminPageNum = g.Cfg().GetInt("adminInfo.pageNum")
	// 设置并启动后台gtoken处理
	initTopicGfToken()
}

func initTopicGfToken() {
	//多端登陆配置 
	TopicGfToken = &gtoken.GfToken{
		CacheMode:        g.Cfg().GetInt8("gToken.CacheMode"),
		CacheKey:         g.Cfg().GetString("gToken.CacheKey"),
		Timeout:          g.Cfg().GetInt("gToken.Timeout"),
		MaxRefresh:       g.Cfg().GetInt("gToken.MaxRefresh"),
		TokenDelimiter:   g.Cfg().GetString("gToken.TokenDelimiter"),
		EncryptKey:       g.Cfg().GetBytes("gToken.EncryptKey"),
		AuthFailMsg:      g.Cfg().GetString("gToken.AuthFailMsg"),
		MultiLogin:       true,
		LoginPath:        "/pub6Login/login",
		LoginBeforeFunc:  service.UserLogin,
		LoginAfterFunc:   service.UserLoginAfter,
		LogoutPath:       "/pub6Login/logout",
		AuthPaths:        g.SliceStr{"/api.v2/*"},
		AuthAfterFunc:    service.AuthAfterFunc,
		LogoutBeforeFunc: service.LoginOut,
	}
	TopicGfToken.Start()
}
