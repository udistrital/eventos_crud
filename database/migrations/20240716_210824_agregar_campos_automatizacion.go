package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AgregarCamposAutomatizacion_20240716_210824 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AgregarCamposAutomatizacion_20240716_210824{}
	m.Created = "20240716_210824"

	migration.Register("AgregarCamposAutomatizacion_20240716_210824", m)
}

// Run the migrations
func (m *AgregarCamposAutomatizacion_20240716_210824) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

	m.SQL("ALTER TABLE eventos.calendario_evento ADD COLUMN automatizacion boolean NOT NULL;")
	m.SQL("COMMENT ON COLUMN eventos.calendario_evento.automatizacion IS 'Se usa para saber si el evento es automatizable o no';")
}

// Reverse the migrations
func (m *AgregarCamposAutomatizacion_20240716_210824) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

	m.SQL("ALTER TABLE eventos.calendario_evento DROP COLUMN automatizacion;")

}
