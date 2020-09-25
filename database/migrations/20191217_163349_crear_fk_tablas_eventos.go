package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearFkTablasEventos_20191217_163349 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearFkTablasEventos_20191217_163349{}
	m.Created = "20191217_163349"

	migration.Register("CrearFkTablasEventos_20191217_163349", m)
}

// Run the migrations
func (m *CrearFkTablasEventos_20191217_163349) Up() {
	m.SQL("ALTER TABLE eventos.encargado_evento ADD CONSTRAINT fk_encargado_evento_rol_encargado_evento FOREIGN KEY (rol_encargado_evento_id) REFERENCES eventos.rol_encargado_evento (id) MATCH FULL ON DELETE SET NULL ON UPDATE CASCADE;")
	m.SQL("ALTER TABLE eventos.tipo_evento ADD CONSTRAINT fk_tipo_evento_tipo_recurrencia FOREIGN KEY (tipo_recurrencia_id) REFERENCES eventos.tipo_recurrencia (id) MATCH FULL ON DELETE SET NULL ON UPDATE CASCADE;")
	m.SQL("ALTER TABLE eventos.tipo_evento ADD CONSTRAINT tipo_evento_uq UNIQUE (tipo_recurrencia_id);")
	m.SQL("ALTER TABLE eventos.calendario_evento ADD CONSTRAINT fk_calendario_evento_tipo_evento FOREIGN KEY (tipo_evento_id) REFERENCES eventos.tipo_evento (id) MATCH FULL ON DELETE SET NULL ON UPDATE CASCADE;")
	m.SQL("ALTER TABLE eventos.encargado_evento ADD CONSTRAINT fk_encargado_evento_calendario_evento FOREIGN KEY (calendario_evento_id) REFERENCES eventos.calendario_evento (id) MATCH FULL ON DELETE SET NULL ON UPDATE CASCADE;")
	m.SQL("ALTER TABLE eventos.tipo_publico ADD CONSTRAINT fk_tipo_publico_calendario_evento FOREIGN KEY (calendario_evento_id) REFERENCES eventos.calendario_evento (id) MATCH FULL ON DELETE SET NULL ON UPDATE CASCADE;")
	m.SQL("ALTER TABLE eventos.calendario_evento ADD CONSTRAINT fk_evento_padre FOREIGN KEY (evento_padre_id) REFERENCES eventos.calendario_evento (id) MATCH FULL ON DELETE NO ACTION ON UPDATE NO ACTION;")

}

// Reverse the migrations
func (m *CrearFkTablasEventos_20191217_163349) Down() {
	m.SQL("ALTER TABLE eventos.encargado_evento DROP CONSTRAINT IF EXISTS fk_encargado_evento_rol_encargado_evento CASCADE;")
	m.SQL("ALTER TABLE eventos.tipo_evento DROP CONSTRAINT IF EXISTS fk_tipo_evento_tipo_recurrencia CASCADE;")
	m.SQL("ALTER TABLE eventos.tipo_evento DROP CONSTRAINT IF EXISTS tipo_evento_uq CASCADE;")
	m.SQL("ALTER TABLE eventos.calendario_evento DROP CONSTRAINT IF EXISTS fk_calendario_evento_tipo_evento CASCADE;")
 	m.SQL("ALTER TABLE eventos.encargado_evento DROP CONSTRAINT IF EXISTS fk_encargado_evento_calendario_evento CASCADE;")
	m.SQL("ALTER TABLE eventos.tipo_publico DROP CONSTRAINT IF EXISTS fk_tipo_publico_calendario_evento CASCADE;")
	m.SQL("ALTER TABLE eventos.calendario_evento DROP CONSTRAINT IF EXISTS fk_evento_padre CASCADE;")
}
