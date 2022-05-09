 
package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AgregarCamposExtension_20220507_093403 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AgregarCamposExtension_20220507_093403{}
	m.Created = "20220507_093403"

	migration.Register("AgregarCamposExtension_20220507_093403", m)
}

// Run the migrations
func (m *AgregarCamposExtension_20220507_093403) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20220507_093403_agregar_campos_extension.up.sql")

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
func (m *AgregarCamposExtension_20220507_093403) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20220507_093403_agregar_campos_extension.down.sql")

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
