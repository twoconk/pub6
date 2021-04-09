//topictaskprogress.go

package topic

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
	taskprogressModel "gfast/app/model/topic/topictasksprogress"
	taskprogressService "gfast/app/service/topic/taskprogress_service" 
	"gfast/library/response"
)

//控制器
type Taskprogress struct{}


//列表页
func (c *Taskprogress) List(r *ghttp.Request) {
	// 定义一个结构体存储请求参数
	var req *taskprogressModel.SelectPageReq
	// 获取参数
	err := r.Parse(&req)
	if err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}
	total, page, list, err := taskprogressService.SelectListByPage(req)
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
func (c *Taskprogress) Add(r *ghttp.Request) {
	if r.Method == "POST" {
		var req *taskprogressModel.AddReq
		// 通过Parse方法解析获取参数
		err := r.Parse(&req)
		if err != nil {
			response.FailJson(true, r, err.(*gvalid.Error).FirstString())
		}
		//填写用户id
		//req.UserId = users_service.GetLoginID(r)
		
		// 调用service中的添加函数添加
		err = taskprogressService.AddSave(req)
		if err != nil {
			response.FailJson(true, r, err.Error())
		}
		response.SusJson(true, r, "添加成功")
	}
}


// 修改
func (c *Taskprogress) Edit(r *ghttp.Request) {
	// 如果是post提交的请求就执行修改操作
	if r.Method == "POST" {
		var editReq *taskprogressModel.EditReq
		// 通过Parse方法解析获取参数
		err := r.Parse(&editReq)
		if err != nil {
			response.FailJson(true, r, err.(*gvalid.Error).FirstString())
		}
		err = taskprogressService.EditSave(editReq)
		if err != nil {
		response.FailJson(true, r, err.Error())
		}
		response.SusJson(true, r, "修改参数成功")
	}
	// 不是post提交的请求就到修改页面后查询出要修改的记录
	id := r.GetInt("id")
	params, err := taskprogressService.GetByID(int64(id))
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "ok", params)
}


// 删除
func (c *Taskprogress) Delete(r *ghttp.Request) {
	var req *taskprogressModel.RemoveReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}
	err := taskprogressService.DeleteByIds(req.Ids)
	if err != nil {
	response.FailJson(true, r, "删除失败")
	}
	response.SusJson(true, r, "删除成功")
}