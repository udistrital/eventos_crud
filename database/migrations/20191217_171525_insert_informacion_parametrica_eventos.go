package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertInformacionParametricaEventos_20191217_171525 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertInformacionParametricaEventos_20191217_171525{}
	m.Created = "20191217_171525"

	migration.Register("InsertInformacionParametricaEventos_20191217_171525", m)
}

// Run the migrations
func (m *InsertInformacionParametricaEventos_20191217_171525) Up() {
	m.SQL("INSERT INTO EVENTOS.ROL_ENCARGADO_EVENTO (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion, numero_de_orden) VALUES (1,'Creador','Creador', 'CRE', true, now(), now(), 1);")
	m.SQL("INSERT INTO EVENTOS.ROL_ENCARGADO_EVENTO (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion, numero_de_orden) VALUES (2,'Actualizador','Actualizador', 'act', true, now(), now(), 2);")
	m.SQL("INSERT INTO EVENTOS.TIPO_RECURRENCIA (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (1,'Anual','Anual', 'ANO', true, now(), now());")
	m.SQL("INSERT INTO EVENTOS.TIPO_RECURRENCIA (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (2,'Semestral','Semestral', 'SEM', true, now(), now());")
	m.SQL("INSERT INTO EVENTOS.TIPO_EVENTO (id, nombre, descripcion, codigo_abreviacion, activo, dependencia_id, fecha_creacion, fecha_modificacion, tipo_recurrencia_id) VALUES (1,'Inscripcion','Inscripcion de espacios', 'INS', true, 122, now(), now(), 1);")
	m.SQL("INSERT INTO EVENTOS.TIPO_EVENTO (id, nombre, descripcion, codigo_abreviacion, activo, dependencia_id, fecha_creacion, fecha_modificacion, tipo_recurrencia_id) VALUES (2,'Grado','Grado', 'GRD', true, 122, now(), now(), 2);")

}

// Reverse the migrations
func (m *InsertInformacionParametricaEventos_20191217_171525) Down() {
	m.SQL("DELETE FROM EVENTOS.ROL_ENCARGADO_EVENTO")
	m.SQL("DELETE FROM EVENTOS.TIPO_RECURRENCIA")
	m.SQL("DELETE FROM EVENTOS.TIPO_EVENTO")
}
