// ==========================================================================
// 生成日期：2021-01-23 14:24:09
// 生成人：gfast
// ==========================================================================
package category_service

import (
	categoryModel "gfast/app/model/topic/bookcategory"
)

// 添加
func AddSave(req *categoryModel.AddReq) error {
	return categoryModel.AddSave(req)
}

// 删除
func DeleteByIds(Ids []int) error {
	return categoryModel.DeleteByIds(Ids)
}

//修改
func EditSave(editReq *categoryModel.EditReq) error {
	return categoryModel.EditSave(editReq)
}

// 根据ID查询
func GetByID(id int64) (*categoryModel.Entity, error) {
	return categoryModel.GetByID(id)
}

// 分页查询
func SelectListByPage(req *categoryModel.SelectPageReq) (total int, page int64, list []*categoryModel.Entity, err error) {
	return categoryModel.SelectListByPage(req)
}
