package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type EliminarCamposTablaTipoEvento_20201028_230638 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &EliminarCamposTablaTipoEvento_20201028_230638{}
	m.Created = "20201028_230638"

	migration.Register("EliminarCamposTablaTipoEvento_20201028_230638", m)
}

// Run the migrations
func (m *EliminarCamposTablaTipoEvento_20201028_230638) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20201028_230638_eliminar_campos_tabla_tipo_evento.up.sql")

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
func (m *EliminarCamposTablaTipoEvento_20201028_230638) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20201028_230638_eliminar_campos_tabla_tipo_evento.down.sql")

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
