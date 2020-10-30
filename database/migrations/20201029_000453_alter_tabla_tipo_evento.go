package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AlterTablaTipoVento_20201029_000453 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AlterTablaTipoVento_20201029_000453{}
	m.Created = "20201029_000453"

	migration.Register("AlterTablaTipoVento_20201029_000453", m)
}

// Run the migrations
func (m *AlterTablaTipoVento_20201029_000453) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20201029_000453_alter_tabla_tipo_evento.up.sql")

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
func (m *AlterTablaTipoVento_20201029_000453) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20201029_000453_alter_tabla_tipo_evento.down.sql")

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
