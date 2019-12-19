package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AlterColumnaAplicacionId_20191219_111609 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AlterColumnaAplicacionId_20191219_111609{}
	m.Created = "20191219_111609"

	migration.Register("AlterColumnaAplicacionId_20191219_111609", m)
}

// Run the migrations
func (m *AlterColumnaAplicacionId_20191219_111609) Up() {
	m.SQL("ALTER TABLE eventos.calendario_evento ADD aplicacion_id integer NOT NULL;")
	m.SQL("COMMENT ON COLUMN eventos.calendario_evento.aplicacion_id IS 'Campo obligatorio para relacionar el evento creado a una aplicación. Es un id relacionado a la tabla aplicacion en el esquema configuracion';")

}

// Reverse the migrations
func (m *AlterColumnaAplicacionId_20191219_111609) Down() {
	m.SQL("ALTER TABLE eventos.calendario_evento DROP COLUMN aplicacion_id;")

}
