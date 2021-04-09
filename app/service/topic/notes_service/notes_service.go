// ==========================================================================
// 生成日期：2021-01-23 14:24:09
// 生成人：gfast
// ==========================================================================
package notes_service

import (
	notesModel "gfast/app/model/topic/topicnotes"

	"github.com/gogf/gf/os/gtime"
)

// 添加
func AddSave(req *notesModel.AddReq) error {
	req.CreateTime = gtime.Now()
	req.ModifyTime = gtime.Now()
	return notesModel.AddSave(req)
}

// 删除
func DeleteByIds(Ids []int) error {
	return notesModel.DeleteByIds(Ids)
}

//修改
func EditSave(editReq *notesModel.EditReq) error {
	editReq.ModifyTime = gtime.Now()
	return notesModel.EditSave(editReq)
}

// 根据ID查询
func GetByID(id int64) (*notesModel.Entity, error) {
	return notesModel.GetByID(id)
}

// 分页查询
func SelectListByPage(req *notesModel.SelectPageReq) (total int, page int64, list []*notesModel.ShowEntity, err error) {
	return notesModel.SelectListByPage(req)
}
