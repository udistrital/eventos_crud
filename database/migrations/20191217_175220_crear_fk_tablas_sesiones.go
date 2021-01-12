package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearFkTablasSesiones_20191217_175220 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearFkTablasSesiones_20191217_175220{}
	m.Created = "20191217_175220"

	migration.Register("CrearFkTablasSesiones_20191217_175220", m)
}

// Run the migrations
func (m *CrearFkTablasSesiones_20191217_175220) Up() {
	m.SQL("ALTER TABLE eventos.sesion ADD CONSTRAINT fk_sesion_tipo_sesion FOREIGN KEY (tipo_sesion) REFERENCES eventos.tipo_sesion (id) MATCH FULL ON DELETE RESTRICT ON UPDATE CASCADE;")
	m.SQL("ALTER TABLE eventos.participante_sesion ADD CONSTRAINT fk_participante_sesion_sesion FOREIGN KEY (sesion) REFERENCES eventos.sesion (id) MATCH FULL ON DELETE RESTRICT ON UPDATE CASCADE;")
	m.SQL("ALTER TABLE eventos.participante_sesion ADD CONSTRAINT fk_participante_sesion_rol_participante_sesion FOREIGN KEY (rol_participante_sesion)REFERENCES eventos.rol_participante_sesion (id) MATCH FULL 	ON DELETE RESTRICT ON UPDATE CASCADE;")
	m.SQL("ALTER TABLE eventos.sesion_patron_recurrencia ADD CONSTRAINT fk_sesion_patron_recurrencia_tipo_recurrencia FOREIGN KEY (tipo_recurrencia) REFERENCES eventos.tipo_recurrencia (id) MATCH FULL ON DELETE RESTRICT ON UPDATE CASCADE;")
	m.SQL("ALTER TABLE eventos.sesion_patron_recurrencia ADD CONSTRAINT fk_sesion_patron_recurrencia_sesion FOREIGN KEY (sesion) REFERENCES eventos.sesion (id) MATCH FULL ON DELETE RESTRICT ON UPDATE CASCADE;")
	m.SQL("ALTER TABLE eventos.relacion_sesiones ADD CONSTRAINT fk_relacion_sesiones_sesion_padre FOREIGN KEY (sesion_padre) REFERENCES eventos.sesion (id) MATCH FULL ON DELETE RESTRICT ON UPDATE CASCADE;")
	m.SQL("ALTER TABLE eventos.participante_sesion ADD CONSTRAINT uq_participante_sesion UNIQUE (sesion,rol_participante_sesion,ente);")
	m.SQL("ALTER TABLE eventos.relacion_sesiones ADD CONSTRAINT uq_sesion_padre_sesion_hijo UNIQUE (sesion_padre,sesion_hijo);")
	m.SQL("ALTER TABLE eventos.sesion_patron_recurrencia ADD CONSTRAINT uq_sesion_patron_recurrencia UNIQUE (tipo_recurrencia,sesion);")
	m.SQL("ALTER TABLE eventos.relacion_sesiones ADD CONSTRAINT fk_relacion_sesiones_sesion_hijo FOREIGN KEY (sesion_hijo) REFERENCES eventos.sesion (id) MATCH FULL ON DELETE RESTRICT ON UPDATE CASCADE;")



}

// Reverse the migrations
func (m *CrearFkTablasSesiones_20191217_175220) Down() {
	m.SQL("ALTER TABLE eventos.sesion DROP CONSTRAINT IF EXISTS fk_sesion_tipo_sesion CASCADE;")
	m.SQL("ALTER TABLE eventos.participante_sesion DROP CONSTRAINT IF EXISTS fk_participante_sesion_sesion CASCADE;")
	m.SQL("ALTER TABLE eventos.participante_sesion DROP CONSTRAINT IF EXISTS fk_participante_sesion_rol_participante_sesion CASCADE;")
	m.SQL("ALTER TABLE eventos.sesion_patron_recurrencia DROP CONSTRAINT IF EXISTS fk_sesion_patron_recurrencia_tipo_recurrencia CASCADE;")
	m.SQL("ALTER TABLE eventos.sesion_patron_recurrencia DROP CONSTRAINT IF EXISTS fk_sesion_patron_recurrencia_sesion CASCADE;")
	m.SQL("ALTER TABLE eventos.relacion_sesiones DROP CONSTRAINT IF EXISTS fk_relacion_sesiones_sesion_padre CASCADE;")
	m.SQL("ALTER TABLE eventos.participante_sesion DROP CONSTRAINT IF EXISTS uq_participante_sesion CASCADE;")
	m.SQL("ALTER TABLE eventos.relacion_sesiones DROP CONSTRAINT IF EXISTS uq_sesion_padre_sesion_hijo CASCADE;")
	m.SQL("ALTER TABLE eventos.sesion_patron_recurrencia DROP CONSTRAINT IF EXISTS uq_sesion_patron_recurrencia CASCADE;")
	m.SQL("ALTER TABLE eventos.relacion_sesiones DROP CONSTRAINT IF EXISTS fk_relacion_sesiones_sesion_hijo CASCADE;")


}
