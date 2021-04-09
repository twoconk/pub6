// ==========================================================================
// 生成日期：2021-01-23 14:24:09
// 生成人：gfast
// ==========================================================================
package topic

import (
	resourceModel "gfast/app/model/topic/resource"
	resourceService "gfast/app/service/topic/resource_service"
	topicService "gfast/app/service/topic/topic_service"
	topicUserService "gfast/app/service/topic/users_service"
	"gfast/library/response"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

//控制器
type Resource struct{}

//列表页
func (c *Resource) List(r *ghttp.Request) {
	// 定义一个结构体存储请求参数
	var req *resourceModel.SelectPageReq
	// 获取参数
	err := r.Parse(&req)
	if err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}
	total, page, list, err := resourceService.SelectListByPage(req)
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
func (c *Resource) Add(r *ghttp.Request) {
	if r.Method == "POST" {
		var req *resourceModel.AddReq
		// 通过Parse方法解析获取参数
		err := r.Parse(&req)
		if err != nil {
			response.FailJson(true, r, err.(*gvalid.Error).FirstString())
		}
		// 判断主题是不是存在
		id := r.GetInt64("topicId")
		if (id == 0){ 
			response.FailJson(true, r, "参数错误")
			return
		}
		topic, err := topicService.GetByID(int64(id))
		if err != nil || topic == nil || topic.Id != id {
			//判断主题是不是存在
			response.FailJson(true, r, err.Error())
			return
		}
		//填写用户id
		req.UserId = topicUserService.GetLoginID(r)
		// 调用service中的添加函数添加
		err = resourceService.AddSave(req)
		if err != nil {
			response.FailJson(true, r, err.Error())
		}
		response.SusJson(true, r, "添加成功")
	}
}

// 修改
func (c *Resource) Edit(r *ghttp.Request) {
	// 如果是post提交的请求就执行修改操作
	if r.Method == "POST" {
		var editReq *resourceModel.EditReq
		// 通过Parse方法解析获取参数
		err := r.Parse(&editReq)
		if err != nil {
			response.FailJson(true, r, err.(*gvalid.Error).FirstString())
		}
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
			params, err := resourceService.GetByID(int64(id))
			if err != nil {
				response.FailJson(true, r, err.Error())
			}
			if (userId != params.UserId){
				response.FailJson(true, r, err.Error()) 
			}
		}

		err = resourceService.EditSave(editReq)
		if err != nil {
			response.FailJson(true, r, err.Error())
		}
		response.SusJson(true, r, "修改参数成功")
	}
	// 不是post提交的请求就到修改页面后查询出要修改的记录
	id := r.GetInt("id")
	params, err := resourceService.GetByID(int64(id))
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "ok", params)
}

// 删除
func (c *Resource) Delete(r *ghttp.Request) {
	if r.Method == "POST" {
		response.SusJson(true, r, "error!")
		return;
	}
	var req *resourceModel.RemoveReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}
	//获取登录ID 
	params, _ := resourceService.GetByID(int64(req.Ids[0]))

	//id := r.GetInt("id")
	//获取登录ID
	userId := topicUserService.GetLoginID(r)
	if params != nil && userId!= params.UserId{ 
		response.FailJson(true, r, "删除失败")
		return
	}

	err := resourceService.DeleteByIds(req.Ids)
	if err != nil {
		response.FailJson(true, r, "删除失败")
	}
	response.SusJson(true, r, "删除成功")
}
