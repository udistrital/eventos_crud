ALTER TABLE eventos.calendario ADD COLUMN aplica_edicion_actividades boolean;

COMMENT ON COLUMN eventos.calendario.aplica_edicion_actividades IS 'Campo para verficar si al calendario aplica edicion de fechas actividades';


ALTER TABLE eventos.calendario_evento DROP COLUMN aplica_edicion_actividades;

ALTER TABLE eventos.calendario_evento DROP COLUMN poster_url;