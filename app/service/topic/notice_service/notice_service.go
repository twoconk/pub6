// ==========================================================================
// 生成日期：2021-01-23 14:24:09
// 生成人：gfast
// ==========================================================================
package notices_service

import (
	noticesModel "gfast/app/model/topic/topicnotices"

	"github.com/gogf/gf/os/gtime"
)

// 添加
func AddSave(req *noticesModel.AddReq) error {
	req.CreateTime = gtime.Now()

	return noticesModel.AddSave(req)
}

// 删除
func DeleteByIds(Ids []int) error {
	return noticesModel.DeleteByIds(Ids)
}

//修改
func EditSave(editReq *noticesModel.EditReq) error {
	return noticesModel.EditSave(editReq)
}

// 根据ID查询
func GetByID(id int64) (*noticesModel.Entity, error) {
	return noticesModel.GetByID(id)
}

// 分页查询
func SelectListByPage(req *noticesModel.SelectPageReq) (total int, page int64, list []*noticesModel.Entity, err error) {
	return noticesModel.SelectListByPage(req)
}
