package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type AuditoriaEventos struct {
	Id                int       `orm:"column(id);pk;auto" json:"Id"`
	EntidadTipo       string    `orm:"column(entidad_tipo)" json:"EntidadTipo"`
	Operacion         string    `orm:"column(operacion)" json:"Operacion"`
	ValorAnterior     *string   `orm:"column(valor_anterior);type(jsonb);null" json:"ValorAnterior,omitempty"`
	ValorNuevo        *string   `orm:"column(valor_nuevo);type(jsonb);null" json:"ValorNuevo,omitempty"`
	TerceroId         *int      `orm:"column(tercero_id);null" json:"TerceroId,omitempty"`
	Endpoint          *string   `orm:"column(endpoint);null" json:"Endpoint,omitempty"`
	Activo            bool      `orm:"column(activo)" json:"Activo"`
	FechaCreacion     time.Time `orm:"column(fecha_creacion);type(timestamp without time zone)" json:"FechaCreacion"`
	FechaModificacion time.Time `orm:"column(fecha_modificacion);type(timestamp without time zone)" json:"FechaModificacion"`
}

func (t *AuditoriaEventos) TableName() string {
	return "auditoria_eventos"
}

func init() {
	orm.RegisterModel(new(AuditoriaEventos))
}

func AddAuditoriaEventos(m *AuditoriaEventos) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}
