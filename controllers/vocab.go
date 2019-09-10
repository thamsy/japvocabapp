package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"japapp/secret"
	"japapp/sheetsapi"
	"japapp/util"
	"log"
	"strconv"
)

type VocabController struct {
	beego.Controller
}

type IncrQuestionRequest struct {
	IsKanaCorr bool
	IsKanjiCorr bool
	IsGrpCorr bool
}

func (c *VocabController) GetAllQuestions() {
	questions := sheetsapi.GetAllNewQuestions()
	c.Data["json"] = &questions
	c.ServeJSON()
}

func (c *VocabController) GetSomeQuestions() {
	num, _ := strconv.Atoi(c.Ctx.Input.Param(":num"))
	questions := sheetsapi.GetAllNewQuestions()
	util.Shuffle(questions)
	if len(questions) > num {
		questions = questions[:num]
	}
	c.Data["json"] = &questions
	c.ServeJSON()
}



func (c *VocabController) IncrNewQuestion() {
	// --- TESTACC
	uid := fmt.Sprintf("%+v", c.Ctx.Input.Session("uid"))
	if uid == secret.TESTACC_UID {
		c.Ctx.Output.Body([]byte("nothing done!"))
		return
	}
	// TESTACC ---

	var ob IncrQuestionRequest
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	if err != nil {
		log.Printf("%+v", c.Ctx.Input.RequestBody)
		log.Printf("%+v", err)
		c.Ctx.Output.Status = 400
		c.Ctx.Output.Body([]byte("Bad JSON"))
		return
	}
	sheetsapi.IncrNewQuestion(c.Ctx.Input.Param(":row"), ob.IsKanaCorr, ob.IsKanjiCorr, ob.IsGrpCorr)
	c.Ctx.Output.Body([]byte("ok!"))
}

func (c *VocabController) CheckUnderstoodNewVocab() {
	// --- TESTACC
	uid := fmt.Sprintf("%+v", c.Ctx.Input.Session("uid"))
	if uid == secret.TESTACC_UID {
		c.Ctx.Output.Body([]byte("nothing done!"))
		return
	}
	// TESTACC ---

	sheetsapi.CheckUnderstoodNewQuestions()
	c.Ctx.Output.Body([]byte("checked whether understood questions!"))
}