package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaRolParticipanteSesion_20191217_174436 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaRolParticipanteSesion_20191217_174436{}
	m.Created = "20191217_174436"

	migration.Register("CrearTablaRolParticipanteSesion_20191217_174436", m)
}

// Run the migrations
func (m *CrearTablaRolParticipanteSesion_20191217_174436) Up() {
	m.SQL("CREATE TABLE eventos.rol_participante_sesion(id serial NOT NULL,nombre varchar(50) NOT NULL,descripcion varchar(250),codigo_abreviacion varchar(20),activo boolean NOT NULL,numero_orden numeric(5,2),CONSTRAINT pk_rol_participante_sesion PRIMARY KEY (id),CONSTRAINT uq_nombre_rol_participante_sesion UNIQUE (nombre));")
	m.SQL("ALTER TABLE eventos.rol_participante_sesion OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE eventos.rol_participante_sesion IS 'Permite almacenar los diferentes roles que puede tener un participante en una sesión';")
	m.SQL("COMMENT ON COLUMN eventos.rol_participante_sesion.id IS 'Identificador de la tabla rol_participante_sesion';")
	m.SQL("COMMENT ON COLUMN eventos.rol_participante_sesion.nombre IS 'Nombre del rol que puede tener un participante de una sesion';")
	m.SQL("COMMENT ON COLUMN eventos.rol_participante_sesion.descripcion IS 'Descripcion que permite especificar el rol y las funciones en la sesion';")
	m.SQL("COMMENT ON COLUMN eventos.rol_participante_sesion.codigo_abreviacion IS 'Codigo de abreviacion del nombre del rol que cumple un participante en la sesion';")
	m.SQL("COMMENT ON COLUMN eventos.rol_participante_sesion.activo IS 'Booleano que permite identificar si el rol se encuentra activo o no';")
	m.SQL("COMMENT ON COLUMN eventos.rol_participante_sesion.numero_orden IS 'Numero de orden para el rol_participante_sesion';")
	m.SQL("COMMENT ON CONSTRAINT uq_nombre_rol_participante_sesion ON eventos.rol_participante_sesion  IS 'UK para garantizar unicidad del nombre';")
}

// Reverse the migrations
func (m *CrearTablaRolParticipanteSesion_20191217_174436) Down() {
	m.SQL("DROP TABLE IF EXISTS eventos.rol_participante_sesion")


}
