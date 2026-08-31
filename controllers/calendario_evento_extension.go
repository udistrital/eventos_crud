package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/eventos_crud/models"
)

type CalendarioEventoExtensionController struct {
	beego.Controller
}

func (c *CalendarioEventoExtensionController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @router / [post]
func (c *CalendarioEventoExtensionController) Post() {
	var v models.CalendarioEventoExtension
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		v.FechaCreacion = fechaActual()
		v.FechaModificacion = fechaActual()
		if _, err := models.AddCalendarioEventoExtension(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = normalizarCalendarioEventoExtensionRespuesta(v)
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
func (c *CalendarioEventoExtensionController) GetOne() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	v, err := models.GetCalendarioEventoExtensionById(id)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
		c.Ctx.Output.SetStatus(400)
	} else {
		c.Data["json"] = normalizarCalendarioEventoExtensionRespuesta(v)
	}
	c.ServeJSON()
}

// @router / [get]
func (c *CalendarioEventoExtensionController) GetAll() {
	fields, sortby, order, query, limit, offset, ok := parseCrudQuery(&c.Controller)
	if !ok {
		return
	}
	l, err := models.GetAllCalendarioEventoExtension(query, fields, sortby, order, offset, limit)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
		c.Ctx.Output.SetStatus(400)
	} else {
		if l == nil {
			l = append(l, map[string]interface{}{})
		}
		c.Data["json"] = normalizarCalendarioEventoExtensionRespuesta(l)
	}
	c.ServeJSON()
}

// @router /:id [put]
func (c *CalendarioEventoExtensionController) Put() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	v := models.CalendarioEventoExtension{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		v.FechaCreacion = fechaCorreccion(v.FechaCreacion)
		v.FechaModificacion = fechaActual()
		if err := models.UpdateCalendarioEventoExtensionById(&v); err == nil {
			c.Data["json"] = normalizarCalendarioEventoExtensionRespuesta(v)
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
func (c *CalendarioEventoExtensionController) Delete() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err := models.DeleteCalendarioEventoExtension(id); err == nil {
		c.Data["json"] = map[string]interface{}{"Id": id}
	} else {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
		c.Ctx.Output.SetStatus(400)
	}
	c.ServeJSON()
}

func parseCrudQuery(c *beego.Controller) (fields []string, sortby []string, order []string, query map[string]string, limit int64, offset int64, ok bool) {
	query = make(map[string]string)
	limit = 10
	ok = true
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return fields, sortby, order, query, limit, offset, false
			}
			query[kv[0]] = kv[1]
		}
	}
	return
}
