package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaSesion_20191217_172447 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaSesion_20191217_172447{}
	m.Created = "20191217_172447"

	migration.Register("CrearTablaSesion_20191217_172447", m)
}

// Run the migrations
func (m *CrearTablaSesion_20191217_172447) Up() {
	m.SQL("CREATE TABLE eventos.sesion(id serial NOT NULL,descripcion varchar(250),fecha_creacion timestamp NOT NULL,fecha_modificacion timestamp,fecha_inicio timestamp NOT NULL,fecha_fin timestamp,periodo integer,recurrente boolean NOT NULL,numero_recurrencias integer,tipo_sesion integer NOT NULL,CONSTRAINT pk_sesion PRIMARY KEY (id));")
	m.SQL("ALTER TABLE eventos.sesion OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE eventos.sesion IS 'Tabla que almacena cada uno de las sesiones y sus datos';")
	m.SQL("COMMENT ON COLUMN eventos.sesion.id IS 'Identificador de la sesión';")
	m.SQL("COMMENT ON COLUMN eventos.sesion.descripcion IS 'Descripción de la sesión qeu se registra';")
	m.SQL("COMMENT ON COLUMN eventos.sesion.fecha_creacion IS 'Fecha de creacion de una sesión para la organización';")
	m.SQL("COMMENT ON COLUMN eventos.sesion.fecha_modificacion IS 'Fecha de modifiación del evento para la organización';")
	m.SQL("COMMENT ON COLUMN eventos.sesion.fecha_inicio IS 'Fecha de inicio de la sesión para la organización';")
	m.SQL("COMMENT ON COLUMN eventos.sesion.fecha_fin IS 'Fecha de finalización de la sesión para la organización';")
	m.SQL("COMMENT ON COLUMN eventos.sesion.periodo IS 'Se almacena el periodo al cual pertenece el evento';")
	m.SQL("COMMENT ON COLUMN eventos.sesion.recurrente IS 'Booleano qeu permite identificar si la sesion es recurrente o no';")
	m.SQL("COMMENT ON COLUMN eventos.sesion.numero_recurrencias IS 'Numero de veces que va a ser necesarioqeu se repita la sesión';")
	m.SQL("COMMENT ON CONSTRAINT pk_sesion ON eventos.sesion  IS 'Restriccion pk de la tabla sesión';")
}

// Reverse the migrations
func (m *CrearTablaSesion_20191217_172447) Down() {
	m.SQL("DROP TABLE IF EXISTS eventos.sesion")


}
