package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaTipoEvento_20191217_160342 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaTipoEvento_20191217_160342{}
	m.Created = "20191217_160342"

	migration.Register("CrearTablaTipoEvento_20191217_160342", m)
}

// Run the migrations
func (m *CrearTablaTipoEvento_20191217_160342) Up() {
	m.SQL("CREATE TABLE IF NOT EXISTS eventos.tipo_evento (id serial NOT NULL,nombre varchar(50) NOT NULL,descripcion varchar(250),codigo_abreviacion varchar(20),activo boolean NOT NULL,dependencia_id integer NOT NULL,fecha_creacion timestamp NOT NULL,fecha_modificacion timestamp NOT NULL,tipo_recurrencia_id integer,CONSTRAINT pk_tipo_evento PRIMARY KEY (id));")
	m.SQL("ALTER TABLE eventos.tipo_evento OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE eventos.tipo_evento IS 'Tabla que permite registrar los diferentes tipos de sesion que tienen las diferentes organizaciones';")
	m.SQL("COMMENT ON COLUMN eventos.tipo_evento.id IS 'Identificador de la tabla tipo sesion';")
	m.SQL("COMMENT ON COLUMN eventos.tipo_evento.nombre IS 'Nombre del tipo de sesion';")
	m.SQL("COMMENT ON COLUMN eventos.tipo_evento.descripcion IS 'Descripcion del tipo de sesion que se referencia';")
	m.SQL("COMMENT ON COLUMN eventos.tipo_evento.codigo_abreviacion IS 'Codigo de abreviación del tipo de sesión que se maneja';")
	m.SQL("COMMENT ON COLUMN eventos.tipo_evento.activo IS 'Booleano que define si el tipo de sesión se encuentra actualmente activo o no';")
	m.SQL("COMMENT ON COLUMN eventos.tipo_evento.dependencia_id IS 'se referencia al esquema de dependencia del sistema OKIOS';")
	m.SQL("COMMENT ON COLUMN eventos.tipo_evento.fecha_creacion IS 'Fecha de creacion de un tipo de evento';")
	m.SQL("COMMENT ON COLUMN eventos.tipo_evento.fecha_modificacion IS 'Fecha de modifiación de un tipo de evento';")
	m.SQL("COMMENT ON CONSTRAINT pk_tipo_evento ON eventos.tipo_evento  IS 'Contrait para pk de la tabla';")

}

// Reverse the migrations
func (m *CrearTablaTipoEvento_20191217_160342) Down() {
	m.SQL("DROP TABLE IF EXISTS eventos.tipo_evento")

}
