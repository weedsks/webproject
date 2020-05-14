package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
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
	IsDel int
}

func (m *GoodsModel) TableName() string {
	return TableNmae("goods")
}

func NewGoods() *GoodsModel {
	return &GoodsModel{}
}

type GetGoodsListPageResult struct {
	GoodsModel
	DownUrl string
	MovieUrl string
	TypeDesc string
	TypeCommon map[int]TypeCommonModel
}


func (m *GoodsModel) GetListPage(where map[string]string, order string, pageIndex int, pageNum int) (totalCount int, list []*GetGoodsListPageResult, err error) {
	offset := (pageIndex - 1) * pageNum
	o := orm.NewOrm()
	sql_list :="select g.*,m.download_url as movie_url from ks_goods as g left join ks_movie as m on m.goods_id=g.id where g.is_del =0"
	sql_count :="select count(*) from ks_goods as g where g.is_del =0"
	for key, value := range where {
		switch key {
		case "name":
			sql_list += " and g.name like '%"+ value +"%'"
			sql_count += " and g.name like '%"+ value +"%'"
			break
		case "is_verify":
			sql_list += " and g.is_verify ="+ value
			sql_count += " and g.is_verify ="+ value
			break
		case "type":
			sql_list += " and g.type ="+ value
			sql_count += " and g.type ="+ value
			break
		case "tag":
			common := []*GoodsTypeCommonRelationModel{}
			o.QueryTable("ks_goods_type_relation").Filter("TypeCommonId", value).All(&common)
			goodsId := "";
			for _,value := range common {
				goodsId += (strconv.Itoa(value.GoodsId)+",")
			}
			if(len(goodsId)>1){
				goodsId = strings.TrimRight(goodsId, ",")
			}
			sql_list += " and g.id in ("+ goodsId +" )"
			sql_count += " and g.id in ("+ goodsId + " )"
			break
		default:

		}
	}
	sql_list +=" order by "+ order +" limit ?,?"
	o.Raw(sql_list,offset,pageNum).QueryRows(&list)
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



func (m *GoodsModel) DelAll(ids []string) (num int64, err error){
	o:=orm.NewOrm()
	num, err = o.QueryTable("ks_goods").Filter("id__in", ids).Update(orm.Params{
		"isDel": 1,
	})
	return
}