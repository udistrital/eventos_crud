package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaEncargadoEvento_20191217_161358 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaEncargadoEvento_20191217_161358{}
	m.Created = "20191217_161358"

	migration.Register("CrearTablaEncargadoEvento_20191217_161358", m)
}

// Run the migrations
func (m *CrearTablaEncargadoEvento_20191217_161358) Up() {
	m.SQL("CREATE TABLE eventos.encargado_evento (id serial NOT NULL,encargado_id integer NOT NULL,fecha_creacion timestamp NOT NULL,fecha_modificacion timestamp NOT NULL,activo boolean NOT NULL,rol_encargado_evento_id integer,calendario_evento_id integer,CONSTRAINT pk_encargado_evento PRIMARY KEY (id));")
	m.SQL("ALTER TABLE eventos.encargado_evento OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE eventos.encargado_evento IS 'Tabla de rompimiento entre los diferentes encargados y el evento';")
	m.SQL("COMMENT ON COLUMN eventos.encargado_eventos.id IS 'Identificador de la tabla responsable_sesion';")
	m.SQL("COMMENT ON COLUMN eventos.encargado_eventos.encargado_id IS 'Se hace referencia al esquema de persona';")
	m.SQL("COMMENT ON COLUMN eventos.encargado_eventos.fecha_creacion IS 'Fecha de creacion de un participante del evento';")
	m.SQL("COMMENT ON COLUMN eventos.encargado_eventos.fecha_modificacion IS 'Fecha de modifiación de un participante del evento';")
	m.SQL("COMMENT ON COLUMN eventos.encargado_eventos.activo IS 'Campo para registrar si el  registro se encuentra activo o no, solo a nivel de registro.';")
	m.SQL("COMMENT ON CONSTRAINT pk_encargado_evento ON eventos.encargado_evento  IS 'Identificador de la tabla responsable evento';")
	

}

// Reverse the migrations
func (m *CrearTablaEncargadoEvento_20191217_161358) Down() {
	m.SQL("DROP TABLE IF EXISTS eventos.encargado_evento")

}
