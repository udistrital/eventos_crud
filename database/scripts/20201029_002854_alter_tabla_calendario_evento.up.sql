ALTER TABLE eventos.calendario_evento ADD nombre varchar(50) NOT NULL;
COMMENT ON COLUMN eventos.calendario_evento.nombre IS 'Nombre del tipo de sesion';

ALTER TABLE eventos.calendario_evento DROP COLUMN periodo_id;

ALTER TABLE eventos.calendario_evento DROP COLUMN aplicacion_id;

ALTER TABLE eventos.calendario_evento ADD dependencia_id json;
COMMENT ON COLUMN eventos.calendario_evento.dependencia_id IS 'se referencia al esquema de dependencia del sistema OKIOS';
