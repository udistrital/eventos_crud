package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaRolEncargadoEvento_20191217_161839 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaRolEncargadoEvento_20191217_161839{}
	m.Created = "20191217_161839"

	migration.Register("CrearTablaRolEncargadoEvento_20191217_161839", m)
}

// Run the migrations
func (m *CrearTablaRolEncargadoEvento_20191217_161839) Up() {
	m.SQL("CREATE TABLE eventos.rol_encargado_evento (id serial NOT NULL,nombre varchar(50) NOT NULL,descripcion varchar(250),codigo_abreviacion varchar(20),activo boolean NOT NULL,numero_de_orden numeric(5,2),fecha_creacion timestamp NOT NULL,fecha_modificacion timestamp NOT NULL,CONSTRAINT pk_rol_encargado_evento PRIMARY KEY (id));")
	m.SQL("ALTER TABLE eventos.rol_encargado_evento OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE eventos.rol_encargado_evento IS 'Permite almacenar los diferentes roles que puede tener un encargado en un evento';")
	m.SQL("COMMENT ON COLUMN eventos.rol_encargado_evento.id IS 'Identificador de la tabla rol_participante_sesion';")
	m.SQL("COMMENT ON COLUMN eventos.rol_encargado_evento.nombre IS 'Nombre del rol que puede tener un participante de una sesion';")
	m.SQL("COMMENT ON COLUMN eventos.rol_encargado_evento.descripcion IS 'Descripcion que permite especificar el rol y las funciones en la sesion';")
	m.SQL("COMMENT ON COLUMN eventos.rol_encargado_evento.codigo_abreviacion IS 'Codigo de abreviacion del nombre del rol que cumple un participante en la sesion';")
	m.SQL("COMMENT ON COLUMN eventos.rol_encargado_evento.activo IS 'Booleano que permite identificar si el rol se encuentra activo o no';")
	m.SQL("COMMENT ON COLUMN eventos.rol_encargado_evento.fecha_creacion IS 'Fecha de creacion de un rol participante del evento';")
	m.SQL("COMMENT ON COLUMN eventos.rol_encargado_evento.fecha_modificacion IS 'Fecha de modifiación de un rol participante del evento';")
}

// Reverse the migrations
func (m *CrearTablaRolEncargadoEvento_20191217_161839) Down() {
	m.SQL("DROP TABLE IF EXISTS eventos.rol_encargado_evento")

}
