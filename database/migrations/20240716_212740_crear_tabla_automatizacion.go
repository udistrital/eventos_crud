package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaAutomatizacion_20240716_212740 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaAutomatizacion_20240716_212740{}
	m.Created = "20240716_212740"

	migration.Register("CrearTablaAutomatizacion_20240716_212740", m)
}

// Run the migrations
func (m *CrearTablaAutomatizacion_20240716_212740) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

	m.SQL(`
		CREATE TABLE IF NOT EXISTS eventos.automatizacion (
			id serial NOT NULL,
			ejecucion_unica boolean NOT NULL,
			endpoint json NOT NULL,
			activo boolean NOT NULL,
			fecha_creacion timestamp NOT NULL DEFAULT now(),
			fecha_modificacion timestamp NOT NULL DEFAULT now(),
			calendario_evento_id integer NOT NULL,
			CONSTRAINT pk_automatizacion PRIMARY KEY (id),
			CONSTRAINT fk_automatizacion_calendario_evento FOREIGN KEY (calendario_evento_id)
			REFERENCES eventos.calendario_evento (id) ON DELETE CASCADE
		);
	`)

}

// Reverse the migrations
func (m *CrearTablaAutomatizacion_20240716_212740) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

	m.SQL("DROP TABLE IF EXISTS eventos.automatizacion")
}
