package main
import (
	"database/sql"
	"github.com/astaxie/beego/toolbox"
)
type DatabaseCheck struct{
}
func (*DatabaseCheck) Check() error{
	_,err:=sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/yy?charset=utf8")
	if err!=nil{
		return err
	}
	return nil
}

func init(){
	toolbox.AddHealthCheck("database",&DatabaseCheck{})
}



