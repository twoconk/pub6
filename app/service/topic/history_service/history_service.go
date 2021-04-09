// ==========================================================================
// 生成日期：2021-01-23 14:24:09
// 生成人：gfast
// ==========================================================================
package history_service

import (
	historyModel "gfast/app/model/topic/answerhistory"
)

// 添加
func AddSave(req *historyModel.AddReq) error {
	return historyModel.AddSave(req)
}

// 删除
func DeleteByIds(Ids []int) error {
	return historyModel.DeleteByIds(Ids)
}

//修改
func EditSave(editReq *historyModel.EditReq) error {
	return historyModel.EditSave(editReq)
}

// 根据ID查询
func GetByID(id int64) (*historyModel.Entity, error) {
	return historyModel.GetByID(id)
}

// 分页查询
func SelectListByPage(req *historyModel.SelectPageReq) (total int, page int64, list []*historyModel.Entity, err error) {
	return historyModel.SelectListByPage(req)
}
