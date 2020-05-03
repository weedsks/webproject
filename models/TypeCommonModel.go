package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type TypeCommonModel struct {
	Id          int
	Type        int
	Name        string
	Order       int
	IsVerify    int
	Icon        string
	CreatedAt   string
	UpdatedAt   string
}

func (m *TypeCommonModel) TableName() string {
	return TableNmae("type_common")
}

func NewTypeCommon() *TypeCommonModel {
	return &TypeCommonModel{}
}

type typeCommonList struct {
	Name string
	Icon string
	IsChecked int
	Id int
}

func (m *TypeCommonModel) GetAll(typeId int) (lists map[int]*typeCommonList, err error) {
	var data []*TypeCommonModel
	o := orm.NewOrm().QueryTable("ks_type_common").Filter("is_verify", 1).Filter("type", typeId)
	o.All(&data)
	if err == orm.ErrNoRows {
		beego.Warning("数据查找不到", err)
	}

	lists = make(map[int]*typeCommonList)
	for _, item := range data {
		lists[item.Id] = &typeCommonList{
			Id:item.Id,
			Name:      item.Name,
			Icon:      item.Icon,
			IsChecked: 0,
		}
	}

	return
}

type ListTypeCommonPageResult struct {
	TypeCommonModel
	TypeDesc string
}

/**
分页查询
*/
func (m *TypeCommonModel) GetListPage(where map[string]string, order string, pageIndex int, pageNum int) (totalCount int, data map[int]ListTypeCommonPageResult, err error) {
	var typeComments []*TypeCommonModel
	o := orm.NewOrm()
	offset := (pageIndex - 1) * pageNum

	query := o.QueryTable("ks_type_common")

	for key, value := range where {
		switch key {
		case "IsVerify":
			is_del,err := strconv.Atoi(value)
			if err==nil {
				query = query.Filter("IsVerify",is_del)
			}
			break
		case "Type":
			typeId,err := strconv.Atoi(value)
			if err==nil {
				query = query.Filter("Type",typeId)
			}
			break
		default:
			query = query.Filter(key, value)

		}
	}

	if len(order) > 0 {
		query = query.OrderBy(order)
	}

	_, err = query.Offset(offset).Limit(pageNum).All(&typeComments)


	if err != nil {
		beego.Error("获取商品类型列表失败->", err)
		return
	}

	type_desc := GoodsTypeDesc();

	data = make(map[int]ListTypeCommonPageResult)
	for _,value:=range typeComments {
		var tmp = new(ListTypeCommonPageResult)
		tmp.TypeCommonModel = *value
		if tmp_desc,ok := type_desc[value.Type];ok{
			tmp.TypeDesc = tmp_desc
		}
		data[value.Id]=*tmp
	}

	count, err := query.Count()
	if err != nil {
		beego.Error("获取商品类型列表数量失败->", err)
		return
	}
	totalCount = int(count)

	return
}

func (m *TypeCommonModel) GetInfoById(id int) (typeComment TypeCommonModel, err error) {
	typeComment = TypeCommonModel{Id: id}
	o := orm.NewOrm()
	err =o.Read(&typeComment)

	if err==orm.ErrNoRows {
		beego.Error("数据查找不到", err)
	}

	return
}
