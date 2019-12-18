package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaCalendarioEvento_20191217_161004 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaCalendarioEvento_20191217_161004{}
	m.Created = "20191217_161004"

	migration.Register("CrearTablaCalendarioEvento_20191217_161004", m)
}

// Run the migrations
func (m *CrearTablaCalendarioEvento_20191217_161004) Up() {
	m.SQL("CREATE TABLE eventos.calendario_evento (id serial NOT NULL,descripcion varchar(250),fecha_creacion timestamp NOT NULL,fecha_modificacion timestamp,fecha_inicio timestamp NOT NULL,fecha_fin timestamp,periodo_id integer NOT NULL,activo boolean NOT NULL,evento_padre_id integer,tipo_evento_id integer,CONSTRAINT pk_calendario_evento PRIMARY KEY (id));")
	m.SQL("ALTER TABLE eventos.calendario_evento OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE eventos.calendario_evento IS 'Tabla que almacena cada uno de las sesiones y sus datos';")
	m.SQL("COMMENT ON COLUMN eventos.calendario_evento.id IS 'Identificador de la sesión';")
	m.SQL("COMMENT ON COLUMN eventos.calendario_evento.descripcion IS 'Descripción de la sesión qeu se registra';")
	m.SQL("COMMENT ON COLUMN eventos.calendario_evento.fecha_creacion IS 'Fecha de creacion de una sesión para la organización';")
	m.SQL("COMMENT ON COLUMN eventos.calendario_evento.fecha_modificacion IS 'Fecha de modifiación del evento para la organización';")
	m.SQL("COMMENT ON COLUMN eventos.calendario_evento.fecha_inicio IS 'Fecha de inicio de la sesión para la organización';")
	m.SQL("COMMENT ON COLUMN eventos.calendario_evento.fecha_fin IS 'Fecha de finalización de la sesión para la organización';")
	m.SQL("COMMENT ON COLUMN eventos.calendario_evento.periodo_id IS 'Se almacena el periodo al cual pertenece el evento';")
	m.SQL("COMMENT ON COLUMN eventos.calendario_evento.activo IS 'Campo para registrar si el  registro se encuentra activo o no, solo a nivel de registro.';")
	m.SQL("COMMENT ON CONSTRAINT pk_calendario_evento ON eventos.calendario_evento  IS 'Restriccion pk de la tabla calendario evento';")

}

// Reverse the migrations
func (m *CrearTablaCalendarioEvento_20191217_161004) Down() {
	m.SQL("DROP TABLE IF EXISTS eventos.calendario_evento")

}
