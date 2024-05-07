package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CambioTablaCalendarioNuevaColumnaMultipleperiodoid_20240507_091020 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CambioTablaCalendarioNuevaColumnaMultipleperiodoid_20240507_091020{}
	m.Created = "20240507_091020"

	migration.Register("CambioTablaCalendarioNuevaColumnaMultipleperiodoid_20240507_091020", m)
}

// Run the migrations
func (m *CambioTablaCalendarioNuevaColumnaMultipleperiodoid_20240507_091020) Up() {
	m.SQL(`ALTER TABLE IF EXISTS eventos.calendario
    		ADD COLUMN multiple_periodo_id json;`)
}

// Reverse the migrations
func (m *CambioTablaCalendarioNuevaColumnaMultipleperiodoid_20240507_091020) Down() {
	m.SQL(`ALTER TABLE IF EXISTS eventos.calendario DROP COLUMN IF EXISTS multiple_periodo_id;
			select * from eventos.calendario;`)
}
