// ==========================================================================
// 生成日期：2021-01-23 14:24:09
// 生成人：gfast
// ==========================================================================
package common_service

import (
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
)

const OPERATOR_ADD int = 0
const OPERATOR_EDIT int = 1
const OPERATOR_DELETE int = 2 
const OPERATOR_SEARCH int = 3

const MODULE_TOPIC string = "topic"
const MODULE_TOPIC_MEMBERS string = "memebers"
const MODULE_TOPIC_NOTES string = "notes"
const MODULE_TOPIC_QUESTIONS string = "questions"
 
type TopicMemeberOperatorEntity struct {
    Id              int64       `orm:"id,primary"        json:"id"`                // 编号                          
    TopicId         int64       `orm:"topic_id"          json:"topic_id"`          // 主题id                        
    UserId          int64       `orm:"user_id"           json:"user_id"`           // 成员id                        
	ModuleName       string      `orm:"module"        json:"module"`        // 笔记标题
    OperCode      int         `orm:"operatorion"        json:"operatorion"`        // 操作(0, join, 1:quit, 2:star, 3:）  
    CreateTime      *gtime.Time `orm:"create_time"       json:"create_time"`       // 时间                          
    TopicTableIndex int         `orm:"topic_table_index" json:"topic_table_index"` // 表id 与表topic 相关           
}
 

func AddToOperatorDb(topicId int64, userId int64, module string, operCode int) error {

	entity := new(TopicMemeberOperatorEntity)

	entity.TopicId = topicId
	entity.UserId = userId
	entity.ModuleName = module
	entity.OperCode = operCode
	entity.CreateTime = gtime.Now()
	
	operHistoryModel := g.DB().Table("guaniu_study_topic_members_history") 
	_, err := operHistoryModel.Save(entity)
	if err != nil {
		g.Log().Error(err)
		return gerror.New("添加失败")
	}
	return nil
}