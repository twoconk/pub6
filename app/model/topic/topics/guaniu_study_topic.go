// ==========================================================================
// 生成日期：2021-01-13 15:10:25
// 生成人：gfast
// ==========================================================================
package topics

import (
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/os/gtime"
    "github.com/gogf/gf/os/gcache"
	"github.com/gogf/gf/util/gconv"
	"database/sql"
    "time" 
)

// AddReq 用于存储新增请求的请求参数
type AddReq struct {
	ParentId        int64       `p:"parentId" `
	TopicType       int         `p:"topicType" `
	TopicOwnerType       int         `p:"topicOwnerType" `
	TopicName       string      `p:"topicName" v:"required#主题名称不能为空"`
	TopicPinyin     string      `p:"topicPinyin" `
	TopicImg        string      `p:"topicImg" `
	TopicContent    string      `p:"topicContent v:"required#主题内容不能为空"" `
	OrderNum        int         `p:"orderNum" `
	Status          int         `p:"status" v:"required#是否开放状态（1正常 0 关闭）不能为空"`
	CreateOwnerId   int64      `p:"createOwnerId" `
	ModifyTime      *gtime.Time `p:"modifyTime" `
	CreateTime      *gtime.Time `p:"createTime" `
	MembersNumber   int         `p:"membersNumber" `
	TopicTableIndex int         `p:"topicTableIndex" `
	NeedVerifyCode int         `p:"needVerifyCode" `
}

type CommentReq struct {
	TopicId              int64       `p:"id" v:"required#topicId不能为空"`
	UserId              int64       `p:"userId"`
	Status          int         `p:"status"` //0:显示 1:不显示
	Content    string      `p:"content" `
	CreateTime      *gtime.Time `orm:"create_time"       json:"create_time"`       // 时间
	NeedVerifyCode int         `p:"needVerifyCode" ` 
}

type NoticeReq struct {
	TopicId              int64       `p:"topicId" v:"required#topicId不能为空"`
	Content    string      `p:"content" `

}

type CommenTopicReq struct {
	TopicId              int64       `p:"topicId" v:"required#topicId不能为空"`
	PageNum         int         `p:"pageNum"`         //当前页码
	PageSize        int         `p:"pageSize"`        //每页数
	
}
type UserTopicReq struct {
	UserId              int64       `p:"userId" v:"required#userId不能为空"`
	PageNum         int         `p:"pageNum"`         //当前页码
	PageSize        int         `p:"pageSize"`        //每页数
	
}
 

// EditReq 用于存储修改请求参数
type EditReq struct {
	Id              int64       `p:"id" v:"required#主键ID不能为空"`
	ParentId        int64       `p:"parentId" `
	TopicType       int         `p:"topicType" `
	TopicName       string      `p:"topicName" v:"required#主题名称不能为空"`
	TopicPinyin     string      `p:"topicPinyin" `
	TopicImg        string      `p:"topicImg" `
	TopicContent    string      `p:"topicContent" `
	OrderNum        int         `p:"orderNum" `
	Status          int         `p:"status" v:"required#是否开放状态（0正常 1 关闭）不能为空"`
	CreateOwnerId   int64      `p:"createOwnerId" `
	ModifyTime      *gtime.Time `p:"modifyTime" `
	MembersNumber   int         `p:"membersNumber" `
	TopicTableIndex int         `p:"topicTableIndex" `
	NeedVerifyCode int         `p:"needVerifyCode" `
}
type RemoveReq struct {
	Ids []int `p:"ids"` //删除id
}

// SelectPageReq 用于存储分页查询的请求参数
type SelectPageReq struct {
	ParentId        int64       `p:"parentId"`        //父主题id
	TopicType       int         `p:"topicType"`       //主题类型
	TopicName       string      `p:"topicName"`       //主题名称
	TopicPinyin     string      `p:"topicPinyin"`     //主题名称拼音
	TopicImg        string      `p:"topicImg"`        //主题图片路径
	TopicContent    string      `p:"topicContent"`    //主题描述
	OrderNum        int         `p:"orderNum"`        //显示顺序
	Status          string      `p:"status"`          //是否开放状态（0正常 1 关闭）
	CreateOwnerId   int64      `p:"createOwnerId"`   //创建者
	ModifyTime      *gtime.Time `p:"modifyTime"`      //时间
	MembersNumber   int         `p:"membersNumber"`   //加入成员数
	TopicTableIndex int         `p:"topicTableIndex"` //表id 与表topic 相关
	BeginTime       string      `p:"beginTime"`       //开始时间
	EndTime         string      `p:"endTime"`         //结束时间
	PageNum         int         `p:"pageNum"`         //当前页码
	PageSize        int         `p:"pageSize"`        //每页数
	Name   string      `orm:"name"       json:"name"`   // 用户名
	Ext    string      `orm:"ext"        json:"ext"`    // 扩展字段
}


// Entity is the golang structure for table guaniu_study_users.
type EntityUser struct {
	Id     int64      `orm:"id,primary" json:"id"`     // 主键
	Name   string      `orm:"name"       json:"name"`   // 用户名
	Ext    string      `orm:"ext"        json:"ext"`    // 扩展字段
	Ctime  int         `orm:"ctime"      json:"ctime"`  // 创建时间 
	UserProfile interface{} `json:"user_profile"` //临时增加对象
}
// Entity is the golang structure for table guaniu_study_users.
type EntityUserProfile struct {
	Id     int64      `orm:"id,primary" json:"id"`     // 主键
	UserId   int64      `orm:"user_id"       json:"user_id"`   // 用户id
	AvaterPath    string      `orm:"avater_path"        json:"avater_path"`    // 头像路径
	AvaterThumbnailPath    string      `orm:"avater_thumbnail_path"        json:"avater_thumbnail_path"`    // 头像路径
	CurrentJobCompany    string      `orm:"current_job_company"        json:"current_job_company"`    //  
	HomebgPath    string      `orm:"home_bg_path"        json:"home_bg_path"`    //  
	SignNotes    string      `orm:"sign_notes"        json:"sign_notes"`    //  
	WeiboId    string      `orm:"weibo_id"        json:"weibo_id"`    //  
	WeixinId    string      `orm:"weixin_id"        json:"weixin_id"`    //  
	GithubId    string      `orm:"github_id"        json:"github_id"`    //  
	QQId    string      `orm:"qq_id"        json:"qq_id"`    //  
	DoubanId    string      `orm:"douban_id"        json:"douban_id"`    //  
	CsdnId    string      `orm:"csdn_id"        json:"csdn_id"`    //  
	Site    string      `orm:"blog_site"        json:"blog_site"`    //  
	UserStars  int         `orm:"user_stars"      json:"user_stars"`  // 用户积分 
}

type EntityTopic struct {
	Id              int64       `orm:"id,primary"        json:"id"`                // 编号
	ParentId        int64       `orm:"parent_id"         json:"parent_id"`         // 父主题id
	TopicType       int         `orm:"topic_type"        json:"topic_type"`        // 主题类型
	TopicOwnerType       int         `orm:"topic_owner_type"        json:"topic_owner_type"`        // 主题类型
	TopicName       string      `orm:"topic_name"        json:"topic_name"`        // 主题名称
	TopicPinyin     string      `orm:"topic_pinyin"      json:"topic_pinyin"`      // 主题名称拼音
	TopicImg        string      `orm:"topic_img"         json:"topic_img"`         // 主题图片路径
	TopicContent    string      `orm:"topic_content"     json:"topic_content"`     // 主题描述
	OrderNum        int         `orm:"order_num"         json:"order_num"`         // 显示顺序
	Status          int         `orm:"status"            json:"status"`            // 是否开放状态（0正常 1 关闭）
	CreateOwnerId   int64      `orm:"create_owner_id"   json:"create_owner_id"`   // 创建者
	CreateTime      *gtime.Time `orm:"create_time"       json:"create_time"`       // 时间
	ModifyTime      *gtime.Time `orm:"modify_time"       json:"modify_time"`       // 时间
	MembersNumber   int         `orm:"members_number"    json:"members_number"`    // 加入成员数
	TopicTableIndex int         `orm:"topic_table_index" json:"topic_table_index"` // 表id 与表topic 相关 
	User interface{} `json:"create_owner"` //临时增加对象
}
// Entity is the golang structure for table guaniu_study_topic.
type CommentShowEntity struct {
	Id              int64       `orm:"id,primary"        json:"id"`                // 编号 
	TopicId       int64         `orm:"topic_id"        json:"topic_id"`        // 主题id
	UserId       int64         `orm:"user_id"        json:"user_id"`        // 主题类型
	Content    string      `orm:"content"     json:"content"`     // 主题描述
	OrderSeq        int         `orm:"orderSeq"         json:"orderSeq"`         // 显示顺序
	LikeNum        int         `orm:"likeNum"         json:"likeNum"`         // 显示顺序
	Status          int         `orm:"status"            json:"status"`            // 是否开放状态（0正常 1 关闭）
	CreateTime      *gtime.Time `orm:"create_time"       json:"create_time"`       // 时间
	TopicTableIndex int         `orm:"topic_table_index" json:"topic_table_index"` // 表id 与表topic 相关
	
 
	Name   string      `orm:"name"       json:"name"`   // 用户名
	Ext    string      `orm:"ext"        json:"ext"`    // 扩展字段 
}
type MemberEntity struct {
    Id              int64       `orm:"id,primary"        json:"id"`                // 编号                          
    TopicId         int64       `orm:"topic_id"          json:"topic_id"`          // 主题id                        
    UserId          int64       `orm:"user_id"           json:"user_id"`           // 成员id                        
    Status          int         `orm:"status"            json:"status"`            // 是否开放状态（0正常 1 禁言）  
    AdminFlag       int         `orm:"admin_flag"        json:"admin_flag"`        // 是否管理员（0不是 1 管理员）  
    CreateTime      *gtime.Time `orm:"create_time"       json:"create_time"`       // 时间                          
    TopicTableIndex int         `orm:"topic_table_index" json:"topic_table_index"` // 表id 与表topic 相关           
}
// GetPlugAdByID 根据ID查询记录
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

func GetByIDForDetail(id int64) (entity *ShowEntity, err error) {
	//entity, err := Model.FindOne(id)

    one, _ := g.DB().Table("guaniu_study_topic topics").Where("topics.id = ?", id).InnerJoin("guaniu_study_users user", "topics.create_owner_id=user.id").Fields("topics.*, user.name, user.ext").FindOne()

 
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("根据ID查询记录出错")
	}
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if entity == nil {
		err = gerror.New("根据ID未能查询到记录")
		return nil, err
	}

	entity2, err := Model.FindOne(id)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("根据ID查询记录出错")
	}
	//修改
	entity2.SeeNum += 1
	_, err = Model.Save(entity2) 

	// user := g.DB().Model("guaniu_study_users") 
	// entity.User, err  = user.Fields("id, name, ext").FindOne(entity.CreateOwnerId);

	return entity, nil
}

//增加成员数
func AddMember(id int64) error {

	entity2, err := Model.FindOne(id)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("根据ID查询记录出错")
	}
	//修改
	entity2.MembersNumber += 1
	_, err = Model.Save(entity2) 
	return nil
}
//减少成员数
func RemoveMember(id int64) error {

	entity2, err := Model.FindOne(id)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("根据ID查询记录出错")
	}
	//修改
	entity2.MembersNumber -= 1
	if (entity2.MembersNumber <= 1){
		entity2.MembersNumber = 1;
	}
	_, err = Model.Save(entity2) 
	return nil
}

func GetByIDForEdit(id int64) (*Entity, error) {
	entity, err := Model.FindOne(id)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("根据ID查询记录出错")
	}
	if entity == nil {
		err = gerror.New("根据ID未能查询到记录")
	} 
	//entity.User = nil
	
	return entity, nil
}

// AddSave 添加
func AddSave(req *AddReq) (id int64, err error) {
	entity := new(SaveEntity)
	//entity.Id = req.Id
	entity.ParentId = req.ParentId
	entity.TopicType = req.TopicType
	entity.TopicOwnerType = req.TopicOwnerType
	entity.TopicName = req.TopicName
	entity.TopicPinyin = req.TopicPinyin
	entity.TopicImg = req.TopicImg
	entity.TopicContent = req.TopicContent
	entity.OrderNum = req.OrderNum
	entity.SeeNum = 1
	entity.Status = req.Status
	entity.CreateOwnerId = req.CreateOwnerId
	entity.ModifyTime = req.ModifyTime
	entity.CreateTime = req.CreateTime
	entity.MembersNumber = req.MembersNumber
	entity.TopicTableIndex = req.TopicTableIndex
	result, err := Model.Save(entity)
	if err != nil {
		return 0, err
	}
	id, err = result.LastInsertId()
	if err != nil {
		return id, err
	}
	return id, nil
}

// 删除
func DeleteByIds(Ids []int) error {
	// _, err := Model.Delete("id in(?)", Ids)
	// if err != nil {
	// 	g.Log().Error(err)
	// 	return gerror.New("删除失败")
	// }
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
	entity, err := GetByIDForEdit(req.Id)
	if err != nil {
		return err
	}
	// 修改实体
	entity.ParentId = req.ParentId
	entity.TopicType = req.TopicType
	entity.TopicName = req.TopicName
	entity.TopicPinyin = req.TopicPinyin
	entity.TopicImg = req.TopicImg
	entity.TopicContent = req.TopicContent
	if (req.OrderNum != 0){
		entity.OrderNum = req.OrderNum
	}
	if (req.Status != 0){
		entity.Status = req.Status 
	}
	//entity.CreateOwnerId = req.CreateOwnerId
	entity.ModifyTime = req.ModifyTime
	//entity.MembersNumber = req.MembersNumber
	entity.TopicTableIndex = req.TopicTableIndex
	_, err = Model.Save(entity)
	if err != nil {
		g.Log().Error(err)
		return gerror.New("修改失败")
	}
	return nil
}

func TopicNotice(req *NoticeReq) error {
	//guaniu_study_topic_notices
	//主题公告

	return nil
}

func TopicComment(req *CommentReq) error {
	//guaniu_study_topic_notices
	//主题评论
	//guaniu_study_topic_comments
	commentEntity := new(CommentEntity)

	commentEntity.TopicId = req.TopicId
	commentEntity.UserId = req.UserId
	commentEntity.Content = req.Content
	commentEntity.CreateTime = req.CreateTime

	commentModel := g.DB().Table("guaniu_study_topic_comments") 
	commentModel.Save(commentEntity)

	return nil
}

func TopicNoticeLike(req *CommenTopicReq) error {
	//guaniu_study_topic_notices
	//主题评论

	return nil
}

func TopicJoin(req *CommenTopicReq) error {
	//guaniu_study_topic_notices
	//主题评论

	return nil
}

// 分页查询,返回值total总记录数,page当前页
func SelectCommentListByPage(req *CommenTopicReq) (total int, page int64, list []*CommentShowEntity, err error) {
	model := g.DB().Table("guaniu_study_topic_comments") 
 

	if req != nil { 
		model = model.Where("status = 0 and topic_id = ?", req.TopicId) 
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
	//all, _ := model.Page(int(page), int(req.PageSize)).Order("id DESC").FindAll()
 
  	//关联查询 , todo:性能优化 
    all, _ := g.DB().Table("guaniu_study_topic_comments comments").Where("comments.topic_id = ? and comments.status = 0", req.TopicId).InnerJoin("guaniu_study_users user", "comments.user_id=user.id").Fields("comments.*, user.name, user.ext").Page(int(page), int(req.PageSize)).Order("comments.id asc").FindAll()
 

	if err = all.Structs(&list); err != nil && err != sql.ErrNoRows {
		return
	}
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("分页查询失败")
		return
	}
   

	return
}

// 分页查询,返回值total总记录数,page当前页
func SelectListByPageAndUserId(req *UserTopicReq) (total int, page int64, list []*ShowEntity, err error) {
	if req == nil || req.UserId == 0 {
		//不能没有Topicid
		err = gerror.New("获取总记录数失败")
		return
	}
	model := g.DB().Table("guaniu_study_topic_members")
 
	if req != nil { 
		model = model.Where("user_id = ?", req.UserId) 
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

	topiclist := ([]*MemberEntity)(nil)
  	//关联查询 , todo:性能优化 
	topiclistall, _ := model.Page(int(page), int(req.PageSize)).Order("id DESC").FindAll()
	err = topiclistall.Structs(&topiclist)  	//关联查询 , todo:性能优化 
	if topiclist == nil { 
		return
	}

	//all, _ := g.DB().Table("guaniu_study_topic topics").Where("topics.status = 1 or topics.status = 3").InnerJoin("guaniu_study_users user", "topics.create_owner_id=user.id").InnerJoin("guaniu_study_topic_members members", "topics.create_owner_id = members.topic_id").Where("members.user_id = ?", req.UserId).Fields("topics.id, topics.topic_type, topics.topic_owner_type, topics.topic_name, topics.topic_tag, topics.topic_img, topics.status, topics.create_owner_id, topics.create_time, topics.modify_time, topics.members_number, topics.see_num, user.name, user.ext").Page(int(page), int(req.PageSize)).Order("topics.id desc").FindAll()

	all, _ := g.DB().Table("guaniu_study_topic topics").Where("topics.status = 1 or topics.status = 3").Where("topics.id", gdb.ListItemValuesUnique(topiclist, "TopicId")).InnerJoin("guaniu_study_users user", "topics.create_owner_id=user.id").Fields("topics.id, topics.topic_type, topics.topic_owner_type, topics.topic_name, topics.topic_tag, topics.topic_img, topics.status, topics.create_owner_id, topics.create_time, topics.modify_time, topics.members_number, topics.see_num, user.name, user.ext").Order("topics.id desc").FindAll()
	if err = all.Structs(&list); err != nil && err != sql.ErrNoRows {
		return
	}
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("分页查询失败")
		return
	}

	return
}

// 分页查询,返回值total总记录数,page当前页
func SelectListByPage(req *SelectPageReq) (total int, page int64, list []*ShowEntity, err error) {
	model := Model

	//增加约束
	//model = model.Fileds("id,topic_type,topic_owner_type,topic_name, topic_content, order_num, topic_img, create_time, modify_time")

	setUserId := false

	status :=  0
	orderNum :=  0
	if req != nil {
		if req.TopicName != "" {
			model = model.Where("topic_name like ?", "%"+req.TopicName+"%")
		}
		if req.TopicPinyin != "" {
			model = model.Where("topic_pinyin = ?", req.TopicPinyin)
		}
		if req.OrderNum != 0 {
			model = model.Where("order_num = ?", req.OrderNum)
			orderNum =  req.OrderNum
		} 
		if req.Status != "" {
			model = model.Where("status = ?", req.Status)
			status =  gconv.Int(req.Status)
		}else{
			model = model.Where("status = 1")
			status =  1
		}

		if (req.CreateOwnerId != 0){
			//
			model = model.Where("create_owner_id = ?", req.CreateOwnerId)
			setUserId = true
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

	//设置缓存

	
	// 分页排序查询
	//list, err = model.Page(int(page), int(req.PageSize)).Order("id DESC").All()  	//关联查询 , todo:性能优化 
	if (status == 0){
		if (orderNum != 0){
			if (!setUserId){
	    		all, _ := g.DB().Table("guaniu_study_topic topics").Where("(topics.status = 1 or topics.status = 3) and topics.order_num = ?", orderNum).InnerJoin("guaniu_study_users user", "topics.create_owner_id=user.id").Fields("topics.id, topics.topic_type, topics.topic_owner_type, topics.topic_name, topics.topic_tag, topics.topic_img, topics.status, topics.create_owner_id, topics.create_time, topics.modify_time, topics.members_number, topics.see_num, user.name, user.ext").Page(int(page), int(req.PageSize)).Order("topics.id desc").FindAll()
			
				if err = all.Structs(&list); err != nil && err != sql.ErrNoRows {
					return
				}
			}else{
		    	all, _ := g.DB().Table("guaniu_study_topic topics").Where("topics.create_owner_id = ?  and (topics.status = 1 or topics.status = 3) and topics.order_num = ?",req.CreateOwnerId, orderNum).InnerJoin("guaniu_study_users user", "topics.create_owner_id=user.id").Fields("topics.id, topics.topic_type, topics.topic_owner_type, topics.topic_name, topics.topic_tag, topics.topic_img, topics.status, topics.create_owner_id, topics.create_time, topics.modify_time, topics.members_number, topics.see_num, user.name, user.ext").Page(int(page), int(req.PageSize)).Order("topics.id desc").FindAll()
			
				if err = all.Structs(&list); err != nil && err != sql.ErrNoRows {
					return
				}
			}
		}else{

			if (!setUserId){
	    		all, _ := g.DB().Table("guaniu_study_topic topics").Where("topics.status = 1 or topics.status = 3").InnerJoin("guaniu_study_users user", "topics.create_owner_id=user.id").Fields("topics.id, topics.topic_type, topics.topic_owner_type, topics.topic_name, topics.topic_tag, topics.topic_img, topics.status, topics.create_owner_id, topics.create_time, topics.modify_time, topics.members_number, topics.see_num, user.name, user.ext").Page(int(page), int(req.PageSize)).Order("topics.id desc").FindAll()
			
				if err = all.Structs(&list); err != nil && err != sql.ErrNoRows {
					return
				}
			}else{
		    	all, _ := g.DB().Table("guaniu_study_topic topics").Where("topics.create_owner_id = ?  and (topics.status = 1 or topics.status = 3)",req.CreateOwnerId).InnerJoin("guaniu_study_users user", "topics.create_owner_id=user.id").Fields("topics.id, topics.topic_type, topics.topic_owner_type, topics.topic_name, topics.topic_tag, topics.topic_img, topics.status, topics.create_owner_id, topics.create_time, topics.modify_time, topics.members_number, topics.see_num, user.name, user.ext").Page(int(page), int(req.PageSize)).Order("topics.id desc").FindAll()
			
				if err = all.Structs(&list); err != nil && err != sql.ErrNoRows {
					return
				}
			}
		}
	}else{
		if (orderNum != 0){

			if (!setUserId){
	    		all, _ := g.DB().Table("guaniu_study_topic topics").Where("topics.status = ? and topics.order_num = ?", status, orderNum).InnerJoin("guaniu_study_users user", "topics.create_owner_id=user.id").Fields("topics.id, topics.topic_type, topics.topic_owner_type, topics.topic_name, topics.topic_tag, topics.topic_img, topics.status, topics.create_owner_id, topics.create_time, topics.modify_time, topics.members_number, topics.see_num,user.name, user.ext").Page(int(page), int(req.PageSize)).Order("topics.id desc").FindAll()
		 	
				if err = all.Structs(&list); err != nil && err != sql.ErrNoRows {
					return
				}
			}else{
				all, _ := g.DB().Table("guaniu_study_topic topics").Where("topics.create_owner_id = ?  and topics.status = ? and topics.order_num = ?", req.CreateOwnerId,  status, orderNum).InnerJoin("guaniu_study_users user", "topics.create_owner_id=user.id").Fields("topics.id, topics.topic_type, topics.topic_owner_type, topics.topic_name, topics.topic_tag, topics.topic_img, topics.status, topics.create_owner_id, topics.create_time, topics.modify_time, topics.members_number, topics.see_num,user.name, user.ext").Page(int(page), int(req.PageSize)).Order("topics.id desc").FindAll()
		 	
				if err = all.Structs(&list); err != nil && err != sql.ErrNoRows {
					return
				}
			}
		}else{

			if (!setUserId){
	    		all, _ := g.DB().Table("guaniu_study_topic topics").Where("topics.status = ?", status).InnerJoin("guaniu_study_users user", "topics.create_owner_id=user.id").Fields("topics.id, topics.topic_type, topics.topic_owner_type, topics.topic_name, topics.topic_tag, topics.topic_img, topics.status, topics.create_owner_id, topics.create_time, topics.modify_time, topics.members_number, topics.see_num,user.name, user.ext").Page(int(page), int(req.PageSize)).Order("topics.id desc").FindAll()
		 	
				if err = all.Structs(&list); err != nil && err != sql.ErrNoRows {
					return
				}
			}else{
				all, _ := g.DB().Table("guaniu_study_topic topics").Where("topics.create_owner_id = ?  and topics.status = ?", req.CreateOwnerId,  status).InnerJoin("guaniu_study_users user", "topics.create_owner_id=user.id").Fields("topics.id, topics.topic_type, topics.topic_owner_type, topics.topic_name, topics.topic_tag, topics.topic_img, topics.status, topics.create_owner_id, topics.create_time, topics.modify_time, topics.members_number, topics.see_num,user.name, user.ext").Page(int(page), int(req.PageSize)).Order("topics.id desc").FindAll()
		 	
				if err = all.Structs(&list); err != nil && err != sql.ErrNoRows {
					return
				}
			}
		}
	}

	if err != nil {
		g.Log().Error(err)
		err = gerror.New("分页查询失败")
		return
	}
  
  	//关联查询 , todo:性能优化
	// userlist := ([]*EntityUser)(nil)
	// user := g.DB().Table("guaniu_study_users") 
	// user.Fields("id, name, ext").Where("id", gdb.ListItemValuesUnique(list, "CreateOwnerId")).Structs(&userlist);
	// userProfileList := ([]*EntityUserProfile)(nil)
	// userProfile := g.DB().Table("guaniu_study_user_profile") 
	// userProfile.Fields("user_id, user_stars, avater_thumbnail_path, current_job_company").Where("user_id", gdb.ListItemValuesUnique(list, "CreateOwnerId")).Structs(&userProfileList);

	// for _, value := range list{ 
	// 	//
	// 	//t,ok := value.(EntityTopic)
	// 	for _, userItem := range userlist{
	// 		if value.CreateOwnerId == userItem.Id {

	// 			for _, profile := range userProfileList{
	// 				if value.CreateOwnerId == profile.UserId {
	// 					userItem.UserProfile = profile
	// 				}
	// 			} 
	// 			value.User = userItem
	// 		}
	// 	} 
 //   	}
	//g.Dump(list)

	return
}

// 分页查询,返回值total总记录数,page当前页
func SelectListByPageAndType(req *SelectPageReq) (total int, page int64, list []*ShowEntity, err error) {
	model := Model

	//增加约束
	//model = model.Fileds("id,topic_type,topic_owner_type,topic_name, topic_content, order_num, topic_img, create_time, modify_time")
 

	if req == nil || req.TopicType == 0 {
		return
	}
	status :=  0
	orderNum :=  0 
	if req.TopicType != 0 {
		model = model.Where("topic_type = ?", req.TopicType) 
	} 
	if req.Status != "" {
		model = model.Where("status = ?", req.Status)
		status =  gconv.Int(req.Status)
	}else{
		model = model.Where("status = 1")
		status =  1
	}

	if (req.CreateOwnerId != 0){
		//
		model = model.Where("create_owner_id = ?", req.CreateOwnerId) 
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

	//设置缓存

	
	// 分页排序查询
	//list, err = model.Page(int(page), int(req.PageSize)).Order("id DESC").All()  	//关联查询 , todo:性能优化 
	if (status == 0){ 
		all, _ := g.DB().Table("guaniu_study_topic topics").Where("(topics.status = 1 or topics.status = 3) and topics.order_num = ? and topics.topic_type = ?", orderNum, req.TopicType).InnerJoin("guaniu_study_users user", "topics.create_owner_id=user.id").Fields("topics.id, topics.topic_type, topics.topic_owner_type, topics.topic_name, topics.topic_tag, topics.topic_img, topics.status, topics.create_owner_id, topics.create_time, topics.modify_time, topics.members_number, topics.see_num, user.name, user.ext").Page(int(page), int(req.PageSize)).Order("topics.id desc").FindAll()
	
		if err = all.Structs(&list); err != nil && err != sql.ErrNoRows {
			return
		} 
	}else{ 
		all, _ := g.DB().Table("guaniu_study_topic topics").Where("topics.status = ? and topics.order_num = ? and topics.topic_type = ?", status, orderNum, req.TopicType).InnerJoin("guaniu_study_users user", "topics.create_owner_id=user.id").Fields("topics.id, topics.topic_type, topics.topic_owner_type, topics.topic_name, topics.topic_tag, topics.topic_img, topics.status, topics.create_owner_id, topics.create_time, topics.modify_time, topics.members_number, topics.see_num,user.name, user.ext").Page(int(page), int(req.PageSize)).Order("topics.id desc").FindAll()
 	
		if err = all.Structs(&list); err != nil && err != sql.ErrNoRows {
			return
		} 
	}

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
	key := "topiclist_page_all"
	if req != nil {
		if req.TopicName != "" {
			key = key+"_"+ req.TopicName
			model.Where("topic_name like ?", "%"+req.TopicName+"%")
		}
		if req.TopicPinyin != "" {
			key = key+"_"+ req.TopicPinyin
			model.Where("topic_pinyin = ?", req.TopicPinyin)
		}
		if req.TopicImg != "" {
			model.Where("topic_img = ?", req.TopicImg)
		} 
		if req.Status != "" {
			key = key+"_"+ req.Status
			model.Where("status = ?", req.Status)
		}
	}
    b, _ := gcache.Contains(key)
    if (b){
    	cache, _ := gcache.Get(key)
    	//return 
		g.Log().Debug(cache)
    } 


	// 查询
	list, err = model.Order("id asc").All()
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("查询失败")
		return
	}
	gcache.SetIfNotExist(key, gconv.MapDeep(list), 1000*time.Millisecond)
	return
}
