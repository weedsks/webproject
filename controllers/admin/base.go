package admin

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"io"
	"math/rand"
	"strconv"
	"strings"
)

type baseController struct {
	beego.Controller
	controllerName string
	actionName     string
	limit          int
	page           int
}

//获取用户IP地址
func (p *baseController) getClientIp() string {
	s := strings.Split(p.Ctx.Request.RemoteAddr, ":")
	return s[0]
}

//返回静态文件模块目录
func (p *baseController) getAdminFix() string {
	return beego.AppConfig.String("adminfix")
}

func (c *baseController) JsonResult(errCode int, errMsg string, count int, data ...interface{}) {
	jsonData := make(map[string]interface{}, 4)

	jsonData["code"] = errCode
	jsonData["msg"] = errMsg
	jsonData["count"] = count
	//扩展字段
	//for key,value:=range extendField{
	//	jsonData[key]=value
	//}

	if len(data) > 0 && data[0] != nil {
		jsonData["data"] = data[0]
	}

	returnJSON, err := json.Marshal(jsonData)

	if err != nil {
		beego.Error(err)
	}

	c.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	c.Ctx.ResponseWriter.Header().Set("Cache-Control", "no-cache, no-store")
	io.WriteString(c.Ctx.ResponseWriter, string(returnJSON))

	c.StopRun()
}

func (p *baseController) Prepare() {
	p.controllerName, p.actionName = p.GetControllerAndAction()
	requestUrl := p.Ctx.Request.RequestURI
	tmp_page, _ := p.GetInt("p", 1)
	p.page, _ = p.GetInt("page", tmp_page)

	p.limit, _ = p.GetInt("limit", 10)

	//验证是否登录
	p.CheckLogin(requestUrl)
}

func (p *baseController) CheckLogin(requestUrl string) bool {
	//免登录的方法
	unLoginAction := map[string]int{
		"/admin/login": 1,
	}

	user := p.GetSession("userInfo")
	//fmt.Println(p.controllerName)
	//fmt.Println(p.actionName)
	if (user == nil) && (unLoginAction[requestUrl] == 0) {
		p.Redirect("/admin/login", 302)
		p.StopRun()
	} else if (user != nil) && (unLoginAction[requestUrl] != 0) {
		p.Redirect("/admin", 302)
		p.StopRun()
	}
	return true
}

func (p *baseController) getToken(module string) (token string) {
	putPolicy := storage.PutPolicy{
		Scope:beego.AppConfig.String("QiNiu_bucket"),
		ReturnBody: `{
			"hash":"$(etag)",
			"code":"0",
			"msg":"success",
			"data":{
				"src": "$(key)",
				"title": "$(x:name)"
			}
		}`,
		SaveKey:module+"/$(year)-$(mon)-$(day)-$(hour)-$(min)-$(sec)-"+ strconv.Itoa(rand.Intn(10000)) +"$(ext)",
	}

	max := qbox.NewMac(beego.AppConfig.String("QiNiu_AccessKey"), beego.AppConfig.String("QiNiu_SecretKey"))

	token = putPolicy.UploadToken(max)
	return
}

func (this *baseController) SetPaginator(per int, nums int)  *pagination.Paginator {
	p := pagination.NewPaginator(this.Ctx.Request, per, nums)
	this.Data["paginator"] = p
	return p
}