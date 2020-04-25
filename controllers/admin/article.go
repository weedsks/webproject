package admin

import (
	//"github.com/Echosong/beego_blog/util"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"unsafe"
	"webproject/models"
)

type ArticleController struct {
	baseController
}
//@router /articleIndex [get]
func (c *ArticleController) Index() {
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

	totalCount, data, merr := models.NewArticle().GetListPage(where, order, c.page, c.limit)
	if merr != nil {
		beego.Notice(merr)
	}

	c.SetPaginator(c.limit,totalCount)

	c.Data["count"] = totalCount
	c.Data["list"] = data

	c.TplName = c.getAdminFix() + "/article-list.html"
}
//@router /articleDetail [get]
func (c *ArticleController) Detail() {
	id, err := c.GetInt("id")
	if err != nil {
		beego.Warning("参数错误->", err)
	}
	info, err := models.NewArticle().GetInfoById(id)
	c.Data["info"] = info
	typeComon_arr, _ := models.NewTypeCommon().GetAll(1)
	for key, _ := range typeComon_arr {
		if _, ok := info.TypeCommon[key]; ok {
			typeComon_arr[key].IsChecked = 1
		}
	}
	//c.JsonResult(0,"weeds",0,info)
	c.Data["typeCommon"] = typeComon_arr
	c.TplName = c.getAdminFix() + "/article-edit.html"
}
//@router /getQiNiuToken [*]
func (c *ArticleController) GetQiNiuToken() {
	token := c.getToken("article")
	c.Ctx.WriteString(token)
}
//@router /articleUpdate [post]
func (c *ArticleController) Update() {
	id, err := c.GetInt("Id")
	if err != nil {
		c.JsonResult(-1, "参数错误ID！", 0)
	}

	oldInfo, err := models.NewArticle().GetInfoById(id)
	if err != nil {
		c.JsonResult(-1, err.Error(), 0)
	}

	name := c.GetString("Name", oldInfo.Name)
	description := c.GetString("Description", oldInfo.Description)
	smallImage := c.GetString("SmallImage", oldInfo.SmallImage)
	image := c.GetString("Image", oldInfo.Image)

	isComment, err := c.GetInt("IsComment", oldInfo.IsComment)
	isVerify, err := c.GetInt("IsVerify", oldInfo.IsVerify)
	isRecommend, err := c.GetInt("IsRecommend", oldInfo.IsRecommend)
	order, err := c.GetInt("Order", oldInfo.Order)
	recommendStars, err := c.GetInt("RecommendStars", oldInfo.RecommendStars)

	content := c.GetString("Content", oldInfo.Content)

	if err != nil {
		c.JsonResult(-1, err.Error(), 0)
	}

	o := orm.NewOrm()
	//更新附表
	article := models.ArticleModel{GoodsId: id}
	if o.Read(&article, "GoodsId") == nil {
		article.Content = content
		o.Update(&article, "Content")
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
		o.Update(&goods, "Image", "SmallImage", "Description", "Name", "IsComment", "IsVerify", "IsRecommend", "Order", "RecommendStars")
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
//@router /articleCreate [*]
func (c *ArticleController) Create() {
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

		content := c.GetString("Content")

		if err != nil {
			c.JsonResult(-1, err.Error(), 0)
		}

		o := orm.NewOrm()

		//更新主表
		goods := models.GoodsModel{}
		goods.Image = image
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
		article := models.ArticleModel{}
		article.GoodsId = goods_id
		article.Content = content
		o.Insert(&article)

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
	typeComon_arr, _ := models.NewTypeCommon().GetAll(1)
	c.Data["typeCommon"] = typeComon_arr
	c.TplName = c.getAdminFix() + "/article-add.html"

}
