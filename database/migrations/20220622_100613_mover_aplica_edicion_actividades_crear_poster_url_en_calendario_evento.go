package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type MoverAplicaEdicionActividadesCrearPosterUrlEnCalendarioEvento_20220622_100613 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &MoverAplicaEdicionActividadesCrearPosterUrlEnCalendarioEvento_20220622_100613{}
	m.Created = "20220622_100613"

	migration.Register("MoverAplicaEdicionActividadesCrearPosterUrlEnCalendarioEvento_20220622_100613", m)
}

// Run the migrations
func (m *MoverAplicaEdicionActividadesCrearPosterUrlEnCalendarioEvento_20220622_100613) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20220622_100613_mover_aplica_edicion_actividades_crear_poster_url_en_calendario_evento.up.sql")

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
func (m *MoverAplicaEdicionActividadesCrearPosterUrlEnCalendarioEvento_20220622_100613) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20220622_100613_mover_aplica_edicion_actividades_crear_poster_url_en_calendario_evento.down.sql")

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
