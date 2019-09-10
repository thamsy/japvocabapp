package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"japapp/controllers"
	"strings"
)

var FilterUser = func(ctx *context.Context) {
	uid := ctx.Input.Session("uid")
	if strings.HasPrefix(ctx.Input.URL(), "/login") {
		if uid != nil {
			ctx.Redirect(302, "/")
		}
		return
	}
	if uid == nil {
		ctx.Redirect(302, "/login")
	}
}


func init() {
	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")
    beego.Router("/", &controllers.HomeController{})
	beego.Router("/tester", &controllers.TesterController{})
	beego.Router("/vocab/all", &controllers.VocabController{}, "get:GetAllQuestions")
	beego.Router("/vocab/:num([0-9]+)", &controllers.VocabController{}, "get:GetSomeQuestions")
    beego.Router("/vocab/incr/:row([0-9]+)", &controllers.VocabController{}, "post:IncrNewQuestion")
	beego.Router("/vocab/check", &controllers.VocabController{}, "post:CheckUnderstoodNewVocab")
}
