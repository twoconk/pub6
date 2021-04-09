package topic

import (
	membersModel "gfast/app/model/topic/topicmembers"
	membersService "gfast/app/service/topic/members_service"
	topicUserService "gfast/app/service/topic/users_service"
	topicService "gfast/app/service/topic/topic_service"
	"gfast/library/response"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

//控制器
type Members struct{}

//列表页
func (c *Members) List(r *ghttp.Request) {
	// 定义一个结构体存储请求参数
	var req *membersModel.SelectPageReq
	// 获取参数
	err := r.Parse(&req)
	if err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}
	total, page, list, err := membersService.SelectListByPage(req)
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
func (c *Members) Add(r *ghttp.Request) {
	if r.Method == "POST" {
		var req *membersModel.AddReq
		// 通过Parse方法解析获取参数
		err := r.Parse(&req)
		if err != nil {
			response.FailJson(true, r, err.(*gvalid.Error).FirstString())
		}
		if req.TopicId == 0 {
			response.FailJson(true, r, "no topic specific.")
			return
		}

		topic, err := topicService.GetByID(int64(req.TopicId))
		if err != nil || topic == nil || topic.Id != req.TopicId {
			//判断主题是不是存在
			response.FailJson(true, r, err.Error())
			return
		}
		//增加人数
		topicService.AddMember(int64(req.TopicId))

		//填写用户id
		req.UserId = topicUserService.GetLoginID(r)
		
		// 调用service中的添加函数添加
		err = membersService.AddSave(req)
		if err != nil {
			response.FailJson(true, r, err.Error())
		}
		response.SusJson(true, r, "添加成功")
	}
}

// 修改
func (c *Members) Edit(r *ghttp.Request) {
	// 如果是post提交的请求就执行修改操作
	if r.Method == "POST" {
		var editReq *membersModel.EditReq
		// 通过Parse方法解析获取参数
		err := r.Parse(&editReq)
		if err != nil {
			response.FailJson(true, r, err.(*gvalid.Error).FirstString())
		}
		err = membersService.EditSave(editReq)
		if err != nil {
			response.FailJson(true, r, err.Error())
		}
		response.SusJson(true, r, "修改参数成功")
	}
	// 不是post提交的请求就到修改页面后查询出要修改的记录
	id := r.GetInt("id")
	params, err := membersService.GetByID(int64(id))
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "ok", params)
}

// 修改
func (c *Members) CheckJoin(r *ghttp.Request) {
	//post
	if r.Method != "POST" {
		response.FailJson(true, r, "CheckJoin失败！")
		return
	} 
	var getReq *membersModel.GetReqByTopic
	// 通过Parse方法解析获取参数
	err := r.Parse(&getReq)
	if err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}
	//填写用户id
	getReq.UserId = topicUserService.GetLoginID(r)
	params, err := membersService.GetByTopicId(getReq) 

	chkResult := make(map[string]interface{})
	chkResult["result"] = 1
	if err != nil || params == nil {
		chkResult["result"] = 0
		response.SusJson(true, r, "ok", chkResult)
	} 
	response.SusJson(true, r, "ok", chkResult)

}

// 删除
func (c *Members) Delete(r *ghttp.Request) {
	//post
	if r.Method != "POST" {
		response.FailJson(true, r, "删除失败！")
		return
	}

	var req *membersModel.RemoveReqByTopic
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}
	if req.TopicId == 0 {
		response.FailJson(true, r, "no topic specific.")
		return
	}
	topic, err := topicService.GetByID(int64(req.TopicId))
	if err != nil || topic == nil || topic.Id != req.TopicId {
		//判断主题是不是存在
		response.FailJson(true, r, err.Error())
		return
	}
	//减少人数
	topicService.RemoveMember(int64(req.TopicId))

	//填写用户id
	req.UserId = topicUserService.GetLoginID(r)
	err = membersService.DeleteByTopicId(req)
	if err != nil {
		response.FailJson(true, r, "删除失败")
	}
	response.SusJson(true, r, "删除成功")
}
