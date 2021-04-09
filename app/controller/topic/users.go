// ==========================================================================
// 生成日期：2021-01-21 15:14:17
// 生成人：gfast
// ==========================================================================
package topic

import (
	usersModel "gfast/app/model/topic/users"
	usersService "gfast/app/service/topic/users_service"
	"gfast/library/response"
	libraryService "gfast/library/service"

	"github.com/gogf/gf/frame/g"
    "github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"

	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
	"fmt"
	"math/rand"
	"time"
)
 
var (
	accessKey = "12222"//修改为七牛对应的accessKey
	secretKey = "12-1"//修改为七牛对应的secretKey
	bucket    =  "21"
)

//控制器
type Users struct{}

//列表页
func (c *Users) List(r *ghttp.Request) {
	// 定义一个结构体存储请求参数
	var req *usersModel.SelectPageReq
	// 获取参数
	err := r.Parse(&req)
	if err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}
	total, page, list, err := usersService.SelectListByPage(req)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	result := g.Map{
		"currentPage": page,
		"total":       total,
		"list":        list,
	}
	response.SusJson(true, r, "获取列表数据成功", result)
}

// 新增
/*
liyizhang@bogon gfast % curl -d '{"name":"john","email":"123@qq.com", "passwd":"111111"}' "http://127.0.0.1:8200/api.v2/user/add"
{"code":0,"data":null,"msg":"添加成功"}%


curl -d '{"username":"john","email":"123@qq.com", "passwd":"111111", "idValueC":"3244", "idKeyC":"3244"}' "http://127.0.0.1:8200/pub6Login/login"

curl  http://127.0.0.1:8200/api.v2/user/list -H 'Authorization:Bearer vGCiSmcWADHgtiq2XeDTokcTCipNawAG8xZW1F+49fgzaNhaoop+XX1XHvm2qx715BB5VGUxeFlXqTL7WcE+c59q3k7TRk4Ai1P0E01dYeJlsyCA3Ss9vU066gLs4ipB'
{"code":0,"data":{"currentPage":1,"list":[{"id":1,"name":"john","email":"123@qq.com","mobile":"","passwd":"yxhe1kif+kEbTPrEJgdYxQ==","salt":"0","ext":"","status":1,"ctime":1611294881,"mtime":"2021-01-22 14:04:35"}],"total":1},"msg":"获取列表数据成功"}%

liyizhang@bogon gfast % curl "http://127.0.0.1:8200/pub6Login/logout" -H 'Authorization:Bearer vGCiSmcWADHgtiq2XeDTokcTCipNawAG8xZW1F+49fgzaNhaoop+XX1XHvm2qx71wpU/AgCNGbdKNa2e7mn51+bdHmTmtFaMyX5YTm+Y4aLiF0CIES5RibQ4A6cJtmeC'
{"code":0,"msg":"success","data":"Logout success"}%                                                                                                                                      liyizhang@bogon gfast %

*/
func (c *Users) Add(r *ghttp.Request) {
	if r.Method == "POST" {
		var req *usersModel.AddReq
		// 通过Parse方法解析获取参数
		err := r.Parse(&req)
		if err != nil {
			response.FailJson(true, r, err.(*gvalid.Error).FirstString())
		}

		data := r.GetFormMapStrStr()
		//判断验证码是否正确
		if !libraryService.VerifyString(data["idKeyC"], data["idValueC"]) {
			response.JsonExit(r, response.ErrorCode, "验证码输入错误")
		} 

		// 调用service中的添加函数添加
		err = usersService.AddSave(req)
		if err != nil {
			response.FailJson(true, r, err.Error())
		}
		response.SusJson(true, r, "添加成功")
	}
}

// 修改资料
func (c* Users)EditProfile(r *ghttp.Request) {

	// 如果是post提交的请求就执行修改操作
	if r.Method == "POST" {
		var editReq *usersModel.UserProfileEditEntity
		// 通过Parse方法解析获取参数
		err := r.Parse(&editReq)
		if err != nil {
			response.FailJson(true, r, err.(*gvalid.Error).FirstString())
		}
		//找到当前用户的ID 
		editReq.UserId = usersService.GetLoginID(r)
		//保存编辑
		err = usersService.EditUserProfileSave(editReq)
		if err != nil {
			response.FailJson(true, r, err.Error())
		}
		response.SusJson(true, r, "修改参数成功")
	}
    response.FailJson(true, r, "请求错误") 
}

// 修改头像
func (c* Users)EditHeaderImage(r *ghttp.Request) {

	// 如果是post提交的请求就执行修改操作
	if r.Method == "POST" {
		var editReq *usersModel.UserProfileEditEntity
		// 通过Parse方法解析获取参数
		err := r.Parse(&editReq)
		if err != nil {
			response.FailJson(true, r, err.(*gvalid.Error).FirstString())
		}
		//找到当前用户的ID 
		editReq.UserId = usersService.GetLoginID(r)
		//保存编辑
		err = usersService.EditHeaderImageSave(editReq)
		if err != nil {
			response.FailJson(true, r, err.Error())
		}
		response.SusJson(true, r, "修改参数成功")
	}
    response.FailJson(true, r, "请求错误") 
}

// 修改密码
func (c* Users)EditPassword(r *ghttp.Request) {

	// 如果是post提交的请求就执行修改操作
	if r.Method == "POST" {
		var editReq *usersModel.EditReq
		// 通过Parse方法解析获取参数
		err := r.Parse(&editReq)
		if err != nil {
			response.FailJson(true, r, err.(*gvalid.Error).FirstString())
		}
		//判断密码输入是否正确
		if len(editReq.Passwd) != len(editReq.RePasswd) || len(editReq.Passwd) < 6 ||editReq.Passwd != editReq.RePasswd {
			  
			response.FailJson(true, r, "两次输入密码不一致，或者密码长度小于6") 
			return
		}
		//找到当前用户的ID 
		editReq.Id = usersService.GetLoginID(r)
		//保存编辑
		err = usersService.EditPasswordSave(editReq)
		if err != nil {
			response.FailJson(true, r, err.Error())
		}
		response.SusJson(true, r, "修改参数成功")
	}
    response.FailJson(true, r, "请求错误") 
}



// func (c *Users) Edit(r *ghttp.Request) {
// 	// 不是post提交的请求报错
// 	if r.Method != "POST" {
		
// 		response.FailJson(true, r, "no such api.")
// 	}
// 	// 如果是post提交的请求就执行修改操作 
// 	var editReq *usersModel.EditReq
// 	// 通过Parse方法解析获取参数
// 	err := r.Parse(&editReq)
// 	if err != nil {
// 		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
// 	}
// 	err = usersService.EditSave(editReq)
// 	if err != nil {
// 		response.FailJson(true, r, err.Error())
// 	}
// 	response.SusJson(true, r, "修改参数成功") 
// }

func (c *Users) GetUserProfileByUserId(r *ghttp.Request) {
	if r.Method != "POST" {
		
		response.FailJson(true, r, "no such api.")
	}
	id := r.GetInt("userid")
		//id := usersService.GetLoginID(r)

	params, err := usersService.GetUserProfileByUserId(int64(id))
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "ok", params) 
}

func (c *Users) GetUploadToken(r *ghttp.Request) {
	if r.Method != "POST" {
		
		response.FailJson(true, r, "no such api.")
	}
	//filename := ""//r.GetString("filename")


	// 简单上传凭证 

	// 设置上传凭证有效期
	//添加随机数种子
	rand.Seed(time.Now().UnixNano())

	keyToOverwrite := gconv.String(int64(rand.Intn(500000)) + time.Now().UnixNano())//+".png"
	putPolicy := storage.PutPolicy{
		//Scope: fmt.Sprintf("%s:%s", bucket, keyToOverwrite),
		Scope:bucket,
	}
	putPolicy.Expires = 3000 //示例30分钟有效期

	mac := auth.New(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	fmt.Println(upToken)
 
	response.SusJson(true, r, "ok", g.Map{
				"token":   upToken,
				"name": keyToOverwrite,
				"server":"https://voice.it3q.com/",
			}) 
}

// 删除
// func (c *Users) Delete(r *ghttp.Request) {
// 	var req *usersModel.RemoveReq
// 	//获取参数
// 	if err := r.Parse(&req); err != nil {
// 		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
// 	}
// 	err := usersService.DeleteByIds(req.Ids)
// 	if err != nil {
// 		response.FailJson(true, r, "删除失败")
// 	}
// 	response.SusJson(true, r, "删除成功")
// }
