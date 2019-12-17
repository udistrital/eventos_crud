package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaTipoRecurrencia_20191217_162110 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaTipoRecurrencia_20191217_162110{}
	m.Created = "20191217_162110"

	migration.Register("CrearTablaTipoRecurrencia_20191217_162110", m)
}

// Run the migrations
func (m *CrearTablaTipoRecurrencia_20191217_162110) Up() {
	m.SQL("CREATE TABLE eventos.tipo_recurrencia (id serial NOT NULL,nombre varchar(50) NOT NULL,descripcion varchar(250),codigo_abreviacion varchar(20),activo boolean NOT NULL,fecha_creacion timestamp NOT NULL,fecha_modificacion timestamp NOT NULL,CONSTRAINT pk_tipo_recurrencia PRIMARY KEY (id));")
	m.SQL("ALTER TABLE eventos.tipo_recurrencia OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE eventos.tipo_recurrencia IS 'Tabla que almacena los tipos de recurrencia en tiempo que se tienen (por minuto, hora, diario...)';")
	m.SQL("COMMENT ON COLUMN eventos.tipo_recurrencia.id IS 'Identificador de la tabla tipo_recurrencia';")
	m.SQL("COMMENT ON COLUMN eventos.tipo_recurrencia.nombre IS 'Nombre del tipo de recurrencia que se registra en la tabla';")
	m.SQL("COMMENT ON COLUMN eventos.tipo_recurrencia.descripcion IS 'Breve descripcion del tipo de recurrencia que se registra';")
	m.SQL("COMMENT ON COLUMN eventos.tipo_recurrencia.codigo_abreviacion IS 'Codigo de abreviacion del tipo de recurrencia';")
	m.SQL("COMMENT ON COLUMN eventos.tipo_recurrencia.activo IS 'Booleano que permite identificar si el tipo de recurrencia se encuentra activo o no';")
	m.SQL("COMMENT ON COLUMN eventos.tipo_recurrencia.fecha_creacion IS 'Fecha de creacion de un tipo de recurencia';")
	m.SQL("COMMENT ON COLUMN eventos.tipo_recurrencia.fecha_modificacion IS 'Fecha de modifiación de un tipo de recurrencia';")

}

// Reverse the migrations
func (m *CrearTablaTipoRecurrencia_20191217_162110) Down() {
	m.SQL("DROP TABLE IF EXISTS eventos.tipo_recurrencia")


}
