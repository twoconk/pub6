// ==========================================================================
// 生成日期：2021-01-24 17:09:40
// 生成人：gfast
// ==========================================================================
package task_service

import (
	taskModel "gfast/app/model/topic/topictasks"

	"github.com/gogf/gf/os/gtime"
)

// 添加
func AddSave(req *taskModel.AddReq) error {
	req.ModifyTime = gtime.Now()
	req.CreateTime = gtime.Now()

	return taskModel.AddSave(req)
}

// 删除
func DeleteByIds(Ids []int) error {
	return taskModel.DeleteByIds(Ids)
}

//修改
func EditSave(editReq *taskModel.EditReq) error {
	editReq.ModifyTime = gtime.Now()
	return taskModel.EditSave(editReq)
}

// 根据ID查询
func GetByID(id int64) (*taskModel.Entity, error) {
	return taskModel.GetByID(id)
}

// 分页查询
func SelectListByPage(req *taskModel.SelectPageReq) (total int, page int64, list []*taskModel.Entity, err error) {
	return taskModel.SelectListByPage(req)
}
