package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"time"
)

func init(){
	dbuser := beego.AppConfig.String("dbuser")
	dbpwd := beego.AppConfig.String("dbpwd")
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbname := beego.AppConfig.String("dbname")


	orm.RegisterDataBase("default", "mysql", dbuser+":"+ dbpwd +"@tcp("+dbhost+":"+ dbport +")/"+ dbname +"?charset=utf8", 30)
	orm.RegisterModel(new(GoodsModel),new(UsersModel),new(RolesModel),new(MenusModel),new(ArticleModel),new(TypeCommonModel), new(GoodsTypeCommonRelationModel))
}

func TableNmae(str string) string {
	return beego.AppConfig.String("dbprefix") + str
}

func GoodsTypeDesc() (typeDesc map[int]string) {
	typeDesc =map[int]string{
		1:"文章",
		2:"电影",
		3:"MAC下载",
	}
	return
}

func getPrivateQiNiuUrl(Ourl string) (Nurl string) {
	max := qbox.NewMac(beego.AppConfig.String("QiNiu_AccessKey"), beego.AppConfig.String("QiNiu_SecretKey"))
	doma := beego.AppConfig.String("QiNiu_domain")
	deadline := time.Now().Add(time.Second * 3600).Unix() //1小时有效期

	Nurl = storage.MakePrivateURL(max,doma, Ourl,deadline)
	return
}