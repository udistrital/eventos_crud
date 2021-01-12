package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearEsquemaEventos_20191217_160109 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearEsquemaEventos_20191217_160109{}
	m.Created = "20191217_160109"

	migration.Register("CrearEsquemaEventos_20191217_160109", m)
}

// Run the migrations
func (m *CrearEsquemaEventos_20191217_160109) Up() {
	m.SQL("CREATE SCHEMA IF NOT EXISTS eventos;")
	m.SQL("ALTER SCHEMA eventos OWNER TO desarrollooas;")
	m.SQL("SET search_path TO pg_catalog,public,eventos;")

}

// Reverse the migrations
func (m *CrearEsquemaEventos_20191217_160109) Down() {
	m.SQL("DROP SCHEMA IF EXISTS eventos");

}
