package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/eventos_crud/models"
)

type AuditoriaEventosController struct {
	beego.Controller
}

func (c *AuditoriaEventosController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// Post ...
// @Title Post
// @Description registrar un evento de auditoría
// @Param	body		body 	models.AuditoriaEventos	true	"body for auditoria content"
// @Success 201 {int} models.AuditoriaEventos
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *AuditoriaEventosController) Post() {
	var v models.AuditoriaEventos
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		v.Activo = true
		v.FechaCreacion = fechaActual()
		v.FechaModificacion = fechaActual()
		if _, err := models.AddAuditoriaEventos(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			logs.Error(err)
			c.Data["json"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
			c.Ctx.Output.SetStatus(400)
		}
	} else {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
		c.Ctx.Output.SetStatus(400)
	}
	c.ServeJSON()
}
