// ==========================================================================
// 生成日期：2021-01-13 15:10:25
// 生成人：gfast
// ==========================================================================
package topic_service

import (
	topicModel "gfast/app/model/topic/topics"

	"github.com/gogf/gf/os/gtime"
)

// 添加
func AddSave(req *topicModel.AddReq) (id int64, err error) {
	req.ModifyTime = gtime.Now()
	req.CreateTime = gtime.Now()
	//req.Status = 1;
	req.MembersNumber = 1;//default value. 

	return topicModel.AddSave(req)
}

// 删除
func DeleteByIds(Ids []int) error {
	
	return topicModel.DeleteByIds(Ids)
}

//修改
func EditSave(editReq *topicModel.EditReq) error {
	editReq.ModifyTime = gtime.Now()
	return topicModel.EditSave(editReq)
}

// 根据ID查询
func GetByID(id int64) (*topicModel.Entity, error) {
	return topicModel.GetByID(id)
}

func AddMember(id int64) error {
	return topicModel.AddMember(id)
}
func RemoveMember(id int64) error {
	return topicModel.RemoveMember(id)
}

func GetByIDForDetail(id int64) (entity *topicModel.ShowEntity,err error) {
	return topicModel.GetByIDForDetail(id)
}

// 分页查询
func SelectListByPage(req *topicModel.SelectPageReq) (total int, page int64, list []*topicModel.ShowEntity, err error) {
	return topicModel.SelectListByPage(req)
}

func SelectListByPageAndType(req *topicModel.SelectPageReq) (total int, page int64, list []*topicModel.ShowEntity, err error) {
	return topicModel.SelectListByPageAndType(req)
}
 
// 分页查询
func SelectCommentListByPage(req *topicModel.CommenTopicReq) (total int, page int64, list []*topicModel.CommentShowEntity, err error) {
	return topicModel.SelectCommentListByPage(req)
}
 
// 分页查询
func SelectListByPageAndUserId(req *topicModel.UserTopicReq) (total int, page int64, list []*topicModel.ShowEntity, err error) {
	return topicModel.SelectListByPageAndUserId(req)
}
 
 //add comment
func AddComment(req *topicModel.CommentReq) error {
	req.CreateTime = gtime.Now()
	req.Status = 0;
	return topicModel.TopicComment(req)
}
