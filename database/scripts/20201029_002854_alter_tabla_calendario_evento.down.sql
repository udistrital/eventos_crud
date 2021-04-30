ALTER TABLE eventos.calendario_evento DROP COLUMN nombre;

ALTER TABLE eventos.calendario_evento ADD COLUMN periodo_id integer NOT NULL;
COMMENT ON COLUMN eventos.calendario_evento.periodo_id IS 'Se almacena el periodo al cual pertenece el evento';

-- ALTER TABLE eventos.calendario_evento ADD COLUMN aplicacion_id integer NOT NULL;
-- COMMENT ON COLUMN eventos.calendario_evento.aplicacion_id IS 'Campo obligatorio para realcionar el evento creado a una aplicacion. Es un id relacionado a la tabla de aplicacion en el esquema configuracion';

ALTER TABLE eventos.calendario_evento DROP COLUMN dependencia_id;