//topictask.go

package topic

import (
	taskModel "gfast/app/model/topic/topictasks"
	taskService "gfast/app/service/topic/task_service"
	//topicUserService "gfast/app/service/topic/users_service"
	"gfast/library/response"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
    "html"
)

//控制器
type Task struct{}

//列表页
func (c *Task) List(r *ghttp.Request) {
	// 定义一个结构体存储请求参数
	var req *taskModel.SelectPageReq
	// 获取参数
	err := r.Parse(&req)

	//判断是否有topic ID
	
	if err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}
	total, page, list, err := taskService.SelectListByPage(req)
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

curl -d '{"TopicId":1,"TaskTitle":"Go语言学习-1", "TaskContent":"好想学习Go语言，找个团", "status":0}' -H 'Authorization:Bearer vGCiSmcWADHgtiq2XeDTokcTCipNawAG8xZW1F+49fgzaNhaoop+XX1XHvm2qx715BB5VGUxeFlXqTL7WcE+c59q3k7TRk4Ai1P0E01dYeJlsyCA3Ss9vU066gLs4ipB' "http://127.0.0.1:8200/api.v2/topictask/add"


{"code":0,"data":3,"msg":"添加成功"}%
*/
// 新增
func (c *Task) Add(r *ghttp.Request) {
	if r.Method == "POST" {
		var req *taskModel.AddReq
		// 通过Parse方法解析获取参数
		err := r.Parse(&req)
		if err != nil {
			response.FailJson(true, r, err.(*gvalid.Error).FirstString())
		}
		//填写用户id
		//req.UserId = topicUserService.GetLoginID(r)
		//转义关键字段
		req.TaskTitle = html.EscapeString(req.TaskTitle)
		req.TaskContent = html.EscapeString(req.TaskContent) 
		
		// 调用service中的添加函数添加
		err = taskService.AddSave(req)
		if err != nil {
			response.FailJson(true, r, err.Error())
		}
		response.SusJson(true, r, "添加成功")
	}
}

// 修改
func (c *Task) Edit(r *ghttp.Request) {
	// 如果是post提交的请求就执行修改操作
	if r.Method == "POST" {
		var editReq *taskModel.EditReq
		// 通过Parse方法解析获取参数
		err := r.Parse(&editReq)
		if err != nil {
			response.FailJson(true, r, err.(*gvalid.Error).FirstString())
		}
		//判断该topictask是否是这个用户的

		//转义关键字段
		editReq.TaskTitle = html.EscapeString(editReq.TaskTitle)
		editReq.TaskContent = html.EscapeString(editReq.TaskContent) 

		err = taskService.EditSave(editReq)
		if err != nil {
			response.FailJson(true, r, err.Error())
		}
		response.SusJson(true, r, "修改参数成功")
	}
	// 不是post提交的请求就到修改页面后查询出要修改的记录
	id := r.GetInt("id")
	params, err := taskService.GetByID(int64(id))
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "ok", params)
}

// 删除
func (c *Task) Delete(r *ghttp.Request) {
	var req *taskModel.RemoveReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}
	err := taskService.DeleteByIds(req.Ids)
	if err != nil {
		response.FailJson(true, r, "删除失败")
	}
	response.SusJson(true, r, "删除成功")
}
