package models

import (
	"encoding/json"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

func auditCalendarioEventoSnapshot(evento *CalendarioEvento) map[string]interface{} {
	return map[string]interface{}{
		"Id":                evento.Id,
		"FechaCreacion":     auditTime(evento.FechaCreacion),
		"FechaModificacion": auditTime(evento.FechaModificacion),
		"FechaInicio":       auditTime(evento.FechaInicio),
		"FechaFin":          auditTime(evento.FechaFin),
		"Activo":            evento.Activo,
		"DependenciaId":     auditDependencia(evento.DependenciaId),
		"ProcesoId":         auditProcesoID(evento.ProcesoId),
		"EventoCatalogoId":  auditEventoCatalogoID(evento.EventoCatalogoId),
		"UbicacionId":       evento.UbicacionId,
		"PosterUrl":         evento.PosterUrl,
	}
}

func auditTime(value time.Time) interface{} {
	if value.IsZero() {
		return nil
	}
	return value.UTC().Format(time.RFC3339Nano)
}

func auditDependencia(value string) interface{} {
	value = strings.TrimSpace(value)
	if value == "" {
		return nil
	}

	var decoded interface{}
	if err := json.Unmarshal([]byte(value), &decoded); err == nil {
		return decoded
	}
	return value
}

func auditProcesoID(value *Proceso) interface{} {
	if value == nil {
		return nil
	}
	return value.Id
}

func auditEventoCatalogoID(value *EventoCatalogo) interface{} {
	if value == nil {
		return nil
	}
	return value.Id
}

func auditCalendarioEventoChanges(before, after map[string]interface{}) map[string]interface{} {
	changes := make(map[string]interface{})
	globales := make(map[string]interface{})
	for _, field := range []string{"FechaInicio", "FechaFin", "Activo"} {
		if !reflect.DeepEqual(before[field], after[field]) {
			globales[field] = map[string]interface{}{
				"anterior": before[field],
				"nuevo":    after[field],
			}
		}
	}
	if len(globales) > 0 {
		changes["globales"] = globales
	}

	otros := make(map[string]interface{})
	for _, field := range []string{"ProcesoId", "EventoCatalogoId", "UbicacionId", "PosterUrl"} {
		if !reflect.DeepEqual(before[field], after[field]) {
			otros[field] = map[string]interface{}{
				"anterior": before[field],
				"nuevo":    after[field],
			}
		}
	}
	if len(otros) > 0 {
		changes["otros"] = otros
	}

	dependenciaChanges := auditDependenciaChanges(before["DependenciaId"], after["DependenciaId"])
	for key, value := range dependenciaChanges {
		changes[key] = value
	}

	return changes
}

func auditDependenciaChanges(beforeValue, afterValue interface{}) map[string]interface{} {
	before, _ := beforeValue.(map[string]interface{})
	after, _ := afterValue.(map[string]interface{})
	changes := make(map[string]interface{})

	previousProjects := auditProjectSet(before["proyectos"])
	currentProjects := auditProjectSet(after["proyectos"])
	added := differenceProjectSet(currentProjects, previousProjects)
	removed := differenceProjectSet(previousProjects, currentProjects)
	if len(added) > 0 || len(removed) > 0 {
		changes["asociaciones"] = map[string]interface{}{
			"agregados": added,
			"retirados": removed,
		}
	}

	previousDates := auditParticularDates(before["fechas"])
	currentDates := auditParticularDates(after["fechas"])
	programIDs := make(map[int]struct{}, len(previousDates)+len(currentDates))
	for id := range previousDates {
		programIDs[id] = struct{}{}
	}
	for id := range currentDates {
		programIDs[id] = struct{}{}
	}
	ids := make([]int, 0, len(programIDs))
	for id := range programIDs {
		ids = append(ids, id)
	}
	sort.Ints(ids)

	particulares := make([]map[string]interface{}, 0)
	for _, id := range ids {
		previous, previousOK := previousDates[id]
		current, currentOK := currentDates[id]
		if previousOK && currentOK && reflect.DeepEqual(previous, current) {
			continue
		}
		particulares = append(particulares, map[string]interface{}{
			"programa_id": id,
			"anterior":    auditOptionalValue(previous, previousOK),
			"nuevo":       auditOptionalValue(current, currentOK),
		})
	}
	if len(particulares) > 0 {
		changes["particulares"] = particulares
	}

	return changes
}

func auditProjectSet(value interface{}) map[int]struct{} {
	projects := make(map[int]struct{})
	values, _ := value.([]interface{})
	for _, raw := range values {
		if id, ok := auditInteger(raw); ok {
			projects[id] = struct{}{}
		}
	}
	return projects
}

func auditParticularDates(value interface{}) map[int]map[string]interface{} {
	dates := make(map[int]map[string]interface{})
	values, _ := value.([]interface{})
	for _, raw := range values {
		date, ok := raw.(map[string]interface{})
		if !ok {
			continue
		}
		id, ok := auditInteger(date["Id"])
		if ok {
			dates[id] = date
		}
	}
	return dates
}

func auditInteger(value interface{}) (int, bool) {
	switch number := value.(type) {
	case float64:
		return int(number), number == float64(int(number))
	case int:
		return number, true
	case json.Number:
		parsed, err := strconv.Atoi(string(number))
		return parsed, err == nil
	case string:
		parsed, err := strconv.Atoi(number)
		return parsed, err == nil
	default:
		return 0, false
	}
}

func differenceProjectSet(source, excluded map[int]struct{}) []int {
	result := make([]int, 0)
	for id := range source {
		if _, exists := excluded[id]; !exists {
			result = append(result, id)
		}
	}
	sort.Ints(result)
	return result
}

func auditOptionalValue(value map[string]interface{}, exists bool) interface{} {
	if !exists {
		return nil
	}
	return value
}

func nullableString(value []byte) *string {
	if len(value) == 0 {
		return nil
	}
	result := string(value)
	return &result
}
