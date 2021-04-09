// ==========================================================================
// 生成日期：2021-01-23 14:24:09
// 生成人：gfast
// ==========================================================================
package resource_service

import (
	resourceModel "gfast/app/model/topic/resource"
	"github.com/gogf/gf/os/gtime"
)

// 添加
func AddSave(req *resourceModel.AddReq) error {
	req.CreateTime = gtime.Now() 
	return resourceModel.AddSave(req)
}

// 删除
func DeleteByIds(Ids []int) error {
	return resourceModel.DeleteByIds(Ids)
}

//修改
func EditSave(editReq *resourceModel.EditReq) error {
	return resourceModel.EditSave(editReq)
}

// 根据ID查询
func GetByID(id int64) (*resourceModel.Entity, error) {
	return resourceModel.GetByID(id)
}

// 分页查询
func SelectListByPage(req *resourceModel.SelectPageReq) (total int, page int64, list []*resourceModel.ShowEntity, err error) {
	return resourceModel.SelectListByPage(req)
}
