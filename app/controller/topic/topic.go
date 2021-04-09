// ==========================================================================
// 生成日期：2021-01-13 15:10:25
// 生成人：gfast
// ==========================================================================
package topic

import (
	topicModel "gfast/app/model/topic/topics"
	topicService "gfast/app/service/topic/topic_service"
	topicUserService "gfast/app/service/topic/users_service"
	membersService "gfast/app/service/topic/members_service"
	libraryService "gfast/library/service"
	"gfast/app/service/topic/users_service"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"

	"gfast/library/response"
    "html"
)

// Topic分类
type EntityTopicCatgory struct {
    Id     int    `orm:"id"       json:"id"`
    Name    string    `orm:"name"       json:"name"` 
}

//控制器
type Topic struct{}
//目录
func (c *Topic) Catgory(r *ghttp.Request) {

	// 定义一个结构体存储请求参数
	var req *topicModel.SelectPageReq
	// 获取参数
	err := r.Parse(&req)
	if err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}

	cats := ([]EntityTopicCatgory)(nil)
	// 或者 var users []User
	err = g.DB("default").Table("guaniu_study_topic_type").Structs(&cats)
 
	response.SusJson(true, r, "获取目录成功", cats)
}

//列表页
func (c *Topic) List(r *ghttp.Request) {
	// 定义一个结构体存储请求参数
	var req *topicModel.SelectPageReq
	// 获取参数
	err := r.Parse(&req)
	if err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}
	total, page, list, err := topicService.SelectListByPage(req)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	result := g.Map{
		"currentPage": page,
		"total":       total,
		"list":        list,
	}
	response.SusJson(true, r, "获取列表数据成功", result)
}

func (c *Topic) ListByType(r *ghttp.Request) {
	// 定义一个结构体存储请求参数
	var req *topicModel.SelectPageReq
	// 获取参数
	err := r.Parse(&req)
	if err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}
	total, page, list, err := topicService.SelectListByPageAndType(req)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	result := g.Map{
		"currentPage": page,
		"total":       total,
		"list":        list,
	}
	response.SusJson(true, r, "获取列表数据成功", result)
}

//通过用户查询用户加入的主题
//评论列表页
func (c *Topic) UserTopicList(r *ghttp.Request) {
	// 定义一个结构体存储请求参数 SelectCommentListByPage(req *topicModel.CommenTopicReq)
	var req *topicModel.UserTopicReq
	// 获取参数
	err := r.Parse(&req)
	if err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}
	//增加主题查看人数
	
	total, page, list, err := topicService.SelectListByPageAndUserId(req)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	result := g.Map{
		"currentPage": page,
		"total":       total,
		"list":        list,
	}
	response.SusJson(true, r, "获取列表数据成功", result)
}



//评论列表页
func (c *Topic) CommentList(r *ghttp.Request) {
	// 定义一个结构体存储请求参数 SelectCommentListByPage(req *topicModel.CommenTopicReq)
	var req *topicModel.CommenTopicReq
	// 获取参数
	err := r.Parse(&req)
	if err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}
	//增加主题查看人数
	
	total, page, list, err := topicService.SelectCommentListByPage(req)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	result := g.Map{
		"currentPage": page,
		"total":       total,
		"list":        list,
	}
	response.SusJson(true, r, "获取列表数据成功", result)
}

/*
liyizhang@macdeMac-2 kmip4j %

curl -d '{"topicType":1,"topicName":"Go语言学习", "topicContent":"好想学习Go语言，找个团", "status":0}' -H 'Authorization:Bearer vGCiSmcWADHgtiq2XeDTokcTCipNawAG8xZW1F+49fgzaNhaoop+XX1XHvm2qx715BB5VGUxeFlXqTL7WcE+c59q3k7TRk4Ai1P0E01dYeJlsyCA3Ss9vU066gLs4ipB' "http://127.0.0.1:8200/api.v2/topic/add"


{"code":0,"data":3,"msg":"添加成功"}%
*/
// 新增
func (c *Topic) Add(r *ghttp.Request) {
	if r.Method == "POST" {

		var req *topicModel.AddReq
		// 通过Parse方法解析获取参数
		err := r.Parse(&req)
		if err != nil {
			response.FailJson(true, r, err.(*gvalid.Error).FirstString())
		}
		//判断验证码是否正确，如何判断是同一个ip连续提交了几次，降权，判断
		
		data := r.GetFormMapStrStr()
		if !libraryService.VerifyString(data["idKeyC"], data["idValueC"]) {
			response.JsonExit(r, response.ErrorCode, "验证码输入错误")
		} 
		//填写用户id
		req.CreateOwnerId = users_service.GetLoginID(r)

		//转义关键字段
		req.TopicName = html.EscapeString(req.TopicName)
		req.TopicContent = html.EscapeString(req.TopicContent) 

		// 调用service中的添加函数添加
		id, _ := topicService.AddSave(req)
		if err != nil {
			response.FailJson(true, r, err.Error())
		}
		response.SusJson(true, r, "添加成功", id)
	}
}

// 增加评论
func (c *Topic) AddComment(r *ghttp.Request) {
	// 如果是post提交的请求就执行修改操作
	if r.Method == "POST" { 
		//加入该主题才能提交评论
		
		// 不是post提交的请求就到修改页面后查询出要修改的记录 
		id := r.GetInt("topicId")
		if (id == 0){

			response.FailJson(true, r, "参数错误")
			return
		}
		topic, err := topicService.GetByID(int64(id))
		if err != nil {
			//判断主题是不是存在
			response.FailJson(true, r, err.Error())
		}
		//获取登录ID
		userId := topicUserService.GetLoginID(r)

		if (topic.CreateOwnerId != userId){ 
			//判断该用户是否加入了这个组，没有则不能添加评论000
			params, _ := membersService.CheckUserInMembers(int64(id), userId)
			if params == nil {
				response.FailJson(true, r, "没有加入组") 
				return
			} 
		}
		var commentReq *topicModel.CommentReq
		// 通过Parse方法解析获取参数
		err = r.Parse(&commentReq)
		if err != nil {
			response.FailJson(true, r, err.(*gvalid.Error).FirstString())
		} 
		//转义关键字段
		commentReq.Content = html.EscapeString(commentReq.Content)

		commentReq.UserId = int64(userId)
		//存储
		topicService.AddComment(commentReq)

		response.SusJson(true, r, "添加成功")
	}

	response.SusJson(true, r, "无操作")
}

// 修改
func (c *Topic) Edit(r *ghttp.Request) {
	// 如果是post提交的请求就执行修改操作
	if r.Method == "POST" {
		//判断当前用户是否有修改该topic的权限，否则返回错误

		// 不是post提交的请求就到修改页面后查询出要修改的记录
		id := r.GetInt("id")
		params, err := topicService.GetByID(int64(id))
		if err != nil {
			response.FailJson(true, r, err.Error())
		}
		//获取登录ID
		userId := topicUserService.GetLoginID(r)
		if (params.CreateOwnerId != userId){ 
			response.FailJson(true, r, "该用户没有操作权限")
			return;
		}

		var editReq *topicModel.EditReq
		// 通过Parse方法解析获取参数
		err = r.Parse(&editReq)
		if err != nil {
			response.FailJson(true, r, err.(*gvalid.Error).FirstString())
		}
		//填写用户id
		editReq.CreateOwnerId = userId

		err = topicService.EditSave(editReq)
		if err != nil {
			response.FailJson(true, r, err.Error())
		}
		response.SusJson(true, r, "修改参数成功")
	}
	// 不是post提交的请求就到修改页面后查询出要修改的记录
	id := r.GetInt("id")
	params, err := topicService.GetByIDForDetail(int64(id))
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "ok", params)
}

// 删除
func (c *Topic) Delete(r *ghttp.Request) {
	//post
	if r.Method != "POST" {
		response.FailJson(true, r, "删除失败！")
		return
	}


	var req *topicModel.RemoveReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}
	//获取登录ID 
	params, _ := topicService.GetByID(int64(req.Ids[0]))

	//id := r.GetInt("id")
	//获取登录ID
	userId := topicUserService.GetLoginID(r)
	if params != nil && userId!= params.CreateOwnerId{ 
		response.FailJson(true, r, "删除失败")
		return
	}
	
	err := topicService.DeleteByIds(req.Ids)
	if err != nil {
		response.FailJson(true, r, "删除失败")
	}
	response.SusJson(true, r, "删除成功")
}
