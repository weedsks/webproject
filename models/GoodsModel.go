package models

import (
	"github.com/astaxie/beego/orm"
	"strings"
	"time"
)

type GoodsModel struct {
	Id int
	Type int
	Name string
	Description string
	Image string
	SmallImage string
	Tag int
	Price float32
	CoverTag int
	ViewNum int
	FavoriteNum int
	PraiseNum int
	IsRecommend int
	Order int
	IsVerify int
	DeletedAt time.Time
	RecommendStars int
	IsComment int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *GoodsModel) TableName() string {
	return TableNmae("goods")
}

func NewGoods() *GoodsModel {
	return &GoodsModel{}
}

type GetGoodsListPageResult struct {
	GoodsModel
	TypeDesc string
	TypeCommon map[int]TypeCommonModel
}


func (m *GoodsModel) GetListPage(where map[string]string, order string, pageIndex int, pageNum int) (totalCount int, list []*GetGoodsListPageResult, err error) {
	offset := (pageIndex - 1) * pageNum
	o := orm.NewOrm()
	sql_list :="select g.* from ks_goods as g where g.is_del =0"
	sql_count :="select count(*) from ks_goods as g where g.is_del =0"
	for key, value := range where {
		switch key {
		case "name":
			sql_list += " and g.name like %"+ value +"%"
			sql_count += " and g.name like %"+ value +"%"
			break
		case "is_verify":
			sql_list += " and g.is_verify ="+ value
			sql_count += " and g.is_verify ="+ value
			break
		default:

		}
	}
	sql_list +=" order by ? limit ?,?"
	o.Raw(sql_list, order,offset,pageNum).QueryRows(&list)
	o.Raw(sql_count).QueryRow(&totalCount)

	GoodsTypeDesc := GoodsTypeDesc()

	for key, value :=range list {
		if tmp_desc,ok :=GoodsTypeDesc[value.Type];ok{
			list[key].TypeDesc = tmp_desc
		}
		if !strings.Contains(value.Image,"http") {
			list[key].Image = getPrivateQiNiuUrl(value.Image)
			list[key].SmallImage = getPrivateQiNiuUrl(value.SmallImage)
		}

	}

	return
}



