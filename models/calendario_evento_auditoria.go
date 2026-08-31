package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type CalendarioEventoAuditoria struct {
	Id                 int               `orm:"column(id);pk;auto" json:"Id"`
	CalendarioEventoId *CalendarioEvento `orm:"column(calendario_evento_id);rel(fk)" json:"CalendarioEventoId"`
	Operacion          string            `orm:"column(operacion)" json:"Operacion"`
	TerceroId          int               `orm:"column(tercero_id)" json:"TerceroId"`
	FechaCreacion      time.Time         `orm:"column(fecha_creacion);type(timestamp without time zone)" json:"FechaCreacion"`
	FechaModificacion  time.Time         `orm:"column(fecha_modificacion);type(timestamp without time zone)" json:"FechaModificacion"`
	Activo             bool              `orm:"column(activo)" json:"Activo"`
	EstadoAnterior     *string           `orm:"column(estado_anterior);type(jsonb);null" json:"EstadoAnterior,omitempty"`
	EstadoNuevo        *string           `orm:"column(estado_nuevo);type(jsonb)" json:"EstadoNuevo"`
	Cambios            *string           `orm:"column(cambios);type(jsonb)" json:"Cambios"`
}

type CalendarioEventoRequest struct {
	FechaInicio      time.Time       `json:"FechaInicio"`
	FechaFin         time.Time       `json:"FechaFin"`
	Activo           bool            `json:"Activo"`
	DependenciaId    string          `json:"DependenciaId"`
	ProcesoId        *Proceso        `json:"ProcesoId"`
	EventoCatalogoId *EventoCatalogo `json:"EventoCatalogoId"`
	UbicacionId      int             `json:"UbicacionId"`
	PosterUrl        string          `json:"PosterUrl"`
	TerceroId        int             `json:"TerceroId"`
}

type TerceroRequest struct {
	TerceroId int `json:"TerceroId"`
}

func (t *CalendarioEventoAuditoria) TableName() string {
	return "calendario_evento_auditoria"
}

func init() {
	orm.RegisterModel(new(CalendarioEventoAuditoria))
}

func AddCalendarioEventoAuditoria(m *CalendarioEventoAuditoria) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func GetCalendarioEventoAuditoriaById(id int) (v *CalendarioEventoAuditoria, err error) {
	o := orm.NewOrm()
	v = &CalendarioEventoAuditoria{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetAllCalendarioEventoAuditoria(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CalendarioEventoAuditoria)).RelatedSel()
	for k, v := range query {
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, v == "true" || v == "1")
		} else {
			qs = qs.Filter(k, v)
		}
	}

	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(order) == 1 {
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else if len(order) != 0 {
		return nil, errors.New("Error: unused 'order' fields")
	}

	var l []CalendarioEventoAuditoria
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

func UpdateCalendarioEventoAuditoriaById(m *CalendarioEventoAuditoria) (err error) {
	o := orm.NewOrm()
	v := CalendarioEventoAuditoria{Id: m.Id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

func DeleteCalendarioEventoAuditoria(id int) (err error) {
	o := orm.NewOrm()
	v := CalendarioEventoAuditoria{Id: id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CalendarioEventoAuditoria{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
