package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type UsersModel struct {
	Id         int
	Name       string `orm:"size(255)"`
	Mobile     string `orm:"size(11)"`
	Password   string
	Email      string
	Created_at string
	Updated_at string
	Salt       string
	RoleId		int
	IsDel    int
}

func (m *UsersModel) TableName() string {
	return TableNmae("users")
}

func NewUser() *UsersModel  {
	return &UsersModel{}
}


func (m *UsersModel) VerificationLogin(keyNmae string) UsersModel {
	orm.Debug = true
	user := UsersModel{Name: keyNmae,IsDel:0}
	o := orm.NewOrm()
	o.Read(&user, "Name","IsDel")
	return user
}



type ListUserPageResult struct {
	UsersModel
	Role_name string
}

/**
分页查询
*/
func (m *UsersModel) GetListPage(where map[string]string, order string, pageIndex int, pageNum int) (totalCount int, data map[int]ListUserPageResult, err error) {
	var users []*UsersModel
	o := orm.NewOrm()
	offset := (pageIndex - 1) * pageNum

	query := o.QueryTable("ks_users")

	for key, value := range where {
		switch key {
		case "IsDel":
			is_del,err := strconv.Atoi(value)
			if err==nil {
				query = query.Filter("IsDel",is_del)
			}
			break
		default:
			query = query.Filter(key, value)

		}
	}

	if len(order) > 0 {
		query = query.OrderBy(order)
	}

	_, err = query.Offset(offset).Limit(pageNum).All(&users)


	if err != nil {
		beego.Error("获取用户列表失败->", err)
		return
	}

	role_arr := NewRole().GetAll();

	data = make(map[int]ListUserPageResult)
	for _,value:=range users {
		var tmp = new(ListUserPageResult)
		tmp.UsersModel = *value
		if tmp_role,ok := role_arr[value.RoleId];ok{
			tmp.Role_name = tmp_role.Name
		}
		data[value.Id]=*tmp
	}

	count, err := query.Count()
	if err != nil {
		beego.Error("获取用户列表数量失败->", err)
		return
	}
	totalCount = int(count)

	return
}

func (m *UsersModel) GetInfoById(id int) (user UsersModel, err error) {
	user = UsersModel{Id: id}
	o := orm.NewOrm()
	err =o.Read(&user)

	if err==orm.ErrNoRows {
		beego.Error("数据查找不到", err)
	}

	return
}
