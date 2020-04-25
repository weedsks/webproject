package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type RolesModel struct {
	Id         int
	Name       string
	Fun_ids    string
	Is_del     int8
	Updated_at string
	Created_at string

}

func (m *RolesModel) TableName() string {
	return TableNmae("roles")
}

func NewRole() *RolesModel  {
	return &RolesModel{}
}

func (m *RolesModel) GetAll() (data map[int]RolesModel) {
	role := []*RolesModel{}
	query := orm.NewOrm().QueryTable("ks_roles")
	query = query.Filter("is_del",0)
	_,err :=query.All(&role)
	if err!=nil {
		beego.Error("获取角色列表失败->", err)
		return
	}
	data = make(map[int]RolesModel)
	for _,value:=range role {
		data[value.Id] = *value
	}
	return
}
