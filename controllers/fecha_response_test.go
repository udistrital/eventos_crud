package controllers

import "testing"

func TestFechaRFC3339VisibleLimpiaFechaGoDobleZona(t *testing.T) {
	fecha := fechaRFC3339Visible("2027-05-27 23:59:00 +0000 +0000")
	if fecha != "2027-05-27T23:59:00Z" {
		t.Fatalf("se esperaba RFC3339 preservando hora visible, obtuvo %s", fecha)
	}
}

func TestFechaRFC3339VisiblePreservaFracciones(t *testing.T) {
	fecha := fechaRFC3339Visible("2026-07-10 10:05:50.821618 +0000 +0000")
	if fecha != "2026-07-10T10:05:50.821618Z" {
		t.Fatalf("se esperaba RFC3339 con fracciones, obtuvo %s", fecha)
	}
}
