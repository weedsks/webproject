package main

import (
	"github.com/astaxie/beego/orm"
	_ "webproject/routers"
	"github.com/astaxie/beego"
)

func main() {
	orm.Debug = true
	beego.Run()
}

