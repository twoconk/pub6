// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package plug_link

import (
	"gfast/app/model/admin/plug_linktype"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
)

// AddReq 用于存储新增链接请求的请求参数
type AddReq struct {
	LinkName     string `p:"linkName" v:"required#名称不能为空"`     // 链接名称
	LinkUrl      string `p:"linkUrl" v:"required#名称不能为空"`      // 链接URL
	LinkTarget   string `p:"linkTarget" `                      // 打开方式
	LinkTypeID   int    `p:"linkTypeID"`                       // 所属栏目ID
	LinkQQ       string `p:"linkQQ" v:"required#名称不能为空"`       // 联系QQ
	LinkOrder    int64  `p:"linkOrder"`                        // 排序
	LinkOpen     int    `p:"linkOpen"`                         // 0禁用1启用(是否审核)
	LinkUsername string `p:"linkUsername" v:"required#名称不能为空"` // 申请友情链接的联系人
	LinkEmail    string `p:"linkEmail"`                        // 联系邮箱
	LinkRemark   string `p:"linkRemark"`                       // 申请友情链接时的备注
}

// EditReq 用于存储修改广告位请求参数
type EditReq struct {
	PlugLinkID int64 `p:"plugLinkID" v:"required|min:1#广告id不能为空|广告id参数错误"`
	AddReq
}

// SelectPageReq 用于存储分页查询广告的请求参数
type SelectPageReq struct {
	LinkName string `p:"linkName"` // 广告名称
	PageNo   int64  `p:"pageNum"`  // 当前页
	PageSize int64  `p:"pageSize"` // 每页显示记录数
}

// 用于存储分页查询的数据
type ListEntity struct {
	Entity
	LinkTypeName string `orm:"linktype_name"      json:"linktype_name" ` // 友情链接所属分类
}

// GetPlugLinkByID 根据ID查询链接记录
func GetPlugLinkByID(id int64) (*Entity, error) {
	entity, err := Model.FindOne("link_id", id)
	if err != nil {
		g.Log().Error(err)
		return nil, gerror.New("根据ID查询链接记录出错")
	}
	if entity == nil {
		return nil, gerror.New("根据ID未能查询到链接记录")
	}
	return entity, nil
}

// AddSave 添加友情链接
func AddSave(req *AddReq) error {
	var entity Entity
	entity.LinkName = req.LinkName
	entity.LinkUrl = req.LinkUrl
	entity.LinkTarget = req.LinkTarget
	entity.LinkTypeid = req.LinkTypeID
	entity.LinkQq = req.LinkQQ
	entity.LinkAddtime = int(gtime.Timestamp()) // 添加时间
	entity.LinkOrder = req.LinkOrder
	entity.LinkOpen = req.LinkOpen
	entity.LinkUsername = req.LinkUsername
	entity.LinkEmail = req.LinkEmail
	entity.LinkRemark = req.LinkRemark
	// 保存实体
	_, err := entity.Insert()
	if err != nil {
		g.Log().Error(err)
		return gerror.New("保存失败")
	}
	return nil
}

// 根据ID批量删除链接
func DeleteByIDs(ids []int) error {
	_, err := Model.Delete("link_id in(?)", ids)
	if err != nil {
		g.Log().Error(err)
		return gerror.New("删除链接失败")
	}
	return nil
}

// 根据ID来修改链接信息
func EditSave(editReq *EditReq) error {
	// 先根据ID来查询要修改的链接记录
	entity, err := GetPlugLinkByID(editReq.PlugLinkID)
	if err != nil {
		return err
	}
	// 修改实体
	entity.LinkName = editReq.LinkName
	entity.LinkUrl = editReq.LinkUrl
	entity.LinkTarget = editReq.LinkTarget
	entity.LinkTypeid = editReq.LinkTypeID
	entity.LinkQq = editReq.LinkQQ
	entity.LinkOrder = editReq.LinkOrder
	entity.LinkOpen = editReq.LinkOpen
	entity.LinkUsername = editReq.LinkUsername
	entity.LinkEmail = editReq.LinkEmail
	entity.LinkRemark = editReq.LinkRemark
	_, err = Model.Save(entity)
	if err != nil {
		g.Log().Error(err)
		return gerror.New("修改栏目失败")
	}
	return nil
}

// 分页查询,返回值total总记录数,page当前页
func SelectListByPage(req *SelectPageReq) (total int, page int64, list []*ListEntity, err error) {
	model := g.DB().Table(Table + " link")
	if req != nil {
		if req.LinkName != "" {
			model.Where("link.link_name like ?", "%"+req.LinkName+"%")
		}
	}
	model = model.LeftJoin(plug_linktype.Table+" type", "type.linktype_id=link.link_typeid")
	// 查询友情链接总记录数(总行数)
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
	var res gdb.Result
	res, err = model.Fields("link.*,type.linktype_name").Page(int(page), int(req.PageSize)).Order("link.link_order asc,link.link_id asc").All()
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("分页查询友情链接失败")
		return 0, 0, nil, err
	}
	err = res.Structs(&list)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("分页查询广告失败")
		return 0, 0, nil, err
	}
	return total, page, list, nil
}

// 按链接分类查询当前分类下的size条最新链接(status:1启用,0未启用,优先序号排序，其次时间倒序)
func ListByTypeId(typeId int, size int, status int) (list []*Entity, err error) {
	list, err = Model.Where("link_typeid = ?", typeId).And("link_open = ?", status).Fields("link_name,link_url,link_target").Order("link_order asc,link_addtime desc").Limit(size).All()
	if err != nil {
		g.Log().Error(err)
		return nil, gerror.New("按分类查询链接出错")
	}
	return list, nil
}