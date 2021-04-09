// ==========================================================================
// 生成日期：2021-01-23 14:24:09
// 生成人：gfast
// ==========================================================================
package questions_service

import (
	questionsModel "gfast/app/model/topic/questions"

	"github.com/gogf/gf/os/gtime"
)

// 添加
func AddSave(req *questionsModel.AddReq) error {
	req.CreateTime = gtime.Now()
	req.ModifyTime = gtime.Now()
	return questionsModel.AddSave(req)
}

// 删除
func DeleteByIds(Ids []int) error {
	return questionsModel.DeleteByIds(Ids)
}

//修改
func EditSave(editReq *questionsModel.EditReq) error {
	editReq.ModifyTime = gtime.Now()
	return questionsModel.EditSave(editReq)
}

// 根据ID查询
func GetByID(id int64) (*questionsModel.Entity, error) {
	return questionsModel.GetByID(id)
}

// 分页查询
func SelectListByPage(req *questionsModel.SelectPageReq) (total int, page int64, list []*questionsModel.Entity, err error) {
	return questionsModel.SelectListByPage(req)
}
