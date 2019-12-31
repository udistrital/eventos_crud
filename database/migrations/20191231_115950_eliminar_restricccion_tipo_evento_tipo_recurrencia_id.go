package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type EliminarRestricccionTipoEventoTipoRecurrenciaId_20191231_115950 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &EliminarRestricccionTipoEventoTipoRecurrenciaId_20191231_115950{}
	m.Created = "20191231_115950"

	migration.Register("EliminarRestricccionTipoEventoTipoRecurrenciaId_20191231_115950", m)
}

// Run the migrations
func (m *EliminarRestricccionTipoEventoTipoRecurrenciaId_20191231_115950) Up() {
	m.SQL("ALTER TABLE eventos.tipo_evento DROP CONSTRAINT IF EXISTS tipo_evento_uq;") 

}

// Reverse the migrations
func (m *EliminarRestricccionTipoEventoTipoRecurrenciaId_20191231_115950) Down() {
	m.SQL("ALTER TABLE eventos.tipo_evento ADD CONSTRAINT tipo_evento_uq UNIQUE (tipo_recurrencia_id);") 

}
