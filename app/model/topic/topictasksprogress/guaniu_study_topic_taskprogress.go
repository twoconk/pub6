// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package topictasksprogress

// Fill with you ideas below.
import (
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
)

// AddReq 用于存储新增请求的请求参数
type AddReq struct {
	TopicId         int64       `p:"topicId" `
	TaskId          int64       `p:"taskId" `
	UserId          int64       `p:"userId" `
	Content         string      `p:"content" `
	Status          int         `p:"status" v:"required#状态（0开始 1 完成）不能为空"`
	CreateOwnerId   int         `p:"createOwnerId" `
	ModifyTime      *gtime.Time `p:"modifyTime" `
	CreateTime      *gtime.Time `p:"createTime" `
	TopicTableIndex int         `p:"topicTableIndex" `
}

// EditReq 用于存储修改请求参数
type EditReq struct {
	Id              int64       `p:"id" v:"required#主键ID不能为空"`
	TopicId         int64       `p:"topicId" `
	TaskId          int64       `p:"taskId" `
	UserId          int64       `p:"userId" `
	Content         string      `p:"content" `
	Status          int         `p:"status" v:"required#状态（0开始 1 完成）不能为空"`
	CreateOwnerId   int         `p:"createOwnerId" `
	ModifyTime      *gtime.Time `p:"modifyTime" `
	TopicTableIndex int         `p:"topicTableIndex" `
}
type RemoveReq struct {
	Ids []int `p:"ids"` //删除id
}

// SelectPageReq 用于存储分页查询的请求参数
type SelectPageReq struct {
	TopicId         int64       `p:"topicId"`         //主题id
	TaskId          int64       `p:"taskId"`          //任务id
	UserId          int64       `p:"userId"`          //用户id
	Content         string      `p:"content"`         //进度补充描述
	Status          int         `p:"status"`          //状态（0开始 1 完成）
	CreateOwnerId   int         `p:"createOwnerId"`   //创建者
	ModifyTime      *gtime.Time `p:"modifyTime"`      //时间
	TopicTableIndex int         `p:"topicTableIndex"` //表id 与表topic 相关
	BeginTime       string      `p:"beginTime"`       //开始时间
	EndTime         string      `p:"endTime"`         //结束时间
	PageNum         int64       `p:"pageNum"`         //当前页码
	PageSize        int         `p:"pageSize"`        //每页数
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
	entity.TaskId = req.TaskId
	entity.UserId = req.UserId
	entity.Content = req.Content
	entity.Status = req.Status
	entity.CreateTime = req.CreateTime
	entity.CreateOwnerId = req.CreateOwnerId
	entity.ModifyTime = req.ModifyTime
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
	_, err := Model.Delete("id in(?)", Ids)
	if err != nil {
		g.Log().Error(err)
		return gerror.New("删除失败")
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
	entity.TaskId = req.TaskId
	entity.UserId = req.UserId
	entity.Content = req.Content
	entity.Status = req.Status
	entity.CreateOwnerId = req.CreateOwnerId
	entity.ModifyTime = req.ModifyTime
	entity.TopicTableIndex = req.TopicTableIndex
	_, err = Model.Save(entity)
	if err != nil {
		g.Log().Error(err)
		return gerror.New("修改失败")
	}
	return nil
}

// 分页查询,返回值total总记录数,page当前页
func SelectListByPage(req *SelectPageReq) (total int, page int64, list []*Entity, err error) {
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
	page = req.PageNum
	if req.PageSize == 0 {
		req.PageSize = 10
	}
	// 分页排序查询
	list, err = model.Page(int(page), int(req.PageSize)).Order("id asc").All()
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
		if req.Content != "" {
			model.Where("content = ?", req.Content)
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
