package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MenusModel struct {
	Id        int
	ParentId  int
	Order     int
	Title     string
	Icon      string
	Url       string
	IsDel     int
	UpdatedAt string
	CreatedAt string
}

func (m *MenusModel) TableName() string {
	return TableNmae("menus")
}

func NewMenus() *MenusModel {
	return &MenusModel{}
}

type MenuTree struct {
	Id       int
	ParentId int
	Order    int
	Title    string
	Icon     string
	Url      string
	Child    []*MenuTree
}


func (m *MenusModel) GetMenu(pid int) (treeList []*MenuTree) {
	list := []*MenusModel{}
	query := orm.NewOrm().QueryTable("ks_menus")
	query = query.Filter("is_del", 0).Filter("ParentId", pid).OrderBy("parentid", "-order")
	_, err := query.All(&list)
	if err != nil {
		beego.Error("获取角色列表失败->", err)
		return
	}
	for _, v := range list {
		Child := v.GetMenu(v.Id)
		node := &MenuTree{
			Id:       v.Id,
			ParentId: v.ParentId,
			Order:    v.Order,
			Title:    v.Title,
			Icon:     v.Icon,
			Url:      v.Url,
		}
		node.Child = Child
		treeList = append(treeList, node)
	}
	return
}

/**
列表重组
*/
func (m *MenusModel) GetAll() (data []*MenusModel) {
	menuTree := m.GetMenu(0)
	for _, value := range menuTree {
		//父级
		ptmp := &MenusModel{
			Id:       value.Id,
			ParentId: value.ParentId,
			Order:    value.Order,
			Title:    value.Title,
			Icon:     value.Icon,
			Url:      value.Url,
		}
		data = append(data, ptmp)

		//子集 目前只有2级
		if len(value.Child) > 0 {
			for _, v := range value.Child {
				tmp := &MenusModel{
					Id:       v.Id,
					ParentId: v.ParentId,
					Order:    v.Order,
					Title:    v.Title,
					Icon:     v.Icon,
					Url:      v.Url,
				}
				data = append(data, tmp)
			}
		}
	}
	fmt.Println(data)
	return
}


func (m *MenusModel) GetInfoById(id int) (data MenusModel, err error) {
	data = MenusModel{Id: id}
	o := orm.NewOrm()
	err = o.Read(&data)
	if err == orm.ErrNoRows {
		beego.Warning("数据查找不到", err)
	}

	return
}

/**
返回父级
*/
func (m *MenusModel) GetParents() (data map[int]string) {
	list := []*MenusModel{}
	query := orm.NewOrm().QueryTable("ks_menus")
	query.Filter("is_del", 0)
	_, err := query.All(&list)
	if err != nil {
		beego.Error("获取角色列表失败->", err)
		return
	}
	data = make(map[int]string)
	for _, value := range list {
		data[value.Id] = value.Title
	}
	return
}
