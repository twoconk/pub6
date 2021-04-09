// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package bookresource

import (
	"database/sql"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/os/gtime"
)

// Entity is the golang structure for table guaniu_study_book_resource.
type Entity struct {
	Id              int64       `orm:"id,primary"        json:"id"`                // 编号
	BookName        string      `orm:"book_name"         json:"book_name"`         // 书的名称
	BookAuthor      string      `orm:"book_author"       json:"book_author"`       // 书的作者
	BookDesc        string      `orm:"book_desc"         json:"book_desc"`         // 书的描述
	BookTags        string      `orm:"book_tags"         json:"book_tags"`         // 书的标签
	BookCbs         string      `orm:"book_cbs"          json:"book_cbs"`          // 出版社
	PrintTime       *gtime.Time `orm:"print_time"        json:"print_time"`        // 发行时间
	BookIsbn        string      `orm:"book_isbn"         json:"book_isbn"`         // isbn
	Status          int         `orm:"status"            json:"status"`            // 是否开放状态（0正常 1 删除）
	CreateTime      *gtime.Time `orm:"create_time"       json:"create_time"`       // 时间
	TopicTableIndex int         `orm:"topic_table_index" json:"topic_table_index"` // 表id 与表topic 相关
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
// Deprecated.
func (r *Entity) OmitEmpty() *arModel {
	return Model.Data(r).OmitEmpty()
}

// Inserts does "INSERT...INTO..." statement for inserting current object into table.
// Deprecated.
func (r *Entity) Insert() (result sql.Result, err error) {
	return Model.Data(r).Insert()
}

// InsertIgnore does "INSERT IGNORE INTO ..." statement for inserting current object into table.
// Deprecated.
func (r *Entity) InsertIgnore() (result sql.Result, err error) {
	return Model.Data(r).InsertIgnore()
}

// Replace does "REPLACE...INTO..." statement for inserting current object into table.
// If there's already another same record in the table (it checks using primary key or unique index),
// it deletes it and insert this one.
// Deprecated.
func (r *Entity) Replace() (result sql.Result, err error) {
	return Model.Data(r).Replace()
}

// Save does "INSERT...INTO..." statement for inserting/updating current object into table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
// Deprecated.
func (r *Entity) Save() (result sql.Result, err error) {
	return Model.Data(r).Save()
}

// Update does "UPDATE...WHERE..." statement for updating current object from table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
// Deprecated.
func (r *Entity) Update() (result sql.Result, err error) {
	where, args, err := gdb.GetWhereConditionOfStruct(r)
	if err != nil {
		return nil, err
	}
	return Model.Data(r).Where(where, args).Update()
}

// Delete does "DELETE FROM...WHERE..." statement for deleting current object from table.
// Deprecated.
func (r *Entity) Delete() (result sql.Result, err error) {
	where, args, err := gdb.GetWhereConditionOfStruct(r)
	if err != nil {
		return nil, err
	}
	return Model.Where(where, args).Delete()
}
