ALTER TABLE eventos.calendario DROP COLUMN aplica_edicion_actividades;


ALTER TABLE eventos.calendario_evento ADD COLUMN aplica_edicion_actividades boolean;

COMMENT ON COLUMN eventos.calendario_evento.aplica_edicion_actividades IS 'Campo para verficar si al calendario aplica edicion de fechas actividades';

ALTER TABLE eventos.calendario_evento ADD COLUMN poster_url character varying(250);

COMMENT ON COLUMN eventos.calendario_evento.poster_url IS 'Campo string para url poster';