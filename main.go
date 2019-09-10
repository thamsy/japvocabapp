package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	_ "japapp/routers"
	"log"
)

func main() {
	var err error
	beego.GlobalSessions, err = session.NewManager("memory",
		&session.ManagerConfig{
			CookieName:              "gosessionid",
			EnableSetCookie:         true,
			Gclifetime:              3600,
			Maxlifetime:             3600,
			DisableHTTPOnly:         true,
			Secure:                  true,
			CookieLifeTime:          3600,
			ProviderConfig:          "{}",
			Domain:                  "/",
		})
	go beego.GlobalSessions.GC()

	if err != nil {
		log.Fatalf("%+v", err)
	}

	beego.Run()
}

