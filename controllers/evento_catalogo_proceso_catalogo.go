package controllers

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/udistrital/eventos_crud/models"
)

type EventoCatalogoProcesoCatalogoController struct{ beego.Controller }

func (c *EventoCatalogoProcesoCatalogoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create EventoCatalogoProcesoCatalogo
// @Param	body		body 	models.EventoCatalogoProcesoCatalogo	true		"body for EventoCatalogoProcesoCatalogo content"
// @Success 201 {int} models.EventoCatalogoProcesoCatalogo
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *EventoCatalogoProcesoCatalogoController) Post() {
	var v models.EventoCatalogoProcesoCatalogo
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		v.FechaCreacion = fechaActual()
		v.FechaModificacion = fechaActual()
		if _, err := models.AddEventoCatalogoProcesoCatalogo(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get EventoCatalogoProcesoCatalogo by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.EventoCatalogoProcesoCatalogo
// @Failure 404 not found resource
// @router /:id [get]
func (c *EventoCatalogoProcesoCatalogoController) GetOne() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	v, err := models.GetEventoCatalogoProcesoCatalogoById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get EventoCatalogoProcesoCatalogo
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.EventoCatalogoProcesoCatalogo
// @Failure 404 not found resource
// @router / [get]
func (c *EventoCatalogoProcesoCatalogoController) GetAll() {
	query := make(map[string]string)
	fields := []string{}
	sortby := []string{}
	order := []string{}
	var limit int64 = 10
	var offset int64

	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) == 2 {
				query[kv[0]] = kv[1]
			}
		}
	}
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}

	l, err := models.GetAllEventoCatalogoProcesoCatalogo(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the EventoCatalogoProcesoCatalogo
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.EventoCatalogoProcesoCatalogo	true		"body for EventoCatalogoProcesoCatalogo content"
// @Success 200 {object} models.EventoCatalogoProcesoCatalogo
// @Failure 400 the request contains incorrect syntax
// @router /:id [put]
func (c *EventoCatalogoProcesoCatalogoController) Put() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	v := models.EventoCatalogoProcesoCatalogo{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		v.FechaCreacion = fechaCorreccion(v.FechaCreacion)
		v.FechaModificacion = fechaActual()
		if err := models.UpdateEventoCatalogoProcesoCatalogoById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the EventoCatalogoProcesoCatalogo
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 404 not found resource
// @router /:id [delete]
func (c *EventoCatalogoProcesoCatalogoController) Delete() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err := models.DeleteEventoCatalogoProcesoCatalogo(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
