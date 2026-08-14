package models

import (
	"encoding/json"
	"time"

	"github.com/astaxie/beego/orm"
)

const (
	calendarioEventoOperacionInsert     = "INSERT"
	calendarioEventoOperacionUpdate     = "UPDATE"
	calendarioEventoOperacionDeactivate = "DEACTIVATE"
)

func AddCalendarioEventoAuditado(evento *CalendarioEvento, terceroID int) error {
	o := orm.NewOrm()
	if err := o.Begin(); err != nil {
		return err
	}
	if _, err := o.Insert(evento); err != nil {
		_ = o.Rollback()
		return err
	}
	if err := persistCalendarioEventoAudit(o, evento.Id, calendarioEventoOperacionInsert, terceroID, nil, evento); err != nil {
		_ = o.Rollback()
		return err
	}
	if err := o.Commit(); err != nil {
		_ = o.Rollback()
		return err
	}
	return nil
}

func UpdateCalendarioEventoAuditado(
	id int,
	incoming *CalendarioEvento,
	fields map[string]bool,
	terceroID int,
	fechaModificacion time.Time,
) (*CalendarioEvento, error) {
	return updateCalendarioEventoAuditado(id, terceroID, fechaModificacion, "", func(evento *CalendarioEvento) {
		applyCalendarioEventoFields(evento, incoming, fields)
	})
}

func InactivarCalendarioEventoAuditado(id int, terceroID int, fechaModificacion time.Time) error {
	_, err := updateCalendarioEventoAuditado(id, terceroID, fechaModificacion, calendarioEventoOperacionDeactivate, func(evento *CalendarioEvento) {
		evento.Activo = false
	})
	return err
}

func updateCalendarioEventoAuditado(
	id int,
	terceroID int,
	fechaModificacion time.Time,
	operacion string,
	mutate func(*CalendarioEvento),
) (*CalendarioEvento, error) {
	o := orm.NewOrm()
	if err := o.Begin(); err != nil {
		return nil, err
	}

	before, err := getCalendarioEventoByIDForUpdate(o, id)
	if err != nil {
		_ = o.Rollback()
		return nil, err
	}

	updated := *before
	mutate(&updated)
	updated.Id = id
	updated.FechaCreacion = before.FechaCreacion
	updated.FechaModificacion = fechaModificacion

	if _, err := o.Update(&updated); err != nil {
		_ = o.Rollback()
		return nil, err
	}

	operacion = calendarioEventoOperacionCambio(operacion, before, &updated)
	if err := persistCalendarioEventoAudit(o, id, operacion, terceroID, before, &updated); err != nil {
		_ = o.Rollback()
		return nil, err
	}
	if err := o.Commit(); err != nil {
		_ = o.Rollback()
		return nil, err
	}

	return &updated, nil
}

func calendarioEventoOperacionCambio(operacion string, before, after *CalendarioEvento) string {
	if operacion != "" {
		return operacion
	}
	if before.Activo && !after.Activo {
		return calendarioEventoOperacionDeactivate
	}
	return calendarioEventoOperacionUpdate
}

func getCalendarioEventoByIDForUpdate(o orm.Ormer, id int) (*CalendarioEvento, error) {
	var lockedID int
	if err := o.Raw("SELECT id FROM eventos.calendario_evento WHERE id = ? FOR UPDATE", id).QueryRow(&lockedID); err != nil {
		return nil, err
	}
	evento := &CalendarioEvento{Id: lockedID}
	if err := o.Read(evento); err != nil {
		return nil, err
	}
	return evento, nil
}

func applyCalendarioEventoFields(current, incoming *CalendarioEvento, fields map[string]bool) {
	if fields["FechaInicio"] {
		current.FechaInicio = incoming.FechaInicio
	}
	if fields["FechaFin"] {
		current.FechaFin = incoming.FechaFin
	}
	if fields["Activo"] {
		current.Activo = incoming.Activo
	}
	if fields["DependenciaId"] {
		current.DependenciaId = incoming.DependenciaId
	}
	if fields["ProcesoId"] {
		current.ProcesoId = incoming.ProcesoId
	}
	if fields["EventoCatalogoId"] {
		current.EventoCatalogoId = incoming.EventoCatalogoId
	}
	if fields["UbicacionId"] {
		current.UbicacionId = incoming.UbicacionId
	}
	if fields["PosterUrl"] {
		current.PosterUrl = incoming.PosterUrl
	}
}

func persistCalendarioEventoAudit(
	o orm.Ormer,
	calendarioEventoID int,
	operacion string,
	terceroID int,
	before *CalendarioEvento,
	after *CalendarioEvento,
) error {
	var beforeJSON []byte
	var err error
	if before != nil {
		beforeJSON, err = json.Marshal(auditCalendarioEventoSnapshot(before))
		if err != nil {
			return err
		}
	}
	afterJSON, err := json.Marshal(auditCalendarioEventoSnapshot(after))
	if err != nil {
		return err
	}

	changes := make(map[string]interface{})
	if before != nil {
		changes = auditCalendarioEventoChanges(
			auditCalendarioEventoSnapshot(before),
			auditCalendarioEventoSnapshot(after),
		)
	} else {
		changes["creacion"] = true
	}
	changesJSON, err := json.Marshal(changes)
	if err != nil {
		return err
	}

	return insertCalendarioEventoAuditoria(
		o,
		calendarioEventoID,
		operacion,
		terceroID,
		beforeJSON,
		afterJSON,
		changesJSON,
	)
}

func insertCalendarioEventoAuditoria(
	o orm.Ormer,
	calendarioEventoID int,
	operacion string,
	terceroID int,
	estadoAnterior []byte,
	estadoNuevo []byte,
	cambios []byte,
) error {
	fechaOperacion := time.Now().UTC()
	estadoNuevoString := string(estadoNuevo)
	cambiosString := string(cambios)
	auditoria := &CalendarioEventoAuditoria{
		CalendarioEventoId: &CalendarioEvento{Id: calendarioEventoID},
		Operacion:          operacion,
		TerceroId:          terceroID,
		FechaCreacion:      fechaOperacion,
		FechaModificacion:  fechaOperacion,
		Activo:             true,
		EstadoAnterior:     nullableString(estadoAnterior),
		EstadoNuevo:        &estadoNuevoString,
		Cambios:            &cambiosString,
	}
	_, err := o.Insert(auditoria)
	return err
}
