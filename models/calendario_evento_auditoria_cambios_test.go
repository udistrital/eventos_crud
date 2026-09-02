package models

import "testing"

func TestAuditCalendarioEventoSnapshotIncluyeNumeroOcurrencia(t *testing.T) {
	evento := &CalendarioEvento{Id: 7, NumeroOcurrencia: 3}

	snapshot := auditCalendarioEventoSnapshot(evento)

	if snapshot["NumeroOcurrencia"] != 3 {
		t.Fatalf("numero de ocurrencia inesperado: %v", snapshot["NumeroOcurrencia"])
	}
}

func TestAuditCalendarioEventoChangesRegistraNumeroOcurrencia(t *testing.T) {
	before := auditCalendarioEventoSnapshot(&CalendarioEvento{NumeroOcurrencia: 1})
	after := auditCalendarioEventoSnapshot(&CalendarioEvento{NumeroOcurrencia: 2})

	changes := auditCalendarioEventoChanges(before, after)
	otros, ok := changes["otros"].(map[string]interface{})
	if !ok {
		t.Fatal("se esperaba el grupo de cambios otros")
	}
	if _, ok := otros["NumeroOcurrencia"]; !ok {
		t.Fatal("se esperaba el cambio de NumeroOcurrencia")
	}
}

func TestApplyCalendarioEventoFieldsIgnoraNumeroOcurrencia(t *testing.T) {
	current := &CalendarioEvento{NumeroOcurrencia: 4}
	incoming := &CalendarioEvento{NumeroOcurrencia: 99}

	applyCalendarioEventoFields(current, incoming, map[string]bool{"NumeroOcurrencia": true})

	if current.NumeroOcurrencia != 4 {
		t.Fatalf("NumeroOcurrencia es administrado por el servidor: %d", current.NumeroOcurrencia)
	}
}
