package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaCalendarioEventoTipoPublico_20201109_190837 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaCalendarioEventoTipoPublico_20201109_190837{}
	m.Created = "20201109_190837"

	migration.Register("CrearTablaCalendarioEventoTipoPublico_20201109_190837", m)
}

// Run the migrations
func (m *CrearTablaCalendarioEventoTipoPublico_20201109_190837) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20201109_190837_crear_tabla_calendario_evento_tipo_publico.up.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}

}

// Reverse the migrations
func (m *CrearTablaCalendarioEventoTipoPublico_20201109_190837) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20201109_190837_crear_tabla_calendario_evento_tipo_publico.down.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}
}
