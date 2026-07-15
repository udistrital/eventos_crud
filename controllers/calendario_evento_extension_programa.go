package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/eventos_crud/models"
)

type CalendarioEventoExtensionProgramaController struct {
	beego.Controller
}

func (c *CalendarioEventoExtensionProgramaController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @router / [post]
func (c *CalendarioEventoExtensionProgramaController) Post() {
	var v models.CalendarioEventoExtensionPrograma
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		v.FechaCreacion = fechaActual()
		v.FechaModificacion = fechaActual()
		if _, err := models.AddCalendarioEventoExtensionPrograma(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = normalizarCalendarioEventoExtensionProgramaRespuesta(v)
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

// @router /:id [get]
func (c *CalendarioEventoExtensionProgramaController) GetOne() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	v, err := models.GetCalendarioEventoExtensionProgramaById(id)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
		c.Ctx.Output.SetStatus(400)
	} else {
		c.Data["json"] = normalizarCalendarioEventoExtensionProgramaRespuesta(v)
	}
	c.ServeJSON()
}

// @router / [get]
func (c *CalendarioEventoExtensionProgramaController) GetAll() {
	fields, sortby, order, query, limit, offset, ok := parseCrudQuery(&c.Controller)
	if !ok {
		return
	}
	l, err := models.GetAllCalendarioEventoExtensionPrograma(query, fields, sortby, order, offset, limit)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
		c.Ctx.Output.SetStatus(400)
	} else {
		if l == nil {
			l = append(l, map[string]interface{}{})
		}
		c.Data["json"] = normalizarCalendarioEventoExtensionProgramaRespuesta(l)
	}
	c.ServeJSON()
}

// @router /:id [put]
func (c *CalendarioEventoExtensionProgramaController) Put() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	v := models.CalendarioEventoExtensionPrograma{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		v.FechaCreacion = fechaCorreccion(v.FechaCreacion)
		v.FechaModificacion = fechaActual()
		if err := models.UpdateCalendarioEventoExtensionProgramaById(&v); err == nil {
			c.Data["json"] = normalizarCalendarioEventoExtensionProgramaRespuesta(v)
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

// @router /:id [delete]
func (c *CalendarioEventoExtensionProgramaController) Delete() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err := models.DeleteCalendarioEventoExtensionPrograma(id); err == nil {
		c.Data["json"] = map[string]interface{}{"Id": id}
	} else {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
		c.Ctx.Output.SetStatus(400)
	}
	c.ServeJSON()
}
