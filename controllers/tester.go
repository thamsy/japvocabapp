package controllers

import (
	"github.com/astaxie/beego"
)

type TesterController struct {
	beego.Controller
}

func (c *TesterController) Get() {
	c.TplName = "tester.html"
}