ALTER TABLE eventos.tipo_evento ADD COLUMN dependencia_id integer NOT NULL;
COMMENT ON COLUMN eventos.tipo_evento.dependencia_id IS 'se referencia al esquema de dependencia del sistema OKIOS';