package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaRelacionSesiones_20191217_175027 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaRelacionSesiones_20191217_175027{}
	m.Created = "20191217_175027"

	migration.Register("CrearTablaRelacionSesiones_20191217_175027", m)
}

// Run the migrations
func (m *CrearTablaRelacionSesiones_20191217_175027) Up() {
	m.SQL("CREATE TABLE eventos.relacion_sesiones(id serial NOT NULL,sesion_padre integer NOT NULL,sesion_hijo integer NOT NULL,CONSTRAINT pk_relacion_sesiones PRIMARY KEY (id));")
	m.SQL("ALTER TABLE eventos.relacion_sesiones OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE eventos.relacion_sesiones IS 'Tabla que almacena la relación entre las sesión padre y la sesión hijo.';")
	m.SQL("COMMENT ON COLUMN eventos.relacion_sesiones.id IS 'Identificador de la tabla relacion_sesiones';")
	m.SQL("COMMENT ON COLUMN eventos.relacion_sesiones.sesion_hijo IS 'Hijo de la relacion de sesiones';")


}

// Reverse the migrations
func (m *CrearTablaRelacionSesiones_20191217_175027) Down() {
	m.SQL("DROP TABLE IF EXISTS eventos.relacion_sesiones")


}
