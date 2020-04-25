package admin

import (
	"github.com/Echosong/beego_blog/util"
	"github.com/astaxie/beego"
	"strings"
	"time"
	"webproject/models"
)
// CMS API
type IndexController struct {
	baseController
}
//@router / [get]
func (c *IndexController) Index()  {
	//获取左侧菜单
	c.Data["menus"] = models.NewMenus().GetMenu(0)
	c.Data["userInfo"] = c.GetSession("userInfo")
	c.TplName = c.getAdminFix() + "/index.html"
}
//@router /welcome
func (c *IndexController) Welcome()  {
	c.Data["userInfo"] = c.GetSession("userInfo")
	c.Data["nowTime"] = time.Now().Format("2006-01-02 15:04:05")
	c.TplName = c.getAdminFix() + "/welcome.html"
}
//@router /login [*]
func (c *IndexController) Login() {
	if c.Ctx.Request.Method == "POST" {
		keyNmae := strings.TrimSpace(c.GetString("username"))
		keyPwd := strings.TrimSpace(c.GetString("password"))

		if len(keyNmae) > 1 && len(keyPwd) > 1 {
			keyPwd = util.Md5(keyPwd+beego.AppConfig.String("pwdsalt"));
			user := models.NewUser().VerificationLogin(keyNmae)
			if user.Password == keyPwd {
				c.SetSession("userInfo", user)
				c.JsonResult(0, "登录成功！", 1)
				//c.Redirect("/", 302)
			}
		}
		c.JsonResult(0, "账号或密码错误！",1)

	}
	c.TplName = c.getAdminFix() + "/login.html"
}
//@router /logout
func (c *IndexController) Logout() {
	c.DelSession("userInfo")
	c.Redirect("/admin/login", 302)
}
