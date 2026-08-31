package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type RolGestionEvento struct {
	Id                int       `orm:"column(id);pk;auto"`
	Nombre            string    `orm:"column(nombre)"`
	Descripcion       string    `orm:"column(descripcion);null"`
	CodigoAbreviacion string    `orm:"column(codigo_abreviacion);null"`
	Activo            bool      `orm:"column(activo)"`
	NumeroDeOrden     float64   `orm:"column(numero_de_orden);null"`
	FechaCreacion     time.Time `orm:"column(fecha_creacion);type(timestamp without time zone)"`
	FechaModificacion time.Time `orm:"column(fecha_modificacion);type(timestamp without time zone)"`
}

func (t *RolGestionEvento) TableName() string { return "rol_gestion_evento" }

func init() { orm.RegisterModel(new(RolGestionEvento)) }

func AddRolGestionEvento(m *RolGestionEvento) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func GetRolGestionEventoById(id int) (v *RolGestionEvento, err error) {
	o := orm.NewOrm()
	v = &RolGestionEvento{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetAllRolGestionEvento(query map[string]string, fields []string, sortby []string, order []string, offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(RolGestionEvento))
	for k, v := range query {
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			for i, v := range sortby {
				if order[i] == "desc" {
					sortFields = append(sortFields, "-"+v)
				} else if order[i] == "asc" {
					sortFields = append(sortFields, v)
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
			}
		} else if len(order) == 1 {
			for _, v := range sortby {
				if order[0] == "desc" {
					sortFields = append(sortFields, "-"+v)
				} else if order[0] == "asc" {
					sortFields = append(sortFields, v)
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
			}
		} else {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else if len(order) != 0 {
		return nil, errors.New("Error: unused 'order' fields")
	}

	var l []RolGestionEvento
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

func UpdateRolGestionEventoById(m *RolGestionEvento) (err error) {
	o := orm.NewOrm()
	v := RolGestionEvento{Id: m.Id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

func DeleteRolGestionEvento(id int) (err error) {
	o := orm.NewOrm()
	v := RolGestionEvento{Id: id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&RolGestionEvento{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
