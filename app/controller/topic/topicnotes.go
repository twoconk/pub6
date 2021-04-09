package topic

import (
	notesModel "gfast/app/model/topic/topicnotes"
	notesService "gfast/app/service/topic/notes_service"
	topicService "gfast/app/service/topic/topic_service"
	topicUserService "gfast/app/service/topic/users_service"
	membersService "gfast/app/service/topic/members_service"
	"gfast/library/response"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
    "html"
)

//控制器
type Notes struct{}

//列表页
func (c *Notes) List(r *ghttp.Request) {
	// 定义一个结构体存储请求参数
	var req *notesModel.SelectPageReq
	// 获取参数
	err := r.Parse(&req)
	if err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}
	total, page, list, err := notesService.SelectListByPage(req)
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

// 新增
func (c *Notes) Add(r *ghttp.Request) {
	if r.Method == "POST" {
		var req *notesModel.AddReq
		// 通过Parse方法解析获取参数
		err := r.Parse(&req)
		if err != nil {
			response.FailJson(true, r, err.(*gvalid.Error).FirstString())
		}
		if req.TopicId == 0 {
			response.FailJson(true, r, "no topic specific.")
			return
		}
		//填写用户id
		req.UserId = topicUserService.GetLoginID(r)
		

		topic, err := topicService.GetByID(int64(req.TopicId))
		if err != nil || topic == nil || topic.Id != req.TopicId {
			//判断主题是不是存在
			response.FailJson(true, r, err.Error())
			return
		}
		//该用户不是这个主题的owner
		if(topic.CreateOwnerId != req.UserId){ 
			//判断该用户是否加入了这个组，没有则不能添加notes
			params, _ := membersService.CheckUserInMembers(int64(topic.Id), req.UserId)
			if params == nil {
				response.FailJson(true, r, "没有加入该学习主题") 
				return
			} 
		}

		//转义关键字段 
		req.Title = html.EscapeString(req.Title) 

		//转义关键字段 
		req.Content = html.EscapeString(req.Content) 
		
		// 调用service中的添加函数添加
		err = notesService.AddSave(req)
		if err != nil {
			response.FailJson(true, r, err.Error())
		}
		response.SusJson(true, r, "添加成功")
	}
}

// 修改
func (c *Notes) Edit(r *ghttp.Request) {
	// 如果是post提交的请求就执行修改操作
	if r.Method == "POST" {
		var editReq *notesModel.EditReq
		// 通过Parse方法解析获取参数
		err := r.Parse(&editReq)
		if err != nil {
			response.FailJson(true, r, err.(*gvalid.Error).FirstString())
		}
		//转义关键字段 
		editReq.Title = html.EscapeString(editReq.Title) 

		//转义关键字段 
		editReq.Content = html.EscapeString(editReq.Content) 
		
		//判断该主题是否是这个用户创建的，或者这个记录是该用户提交的
		topicId := r.GetInt64("topicId")
		topic, err := topicService.GetByID(int64(topicId))
		if err != nil || topic == nil || topic.Id != topicId {
			//判断主题是不是存在
			response.FailJson(true, r, err.Error())
			return
		}
		userId := topicUserService.GetLoginID(r)

		id := r.GetInt64("id")
		if topic.CreateOwnerId != userId {
			//判断是否有加入这个主题 
			params, err := notesService.GetByID(int64(id))
			if err != nil {
				response.FailJson(true, r, err.Error())
			}
			if (userId != params.UserId){
				response.FailJson(true, r,"no permission.") 
			}
		}

	    //点赞，或者取消赞，可以修改LikeNum实现
		err = notesService.EditSave(editReq)
		if err != nil {
			response.FailJson(true, r, err.Error())
		}
		response.SusJson(true, r, "修改参数成功")
	}
	// 不是post提交的请求就到修改页面后查询出要修改的记录
	id := r.GetInt("id")
	params, err := notesService.GetByID(int64(id))
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "ok", params)
}

// 删除
func (c *Notes) Delete(r *ghttp.Request) {
	var req *notesModel.RemoveReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}
	//获取登录ID 
	params, _ := notesService.GetByID(int64(req.Ids[0]))

	//id := r.GetInt("id")
	//获取登录ID
	userId := topicUserService.GetLoginID(r)
	if params != nil && userId!= params.UserId{ 
		response.FailJson(true, r, "删除失败")
		return
	}
	
	err := notesService.DeleteByIds(req.Ids)
	if err != nil {
		response.FailJson(true, r, "删除失败")
	}
	response.SusJson(true, r, "删除成功")
}
