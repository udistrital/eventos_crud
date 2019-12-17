package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaSesionPatronRecurrencia_20191217_174754 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaSesionPatronRecurrencia_20191217_174754{}
	m.Created = "20191217_174754"

	migration.Register("CrearTablaSesionPatronRecurrencia_20191217_174754", m)
}

// Run the migrations
func (m *CrearTablaSesionPatronRecurrencia_20191217_174754) Up() {
	m.SQL("CREATE TABLE eventos.sesion_patron_recurrencia(id serial NOT NULL,tipo_recurrencia integer NOT NULL,sesion integer NOT NULL,valor varchar(100) NOT NULL,CONSTRAINT pk_sesion_patron_recurrencia PRIMARY KEY (id));") 
	m.SQL("ALTER TABLE eventos.sesion_patron_recurrencia OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE eventos.sesion_patron_recurrencia IS 'Permite almacenar los valores del patron de recurrencia de la sesion';")
	m.SQL("COMMENT ON COLUMN eventos.sesion_patron_recurrencia.id IS 'Identificador de la tabla';")
	m.SQL("COMMENT ON COLUMN eventos.sesion_patron_recurrencia.valor IS 'Valor que se establece para las recurrencias ';")

}

// Reverse the migrations
func (m *CrearTablaSesionPatronRecurrencia_20191217_174754) Down() {
	m.SQL("DROP TABLE IF EXISTS eventos.sesion_patron_recurrencia")


}
