// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package users

import (
	"database/sql"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/os/gtime"
)

// Entity is the golang structure for table guaniu_study_users.
type Entity struct {
	Id     uint64      `orm:"id,primary" json:"id"`     // 主键
	Name   string      `orm:"name"       json:"name"`   // 用户名
	Email  string      `orm:"email"      json:"email"`  // 邮箱
	Mobile string      `orm:"mobile"     json:"mobile"` // 手机号
	Passwd string      `orm:"passwd"     json:"passwd"` // 密码
	Salt   string      `orm:"salt"       json:"salt"`   // 盐值
	Ext    string      `orm:"ext"        json:"ext"`    // 扩展字段
	Status int         `orm:"status"     json:"status"` // 状态（0：未审核,1:通过 10删除）
	Ctime  int         `orm:"ctime"      json:"ctime"`  // 创建时间
	Mtime  *gtime.Time `orm:"mtime"      json:"mtime"`  // 修改时间
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
