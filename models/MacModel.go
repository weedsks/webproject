package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strings"
	"time"
)

type MacModel struct {
	Id        int
	GoodsId   int
	DownloadUrl   string
	Content string
	Link   string
	DeletedAt time.Time
}

func (m *MacModel) TableName() string {
	return TableNmae("mac")
}

func NewMac() *MacModel {
	return &MacModel{}
}

type GetMacListPageResult struct {
	GoodsId int
	DownloadUrl string
	Content string
	Id int
	Type int
	Name string
	Description string
	Image string
	ImageUrl string
	SmallImage string
	SmallImageUrl string
	Tag int
	Price float32
	CoverTag int
	ViewNum int
	FavoriteNum int
	PraiseNum int
	IsRecommend int
	Order int
	IsVerify int
	IsDel int
	DeletedAt time.Time
	RecommendStars int
	IsComment int
	CreatedAt time.Time
	UpdatedAt time.Time
	TypeCommon map[int]TypeCommonModel
	TypeDesc string
}


func (m *MacModel) GetInfoById(id int) (info GetMacListPageResult, err error) {
	o := orm.NewOrm()
	sql_info := "select a.goods_id,a.download_url,a.content,g.* from ks_mac as a left join ks_goods as g on g.id=a.goods_id where g.is_del=0 and g.id = ? limit 1"
	o.Raw(sql_info, id).QueryRow(&info)
	if err == orm.ErrNoRows {
		beego.Warning("数据查找不到", err)
		return
	}
	info.TypeCommon,_= NewGoodsTypeCommonRelation().GetRelation(id)
	if !strings.Contains(info.SmallImage,"http") {
		//info.Image = getPrivateQiNiuUrl(info.Image)
		info.SmallImageUrl = getPrivateQiNiuUrl(info.SmallImage)
	}

	GoodsTypeDesc := GoodsTypeDesc()

	if tmp_desc, ok := GoodsTypeDesc[info.Type]; ok {
		info.TypeDesc = tmp_desc
	}
	//阅读量
	info.ViewNum = info.ViewNum+1
	tmp_goods := GoodsModel{Id:id}
	tmp_goods.ViewNum = info.ViewNum+1
	o.Update(&tmp_goods,"ViewNum")
	return
}


func (m *MacModel) GetListPage(where map[string]string, order string, pageIndex int, pageNum int) (totalCount int, list []*GetMacListPageResult, err error) {
	offset := (pageIndex - 1) * pageNum
	o := orm.NewOrm()
	sql_list :="select a.goods_id,a.download_url,g.* from ks_mac as a left join ks_goods as g on g.id=a.goods_id where g.is_del =0"
	sql_count :="select count(*) from ks_mac as a left join ks_goods as g on g.id=a.goods_id where g.is_del =0"
	for key, value := range where {
		switch key {
		case "name":
			sql_list += " and name like %"+ value +"%"
			sql_count += " and name like %"+ value +"%"
			break
		case "is_verify":
			sql_list += " and is_verify ="+ value
			sql_count += " and is_verify ="+ value
			break
		default:

		}
	}
	sql_list +=" order by ? limit ?,?"
	o.Raw(sql_list, order,offset,pageNum).QueryRows(&list)
	o.Raw(sql_count).QueryRow(&totalCount)

	return
}

