// ==========================================================================
// 生成日期：2021-01-23 14:24:09
// 生成人：gfast
// ==========================================================================
package detail_service

import (
	detailModel "gfast/app/model/topic/answerhistorydetail"
)

// 添加
func AddSave(req *detailModel.AddReq) error {
	return detailModel.AddSave(req)
}

// 删除
func DeleteByIds(Ids []int) error {
	return detailModel.DeleteByIds(Ids)
}

//修改
func EditSave(editReq *detailModel.EditReq) error {
	return detailModel.EditSave(editReq)
}

// 根据ID查询
func GetByID(id int64) (*detailModel.Entity, error) {
	return detailModel.GetByID(id)
}

// 分页查询
func SelectListByPage(req *detailModel.SelectPageReq) (total int, page int64, list []*detailModel.Entity, err error) {
	return detailModel.SelectListByPage(req)
}
