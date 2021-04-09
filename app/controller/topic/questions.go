package topic

import (
	questionsModel "gfast/app/model/topic/questions"
	questionsService "gfast/app/service/topic/question_service"
	topicUserService "gfast/app/service/topic/users_service"
	"gfast/library/response"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

//控制器
type Questions struct{}

//列表页
func (c *Questions) List(r *ghttp.Request) {
	// 定义一个结构体存储请求参数
	var req *questionsModel.SelectPageReq
	// 获取参数
	err := r.Parse(&req)
	if err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}
	total, page, list, err := questionsService.SelectListByPage(req)
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
func (c *Questions) Add(r *ghttp.Request) {
	if r.Method == "POST" {
		var req *questionsModel.AddReq
		// 通过Parse方法解析获取参数
		err := r.Parse(&req)
		if err != nil {
			response.FailJson(true, r, err.(*gvalid.Error).FirstString())
		}
		//填写用户id
		req.CreateOwnerId = topicUserService.GetLoginID(r)
		// 调用service中的添加函数添加
		err = questionsService.AddSave(req)
		if err != nil {
			response.FailJson(true, r, err.Error())
		}
		response.SusJson(true, r, "添加成功")
	}
}

// 修改
func (c *Questions) Edit(r *ghttp.Request) {
	// 如果是post提交的请求就执行修改操作
	if r.Method == "POST" {
		var editReq *questionsModel.EditReq
		// 通过Parse方法解析获取参数
		err := r.Parse(&editReq)
		if err != nil {
			response.FailJson(true, r, err.(*gvalid.Error).FirstString())
		}
		err = questionsService.EditSave(editReq)
		if err != nil {
			response.FailJson(true, r, err.Error())
		}
		response.SusJson(true, r, "修改参数成功")
	}
	// 不是post提交的请求就到修改页面后查询出要修改的记录
	id := r.GetInt("id")
	params, err := questionsService.GetByID(int64(id))
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "ok", params)
}

// 删除
func (c *Questions) Delete(r *ghttp.Request) {
	var req *questionsModel.RemoveReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}
	params, _ := questionsService.GetByID(int64(req.Ids[0]))

	//id := r.GetInt("id")
	//获取登录ID
	userId := topicUserService.GetLoginID(r)
	if params != nil && userId!= params.CreateOwnerId{ 
		response.FailJson(true, r, "删除失败")
		return
	}
	
	err := questionsService.DeleteByIds(req.Ids)
	if err != nil {
		response.FailJson(true, r, "删除失败")
	}
	response.SusJson(true, r, "删除成功")
}
