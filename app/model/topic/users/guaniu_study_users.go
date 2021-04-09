// ==========================================================================
// 生成日期：2021-01-21 15:14:17
// 生成人：gfast
// ==========================================================================
package users

import (
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"database/sql"
)

// AddReq 用于存储新增请求的请求参数
type AddReq struct {
	Name   string      `p:"name" v:"required#用户名不能为空"`
	Email  string      `p:"email" `
	Mobile string      `p:"mobile" `
	Passwd string      `p:"passwd" v:"required#密码不能为空"`
	RePasswd string    `p:"repasswd" v:"required#重复密码不能为空"`   
	Salt   string      `p:"salt" `
	Ext    string      `p:"ext" `
	Status int         `p:"status" v:"required#状态（0：未审核,1:通过 10删除）不能为空"`
	Ctime  int         `p:"ctime" `
	Mtime  *gtime.Time `p:"mtime" `
}

// EditReq 用于存储修改请求参数
type EditReq struct {
	Id     int64       `p:"id" v:"required#主键ID不能为空"` 
	Email  string      `p:"email" `
	Mobile string      `p:"mobile" `
	Passwd string      `p:"passwd" `
	RePasswd string    `p:"repasswd"`   
	NowPasswd string    `p:"nowPasswd"`   
	Salt   string      `p:"salt" `
	Ext    string      `p:"ext" `
	Status int         `p:"status" v:"required#状态（0：未审核,1:通过 10删除）不能为空"`
	Ctime  int         `p:"ctime" `
	Mtime  *gtime.Time `p:"mtime" ` 
}

type EditProfileReq struct {
	UserId     int64       `p:"userId" v:"required#主键ID不能为空"` 
	UserStars  int64 `p:"userStars" `   
	AvaterPath  string `p:"avaterPath" `   
	Email  string `p:"email" `   
	Cityname  string `p:"cityname" `   
	Nickname  string `p:"nickname" `   
	AvaterThumbnailPath  string `p:"avaterThumbnailPath" `   
	CurrentJobCompany  string `p:"currentJobCompany" `   
	HomeBgPath  string `p:"homeBgPath" `   
	SignNotes  string `p:"signNotes" `   
	WeiboId  string `p:"weiboId" `   
	WeixinId  string `p:"weixinId" `   
	GithubId  string `p:"githubId" `   
	QqId  string `p:"qqId" `   
	DoubanId  string `p:"doubanId" `   
	CsdnId  string `p:"csdnId" `   
	FacebookId  string `p:"facebookId" `   
	BlogSite  string `p:"blogSite" `   
	Ext1  string `p:"ext1" `   
	Ext2  string `p:"ext2" `   
	Ext3  string `p:"ext3" `   
	Ext4  string `p:"ext4" `   
	Ext5  string `p:"ext5" `   
	Mtime  *gtime.Time `p:"mtime" `  
}
/*

CREATE TABLE `guaniu_study_user_profile` (
  `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键', 
   `user_id`         bigint(20)      default 0                  comment '用户id',
    `user_stars`         bigint(20)      default 0                  comment '用户积分',
  `avater_path` varchar(1024) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '头像',
  `avater_thumbnail_path` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '头像', 
  `current_job_company` varchar(1024) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '当前公司名字',
  `home_bg_path` varchar(1024) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '背景图',
  `sign_notes` varchar(1024) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '签名',
  `weibo_id` varchar(256) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'weibo账号',
  `weixin_id` varchar(256) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'weixin账号',
  `github_id` varchar(256) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'github账号',
  `qq_id` varchar(256) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'qq账号',
  `douban_id` varchar(256) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'douban账号',
  `csdn_id` varchar(256) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'csdn账号',
  `facebook_id` varchar(256) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'facebook_id账号',
  `blog_site` varchar(1024) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '网站',
  `ext_1` varchar(1024) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'ext_1',
  `ext_2` varchar(1024) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'ext_1',
  `ext_3` varchar(1024) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'ext_1',
  `ext_4` varchar(1024) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'ext_1',
  `ext_5` varchar(1024) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'ext_1',
  `mtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间'
)

*/
type UserProfileEditEntity struct {
	Id     int64      `orm:"id,primary" json:"id"`     // 主键
    UserId     int64      `orm:"user_id"       json:"userId"`    
	UserStars  int64`orm:"user_stars"      json:"userStars"`  
	AvaterPath  string `orm:"avater_path"      json:"avaterPath"`  
	AvaterThumbnailPath  string `orm:"avater_thumbnail_path"      json:"avaterThumbnailPath"`  
	Cityname  string `orm:"cityname"      json:"cityname"`  
	Nickname  string `orm:"nickname"      json:"nickname"`  
	Email  string `orm:"email"      json:"email"`  
	CurrentJobCompany  string `orm:"current_job_company"      json:"currentJobCompany"`  
	HomeBgPath  string `orm:"home_bg_path"      json:"homeBgPath"`   
	SignNotes  string `orm:"sign_notes"      json:"signNotes"`  
	WeiboId  string `orm:"weibo_id"      json:"weiboId"`  
	WeixinId  string `orm:"weixin_id"      json:"weixinId"`  
	GithubId  string `orm:"github_id"      json:"githubId"`  
	QqId  string `orm:"qq_id"      json:"qqId"`  
	DoubanId  string `orm:"douban_id"      json:"doubanId"`  
	CsdnId  string `orm:"csdn_id"      json:"csdnId"`  
	FacebookId  string `orm:"facebook_id"      json:"facebookId"`  
	BlogSite  string `orm:"blog_site"      json:"blogSite"`   
	Ext1  string `orm:"ext_1"      json:"ext1"`   
	Ext2  string `orm:"ext_2"      json:"ext2"`   
	Ext3  string `orm:"ext_3"      json:"ext3"`  
	Ext4  string  `orm:"ext_4"      json:"ext4"`  
	Ext5  string  `orm:"ext_5"      json:"ext5"`  
	Mtime  *gtime.Time `orm:"mtime"      json:"mtime"`  // 修改时间
}
type UserProfileAddEntity struct { 
    UserId     int64      `orm:"user_id"       json:"userId"`    
	UserStars  int64`orm:"user_stars"      json:"userStars"`  
	AvaterPath  string `orm:"avater_path"      json:"avaterPath"`  
	Cityname  string `orm:"cityname"      json:"cityname"`  
	Nickname  string `orm:"nickname"      json:"nickname"`  
	Email  string `orm:"email"      json:"email"`  
	AvaterThumbnailPath  string `orm:"avater_thumbnail_path"      json:"avaterThumbnailPath"`  
	CurrentJobCompany  string `orm:"current_job_company"      json:"currentJobCompany"`  
	HomeBgPath  string `orm:"home_bg_path"      json:"homeBgPath"`   
	SignNotes  string `orm:"sign_notes"      json:"signNotes"`  
	WeiboId  string `orm:"weibo_id"      json:"weiboId"`  
	WeixinId  string `orm:"weixin_id"      json:"weixinId"`  
	GithubId  string `orm:"github_id"      json:"githubId"`  
	QqId  string `orm:"qq_id"      json:"qqId"`  
	DoubanId  string `orm:"douban_id"      json:"doubanId"`  
	CsdnId  string `orm:"csdn_id"      json:"csdnId"`  
	FacebookId  string `orm:"facebook_id"      json:"facebookId"`  
	BlogSite  string `orm:"blog_site"      json:"blogSite"`   
	Ext1  string `orm:"ext_1"      json:"ext1"`   
	Ext2  string `orm:"ext_2"      json:"ext2"`   
	Ext3  string `orm:"ext_3"      json:"ext3"`  
	Ext4  string  `orm:"ext_4"      json:"ext4"`  
	Ext5  string  `orm:"ext_5"      json:"ext5"`  
	Mtime  *gtime.Time `orm:"mtime"      json:"mtime"`  // 修改时间
}


type RemoveReq struct {
	Ids []int `p:"ids"` //删除id
}

// SelectPageReq 用于存储分页查询的请求参数
type SelectPageReq struct {
	Name      string      `p:"name"`      //用户名
	Email     string      `p:"email"`     //邮箱
	Mobile    string      `p:"mobile"`    //手机号
	Passwd    string      `p:"passwd"`    //密码
	Salt      string      `p:"salt"`      //盐值
	Ext       string      `p:"ext"`       //扩展字段
	Status    int         `p:"status"`    //状态（0：未审核,1:通过 10删除）
	Ctime     int         `p:"ctime"`     //创建时间
	Mtime     *gtime.Time `p:"mtime"`     //修改时间
	BeginTime string      `p:"beginTime"` //开始时间
	EndTime   string      `p:"endTime"`   //结束时间
	PageNum   int64       `p:"pageNum"`   //当前页码
	PageSize  int         `p:"pageSize"`  //每页数
}

// GetPlugAdByID 根据ID查询记录
func GetByID(id int64) (*Entity, error) {
	entity, err := Model.Fields("name, email, ext").FindOne(id)
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
	entity.Name = req.Name
	entity.Email = req.Email
	entity.Mobile = req.Mobile
	entity.Passwd = req.Passwd
	entity.Salt = req.Salt
	entity.Ext = req.Ext
	entity.Status = req.Status
	entity.Ctime = req.Ctime
	entity.Mtime = req.Mtime
	result, err := Model.Save(entity)
	if err != nil {
		return err
	}
	var id int64
	id, err = result.LastInsertId()
	if err != nil {
		return err
	}

	//同时给profile里面生成一条数据

	//新建一个，保存
	userProfileEditEntity := new (UserProfileAddEntity)
	userProfileEditEntity.UserId = id 
	userProfileEditEntity.Nickname = req.Name 

	model := g.DB().Table("guaniu_study_user_profile") 
	_, err = model.Save(userProfileEditEntity) 
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
func EditPasswordSave(req *EditReq) error {
	// 先根据ID来查询要修改的记录
	entity, err := Model.FindOne(req.Id) 
	if err != nil {
		return err
	}
	//校验老密码是否一致
	if req.NowPasswd != entity.Passwd{
		return gerror.New("旧密码输入错误，修改失败") 
	}

	// 修改实体
	entity.Passwd = req.Passwd

	_, err = Model.Save(entity)
	if err != nil {
		g.Log().Error(err)
		return gerror.New("修改失败")
	}
	return nil
}

//修改头像
func EditHeaderImageSave(editReq *UserProfileEditEntity) error { 
	// 先根据ID来查询要修改的记录
	userEntity, err := Model.FindOne(editReq.UserId) 
	if err != nil {
		return err
	}
	userEntity.Ext = editReq.AvaterPath

	userEntity.Mtime = gtime.Now()
	//保存到用户表中
	_, err = Model.Save(userEntity)
	if err != nil {
		g.Log().Error(err)
		return gerror.New("修改失败")
	}

	// 先根据ID来查询要修改的记录  
	model := g.DB().Table("guaniu_study_user_profile") 
	one, err := model.FindOne("user_id=?", editReq.UserId)  
	if err != nil  || one == nil{
		//新建一个，保存
		userProfileEditEntity := new (UserProfileAddEntity)
		userProfileEditEntity.UserId = editReq.UserId
		userProfileEditEntity.AvaterPath = editReq.AvaterPath
		userProfileEditEntity.AvaterThumbnailPath = editReq.AvaterThumbnailPath
 
		_, err = model.Save(userProfileEditEntity)
		if err != nil {
			return err
		}
		
		return nil
	}
  	var entity *UserProfileEditEntity
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return err
	}
	entity.AvaterPath = editReq.AvaterPath
	entity.AvaterThumbnailPath = editReq.AvaterThumbnailPath
	_, err = model.Save(entity)
	if err != nil {
		return err
	}  
	return nil
} 

func GetUserProfileByUserId(userId int64)(entity *UserProfileEditEntity,err error) {
	model := g.DB().Table("guaniu_study_user_profile") 
	one, err := model.FindOne("user_id=?", userId)  
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("根据ID查询记录出错")
	} 
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return entity, err
	}
	if entity == nil {
		err = gerror.New("根据ID未能查询到记录")
	}
	return entity, nil
}

//修改资料
func EditUserProfileSave(editReq *UserProfileEditEntity) error {
	if len(editReq.Email) > 0{	
		// 先根据ID来查询要修改的记录
		userEntity, err := Model.FindOne(editReq.UserId) 
		if err != nil {
			return err
		}
		userEntity.Email = editReq.Email 

		userEntity.Mtime = gtime.Now()
		//保存到用户表中
		_, err = Model.Save(userEntity)
		if err != nil {
			g.Log().Error(err)
			return gerror.New("修改失败")
		} 
	}

	// 先根据ID来查询要修改的记录  
	model := g.DB().Table("guaniu_study_user_profile") 
	one, err := model.FindOne("user_id=?", editReq.UserId)  
	if err != nil || one == nil {
		//新建一个，保存
		userProfileEditEntity := new (UserProfileAddEntity)
		userProfileEditEntity.UserId = editReq.UserId
		userProfileEditEntity.UserStars = editReq.UserStars
		userProfileEditEntity.Email = editReq.Email
		userProfileEditEntity.Cityname = editReq.Cityname
		userProfileEditEntity.Nickname = editReq.Nickname
		userProfileEditEntity.HomeBgPath = editReq.HomeBgPath
		userProfileEditEntity.CurrentJobCompany = editReq.CurrentJobCompany
		userProfileEditEntity.SignNotes = editReq.SignNotes
		userProfileEditEntity.WeiboId = editReq.WeiboId
		userProfileEditEntity.WeixinId = editReq.WeixinId
		userProfileEditEntity.GithubId = editReq.GithubId
		userProfileEditEntity.QqId = editReq.QqId
		userProfileEditEntity.DoubanId = editReq.DoubanId
		userProfileEditEntity.CsdnId = editReq.CsdnId
		userProfileEditEntity.FacebookId = editReq.FacebookId
		userProfileEditEntity.BlogSite = editReq.BlogSite 
		userProfileEditEntity.Ext1 = editReq.Ext1 
		userProfileEditEntity.Ext2 = editReq.Ext2 
		userProfileEditEntity.Ext3 = editReq.Ext3 
		userProfileEditEntity.Ext4 = editReq.Ext4 
		userProfileEditEntity.Mtime = editReq.Mtime 
 
		model = g.DB().Table("guaniu_study_user_profile") 
		_, err = model.Save(userProfileEditEntity) 
		if err != nil {
			return err
		}
		
		return nil
	}
  	var entity *UserProfileEditEntity
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return err
	}
	if len(editReq.AvaterPath) > 0{
		entity.AvaterPath = editReq.AvaterPath
	}
	if len(editReq.AvaterThumbnailPath) > 0{
		entity.AvaterThumbnailPath = editReq.AvaterThumbnailPath
	}

	entity.UserStars = editReq.UserStars 
	if len(editReq.HomeBgPath) > 0{
		entity.HomeBgPath = editReq.HomeBgPath
	}
	if len(editReq.Email) > 0{
		entity.Email = editReq.Email
	}
	if len(editReq.Cityname) > 0{
		entity.Cityname = editReq.Cityname
	}
	if len(editReq.Nickname) > 0{
		entity.Nickname = editReq.Nickname
	}
	if len(editReq.CurrentJobCompany) > 0{
		entity.CurrentJobCompany = editReq.CurrentJobCompany
	}
	if len(editReq.SignNotes) > 0{
		entity.SignNotes = editReq.SignNotes 
	}
	if len(editReq.WeiboId) > 0{
		entity.WeiboId = editReq.WeiboId
	}
	if len(editReq.WeixinId) > 0{
		entity.WeixinId = editReq.WeixinId
	}
	if len(editReq.GithubId) > 0{
		entity.GithubId = editReq.GithubId
	} 
	if len(editReq.QqId) > 0{
		entity.QqId = editReq.QqId
	} 
	if len(editReq.DoubanId) > 0{
		entity.DoubanId = editReq.DoubanId
	} 
	if len(editReq.CsdnId) > 0{
		entity.CsdnId = editReq.CsdnId
	} 
	if len(editReq.FacebookId) > 0{
		entity.FacebookId = editReq.FacebookId
	} 
	if len(editReq.BlogSite) > 0{
		entity.BlogSite = editReq.BlogSite 
	} 
	if len(editReq.Ext1) > 0{
		entity.Ext1 = editReq.Ext1 
	} 
	if len(editReq.Ext2) > 0{
		entity.Ext2 = editReq.Ext2 
	} 
	if len(editReq.Ext3) > 0{
		entity.Ext3 = editReq.Ext3 
	} 
	if len(editReq.Ext4) > 0{
		entity.Ext4 = editReq.Ext4 
	} 
	entity.Mtime = editReq.Mtime 

	_, err = model.Save(entity)
	if err != nil {
		return err
	}  
	return nil
}

// 根据ID来修改信息
func EditSave(req *EditReq) error {
	// 先根据ID来查询要修改的记录
	entity, err := Model.FindOne(req.Id) 
	if err != nil {
		return err
	}
	// 修改实体
	//entity.Name = req.Name
	entity.Email = req.Email
	entity.Mobile = req.Mobile
	//entity.Passwd = req.Passwd
	entity.Salt = req.Salt
	entity.Ext = req.Ext
	entity.Status = req.Status
	entity.Ctime = req.Ctime
	entity.Mtime = req.Mtime
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

	//增加约束
	//model = model.Fileds("id,name,email,mobile").where("status != 0")

	if req != nil {
		if req.Name != "" {
			model = model.Where("name like ?", "%"+req.Name+"%")
		}
		if req.Email != "" {
			model = model.Where("email = ?", req.Email)
		}
		if req.Mobile != "" {
			model = model.Where("mobile = ?", req.Mobile)
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
		if req.Name != "" {
			model.Where("name like ?", "%"+req.Name+"%")
		}
		if req.Email != "" {
			model.Where("email = ?", req.Email)
		}
		if req.Mobile != "" {
			model.Where("mobile = ?", req.Mobile)
		}
		if req.Passwd != "" {
			model.Where("passwd = ?", req.Passwd)
		}
		if req.Salt != "" {
			model.Where("salt = ?", req.Salt)
		}
		if req.Ext != "" {
			model.Where("ext = ?", req.Ext)
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
