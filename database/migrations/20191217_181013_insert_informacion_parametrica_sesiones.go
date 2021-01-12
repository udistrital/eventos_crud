package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertInformacionParametricaSesiones_20191217_181013 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertInformacionParametricaSesiones_20191217_181013{}
	m.Created = "20191217_181013"

	migration.Register("InsertInformacionParametricaSesiones_20191217_181013", m)
}

// Run the migrations
func (m *InsertInformacionParametricaSesiones_20191217_181013) Up() {

	m.SQL("INSERT INTO eventos.rol_participante_sesion (id, nombre, descripcion, codigo_abreviacion, activo, numero_orden) VALUES (1, 'Creador', 'Creador del evento', 'CE', true, 1.00);")
	m.SQL("INSERT INTO eventos.rol_participante_sesion (id, nombre, descripcion, codigo_abreviacion, activo, numero_orden) VALUES (2, 'Participante', 'Participante básico del evento', 'PE', true, 2.00);")
	m.SQL("INSERT INTO eventos.tipo_sesion (id, nombre, descripcion, codigo_abreviacion, activo, numero_orden) VALUES (1, 'Sesiones de modalidad de materias de posgrado', 'Sesiones relacionadas con la modalidad de grado de materias de posgrado', 'EMP', true, 1.00);")
	m.SQL("INSERT INTO eventos.tipo_sesion (id, nombre, descripcion, codigo_abreviacion, activo, numero_orden) VALUES (2, 'Publicación de materias modalidad posgrado', 'Publicación de los espacios academicos por parte de los proyectos curriculares', 'PMP', true, 2.00);")
	m.SQL("INSERT INTO eventos.tipo_sesion (id, nombre, descripcion, codigo_abreviacion, activo, numero_orden) VALUES (3, 'Solicitud de materias de posgrado', 'Fechas para que los estudiantes realicen la solicitud de los espacios academicos para cursar la modadlidad de materias de posgrado', 'SMP', true, 3.00);")
	m.SQL("INSERT INTO eventos.tipo_sesion (id, nombre, descripcion, codigo_abreviacion, activo, numero_orden) VALUES (4, 'Primera fecha de selección de admitidos a materias', 'Primer corte para que los posgrados realicen la primera selección de estudiantes a los posgrados', 'PFSAMP', true, 4.00);")
	m.SQL("INSERT INTO eventos.tipo_sesion (id, nombre, descripcion, codigo_abreviacion, activo, numero_orden) VALUES (5, 'Primera fecha de formalización de solicitudes', 'Primera fecha para formalizar solicitudes de materias de posgrado', 'PFFMP', true, 5.00);")
	m.SQL("INSERT INTO eventos.tipo_sesion (id, nombre, descripcion, codigo_abreviacion, activo, numero_orden) VALUES (6, 'Segunda fecha de selección de admitidos a materias', 'Segundo corte para que los posgrados realicen la primera selección de estudiantes a los posgrados', 'SCPPE', true, 6.00);")
	m.SQL("INSERT INTO eventos.tipo_sesion (id, nombre, descripcion, codigo_abreviacion, activo, numero_orden) VALUES (7, 'Segunda fecha de formalización de solicitudes', 'Segunda fecha para formalizar solicitudes de materias de posgrado', 'SFFMP', true, 7.00);")
	m.SQL("INSERT INTO eventos.tipo_sesion (id, nombre, descripcion, codigo_abreviacion, activo, numero_orden) VALUES (8, 'Aprobación de solicitudes iniciales de materias de', 'Aprobación de solicitudes iniciales de materias de posgrado', 'ASIMP', true, 8.00);")

}

// Reverse the migrations
func (m *InsertInformacionParametricaSesiones_20191217_181013) Down() {
	m.SQL("DELETE FROM eventos.rol_participante_sesion")
	m.SQL("DELETE FROM eventos.tipo_sesion")
}
