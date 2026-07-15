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

type EventoCatalogoController struct {
	beego.Controller
}

func (c *EventoCatalogoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create EventoCatalogo
// @Param	body		body 	models.EventoCatalogo	true		"body for EventoCatalogo content"
// @Success 201 {int} models.EventoCatalogo
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *EventoCatalogoController) Post() {
	var v models.EventoCatalogo
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		v.FechaCreacion = fechaActual()
		v.FechaModificacion = fechaActual()
		if _, err := models.AddEventoCatalogo(&v); err == nil {
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

// GetOne ...
// @Title Get One
// @Description get EventoCatalogo by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.EventoCatalogo
// @Failure 404 not found resource
// @router /:id [get]
func (c *EventoCatalogoController) GetOne() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	v, err := models.GetEventoCatalogoById(id)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
		c.Ctx.Output.SetStatus(400)
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get EventoCatalogo
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.EventoCatalogo
// @Failure 404 not found resource
// @router / [get]
func (c *EventoCatalogoController) GetAll() {
	var fields, sortby, order []string
	query := make(map[string]string)
	limit := int64(10)
	var offset int64
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
				return
			}
			query[kv[0]] = kv[1]
		}
	}
	l, err := models.GetAllEventoCatalogo(query, fields, sortby, order, offset, limit)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
		c.Ctx.Output.SetStatus(400)
	} else {
		if l == nil {
			l = append(l, map[string]interface{}{})
		}
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the EventoCatalogo
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.EventoCatalogo	true		"body for EventoCatalogo content"
// @Success 200 {object} models.EventoCatalogo
// @Failure 400 the request contains incorrect syntax
// @router /:id [put]
func (c *EventoCatalogoController) Put() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	v := models.EventoCatalogo{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		v.FechaCreacion = fechaCorreccion(v.FechaCreacion)
		v.FechaModificacion = fechaActual()
		if err := models.UpdateEventoCatalogoById(&v); err == nil {
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

// Delete ...
// @Title Delete
// @Description delete the EventoCatalogo
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 404 not found resource
// @router /:id [delete]
func (c *EventoCatalogoController) Delete() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err := models.DeleteEventoCatalogo(id); err == nil {
		c.Data["json"] = map[string]interface{}{"Id": id}
	} else {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
		c.Ctx.Output.SetStatus(400)
	}
	c.ServeJSON()
}
