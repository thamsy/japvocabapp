package controllers

import (
	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"
	"japapp/secret"
	"log"
)

type LoginController struct {
	beego.Controller
}

//Login process
func (this *LoginController) Get() {
	this.Data["WrongPw"] = "hidden"
	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	err := this.Ctx.Request.ParseForm()
	if err != nil {
		log.Fatalf("Parse Form Error")
	}
	username := this.Ctx.Request.Form.Get("email")
	pwd := this.Ctx.Request.Form.Get("password")

	// Generate Hash
	//hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Printf(string(hash))

	err = bcrypt.CompareHashAndPassword([]byte(secret.PW_HASH), []byte(pwd))
	if username == secret.EMAIL && err == nil{
		//Set the session successful login
		sess, err := beego.GlobalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
		if err != nil {
			log.Fatalf("%+v", err)
		}
		sess.Set("uid", 94)
		this.Ctx.Output.Status = 200
	} else {
		this.Ctx.Output.Status = 401
	}
}

func (this *LoginController) Logout()  {
	beego.GlobalSessions.SessionDestroy(this.Ctx.ResponseWriter, this.Ctx.Request)
	this.Ctx.Redirect(302, "/login")
}

////Registration process
//func (this *RegController) Post() {
//	this.TplNames = "reg.tpl"
//	this.Ctx.Request.ParseForm()
//	username := this.Ctx.Request.Form.Get("username")
//	password := this.Ctx.Request.Form.Get("password")
//	usererr := checkUsername(username)
//	fmt.Println(usererr)
//	if usererr == false {
//		this.Data["UsernameErr"] = "Username error, Please to again"
//		return
//	}
//
//	passerr := checkPassword(password)
//	if passerr == false {
//		this.Data["PasswordErr"] = "Password error, Please to again"
//		return
//	}
//
//	md5Password := md5.New()
//	io.WriteString(md5Password, password)
//	buffer := bytes.NewBuffer(nil)
//	fmt.Fprintf(buffer, "%x", md5Password.Sum(nil))
//	newPass := buffer.String()
//
//	now := time.Now().Format("2006-01-02 15:04:05")
//
//	userInfo := models.GetUserInfo(username)
//
//	if userInfo.Username == "" {
//		var users models.User
//		users.Username = username
//		users.Password = newPass
//		users.Created = now
//		users.Last_logintime = now
//		models.AddUser(users)
//
//		//Set the session successful login
//		sess := globalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
//		sess.Set("uid", userInfo.Id)
//		sess.Set("uname", userInfo.Username)
//		this.Ctx.Redirect(302, "/")
//	} else {
//		this.Data["UsernameErr"] = "User already exists"
//	}
//
//}
//
//func checkPassword(password string) (b bool) {
//	if ok, _ := regexp.MatchString("^[a-zA-Z0-9]{4,16}$", password); !ok {
//		return false
//	}
//	return true
//}
//
//func checkUsername(username string) (b bool) {
//	if ok, _ := regexp.MatchString("^[a-zA-Z0-9]{4,16}$", username); !ok {
//		return false
//	}
//	return true
//}
