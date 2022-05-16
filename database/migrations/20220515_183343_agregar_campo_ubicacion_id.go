package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AgregarCampoUbicacionId_20220515_183343 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AgregarCampoUbicacionId_20220515_183343{}
	m.Created = "20220515_183343"

	migration.Register("AgregarCampoUbicacionId_20220515_183343", m)
}

// Run the migrations
func (m *AgregarCampoUbicacionId_20220515_183343) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20220515_183343_agregar_campo_ubicacion_id.up.sql")

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
func (m *AgregarCampoUbicacionId_20220515_183343) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20220515_183343_agregar_campo_ubicacion_id.down.sql")

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