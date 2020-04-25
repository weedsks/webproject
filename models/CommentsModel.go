package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type CommentsModel struct {
	Id          int
	Type        int
	GoodsId     int
	Pid         int
	Name        string
	Email       string
	Description string
	Order       int
	IsVerify    string
	RoleId      int
	Created_at  string
	Updated_at  string
}

func (m *CommentsModel) TableName() string {
	return TableNmae("comments")
}

func NewComment() *CommentsModel {
	return &CommentsModel{}
}

/**
分页查询
*/
func (m *CommentsModel) GetListPage(where map[string]string, order string, pageIndex int, pageNum int) (totalCount int, data map[int]ListUserPageResult, err error) {
	var comments []*CommentsModel
	o := orm.NewOrm()
	offset := (pageIndex - 1) * pageNum

	query := o.QueryTable("ks_comments")

	for key, value := range where {
		switch key {
		case "GoodsId":
			goods_id, err := strconv.Atoi(value)
			if err == nil {
				query = query.Filter("GoodsId", goods_id)
			}
			break
		default:
			query = query.Filter(key, value)

		}
	}

	if len(order) > 0 {
		query = query.OrderBy(order)
	}

	_, err = query.Offset(offset).Limit(pageNum).All(&comments)

	if err != nil {
		beego.Error("获取评论列表失败->", err)
		return
	}

	count, err := query.Count()
	if err != nil {
		beego.Error("获取用户列表数量失败->", err)
		return
	}
	totalCount = int(count)

	return
}
