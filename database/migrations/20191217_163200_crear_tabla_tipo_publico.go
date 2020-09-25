package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaTipoPublico_20191217_163200 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaTipoPublico_20191217_163200{}
	m.Created = "20191217_163200"

	migration.Register("CrearTablaTipoPublico_20191217_163200", m)
}

// Run the migrations
func (m *CrearTablaTipoPublico_20191217_163200) Up() {
	m.SQL("CREATE TABLE eventos.tipo_publico (id serial NOT NULL,nombre varchar(50) NOT NULL,codigo_abreviacion varchar(250),activo boolean NOT NULL,numero_orden numeric(5,2),fecha_creacion timestamp NOT NULL,fecha_modificacion timestamp NOT NULL,calendario_evento_id integer,CONSTRAINT pk_tipo_publico PRIMARY KEY (id));")
	m.SQL("ALTER TABLE eventos.tipo_publico OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE eventos.tipo_publico IS 'Tabla para almacenar el tipo de publico que puede participar en un evento ';")
}

// Reverse the migrations
func (m *CrearTablaTipoPublico_20191217_163200) Down() {
	m.SQL("DROP TABLE IF EXISTS eventos.tipo_publico")

}
