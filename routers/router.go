// @APIVersion 1.0.0
// @Title mobile API
// @Description mobile has every tool to get any job done, so codename for the new mobile APIs.
// @Contact astaxie@gmail.com
package routers

import (
	"github.com/3xxx/engineercms/controllers"
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
	//onlyoffice页面下载doc
	beego.Router("/attachment/onlyoffice/*", &controllers.OnlyController{}, "get:DownloadDoc")
	//文档管理页面下载doc
	beego.Router("/onlyoffice/download/:id:string", &controllers.OnlyController{}, "get:Download")
	// beego.Router("/onlyoffice/changes", &controllers.OnlyController{}, "post:ChangesUrl")
	//*****onlyoffice document权限
	//添加用户和角色权限
	beego.Router("/onlyoffice/addpermission", &controllers.OnlyController{}, "post:Addpermission")
	//取得文档的用户和角色权限列表
	beego.Router("/onlyoffice/getpermission", &controllers.OnlyController{}, "get:Getpermission")
}
