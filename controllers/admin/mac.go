package admin

import (
	//"github.com/Echosong/beego_blog/util"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
	"unsafe"
	"webproject/models"
)

type MacController struct {
	baseController
}
//@router /macIndex [get]
func (c *MacController) Index() {
	where := map[string]string{
		"IsDel": "0",
	}
	name := c.GetString("name")
	if len(name) > 1 {
		where["name"] = name
	}
	is_verify := c.GetString("is_verify")

	if len(is_verify) > 1 {
		where["is_verify"] = is_verify
	}
	order := "g.order desc"

	totalCount, data, merr := models.NewMac().GetListPage(where, order, c.page, c.limit)
	if merr != nil {
		beego.Notice(merr)
	}

	c.SetPaginator(c.limit,totalCount)

	c.Data["count"] = totalCount
	c.Data["list"] = data
	c.TplName = c.getAdminFix() + "/mac-list.html"
}
//@router /macDetail [get]
func (c *MacController) Detail() {
	id, err := c.GetInt("id")
	if err != nil {
		beego.Warning("参数错误->", err)
	}
	info, err := models.NewMac().GetInfoById(id)
	c.Data["info"] = info
	typeComon_arr, _ := models.NewTypeCommon().GetAll(3)
	for key, _ := range typeComon_arr {
		if _, ok := info.TypeCommon[key]; ok {
			typeComon_arr[key].IsChecked = 1
		}
	}
	//c.JsonResult(0,"weeds",0,info)
	c.Data["typeCommon"] = typeComon_arr
	c.TplName = c.getAdminFix() + "/mac-edit.html"
}
//@router /macUpdate [post]
func (c *MacController) Update() {
	id, err := c.GetInt("Id")
	if err != nil {
		c.JsonResult(-1, "参数错误ID！", 0)
	}

	oldInfo, err := models.NewMac().GetInfoById(id)
	if err != nil {
		c.JsonResult(-1, err.Error(), 0)
	}

	name := c.GetString("Name", oldInfo.Name)
	description := c.GetString("Description", oldInfo.Description)
	smallImage := c.GetString("SmallImage", oldInfo.SmallImage)
	image := c.GetString("Image", oldInfo.Image)

	isComment, err := c.GetInt("IsComment", oldInfo.IsComment)
	isVerify, err := c.GetInt("IsVerify", oldInfo.IsVerify)
	is_del, err := c.GetInt("IsDel", oldInfo.IsDel)
	isRecommend, err := c.GetInt("IsRecommend", oldInfo.IsRecommend)
	order, err := c.GetInt("Order", oldInfo.Order)
	recommendStars, err := c.GetInt("RecommendStars", oldInfo.RecommendStars)

	content := c.GetString("Content", oldInfo.Content)
	downurl := c.GetString("DownloadUrl", oldInfo.DownloadUrl)

	if err != nil {
		c.JsonResult(-1, err.Error(), 0)
	}

	o := orm.NewOrm()
	//更新附表
	mac := models.MacModel{GoodsId: id}
	if o.Read(&mac, "GoodsId") == nil {
		mac.Content = content
		mac.DownloadUrl = downurl
		o.Update(&mac, "Content")
	}

	//更新主表
	goods := models.GoodsModel{Id: id}
	if o.Read(&goods) == nil {
		goods.Image = image
		goods.SmallImage = smallImage
		goods.Description = description
		goods.Name = name
		goods.IsComment = isComment
		goods.IsVerify = isVerify
		goods.IsRecommend = isRecommend
		goods.Order = order
		goods.RecommendStars = recommendStars
		goods.IsDel = is_del
		o.Update(&goods, "Image", "SmallImage", "Description", "Name", "IsComment", "IsVerify", "IsRecommend", "Order", "RecommendStars","IsDel")
	}

	//更新 封面标签
	typeCommon := c.GetStrings("TypeCommons")
	relationInfo := models.GoodsTypeCommonRelationModel{GoodsId: id}
	//删除之前数据重新添加
	o.Delete(&relationInfo, "GoodsId")
	multiData := []models.GoodsTypeCommonRelationModel{}
	for i := 0; i < len(typeCommon); i++ {
		tid, _ := strconv.Atoi(typeCommon[i])
		multiData = append(multiData, models.GoodsTypeCommonRelationModel{
			TypeCommonId: tid,
			GoodsId:      id,
		})
	}
	o.InsertMulti(len(typeCommon), multiData)

	c.JsonResult(0, "请求成功！", 0)

}
//@router /macCreate [*]
func (c *MacController) Create() {
	if c.Ctx.Request.Method == "POST" {
		name := c.GetString("Name")
		description := c.GetString("Description")
		smallImage := c.GetString("SmallImage")
		image := c.GetString("Image")

		isComment, err := c.GetInt("IsComment")
		isVerify, err := c.GetInt("IsVerify")
		isRecommend, err := c.GetInt("IsRecommend")
		order, err := c.GetInt("Order")
		recommendStars, err := c.GetInt("RecommendStars")

		downurl := c.GetString("DownloadUrl")
		content := c.GetString("Content")

		if err != nil {
			c.JsonResult(-1, err.Error(), 0)
		}

		o := orm.NewOrm()

		//更新主表
		goods := models.GoodsModel{}
		goods.Image = image
		goods.Type = 3
		goods.SmallImage = smallImage
		goods.Description = description
		goods.Name = name
		goods.IsComment = isComment
		goods.IsVerify = isVerify
		goods.IsRecommend = isRecommend
		goods.Order = order
		goods.RecommendStars = recommendStars
		id, err := o.Insert(&goods)
		if err != nil {
			c.JsonResult(-1, err.Error(), 0)
		}
		goods_id := *(*int)(unsafe.Pointer(&id))
		//更新附表
		mac := models.MacModel{}
		mac.GoodsId = goods_id
		mac.DownloadUrl = downurl
		mac.Content = content
		o.Insert(&mac)

		//更新 封面标签
		typeCommon := c.GetStrings("TypeCommons")
		multiData := []models.GoodsTypeCommonRelationModel{}
		for i := 0; i < len(typeCommon); i++ {
			tid, _ := strconv.Atoi(typeCommon[i])
			multiData = append(multiData, models.GoodsTypeCommonRelationModel{
				TypeCommonId: tid,
				GoodsId:      goods_id,
			})
		}
		o.InsertMulti(len(typeCommon), multiData)

		c.JsonResult(0, "请求成功！", 0)
	}
	typeComon_arr, _ := models.NewTypeCommon().GetAll(3)
	c.Data["typeCommon"] = typeComon_arr
	c.TplName = c.getAdminFix() + "/mac-add.html"

}
//@router /macDelAll [post]
func (c *MacController) DelAll()  {
	ids := c.GetString("ids")
	ids_arr := strings.Split(ids, ",")
	num, err := models.NewGoods().DelAll(ids_arr)
	if err!=nil {
		c.JsonResult(-1, "更新失败！"+err.Error(), 1, num)
	}

	c.JsonResult(0, "请求成功！", 1, num)

}
