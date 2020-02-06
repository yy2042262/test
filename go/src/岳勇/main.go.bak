package main

import (
	_ "岳勇/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	orm.RegisterDataBase("default","mysql","root:root@tcp(127.0.0.1:3306)/yy")
	beego.Run()
}

