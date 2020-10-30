package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaCalendario_20201028_232539 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaCalendario_20201028_232539{}
	m.Created = "20201028_232539"

	migration.Register("CrearTablaCalendario_20201028_232539", m)
}

// Run the migrations
func (m *CrearTablaCalendario_20201028_232539) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20201028_232539_crear_tabla_calendario.up.sql")

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
func (m *CrearTablaCalendario_20201028_232539) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20201028_232539_crear_tabla_calendario.down.sql")

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
