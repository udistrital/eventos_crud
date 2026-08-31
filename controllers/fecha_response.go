package controllers

import (
	"regexp"
	"strings"
	"time"

	"github.com/udistrital/eventos_crud/models"
	"github.com/udistrital/utils_oas/time_bogota"
)

var fechaHoraVisibleRegexp = regexp.MustCompile(`^(\d{4}-\d{2}-\d{2})[T ](\d{2}:\d{2}:\d{2}(?:\.\d+)?)`)

func fechaRFC3339Visible(fecha string) string {
	fecha = strings.TrimSpace(fecha)
	if fecha == "" {
		return fecha
	}
	match := fechaHoraVisibleRegexp.FindStringSubmatch(fecha)
	if len(match) != 3 {
		return fecha
	}
	return match[1] + "T" + match[2] + "Z"
}

func fechaActual() time.Time {
	fecha, err := fechaTimeVisible(time_bogota.TiempoBogotaFormato())
	if err != nil {
		return fechaUTCVisible(time.Now())
	}
	return fecha
}

func fechaCorreccion(fecha time.Time) time.Time {
	if fecha.IsZero() {
		return fechaActual()
	}
	return fechaUTCVisible(fecha)
}

func fechaUTCVisible(fecha time.Time) time.Time {
	if fecha.IsZero() {
		return fecha
	}
	return time.Date(fecha.Year(), fecha.Month(), fecha.Day(), fecha.Hour(), fecha.Minute(), fecha.Second(), fecha.Nanosecond(), time.UTC)
}

func fechaTimeVisible(fecha string) (time.Time, error) {
	fecha = fechaRFC3339Visible(fecha)
	return time.Parse(time.RFC3339Nano, fecha)
}

func normalizarCalendarioEventoExtensionRespuesta(value interface{}) interface{} {
	switch v := value.(type) {
	case *models.CalendarioEventoExtension:
		normalizarCalendarioEventoExtension(v)
		return v
	case models.CalendarioEventoExtension:
		normalizarCalendarioEventoExtension(&v)
		return v
	case []interface{}:
		for i := range v {
			v[i] = normalizarCalendarioEventoExtensionRespuesta(v[i])
		}
		return v
	case map[string]interface{}:
		normalizarFechaMapa(v, "FechaFin")
		normalizarFechaMapa(v, "FechaCreacion")
		normalizarFechaMapa(v, "FechaModificacion")
		return v
	default:
		return value
	}
}

func normalizarCalendarioEventoExtensionProgramaRespuesta(value interface{}) interface{} {
	switch v := value.(type) {
	case *models.CalendarioEventoExtensionPrograma:
		normalizarCalendarioEventoExtensionPrograma(v)
		return v
	case models.CalendarioEventoExtensionPrograma:
		normalizarCalendarioEventoExtensionPrograma(&v)
		return v
	case []interface{}:
		for i := range v {
			v[i] = normalizarCalendarioEventoExtensionProgramaRespuesta(v[i])
		}
		return v
	case map[string]interface{}:
		normalizarFechaMapa(v, "FechaCreacion")
		normalizarFechaMapa(v, "FechaModificacion")
		if extension, ok := v["CalendarioEventoExtensionId"].(map[string]interface{}); ok {
			normalizarCalendarioEventoExtensionRespuesta(extension)
		}
		if extension, ok := v["ExtensionPadreId"].(map[string]interface{}); ok {
			normalizarCalendarioEventoExtensionRespuesta(extension)
		}
		return v
	default:
		return value
	}
}

func normalizarCalendarioEventoExtension(v *models.CalendarioEventoExtension) {
	if v == nil {
		return
	}
	v.FechaFin = fechaUTCVisible(v.FechaFin)
	v.FechaCreacion = fechaUTCVisible(v.FechaCreacion)
	v.FechaModificacion = fechaUTCVisible(v.FechaModificacion)
}

func normalizarCalendarioEventoExtensionPrograma(v *models.CalendarioEventoExtensionPrograma) {
	if v == nil {
		return
	}
	v.FechaCreacion = fechaUTCVisible(v.FechaCreacion)
	v.FechaModificacion = fechaUTCVisible(v.FechaModificacion)
	normalizarCalendarioEventoExtension(v.CalendarioEventoExtensionId)
	normalizarCalendarioEventoExtension(v.ExtensionPadreId)
}

func normalizarFechaMapa(m map[string]interface{}, campo string) {
	if value, ok := m[campo].(string); ok {
		m[campo] = fechaRFC3339Visible(value)
		return
	}
	if value, ok := m[campo].(time.Time); ok {
		m[campo] = fechaUTCVisible(value)
	}
}
