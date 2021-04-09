package router

import (
	"gfast/app/controller/home"
	"gfast/app/controller/topic" 

	"gfast/middleWare"
	"github.com/gogf/gf/frame/g"
	//"github.com/gogf/csrf"
	"github.com/gogf/gf/net/ghttp"
	//"net/http"
	//"time"

)

//前端路由处理
func init() {
	s := g.Server()

	group := s.Group("/")
	//api

	globaltopicTask := new (topic.Task)
	group.POST("/pub6/topic/tasks", globaltopicTask.List)

	globaltopicMembers := new (topic.Members)
	group.POST("/pub6/topic/members", globaltopicMembers.List)

	globaltopicNotes := new (topic.Notes)
	group.POST("/pub6/topic/notes", globaltopicNotes.List)

	globaltopicNotices := new (topic.Notices)
	group.POST("/pub6/topic/notices", globaltopicNotices.List)

	globaltopicResources := new (topic.Resource)
	group.POST("/pub6/topic/resources", globaltopicResources.List)

	globaltopic := new(topic.Topic)
	group.POST("/pub6/topic/all", globaltopic.List)
	group.POST("/pub6/topic/type", globaltopic.ListByType)
	group.GET("/pub6/topic/get", globaltopic.Edit)
	group.POST("/pub6/topic/category", globaltopic.Catgory)

	globaluser := new(topic.Users)
	group.POST("/pub6/user/reg", globaluser.Add)

	group.Group("/api.v2", func(group *ghttp.RouterGroup) { 
		group.ALL("/csrf", func(r *ghttp.Request) {
			r.Response.Writeln(r.Method + ": " + r.RequestURI)
		})
		/*
		group.Middleware(csrf.NewWithCfg(csrf.Config{
			Cookie: &http.Cookie{
				Name: "_csrf",// token name in cookie
			},
			ExpireTime:      time.Hour * 24,
			TokenLength:     32,
			TokenRequestKey: "X-Token",// use this key to read token in request param
		}))
		*/

		group.Middleware(middleWare.UserAuth) //后台权限验证
	 	group.ALL("/topic", new(topic.Topic))
	 	group.ALL("/topicresource", new (topic.Resource))
	 	group.ALL("/topictask", new (topic.Task))
	 	group.ALL("/topicmembers", new (topic.Members))
	 	group.ALL("/topictaskprogress", new (topic.Taskprogress))
	 	group.ALL("/topicnotes", new (topic.Notes))
	 	group.ALL("/topicntices", new (topic.Notices))
	 	group.ALL("/user", new(topic.Users))
	 	group.ALL("/question", new(topic.Questions))
	 	group.ALL("/questionhistory", new(topic.History))
	 	group.ALL("/questionhistorydetail", new(topic.Detail))
	 	group.ALL("/booklist", new (topic.BookResource))
	 	group.ALL("/bookcatetory", new (topic.Category))
	 	

	})

	s.Group("/", func(group *ghttp.RouterGroup) {
		home := new(home.Index)
		//group.GET("/", home.Index)
		group.GET("/cms/", home.Index)
		group.GET("/cms/list/:cateId/*page/*keyWords", home.List)
		group.GET("/cms/show/:cateIds/:newsId", home.Show)
		group.ALL("/cms/search/*page/*keyWords", home.Search)
	})
}
