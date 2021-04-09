// ==========================================================================
// 生成日期：2021-01-23 14:24:09
// 生成人：gfast
// ==========================================================================
package members_service

import (
	membersModel "gfast/app/model/topic/topicmembers"
	commonService "gfast/app/service/topic/common_service"
	"github.com/gogf/gf/os/gtime"
)

// 添加
func AddSave(req *membersModel.AddReq) error {
	req.CreateTime = gtime.Now()
	//记录历史
	commonService.AddToOperatorDb(req.TopicId, req.UserId, commonService.MODULE_TOPIC_MEMBERS, commonService.OPERATOR_ADD)

	return membersModel.AddSave(req)
}

// 删除
func DeleteByIds(Ids []int) error {
	return membersModel.DeleteByIds(Ids)
}

// 删除
func DeleteByTopicId(removeReq *membersModel.RemoveReqByTopic) error {
	//记录历史
	commonService.AddToOperatorDb(removeReq.TopicId, removeReq.UserId, commonService.MODULE_TOPIC_MEMBERS, commonService.OPERATOR_DELETE)

	return membersModel.DeleteByTopicId(removeReq)
}

//修改
func EditSave(editReq *membersModel.EditReq) error {
	return membersModel.EditSave(editReq)
}

// 根据ID查询
func GetByID(id int64) (*membersModel.Entity, error) {
	return membersModel.GetByID(id)
}

// 根据topic id 用户id 判断是否有加入
func GetByTopicId(req *membersModel.GetReqByTopic) (*membersModel.Entity, error) {
	return membersModel.GetByTopicId(req)
}
 
//根据topic id 用户id 判断是否有加入
func CheckUserInMembers(topic_id int64, user_id int64) (*membersModel.Entity, error) {
	return membersModel.CheckUserInMembers(topic_id, user_id)
}

// 分页查询
func SelectListByPage(req *membersModel.SelectPageReq) (total int, page int64, list []*membersModel.EntityUser, err error) {
	return membersModel.SelectListByPage(req)
}
