// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package plug_adtype

import (
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
)

// AddReq 用于存储新增广告位请求参数
type AddReq struct {
	AdtypeName string `p:"adTypeName" v:"required#广告位名称不能为空"` // 广告位名称
	AdtypeSort int    `p:"adTypeSort" v:"required#广告位排序不能为空"` // 广告位排序
}

// EditReq 用于存储修改广告位请求参数
type EditReq struct {
	AdtypeID int64 `p:"adTypeID" v:"required|min:1#主键ID不能为空|主键ID值错误"`
	AddReq
}

// SelectPageReq 用于存储分页查询广告位的请求参数
type SelectPageReq struct {
	AdtypeName string `p:"adTypeName"` // 广告位名称
	PageNo     int64  `p:"pageNum"`    // 当前页
	PageSize   int64  `p:"pageSize"`   // 每页显示记录数
}

// GetAdtypeByID 根据ID查询广告位记录
func GetAdtypeByID(id int64) (*Entity, error) {
	entity, err := Model.FindOne(id)
	if err != nil {
		g.Log().Error(err)
		return nil, gerror.New("根据ID查询广告位记录出错")
	}
	if entity == nil {
		return nil, gerror.New("根据ID未能查询到广告位记录")
	}
	return entity, nil
}

// 根据广告位的名称和ID来判断是否已存在相同名称的广告位
func CheakAdtypeNameUnique(adTypeName string, adTypeId int64) error {
	var (
		entity *Entity
		err    error
	)
	if adTypeId == 0 {
		entity, err = Model.FindOne(Columns.AdtypeName, adTypeName)
	} else {
		entity, err = Model.Where(Columns.AdtypeName, adTypeName).And(Columns.AdtypeId+"!=?", adTypeId).FindOne()
	}
	if err != nil {
		g.Log().Error(err)
		return gerror.New("校验广告位唯一性失败")
	}
	if entity != nil {
		return gerror.New("广告位名称已经存在")
	}
	return nil
}

// AddSave 添加广告位
func AddSave(req *AddReq) error {
	var entity Entity
	entity.AdtypeName = req.AdtypeName
	entity.AdtypeSort = req.AdtypeSort
	_, err := entity.Insert()
	if err != nil {
		g.Log().Error(err)
		return gerror.New("保存广告位失败")
	}
	return nil
}

// 根据广告位ID来删除广告位
func DeleteAdTypeByID(id []int) error {
	_, err := Model.Where("adtype_id in(?)", id).Delete()
	if err != nil {
		g.Log().Error(err)
		return gerror.New("删除广告位失败")
	}
	return nil
}

// 根据广告位ID来修改广告位信息
func EditSave(editReq *EditReq) error {
	// 先根据ID来查询要修改的广告位记录
	entity, err := GetAdtypeByID(editReq.AdtypeID)
	if err != nil {
		g.Log().Error(err)
		return gerror.New("查询要修改的记录时出错")
	}
	// 修改实体
	entity.AdtypeName = editReq.AdtypeName
	entity.AdtypeSort = editReq.AdtypeSort
	_, err = Model.Save(entity)
	if err != nil {
		g.Log().Error(err)
		return gerror.New("修改广告位失败")
	}
	return nil
}

// 分页查询,返回值total总记录数,page当前页
func SelectListByPage(req *SelectPageReq) (total int, page int64, list []*Entity, err error) {
	model := Model
	if req != nil {
		if req.AdtypeName != "" {
			model = model.Where("adtype_name like ?", "%"+req.AdtypeName+"%")
		}
	}
	// 查询广告位总记录数(总行数)
	total, err = model.Count()
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取总记录数失败")
		return 0, 0, nil, err
	}
	if req.PageNo == 0 {
		req.PageNo = 1
	}
	page = req.PageNo
	if req.PageSize == 0 {
		req.PageSize = 10
	}
	// 分页排序查询
	list, err = model.Page(int(page), int(req.PageSize)).Order("adtype_sort asc,adtype_id asc").All()
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("分页查询广告位失败")
		return 0, 0, nil, err
	}
	return total, page, list, nil
}
