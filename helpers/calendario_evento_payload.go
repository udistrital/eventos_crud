package helpers

import (
	"encoding/json"
	"errors"

	"github.com/udistrital/eventos_crud/models"
)

func ParseCalendarioEventoPayload(body []byte) (models.CalendarioEvento, map[string]bool, int, error) {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(body, &raw); err != nil {
		return models.CalendarioEvento{}, nil, 0, err
	}

	var evento models.CalendarioEvento
	if err := json.Unmarshal(body, &evento); err != nil {
		return models.CalendarioEvento{}, nil, 0, err
	}

	terceroID, err := parseTerceroID(raw)
	if err != nil {
		return models.CalendarioEvento{}, nil, 0, err
	}

	fields := make(map[string]bool, len(raw))
	for field := range raw {
		fields[field] = true
	}

	return evento, fields, terceroID, nil
}

func ParseTerceroIDFromBody(body []byte) (int, error) {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(body, &raw); err != nil {
		return 0, err
	}
	return parseTerceroID(raw)
}

func parseTerceroID(raw map[string]json.RawMessage) (int, error) {
	var terceroID int
	if value, ok := raw["TerceroId"]; ok {
		if err := json.Unmarshal(value, &terceroID); err != nil {
			return 0, err
		}
	}
	if terceroID <= 0 {
		return 0, errors.New("TerceroId debe ser un entero mayor que cero")
	}
	return terceroID, nil
}
