// @APIVersion 1.0.0
// @Title mobile API
// @Description mobile has every tool to get any job done, so codename for the new mobile APIs.
// @Contact astaxie@gmail.com
package routers

import (
	"engineercms/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/url-to-callback", &controllers.OnlyController{}, "*:UrltoCallback")
	// beego.Router("/onlyoffice/post", &controllers.OnlyController{}, "post:PostOnlyoffice")
	beego.Router("/onlyoffice", &controllers.OnlyController{}, "get:Get")
	//table获取所有数据给上面界面使用
	beego.Router("/onlyoffice/data", &controllers.OnlyController{}, "*:GetData")
	//添加一个文档
	beego.Router("/onlyoffice/addattachment", &controllers.OnlyController{}, "post:AddOnlyAttachment")
	//在onlyoffice中打开文档协作
	beego.Router("/onlyoffice/:id:string", &controllers.OnlyController{}, "*:OnlyOffice")
	//删除
	beego.Router("/onlyoffice/deletedoc", &controllers.OnlyController{}, "*:DeleteDoc")
	//修改
	beego.Router("/onlyoffice/updatedoc", &controllers.OnlyController{}, "*:UpdateDoc")

	//下载doc
	beego.Router("/attachment/onlyoffice/*", &controllers.OnlyController{}, "get:DownloadDoc")
}
