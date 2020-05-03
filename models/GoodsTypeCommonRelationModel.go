package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type GoodsTypeCommonRelationModel struct {
	Id             int
	TypeCommonId int
	GoodsId       int
}

func (m *GoodsTypeCommonRelationModel) TableName() string {
	return TableNmae("goods_type_relation")
}

func NewGoodsTypeCommonRelation() *GoodsTypeCommonRelationModel {
	return &GoodsTypeCommonRelationModel{}
}

func (m *GoodsTypeCommonRelationModel) GetRelation(goodsId int) (lists map[int]TypeCommonModel, err error) {
	o := orm.NewOrm()
	sql_list := "select t.id,t.name from ks_goods_type_relation as r left join ks_type_common as t on t.id=r.type_common_id where r.goods_id = ?"
	tmp_list := []*TypeCommonModel{}
	_,err = o.Raw(sql_list, goodsId).QueryRows(&tmp_list)
	if err == orm.ErrNoRows {
		beego.Warning("数据查找不到", err)
	}

	lists = make(map[int]TypeCommonModel)
	for _, item := range tmp_list {
		lists[item.Id] = *item
	}

	return
}
