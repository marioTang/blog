package routers

import (
	"devops/controllers"
	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.HomeController{})
	////注册
	beego.Router("/register", &controllers.RegisterController{})
	////登录
	beego.Router("/login", &controllers.LoginController{})
	////退出
	beego.Router("/exit", &controllers.ExitController{})

	//相册
	beego.Router("/album", &controllers.AlbumController{})

	//文件上传
	beego.Router("/upload", &controllers.UploadController{})

	beego.Router("/article", &controllers.EditorController{})

	beego.Router("/page", &controllers.GetpagController{})

	beego.Router("/details", &controllers.DetailsController{})

	beego.Router("/about", &controllers.AboutController{})

	beego.Router("/show", &controllers.ShowController{})

	beego.Router("/editing", &controllers.EditingArticlesController{})

}





