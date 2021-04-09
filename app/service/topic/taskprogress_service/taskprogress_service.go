//taskprogress_service.go


package taskprogress_service

import (
	taskprogressModel "gfast/app/model/topic/topictasksprogress"
)


// 添加
func AddSave(req *taskprogressModel.AddReq) error {
	return taskprogressModel.AddSave(req)
}

// 删除
func DeleteByIds(Ids []int) error {
	return taskprogressModel.DeleteByIds(Ids)
}

//修改
func EditSave(editReq *taskprogressModel.EditReq) error {
	return taskprogressModel.EditSave(editReq)
}

// 根据ID查询
func GetByID(id int64) (*taskprogressModel.Entity, error) {
	return taskprogressModel.GetByID(id)
}

// 分页查询
func SelectListByPage(req *taskprogressModel.SelectPageReq) (total int, page int64, list []*taskprogressModel.Entity, err error) {
	return taskprogressModel.SelectListByPage(req)
}