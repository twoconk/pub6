// ==========================================================================
// 生成日期：2021-01-21 15:14:17
// 生成人：gfast
// ==========================================================================
package users_service

import (
	usersModel "gfast/app/model/topic/users"
	"gfast/boot"
	"gfast/library/utils"

	"github.com/gogf/gf/net/ghttp"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/os/gtime"

	//"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

//获取登陆用户ID
func GetLoginID(r *ghttp.Request) (userId int64) {
	userInfo := GetLoginInfo(r)
	if userInfo != nil {
		userId = int64(userInfo.Id)
	}
	return
}

//获取缓存的用户信息
func GetLoginInfo(r *ghttp.Request) (userInfo *usersModel.Entity) {
	resp := boot.TopicGfToken.GetTokenData(r)
	//g.Dump(resp)
	gconv.Struct(resp.Get("data"), &userInfo)
	return
}

//获取当前登录用户信息，直接从数据库获取
func GetCurrentUser(r *ghttp.Request) (userInfo *usersModel.Entity, err error) {
	id := GetLoginID(r)
	userInfo, err = usersModel.GetUserById(id)
	return
}

// 添加
func AddSave(req *usersModel.AddReq) error {

	if i, _ := usersModel.Model.Where("name=?", req.Name).Count(); i != 0 {
		return gerror.New("用户名已经存在")
	}
	if  len(req.Mobile)!=0 {
		if i, _ := usersModel.Model.Where("mobile=?", req.Mobile).Count(); i != 0 {
			return gerror.New("手机号已经存在")
		}		
	}
	if len(req.Passwd) != len(req.RePasswd) || len(req.Passwd) < 6 ||req.Passwd != req.RePasswd {
		return gerror.New("两次输入密码不一致，或者密码长度小于6") 
	}

	//密码加密
	req.Passwd = utils.EncryptCBC(gconv.String(req.Passwd), utils.AdminCbcPublicKey)

	//保存用户信息
	entity := new(usersModel.Entity)
	entity.Name = req.Name
	entity.Status = 1//默认修改为1
	entity.Ctime = gconv.Int(gtime.Timestamp())
	entity.Mobile = req.Mobile
	entity.Salt = "0"
	entity.Email = req.Email
	entity.Passwd = req.Passwd
	entity.Mtime = gtime.Now()
	res, err := usersModel.Model.Save(entity)
	_, _ = res.LastInsertId()

	return err
}

//修改密码
func EditPasswordSave(editReq *usersModel.EditReq) error {
	if  editReq.Id ==0 { 
	    return gerror.New("用户错误") 
	}
	//密码加密
	editReq.Passwd = utils.EncryptCBC(gconv.String(editReq.Passwd), utils.AdminCbcPublicKey)
	editReq.NowPasswd = utils.EncryptCBC(gconv.String(editReq.NowPasswd), utils.AdminCbcPublicKey)

	return usersModel.EditPasswordSave(editReq)
}

//修改头像
func EditHeaderImageSave(editReq *usersModel.UserProfileEditEntity) error {
	if  editReq.UserId ==0 { 
	    return gerror.New("用户错误") 
	} 

	editReq.Mtime = gtime.Now()
	return usersModel.EditHeaderImageSave(editReq)
}

//修改资料
func EditUserProfileSave(editReq *usersModel.UserProfileEditEntity) error {
	if  editReq.UserId ==0 { 
	    return gerror.New("用户错误") 
	} 

	editReq.Mtime = gtime.Now()
	return usersModel.EditUserProfileSave(editReq)
}

// 删除
func DeleteByIds(Ids []int) error {
	return usersModel.DeleteByIds(Ids)
}

//修改
func EditSave(editReq *usersModel.EditReq) error {

	return usersModel.EditSave(editReq)
}

// 根据ID查询
func GetByID(id int64) (*usersModel.Entity, error) {
	return usersModel.GetByID(id)
}
//用户用户id获取用户详情
func GetUserProfileByUserId(id int64)  ( *usersModel.UserProfileEditEntity, error) {
	return usersModel.GetUserProfileByUserId(id)
}

// 分页查询
func SelectListByPage(req *usersModel.SelectPageReq) (total int, page int64, list []*usersModel.Entity, err error) {
	return usersModel.SelectListByPage(req)
}
