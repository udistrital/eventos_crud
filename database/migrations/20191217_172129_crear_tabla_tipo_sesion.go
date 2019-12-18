package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaTipoSesion_20191217_172129 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaTipoSesion_20191217_172129{}
	m.Created = "20191217_172129"

	migration.Register("CrearTablaTipoSesion_20191217_172129", m)
}

// Run the migrations
func (m *CrearTablaTipoSesion_20191217_172129) Up() {
	m.SQL("CREATE TABLE eventos.tipo_sesion(id serial NOT NULL,nombre varchar(50) NOT NULL,descripcion varchar(250),codigo_abreviacion varchar(20),activo boolean NOT NULL,numero_orden numeric(5,2),CONSTRAINT pk_tipo_sesion PRIMARY KEY (id),CONSTRAINT uq_nombre_tipo_sesion UNIQUE (nombre));")
	m.SQL("ALTER TABLE eventos.tipo_sesion OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE eventos.tipo_sesion IS 'Tabla que permite registrar los diferentes tipos de sesion que tienen las diferentes organizaciones';")
	m.SQL("COMMENT ON COLUMN eventos.tipo_sesion.id IS 'Identificador de la tabla tipo sesion';")
	m.SQL("COMMENT ON COLUMN eventos.tipo_sesion.nombre IS 'Nombre del tipo de sesion';")
	m.SQL("COMMENT ON COLUMN eventos.tipo_sesion.descripcion IS 'Descripcion del tipo de sesion que se referencia';")
	m.SQL("COMMENT ON COLUMN eventos.tipo_sesion.codigo_abreviacion IS 'Codigo de abreviación del tipo de sesión que se maneja';")
	m.SQL("COMMENT ON COLUMN eventos.tipo_sesion.activo IS 'Booleano que define si el tipo de sesión se encuentra actualmente activo o no';")
	m.SQL("COMMENT ON COLUMN eventos.tipo_sesion.numero_orden IS 'Número de orden para la tabla tipo_sesion';")
	m.SQL("COMMENT ON CONSTRAINT pk_tipo_sesion ON eventos.tipo_sesion  IS 'Contrait para pk de la tabla';")
	m.SQL("COMMENT ON CONSTRAINT uq_nombre_tipo_sesion ON eventos.tipo_sesion  IS 'UK para garantizar unicidad del nombre del tipo de la sesión';")

}

// Reverse the migrations
func (m *CrearTablaTipoSesion_20191217_172129) Down() {
	m.SQL("DROP TABLE IF EXISTS eventos.tipo_sesion")
}
