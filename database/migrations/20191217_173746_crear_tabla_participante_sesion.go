package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaParticipanteSesion_20191217_173746 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaParticipanteSesion_20191217_173746{}
	m.Created = "20191217_173746"

	migration.Register("CrearTablaParticipanteSesion_20191217_173746", m)
}

// Run the migrations
func (m *CrearTablaParticipanteSesion_20191217_173746) Up() {
	m.SQL("CREATE TABLE eventos.participante_sesion(id serial NOT NULL,sesion integer NOT NULL,rol_participante_sesion integer NOT NULL,ente integer NOT NULL,CONSTRAINT pk_participante_sesion PRIMARY KEY (id));") 
	m.SQL("ALTER TABLE eventos.participante_sesion OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE eventos.participante_sesion IS 'Tabla de rompimiento entre los difenrentes participantes y la sesión';")
	m.SQL("COMMENT ON COLUMN eventos.participante_sesion.id IS 'Identificador de la tabla responsable_sesion';")
	m.SQL("COMMENT ON COLUMN eventos.participante_sesion.ente IS 'Columna que referencia el Id de la tabla donde se encuentra el ente.';")
	m.SQL("COMMENT ON CONSTRAINT pk_participante_sesion ON eventos.participante_sesion  IS 'Identificador de la tabla responsable sesión';")

}

// Reverse the migrations
func (m *CrearTablaParticipanteSesion_20191217_173746) Down() {
	m.SQL("DROP TABLE IF EXISTS eventos.participante_sesion")

}
