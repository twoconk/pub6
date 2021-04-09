// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package resource

// ==========================================================================
// 生成日期：2021-01-23 14:24:09
// 生成人：gfast
// ==========================================================================
import (
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"database/sql"
)

// AddReq 用于存储新增请求的请求参数
type AddReq struct {
	TopicId         int64       `p:"topicId" `
	ResourceTypeId  int         `p:"resourceTypeId" `
	BookId          int         `p:"bookId" `
	UserId          int64       `p:"userId" `
	ResourceContent string      `p:"resourceContent" `
	ResourceLink string      `p:"resourceLink" `  
	Status          int         `p:"status" v:"required#是否开放状态（0正常 1 禁言）不能为空"`
	CreateTime      *gtime.Time `p:"createTime" `
	TopicTableIndex int         `p:"topicTableIndex" `
}

// EditReq 用于存储修改请求参数
type EditReq struct {
	Id              int64  `p:"id" v:"required#主键ID不能为空"`
	TopicId         int64  `p:"topicId" `
	ResourceTypeId  int    `p:"resourceTypeId" `
	BookId          int    `p:"bookId" `
	UserId          int64       `p:"userId" `
	ResourceContent string `p:"resourceContent" `
	ResourceLink string      `p:"resourceLink" `
	Status          int    `p:"status" v:"required#是否开放状态（0正常 1 禁言）不能为空"`
	TopicTableIndex int    `p:"topicTableIndex" `
}
type RemoveReq struct {
	Ids []int `p:"ids"` //删除id
}

// SelectPageReq 用于存储分页查询的请求参数
type SelectPageReq struct {
	TopicId         int64  `p:"topicId"`         //主题id
	ResourceTypeId  int    `p:"resourceTypeId"`  //关联资源类型，资源类型1.公开网课/2.付费网课/3.专业书籍/4.文档材料/5.其他类型
	BookId          int    `p:"bookId"`          //书的id
	UserId          int64       `p:"userId" `
	ResourceContent string `p:"resourceContent"` //资源内容，可以是资源链接
	ResourceLink string      `p:"resourceLink" `
	Status          int    `p:"status"`          //是否开放状态（0正常 1 禁言）
	TopicTableIndex int    `p:"topicTableIndex"` //表id 与表topic 相关
	BeginTime       string `p:"beginTime"`       //开始时间
	EndTime         string `p:"endTime"`         //结束时间
	PageNum         int    `p:"pageNum"`         //当前页码
	PageSize        int    `p:"pageSize"`        //每页数
}

// GetPlugAdByID 根据ID查询记录
func GetByID(id int64) (*Entity, error) {
	entity, err := Model.FindOne(id)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("根据ID查询记录出错")
	}
	if entity == nil {
		err = gerror.New("根据ID未能查询到记录")
	}
	return entity, nil
}

// AddSave 添加
func AddSave(req *AddReq) error {
	entity := new(Entity)
	//entity.Id = req.Id
	entity.TopicId = req.TopicId
	entity.ResourceTypeId = req.ResourceTypeId
	entity.BookId = req.BookId
	entity.UserId = req.UserId
	entity.ResourceLink = req.ResourceLink
	entity.ResourceContent = req.ResourceContent
	entity.Status = req.Status
	entity.CreateTime = req.CreateTime
	entity.TopicTableIndex = req.TopicTableIndex
	result, err := Model.Save(entity)
	if err != nil {
		return err
	}
	_, err = result.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

// 删除
func DeleteByIds(Ids []int) error {

	for _, id := range Ids {
		// 先根据ID来查询要修改的记录
		entity, err := GetByID(int64(id))
		if err != nil {
			return err
		}
		entity.Status = 3 //删除
		_, err = Model.Save(entity)
		if err != nil {
			g.Log().Error(err)
			return gerror.New("删除失败")
		}
	}
	return nil
}

// 根据ID来修改信息
func EditSave(req *EditReq) error {
	// 先根据ID来查询要修改的记录
	entity, err := GetByID(req.Id)
	if err != nil {
		return err
	}
	// 修改实体
	entity.TopicId = req.TopicId
	entity.ResourceTypeId = req.ResourceTypeId
	entity.BookId = req.BookId
	entity.ResourceLink = req.ResourceLink
	entity.ResourceContent = req.ResourceContent
	entity.Status = req.Status
	entity.TopicTableIndex = req.TopicTableIndex
	_, err = Model.Save(entity)
	if err != nil {
		g.Log().Error(err)
		return gerror.New("修改失败")
	}
	return nil
}

// 分页查询,返回值total总记录数,page当前页
func SelectListByPage(req *SelectPageReq) (total int, page int64, list []*ShowEntity, err error) {
	model := Model
	if req != nil {
		if req.TopicId != 0 {
			model = model.Where("topic_id = ?", req.TopicId)
		}
	}
	// 查询总记录数(总行数)
	total, err = model.Count()
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取总记录数失败")
		return
	}
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	page = int64(req.PageNum)
	if req.PageSize == 0 {
		req.PageSize = 10
	}
	// 分页排序查询
	// list, err = model.Page(int(page), int(req.PageSize)).Order("id desc").All()

  	//关联查询 , todo:性能优化 
    all, _ := g.DB().Table("guaniu_study_topic_resource notes").Where("notes.topic_id = ?", req.TopicId).InnerJoin("guaniu_study_users user", "notes.user_id=user.id").Fields("notes.id, notes.topic_id, notes.user_id, notes.resource_link, notes.resource_type_id, notes.resource_content, notes.create_time, user.name, user.ext").Page(int(page), int(req.PageSize)).Order("notes.id desc").FindAll()
 

	if err = all.Structs(&list); err != nil && err != sql.ErrNoRows {
		return
	}

	// list, err = model.Page(int(page), int(req.PageSize)).Order("id desc").All()
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("分页查询失败")
		return
	} 
	return
}

// 获取所有数据
func SelectListAll(req *SelectPageReq) (list []*Entity, err error) {
	model := Model
	if req != nil {
		if req.ResourceContent != "" {
			model.Where("resource_content = ?", req.ResourceContent)
		}
	}
	// 查询
	list, err = model.Order("id asc").All()
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("查询失败")
		return
	}
	return
}