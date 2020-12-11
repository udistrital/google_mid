package controllers

import (
	"fmt"
	"encoding/json"
	// "errors"
	// "strconv"
	// "strings"
	// "strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	"github.com/udistrital/google_mid/services/correo"
	"github.com/udistrital/google_mid/models"
)



type NotificacionController struct {
	beego.Controller
}

// URLMapping ...
func (c *NotificacionController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// Post ...
// @Title Post
// @Description Enviar una notificacion
// @Param   body        body    {}  true        "body Enviar notificacion content"
// @Success 200 {}
// @Failure 403 body is empty
// @router / [post]
func (c *NotificacionController) Post() {
	fmt.Println("Send email")
	var v models.Notificacion
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {

		correo.OAuthGmailService()
		_, err = correo.SendEmailOAUTH2_V2(v)
		if err != nil{
			logs.Error(err)
			c.Data["system"] = err
			c.Abort("400")
		}
		c.Ctx.Output.SetStatus(201)
	} else {
		logs.Error(err)
		c.Data["mesaage"] = "Error service POST: The request contains an incorrect data type or an invalid parameter"
		c.Abort("400")
	}
	c.ServeJSON()
}


